// File generated from our OpenAPI spec by Scalar. See README.md for details.

package ssestream

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Decoder interface {
	Event() Event
	Next() bool
	Close() error
	Err() error
}

func NewDecoder(res *http.Response) Decoder {
	if res == nil || res.Body == nil {
		return nil
	}

	var decoder Decoder
	contentType := normalizeContentType(res.Header.Get("content-type"))
	if t, ok := decoderTypes[contentType]; ok {
		decoder = t(res.Body)
	} else {
		scn := bufio.NewScanner(res.Body)
		scn.Buffer(nil, bufio.MaxScanTokenSize<<9)
		decoder = &eventStreamDecoder{rc: res.Body, scn: scn}
	}
	return decoder
}

var decoderTypes = map[string](func(io.ReadCloser) Decoder){}

func RegisterDecoder(contentType string, decoder func(io.ReadCloser) Decoder) {
	decoderTypes[normalizeContentType(contentType)] = decoder
}

func init() {
	RegisterDecoder("application/x-ndjson", newJSONLinesDecoder)
	RegisterDecoder("application/x-jsonl", newJSONLinesDecoder)
	RegisterDecoder("application/jsonl", newJSONLinesDecoder)
}

func normalizeContentType(contentType string) string {
	mediaType, _, _ := strings.Cut(contentType, ";")
	return strings.ToLower(strings.TrimSpace(mediaType))
}

type Event struct {
	Type string
	Data []byte
}

// A base implementation of a Decoder for text/event-stream.
type eventStreamDecoder struct {
	evt Event
	rc  io.ReadCloser
	scn *bufio.Scanner
	err error
}

func (s *eventStreamDecoder) Next() bool {
	if s.err != nil {
		return false
	}

	event := ""
	data := bytes.NewBuffer(nil)

	for s.scn.Scan() {
		txt := s.scn.Bytes()

		// Dispatch event on an empty line
		if len(txt) == 0 {
			s.evt = Event{
				Type: event,
				Data: data.Bytes(),
			}
			return true
		}

		// Split a string like "event: bar" into name="event" and value=" bar".
		name, value, _ := bytes.Cut(txt, []byte(":"))

		// Consume an optional space after the colon if it exists.
		if len(value) > 0 && value[0] == ' ' {
			value = value[1:]
		}

		switch string(name) {
		case "":
			// An empty line in the for ": something" is a comment and should be ignored.
			continue
		case "event":
			event = string(value)
		case "data":
			_, s.err = data.Write(value)
			if s.err != nil {
				return false
			}
			_, s.err = data.WriteRune('\n')
			if s.err != nil {
				return false
			}
		}
	}

	if s.scn.Err() != nil {
		s.err = s.scn.Err()
	}

	return false
}

func (s *eventStreamDecoder) Event() Event {
	return s.evt
}

func (s *eventStreamDecoder) Close() error {
	return s.rc.Close()
}

func (s *eventStreamDecoder) Err() error {
	return s.err
}

type jsonLinesDecoder struct {
	evt Event
	rc  io.ReadCloser
	scn *bufio.Scanner
	err error
}

func newJSONLinesDecoder(rc io.ReadCloser) Decoder {
	scn := bufio.NewScanner(rc)
	scn.Buffer(nil, bufio.MaxScanTokenSize<<9)
	return &jsonLinesDecoder{rc: rc, scn: scn}
}

func (s *jsonLinesDecoder) Next() bool {
	if s.err != nil {
		return false
	}
	for s.scn.Scan() {
		line := bytes.TrimSpace(s.scn.Bytes())
		if len(line) == 0 {
			continue
		}
		s.evt = Event{Data: append([]byte(nil), line...)}
		return true
	}
	if s.scn.Err() != nil {
		s.err = s.scn.Err()
	}
	return false
}

func (s *jsonLinesDecoder) Event() Event {
	return s.evt
}

func (s *jsonLinesDecoder) Close() error {
	return s.rc.Close()
}

func (s *jsonLinesDecoder) Err() error {
	return s.err
}

type Stream[T any] struct {
	decoder  Decoder
	handlers []EventHandler
	cur      T
	err      error
}

type EventHandler struct {
	Kind           string
	Handle         string
	EventTypes     []string
	EventTypeNull  bool
	ErrorProperty  string
	DataStartsWith string
}

func NewStream[T any](decoder Decoder, err error, handlers ...EventHandler) *Stream[T] {
	return &Stream[T]{
		decoder:  decoder,
		handlers: handlers,
		err:      err,
	}
}

// Next returns false if the stream has ended or an error occurred.
// Call Stream.Current() to get the current value.
// Call Stream.Err() to get the error.
//
//		for stream.Next() {
//			data := stream.Current()
//		}
//
//	 	if stream.Err() != nil {
//			...
//	 	}
func (s *Stream[T]) Next() bool {
	if s.err != nil {
		return false
	}

	for s.decoder.Next() {
		event := s.decoder.Event()
		handler := streamHandlerForEvent(event, s.handlers)
		if handler != nil {
			switch handler.Handle {
			case "continue", "ignore":
				continue
			case "error":
				s.err = streamError(event, handler)
				return false
			case "yield":
				// handled below
			default:
				if len(s.handlers) > 0 {
					continue
				}
			}
		} else if len(s.handlers) > 0 {
			continue
		}

		var nxt T
		s.err = json.Unmarshal(event.Data, &nxt)
		if s.err != nil {
			return false
		}
		s.cur = nxt
		return true
	}

	// decoder.Next() may be false because of an error
	s.err = s.decoder.Err()

	return false
}

func (s *Stream[T]) Current() T {
	return s.cur
}

func (s *Stream[T]) Err() error {
	return s.err
}

func (s *Stream[T]) Close() error {
	if s.decoder == nil {
		// already closed
		return nil
	}
	return s.decoder.Close()
}

func streamHandlerForEvent(event Event, handlers []EventHandler) *EventHandler {
	if len(handlers) == 0 {
		return nil
	}
	for i := range handlers {
		handler := &handlers[i]
		if handler.Kind == "data" && handler.DataStartsWith != "" && strings.HasPrefix(string(event.Data), handler.DataStartsWith) {
			return handler
		}
	}
	for i := range handlers {
		handler := &handlers[i]
		if handler.Kind == "event" && eventTypeMatches(event.Type, *handler) {
			return handler
		}
		if handler.Kind == "fallthrough" {
			return handler
		}
	}
	return nil
}

func eventTypeMatches(eventType string, handler EventHandler) bool {
	if handler.EventTypeNull {
		return eventType == ""
	}
	for _, expected := range handler.EventTypes {
		if eventType == expected {
			return true
		}
	}
	return false
}

func streamError(event Event, handler *EventHandler) error {
	if handler.ErrorProperty == "" {
		return fmt.Errorf("received error while streaming: %s", string(event.Data))
	}
	var body map[string]any
	if err := json.Unmarshal(event.Data, &body); err == nil {
		if value, ok := body[handler.ErrorProperty]; ok {
			return fmt.Errorf("received error while streaming: %v", value)
		}
	}
	return fmt.Errorf("received error while streaming: %s", string(event.Data))
}
