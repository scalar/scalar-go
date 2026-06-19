// File generated from our OpenAPI spec by Scalar. See README.md for details.

package pagination

import (
	"net/http"
	"strings"

	"github.com/tidwall/gjson"

	"github.com/scalar/scalar-go/internal/apijson"
	"github.com/scalar/scalar-go/internal/requestconfig"
	"github.com/scalar/scalar-go/option"
	"github.com/scalar/scalar-go/packages/param"
	"github.com/scalar/scalar-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// CursorPageConfig is the per-operation cursor pagination descriptor that the
// generated list method bakes into the page. It is the Go view of the shared,
// language-neutral descriptor: where the items array lives, how to read the next
// cursor, which request param carries it, and on which carrier (query vs body).
// Driving extraction from this config — rather than hardcoded `items`/`next_cursor`
// struct tags — lets a single page type serve every cursor scheme, including
// `cursor_id` (cursor read from the last item) and body-carried cursors.
type CursorPageConfig struct {
	// ItemsPath walks the response body to the items array (e.g. ["items"], ["data"]).
	ItemsPath []string
	// CursorKind is "field" (a top-level response field) or "item" (a field on the
	// last item — `cursor_id` pagination).
	CursorKind string
	// CursorPath walks to the cursor value: from the body root when CursorKind is
	// "field", or from the last item when CursorKind is "item".
	CursorPath []string
	// CursorParam is the request param wire name that carries the cursor.
	CursorParam string
	// CursorLocation is "query" or "body": the carrier the cursor rides on the next request.
	CursorLocation string
	// ContinueOnEmptyItems keeps paginating across an empty page when set (some schemes
	// serve empty intermediate pages that still carry a next cursor).
	ContinueOnEmptyItems bool
	// HasMorePath walks the response body to the scheme's `has_more` boolean. When
	// present and the field reads `false`, pagination stops without a trailing
	// empty-page request — some APIs' final page still carries a non-empty last-item
	// cursor. An absent field keeps paginating (cursor presence decides).
	HasMorePath []string
}

type CursorPage[T any] struct {
	// Data is the list of items on this page. It is extracted from the configured
	// items path (CursorPageConfig.ItemsPath) after the response is unmarshalled,
	// so the SDK exposes the items regardless of their wire field name. Callers can
	// still ask for the raw response via RawJSON().
	Data []T `json:"-"`
	// JSON contains the raw response; check fields with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg    *requestconfig.RequestConfig
	res    *http.Response
	config CursorPageConfig
}

// HasNextPage reports whether the page has another cursor to follow.
func (r CursorPage[T]) HasNextPage() bool {
	// An empty page is terminal unless the scheme opts into continuing across empties.
	if !r.config.ContinueOnEmptyItems && len(r.Data) == 0 {
		return false
	}
	if !r.hasMore() {
		return false
	}
	_, ok := r.nextCursor()
	return ok
}

// hasMore consults the scheme's `has_more` body field when configured. Only an
// explicit boolean `false` stops pagination; a missing or non-boolean field keeps
// going so schemes without the field fall back to cursor-presence semantics.
func (r CursorPage[T]) hasMore() bool {
	if len(r.config.HasMorePath) == 0 {
		return true
	}
	result := gjson.Get(r.JSON.raw, gjsonPath(r.config.HasMorePath))
	if !result.Exists() || !result.IsBool() {
		return true
	}
	return result.Bool()
}

// RawJSON returns the unmodified JSON received from the API.
func (r CursorPage[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorPage[T]) GetNextPage() (res *CursorPage[T], err error) {
	if !r.config.ContinueOnEmptyItems && len(r.Data) == 0 {
		return nil, nil
	}
	if !r.hasMore() {
		return nil, nil
	}
	next, ok := r.nextCursor()
	if !ok {
		return nil, nil
	}
	cfg, err := r.cfg.Clone(r.cfg.Context)
	if err != nil {
		return nil, err
	}
	// Place the cursor on the carrier the scheme asked for. A body-located cursor
	// (e.g. `cursor_id` over POST) must ride the request body so it reaches the wire
	// alongside any caller payload; otherwise it goes on the query string. The first
	// request, manual resume, and these auto-advanced pages therefore all agree on the
	// carrier.
	if r.config.CursorLocation == "body" {
		// Merge the cursor into the request body, preserving any caller payload. The raw JSON token
		// (next.Raw, not next.Value()) keeps the value's exact form so a large numeric `cursor_id`
		// is not rounded through gjson's float64. RequestConfig owns the body bookkeeping that a
		// clone implies (re-seeding the body Clone drops, labelling it JSON, rebuilding the request).
		if err = cfg.WriteJSONBodyField(r.config.CursorParam, next.Raw); err != nil {
			return nil, err
		}
	} else {
		// WithQuery needs the string form of the cursor.
		if err = cfg.Apply(option.WithQuery(r.config.CursorParam, next.String())); err != nil {
			return nil, err
		}
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	// Carry the same config forward so every subsequent page extracts identically.
	if err = res.SetPageConfig(cfg, raw, r.config); err != nil {
		return nil, err
	}
	return res, nil
}

// nextCursor reads the next-page cursor per the configured source and reports
// whether a usable cursor is present. A present cursor is a non-empty string or a
// finite number (a numeric `0` id is still a valid `cursor_id` cursor); anything
// else — absent, empty string, null, object — ends pagination.
func (r *CursorPage[T]) nextCursor() (result gjson.Result, ok bool) {
	// An empty cursor path means there is no body-extractable cursor (e.g. a
	// header/link-header cursor, or a non-cursor scheme); stop after this page.
	if len(r.config.CursorPath) == 0 {
		return result, false
	}
	if r.config.CursorKind == "item" {
		// `cursor_id` pagination: the next cursor is the last item's value at CursorPath.
		items := gjson.Get(r.JSON.raw, gjsonPath(r.config.ItemsPath)).Array()
		if len(items) == 0 {
			return result, false
		}
		result = items[len(items)-1].Get(gjsonPath(r.config.CursorPath))
	} else {
		result = gjson.Get(r.JSON.raw, gjsonPath(r.config.CursorPath))
	}
	if !result.Exists() {
		return result, false
	}
	if result.Type == gjson.String {
		return result, result.Str != ""
	}
	if result.Type == gjson.Number {
		return result, true
	}
	return result, false
}

// SetPageConfig stashes the request config, raw response, and pagination descriptor
// so GetNextPage can clone the request and re-extract items/cursor, then populates
// Data from the configured items path. It returns the items-decode error (if any) so
// the caller surfaces a malformed page instead of treating it as empty.
func (r *CursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response, config CursorPageConfig) error {
	if r == nil {
		return nil
	}
	r.cfg = cfg
	r.res = res
	r.config = config
	return r.populateData()
}

// populateData extracts the items array at the configured path from the raw
// response and unmarshals it into Data. The page's struct tags do not bind the
// items field name (it varies per scheme), so the array is resolved here instead.
// A decode failure is returned rather than swallowed: otherwise a page whose items
// fail to parse would look empty and silently stop pagination with no error.
func (r *CursorPage[T]) populateData() error {
	if len(r.config.ItemsPath) == 0 {
		return nil
	}
	node := gjson.Get(r.JSON.raw, gjsonPath(r.config.ItemsPath))
	if !node.Exists() {
		r.Data = nil
		return nil
	}
	var data []T
	if err := apijson.UnmarshalRoot([]byte(node.Raw), &data); err != nil {
		return err
	}
	r.Data = data
	return nil
}

// gjsonPath joins body-path segments into a gjson path, backslash-escaping any
// gjson path operators in a segment so a literal field name (even one containing a
// dot or wildcard) is matched as a key rather than interpreted as path syntax.
func gjsonPath(segments []string) string {
	parts := make([]string, len(segments))
	for i, segment := range segments {
		parts[i] = gjsonEscape(segment)
	}
	return strings.Join(parts, ".")
}

func gjsonEscape(segment string) string {
	var b strings.Builder
	for _, r := range segment {
		switch r {
		case '.', '*', '?', '\\', '#', '@', '|', '[', ']', '(', ')', '!', '>', '<', '=', '~':
			b.WriteByte('\\')
		}
		b.WriteRune(r)
	}
	return b.String()
}

type CursorPageAutoPager[T any] struct {
	page *CursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorPageAutoPager[T any](page *CursorPage[T], err error) *CursorPageAutoPager[T] {
	return &CursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorPageAutoPager[T]) Next() bool {
	if r.page == nil {
		return false
	}
	// Advance past any exhausted or empty page. GetNextPage owns the stop condition and honors
	// ContinueOnEmptyItems, so a scheme that serves empty intermediate pages keeps advancing until
	// a page yields items or the cursor runs out — instead of this loop short-circuiting the first
	// time it sees an empty page. For the default stop-on-empty scheme GetNextPage returns nil on an
	// empty page (without issuing a request), so this still terminates immediately there.
	for r.idx >= len(r.page.Data) {
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil {
			return false
		}
		r.idx = 0
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorPageAutoPager[T]) Index() int {
	return r.run
}
