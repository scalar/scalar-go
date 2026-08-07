# Scalar Go API

Complete reference of every operation, grouped by resource. See [the README](./README.md) for usage and configuration.

## Contents

- [`Registry`](#registry)
  - [List all API Documents](#list-all-api-documents)
  - [List API Documents in a namespace](#list-api-documents-in-a-namespace)
  - [Create API Document](#create-api-document)
  - [Update API Document metadata](#update-api-document-metadata)
  - [Delete API Document](#delete-api-document)
  - [Get API Document](#get-api-document)
  - [Update API Document version](#update-api-document-version)
  - [Delete API Document version](#delete-api-document-version)
  - [Get API Document version metadata](#get-api-document-version-metadata)
  - [Create API Document version](#create-api-document-version)
  - [Add access group](#add-access-group)
  - [Remove access group](#remove-access-group)
- [`Schemas`](#schemas)
  - [List all shared components](#list-all-shared-components)
  - [Create a shared component](#create-a-shared-component)
  - [Update shared component metadata](#update-shared-component-metadata)
  - [Delete a shared component](#delete-a-shared-component)
  - [`Schemas Version`](#schemas-version)
    - [Get a shared component document](#get-a-shared-component-document)
    - [Delete a shared component version](#delete-a-shared-component-version)
    - [Create a shared component version](#create-a-shared-component-version)
  - [`Schemas AccessGroup`](#schemas-accessgroup)
    - [Add shared component access group](#add-shared-component-access-group)
    - [Remove shared component access group](#remove-shared-component-access-group)
- [`LoginPortals`](#loginportals)
  - [Get a login portal](#get-a-login-portal)
  - [Update portal metadata](#update-portal-metadata)
  - [Delete a login portal](#delete-a-login-portal)
  - [Create a portal](#create-a-portal)
  - [List all portals](#list-all-portals)
- [`Rules`](#rules)
  - [List all rules](#list-all-rules)
  - [Create a rule](#create-a-rule)
  - [Update rule metadata](#update-rule-metadata)
  - [Delete a rule](#delete-a-rule)
  - [Get a rule](#get-a-rule)
  - [Add rule access group](#add-rule-access-group)
  - [Remove rule access group](#remove-rule-access-group)
- [`Themes`](#themes)
  - [List all themes](#list-all-themes)
  - [Create a theme](#create-a-theme)
  - [Update theme metadata](#update-theme-metadata)
  - [Update theme document](#update-theme-document)
  - [Delete a theme](#delete-a-theme)
  - [Get a theme](#get-a-theme)
- [`Teams`](#teams)
  - [List teams](#list-teams)
- [`ScalarDocs`](#scalardocs)
  - [List all projects](#list-all-projects)
  - [Create a project](#create-a-project)
  - [Publish a project](#publish-a-project)
- [`Namespaces`](#namespaces)
  - [List namespaces](#list-namespaces)
- [`Authentication`](#authentication)
  - [Exchange token](#exchange-token)
  - [Get current user](#get-current-user)

## Setup

```go
import (
	"context"
	"fmt"

	sdk "github.com/scalar/scalar-go"
)

client := sdk.NewClient()
```

## `Registry`

### List all API Documents

List all API documents across every namespace the caller can access.

| Direction | Type |
| --- | --- |
| Response | [`[]APIDocument`](./registry.go) |

```go
registry, err := client.Registry.ListAllAPIDocuments(context.Background())
if err != nil {
	panic(err)
}

fmt.Println(registry)
```

### List API Documents in a namespace

List API documents in a namespace.

| Direction | Type |
| --- | --- |
| Response | [`[]APIDocument`](./registry.go) |

```go
registry, err := client.Registry.ListAPIDocuments(context.Background(), "namespace")
if err != nil {
	panic(err)
}

fmt.Println(registry)
```

### Create API Document

Create an API document.

| Direction | Type |
| --- | --- |
| Request | [`RegistryNewAPIDocumentParams`](./registry.go) |
| Response | [`RegistryNewAPIDocumentResponse`](./registry.go) |

```go
registry, err := client.Registry.NewAPIDocument(context.Background(), "namespace", sdk.RegistryNewAPIDocumentParams{
	Document: sdk.F[string](""),
	Slug:     sdk.F[string](""),
	Title:    sdk.F[string](""),
	Version:  sdk.F[string]("x"),
})
if err != nil {
	panic(err)
}

fmt.Println(registry)
```

### Update API Document metadata

Update metadata for an API document.

| Direction | Type |
| --- | --- |
| Request | [`RegistryUpdateAPIDocumentParams`](./registry.go) |

```go
registry, err := client.Registry.UpdateAPIDocument(context.Background(), "namespace", "slug", sdk.RegistryUpdateAPIDocumentParams{})
if err != nil {
	panic(err)
}

fmt.Println(registry)
```

### Delete API Document

Delete an API document and all versions.

```go
registry, err := client.Registry.DeleteAPIDocument(context.Background(), "namespace", "slug")
if err != nil {
	panic(err)
}

fmt.Println(registry)
```

### Get API Document

Get a specific API document version.

| Direction | Type |
| --- | --- |
| Response | `string` |

```go
registry, err := client.Registry.GetAPIDocumentVersion(context.Background(), "namespace", "slug", "semver")
if err != nil {
	panic(err)
}

fmt.Println(registry)
```

### Update API Document version

Update the registry file content for an API document version.

| Direction | Type |
| --- | --- |
| Request | [`RegistryUpdateAPIDocumentVersionParams`](./registry.go) |
| Response | [`RegistryUpdateAPIDocumentVersionResponse`](./registry.go) |

```go
registry, err := client.Registry.UpdateAPIDocumentVersion(context.Background(), "namespace", "slug", "semver", sdk.RegistryUpdateAPIDocumentVersionParams{
	Document: sdk.F[string](""),
})
if err != nil {
	panic(err)
}

fmt.Println(registry)
```

### Delete API Document version

Delete a specific API document version.

```go
registry, err := client.Registry.DeleteAPIDocumentVersion(context.Background(), "namespace", "slug", "semver")
if err != nil {
	panic(err)
}

fmt.Println(registry)
```

### Get API Document version metadata

Get metadata (uid, content shas, version sha, tags) for a specific API document version.

| Direction | Type |
| --- | --- |
| Response | [`RegistryListAPIDocumentVersionMetadataResponse`](./registry.go) |

```go
registry, err := client.Registry.ListAPIDocumentVersionMetadata(context.Background(), "namespace", "slug", "semver")
if err != nil {
	panic(err)
}

fmt.Println(registry)
```

### Create API Document version

Create a new API document version.

| Direction | Type |
| --- | --- |
| Request | [`RegistryNewAPIDocumentVersionParams`](./registry.go) |
| Response | [`RegistryNewAPIDocumentVersionResponse`](./registry.go) |

```go
registry, err := client.Registry.NewAPIDocumentVersion(context.Background(), "namespace", "slug", sdk.RegistryNewAPIDocumentVersionParams{
	Document: sdk.F[string](""),
	Version:  sdk.F[string]("x"),
})
if err != nil {
	panic(err)
}

fmt.Println(registry)
```

### Add access group

Add an access group to an API document.

| Direction | Type |
| --- | --- |
| Request | [`RegistryNewAPIDocumentAccessGroupParams`](./registry.go) |

```go
registry, err := client.Registry.NewAPIDocumentAccessGroup(context.Background(), "namespace", "slug", sdk.RegistryNewAPIDocumentAccessGroupParams{
	AccessGroup: sdk.AccessGroupParam{
		AccessGroupSlug: sdk.F[string]("xxx"),
	},
})
if err != nil {
	panic(err)
}

fmt.Println(registry)
```

### Remove access group

Remove an access group from an API document.

| Direction | Type |
| --- | --- |
| Request | [`RegistryDeleteAPIDocumentAccessGroupParams`](./registry.go) |

```go
registry, err := client.Registry.DeleteAPIDocumentAccessGroup(context.Background(), "namespace", "slug", sdk.RegistryDeleteAPIDocumentAccessGroupParams{
	AccessGroup: sdk.AccessGroupParam{
		AccessGroupSlug: sdk.F[string]("xxx"),
	},
})
if err != nil {
	panic(err)
}

fmt.Println(registry)
```

## `Schemas`

### List all shared components

List schemas in a namespace.

| Direction | Type |
| --- | --- |
| Response | [`[]Schema`](./schema.go) |

```go
schema, err := client.Schemas.List(context.Background(), "namespace")
if err != nil {
	panic(err)
}

fmt.Println(schema)
```

### Create a shared component

Create a schema in a namespace.

| Direction | Type |
| --- | --- |
| Request | [`SchemaNewParams`](./schema.go) |
| Response | [`SchemaNewResponse`](./schema.go) |

```go
schema, err := client.Schemas.New(context.Background(), "namespace", sdk.SchemaNewParams{
	Document: sdk.F[string](""),
	Slug:     sdk.F[string](""),
	Title:    sdk.F[string](""),
	Version:  sdk.F[string]("x"),
})
if err != nil {
	panic(err)
}

fmt.Println(schema)
```

### Update shared component metadata

Update schema metadata.

| Direction | Type |
| --- | --- |
| Request | [`SchemaUpdateParams`](./schema.go) |

```go
schema, err := client.Schemas.Update(context.Background(), "namespace", "slug", sdk.SchemaUpdateParams{})
if err != nil {
	panic(err)
}

fmt.Println(schema)
```

### Delete a shared component

Delete a schema and all related versions.

```go
schema, err := client.Schemas.Delete(context.Background(), "namespace", "slug")
if err != nil {
	panic(err)
}

fmt.Println(schema)
```

### `Schemas Version`

#### Get a shared component document

Get a specific schema version document.

| Direction | Type |
| --- | --- |
| Response | `string` |

```go
version, err := client.Schemas.Version.GetSchema(context.Background(), "namespace", "slug", "semver")
if err != nil {
	panic(err)
}

fmt.Println(version)
```

#### Delete a shared component version

Delete a schema version.

```go
version, err := client.Schemas.Version.DeleteSchema(context.Background(), "namespace", "slug", "semver")
if err != nil {
	panic(err)
}

fmt.Println(version)
```

#### Create a shared component version

Create a schema version.

| Direction | Type |
| --- | --- |
| Request | [`SchemaVersionNewSchemaParams`](./schemaversion.go) |
| Response | [`SchemaVersionNewSchemaResponse`](./schemaversion.go) |

```go
version, err := client.Schemas.Version.NewSchema(context.Background(), "namespace", "slug", sdk.SchemaVersionNewSchemaParams{
	Document: sdk.F[string](""),
	Version:  sdk.F[string]("x"),
})
if err != nil {
	panic(err)
}

fmt.Println(version)
```

### `Schemas AccessGroup`

#### Add shared component access group

Add an access group to a schema.

| Direction | Type |
| --- | --- |
| Request | [`SchemaAccessGroupNewSchemaParams`](./schemaaccessgroup.go) |

```go
accessGroup, err := client.Schemas.AccessGroup.NewSchema(context.Background(), "namespace", "slug", sdk.SchemaAccessGroupNewSchemaParams{
	AccessGroup: sdk.AccessGroupParam{
		AccessGroupSlug: sdk.F[string]("xxx"),
	},
})
if err != nil {
	panic(err)
}

fmt.Println(accessGroup)
```

#### Remove shared component access group

Remove an access group from a schema.

| Direction | Type |
| --- | --- |
| Request | [`SchemaAccessGroupDeleteSchemaParams`](./schemaaccessgroup.go) |

```go
accessGroup, err := client.Schemas.AccessGroup.DeleteSchema(context.Background(), "namespace", "slug", sdk.SchemaAccessGroupDeleteSchemaParams{
	AccessGroup: sdk.AccessGroupParam{
		AccessGroupSlug: sdk.F[string]("xxx"),
	},
})
if err != nil {
	panic(err)
}

fmt.Println(accessGroup)
```

## `LoginPortals`

### Get a login portal

Get a login portal by slug.

| Direction | Type |
| --- | --- |
| Response | [`LoginPortalGetResponse`](./loginportal.go) |

```go
loginPortal, err := client.LoginPortals.Get(context.Background(), "slug")
if err != nil {
	panic(err)
}

fmt.Println(loginPortal)
```

### Update portal metadata

Update metadata for a login portal.

| Direction | Type |
| --- | --- |
| Request | [`LoginPortalUpdateParams`](./loginportal.go) |

```go
loginPortal, err := client.LoginPortals.Update(context.Background(), "slug", sdk.LoginPortalUpdateParams{})
if err != nil {
	panic(err)
}

fmt.Println(loginPortal)
```

### Delete a login portal

Delete a login portal.

```go
loginPortal, err := client.LoginPortals.Delete(context.Background(), "slug")
if err != nil {
	panic(err)
}

fmt.Println(loginPortal)
```

### Create a portal

Create a login portal for the current team.

| Direction | Type |
| --- | --- |
| Request | [`LoginPortalNewParams`](./loginportal.go) |
| Response | [`LoginPortalNewResponse`](./loginportal.go) |

```go
loginPortal, err := client.LoginPortals.New(context.Background(), sdk.LoginPortalNewParams{
	Email: sdk.F[sdk.LoginPortalEmailParam](sdk.LoginPortalEmailParam{
		Logo:             sdk.F[string](""),
		LogoSize:         sdk.F[string]("100"),
		ButtonText:       sdk.F[string]("Login"),
		Message:          sdk.F[string]("Click to access private documentation hosted by scalar.com"),
		Title:            sdk.F[string]("Private Docs"),
		MainColor:        sdk.F[string]("#2a2f45"),
		MainBackground:   sdk.F[string]("#f6f6f6"),
		CardColor:        sdk.F[string]("2a2f45"),
		CardBackground:   sdk.F[string]("#fff"),
		ButtonColor:      sdk.F[string]("#fff"),
		ButtonBackground: sdk.F[string]("#0f0f0f"),
	}),
	Page: sdk.F[sdk.LoginPortalPageParam](sdk.LoginPortalPageParam{
		Title:           sdk.F[string]("Scalar Private Docs"),
		Description:     sdk.F[string]("Login to access your documentation"),
		Head:            sdk.F[string](""),
		Script:          sdk.F[string](""),
		Theme:           sdk.F[string](""),
		CompanyName:     sdk.F[string](""),
		Logo:            sdk.F[string](""),
		LogoURL:         sdk.F[string](""),
		Favicon:         sdk.F[string](""),
		TermsLink:       sdk.F[string](""),
		PrivacyLink:     sdk.F[string](""),
		FormTitle:       sdk.F[string]("Scalar Private Docs"),
		FormDescription: sdk.F[string]("Login to access your documentation"),
		FormImage:       sdk.F[string](""),
	}),
	Slug:  sdk.F[string](""),
	Title: sdk.F[string](""),
})
if err != nil {
	panic(err)
}

fmt.Println(loginPortal)
```

### List all portals

List all login portals for the current team.

| Direction | Type |
| --- | --- |
| Response | [`[]LoginPortal`](./loginportal.go) |

```go
loginPortal, err := client.LoginPortals.List(context.Background())
if err != nil {
	panic(err)
}

fmt.Println(loginPortal)
```

## `Rules`

### List all rules

List all rulesets in a namespace.

| Direction | Type |
| --- | --- |
| Response | [`[]Rule`](./rule.go) |

```go
rule, err := client.Rules.ListRulesets(context.Background(), "namespace")
if err != nil {
	panic(err)
}

fmt.Println(rule)
```

### Create a rule

Create a rule in a namespace.

| Direction | Type |
| --- | --- |
| Request | [`RuleNewRulesetParams`](./rule.go) |
| Response | [`RuleNewRulesetResponse`](./rule.go) |

```go
rule, err := client.Rules.NewRuleset(context.Background(), "namespace", sdk.RuleNewRulesetParams{
	Document: sdk.F[string](""),
	Slug:     sdk.F[string](""),
	Title:    sdk.F[string](""),
})
if err != nil {
	panic(err)
}

fmt.Println(rule)
```

### Update rule metadata

Update rule metadata by slug.

| Direction | Type |
| --- | --- |
| Request | [`RuleUpdateRulesetParams`](./rule.go) |

```go
rule, err := client.Rules.UpdateRuleset(context.Background(), "namespace", "slug", sdk.RuleUpdateRulesetParams{})
if err != nil {
	panic(err)
}

fmt.Println(rule)
```

### Delete a rule

Delete a rule by slug.

```go
rule, err := client.Rules.DeleteRuleset(context.Background(), "namespace", "slug")
if err != nil {
	panic(err)
}

fmt.Println(rule)
```

### Get a rule

Get a rule document by slug.

| Direction | Type |
| --- | --- |
| Response | `string` |

```go
rule, err := client.Rules.GetRulesetDocument(context.Background(), "namespace", "slug")
if err != nil {
	panic(err)
}

fmt.Println(rule)
```

### Add rule access group

Grant an access group to a rule.

| Direction | Type |
| --- | --- |
| Request | [`RuleNewRulesetAccessGroupParams`](./rule.go) |

```go
rule, err := client.Rules.NewRulesetAccessGroup(context.Background(), "namespace", "slug", sdk.RuleNewRulesetAccessGroupParams{
	AccessGroup: sdk.AccessGroupParam{
		AccessGroupSlug: sdk.F[string]("xxx"),
	},
})
if err != nil {
	panic(err)
}

fmt.Println(rule)
```

### Remove rule access group

Remove an access group from a rule.

| Direction | Type |
| --- | --- |
| Request | [`RuleDeleteRulesetAccessGroupParams`](./rule.go) |

```go
rule, err := client.Rules.DeleteRulesetAccessGroup(context.Background(), "namespace", "slug", sdk.RuleDeleteRulesetAccessGroupParams{
	AccessGroup: sdk.AccessGroupParam{
		AccessGroupSlug: sdk.F[string]("xxx"),
	},
})
if err != nil {
	panic(err)
}

fmt.Println(rule)
```

## `Themes`

### List all themes

List all team themes.

| Direction | Type |
| --- | --- |
| Response | [`[]Theme`](./theme.go) |

```go
theme, err := client.Themes.List(context.Background())
if err != nil {
	panic(err)
}

fmt.Println(theme)
```

### Create a theme

Create a team theme.

| Direction | Type |
| --- | --- |
| Request | [`ThemeNewParams`](./theme.go) |
| Response | [`ThemeNewResponse`](./theme.go) |

```go
theme, err := client.Themes.New(context.Background(), sdk.ThemeNewParams{
	Document: sdk.F[string](""),
	Name:     sdk.F[string](""),
	Slug:     sdk.F[string](""),
})
if err != nil {
	panic(err)
}

fmt.Println(theme)
```

### Update theme metadata

Update theme metadata.

| Direction | Type |
| --- | --- |
| Request | [`ThemeUpdateParams`](./theme.go) |

```go
theme, err := client.Themes.Update(context.Background(), "slug", sdk.ThemeUpdateParams{})
if err != nil {
	panic(err)
}

fmt.Println(theme)
```

### Update theme document

Replace the theme document.

| Direction | Type |
| --- | --- |
| Request | [`ThemeReplaceDocumentParams`](./theme.go) |

```go
theme, err := client.Themes.ReplaceDocument(context.Background(), "slug", sdk.ThemeReplaceDocumentParams{
	Document: sdk.F[string](""),
})
if err != nil {
	panic(err)
}

fmt.Println(theme)
```

### Delete a theme

Delete a theme by slug.

```go
theme, err := client.Themes.Delete(context.Background(), "slug")
if err != nil {
	panic(err)
}

fmt.Println(theme)
```

### Get a theme

Get the theme document by slug.

| Direction | Type |
| --- | --- |
| Response | `string` |

```go
theme, err := client.Themes.Get(context.Background(), "slug")
if err != nil {
	panic(err)
}

fmt.Println(theme)
```

## `Teams`

### List teams

List all available teams

| Direction | Type |
| --- | --- |
| Response | [`[]Team`](./team.go) |

```go
team, err := client.Teams.List(context.Background())
if err != nil {
	panic(err)
}

fmt.Println(team)
```

## `ScalarDocs`

### List all projects

List all guide projects.

| Direction | Type |
| --- | --- |
| Response | [`[]GithubProject`](./scalardoc.go) |

```go
scalarDoc, err := client.ScalarDocs.ListGuides(context.Background())
if err != nil {
	panic(err)
}

fmt.Println(scalarDoc)
```

### Create a project

Create a guide project.

| Direction | Type |
| --- | --- |
| Request | [`ScalarDocNewGuideParams`](./scalardoc.go) |
| Response | [`ScalarDocNewGuideResponse`](./scalardoc.go) |

```go
scalarDoc, err := client.ScalarDocs.NewGuide(context.Background(), sdk.ScalarDocNewGuideParams{
	AllowedDomains: sdk.F[[]string]([]string{""}),
	AllowedUsers:   sdk.F[[]string]([]string{""}),
	IsPrivate:      sdk.F[bool](false),
	Name:           sdk.F[string](""),
})
if err != nil {
	panic(err)
}

fmt.Println(scalarDoc)
```

### Publish a project

Start a new publish process.

| Direction | Type |
| --- | --- |
| Response | [`ScalarDocPublishGuideResponse`](./scalardoc.go) |

```go
scalarDoc, err := client.ScalarDocs.PublishGuide(context.Background(), "slug")
if err != nil {
	panic(err)
}

fmt.Println(scalarDoc)
```

## `Namespaces`

### List namespaces

Get all namespaces for the current team

| Direction | Type |
| --- | --- |
| Response | `[]string` |

```go
namespace, err := client.Namespaces.List(context.Background())
if err != nil {
	panic(err)
}

fmt.Println(namespace)
```

## `Authentication`

### Exchange token

Exchange an API key for an access token.

| Direction | Type |
| --- | --- |
| Request | [`AuthenticationExchangePersonalTokenParams`](./authentication.go) |
| Response | [`AuthenticationExchangePersonalTokenResponse`](./authentication.go) |

```go
authentication, err := client.Authentication.ExchangePersonalToken(context.Background(), sdk.AuthenticationExchangePersonalTokenParams{
	PersonalToken: sdk.F[string](""),
})
if err != nil {
	panic(err)
}

fmt.Println(authentication)
```

### Get current user

Get the authenticated user, including their available teams and theme.

| Direction | Type |
| --- | --- |
| Response | [`AuthenticationListCurrentUserResponse`](./authentication.go) |

```go
authentication, err := client.Authentication.ListCurrentUser(context.Background())
if err != nil {
	panic(err)
}

fmt.Println(authentication)
```
