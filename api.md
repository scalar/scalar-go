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
- [`AccessGroups`](#accessgroups)
  - [Create an access group](#create-an-access-group)
  - [Get an access group](#get-an-access-group)
  - [Update an access group](#update-an-access-group)
  - [Delete an access group](#delete-an-access-group)
  - [`AccessGroups Domains`](#accessgroups-domains)
    - [Add an allowed email domain](#add-an-allowed-email-domain)
    - [Remove an allowed email domain](#remove-an-allowed-email-domain)
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
  - [`Teams Members`](#teams-members)
    - [List team members](#list-team-members)
    - [Change a member role](#change-a-member-role)
    - [Remove a member](#remove-a-member)
  - [`Teams Invites`](#teams-invites)
    - [Invite a member](#invite-a-member)
    - [Resend an invite](#resend-an-invite)
    - [Cancel an invite](#cancel-an-invite)
- [`ScalarDocs`](#scalardocs)
  - [List all projects](#list-all-projects)
  - [Create a project](#create-a-project)
  - [Publish a project](#publish-a-project)
  - [List all docs projects](#list-all-docs-projects)
  - [Create a docs project](#create-a-docs-project)
  - [Get a docs project](#get-a-docs-project)
  - [Update a docs project](#update-a-docs-project)
  - [Delete a docs project](#delete-a-docs-project)
  - [Publish a docs project](#publish-a-docs-project)
  - [Read the site config](#read-the-site-config)
  - [Write the site config](#write-the-site-config)
  - [Get the site domains](#get-the-site-domains)
  - [Check domain DNS](#check-domain-dns)
- [`Namespaces`](#namespaces)
  - [List namespaces](#list-namespaces)
- [`Authentication`](#authentication)
  - [Exchange token](#exchange-token)
  - [Get current user](#get-current-user)
- [`Sdks`](#sdks)
  - [List all SDKs](#list-all-sdks)
  - [Create an SDK](#create-an-sdk)
  - [Get an SDK](#get-an-sdk)
  - [Update an SDK](#update-an-sdk)
  - [Delete an SDK](#delete-an-sdk)
  - [Build an SDK](#build-an-sdk)
  - [`Sdks Versions`](#sdks-versions)
    - [Create an SDK version](#create-an-sdk-version)
    - [Delete an SDK version](#delete-an-sdk-version)
  - [`Sdks Repositories`](#sdks-repositories)
    - [Link a repository](#link-a-repository)
    - [Unlink a repository](#unlink-a-repository)
    - [Update publishing settings](#update-publishing-settings)
- [`Mcp`](#mcp)
  - [`Mcp Servers`](#mcp-servers)
    - [List all MCP servers](#list-all-mcp-servers)
    - [Create an MCP server](#create-an-mcp-server)
    - [Get an MCP server](#get-an-mcp-server)
    - [Update an MCP server](#update-an-mcp-server)
    - [Delete an MCP server](#delete-an-mcp-server)
    - [`Mcp Servers Installations`](#mcp-servers-installations)
      - [List installations](#list-installations)
      - [Create an installation](#create-an-installation)
      - [Get an installation](#get-an-installation)
      - [Update an installation](#update-an-installation)
      - [Delete an installation](#delete-an-installation)
      - [Add an access group](#add-an-access-group)
      - [Remove an access group](#remove-an-access-group)

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

Registry

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
registry, err := client.Registry.ListAPIDocuments(context.Background(), "acme")
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
registry, err := client.Registry.NewAPIDocument(context.Background(), "acme", sdk.RegistryNewAPIDocumentParams{
	Document: sdk.F[string]("{\"openapi\":\"3.1.0\",\"info\":{\"title\":\"Acme API\",\"version\":\"1.2.0\"},\"paths\":{}}"),
	Slug:     sdk.F[string]("acme-api"),
	Title:    sdk.F[string]("Acme API"),
	Version:  sdk.F[string]("1.2.0"),
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
registry, err := client.Registry.UpdateAPIDocument(context.Background(), "acme", "acme-api", sdk.RegistryUpdateAPIDocumentParams{})
if err != nil {
	panic(err)
}

fmt.Println(registry)
```

### Delete API Document

Delete an API document and all versions.

```go
registry, err := client.Registry.DeleteAPIDocument(context.Background(), "acme", "acme-api")
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
registry, err := client.Registry.GetAPIDocumentVersion(context.Background(), "acme", "acme-api", "1.2.0")
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
registry, err := client.Registry.UpdateAPIDocumentVersion(context.Background(), "acme", "acme-api", "1.2.0", sdk.RegistryUpdateAPIDocumentVersionParams{
	Document: sdk.F[string]("{\"openapi\":\"3.1.0\",\"info\":{\"title\":\"Acme API\",\"version\":\"1.2.0\"},\"paths\":{}}"),
})
if err != nil {
	panic(err)
}

fmt.Println(registry)
```

### Delete API Document version

Delete a specific API document version.

```go
registry, err := client.Registry.DeleteAPIDocumentVersion(context.Background(), "acme", "acme-api", "1.2.0")
if err != nil {
	panic(err)
}

fmt.Println(registry)
```

### Get API Document version metadata

Get metadata (uid, content shas, version sha, tags) for a specific API document version.

| Direction | Type |
| --- | --- |
| Response | [`ManagedDocVersion`](./shared/shared.go) |

```go
registry, err := client.Registry.ListAPIDocumentVersionMetadata(context.Background(), "acme", "acme-api", "1.2.0")
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
| Response | [`ManagedDocVersion`](./shared/shared.go) |

```go
registry, err := client.Registry.NewAPIDocumentVersion(context.Background(), "acme", "acme-api", sdk.RegistryNewAPIDocumentVersionParams{
	Document: sdk.F[string]("{\"openapi\":\"3.1.0\",\"info\":{\"title\":\"Acme API\",\"version\":\"1.2.0\"},\"paths\":{}}"),
	Version:  sdk.F[string]("1.2.0"),
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
registry, err := client.Registry.NewAPIDocumentAccessGroup(context.Background(), "acme", "acme-api", sdk.RegistryNewAPIDocumentAccessGroupParams{
	AccessGroup: sdk.AccessGroupParam{
		AccessGroupSlug: sdk.F[string]("acme-api"),
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
registry, err := client.Registry.DeleteAPIDocumentAccessGroup(context.Background(), "acme", "acme-api", sdk.RegistryDeleteAPIDocumentAccessGroupParams{
	AccessGroup: sdk.AccessGroupParam{
		AccessGroupSlug: sdk.F[string]("acme-api"),
	},
})
if err != nil {
	panic(err)
}

fmt.Println(registry)
```

## `Schemas`

Schemas

### List all shared components

List schemas in a namespace.

| Direction | Type |
| --- | --- |
| Response | [`[]Schema`](./schema.go) |

```go
schema, err := client.Schemas.List(context.Background(), "acme")
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
| Response | [`UID`](./shared/shared.go) |

```go
schema, err := client.Schemas.New(context.Background(), "acme", sdk.SchemaNewParams{
	Document: sdk.F[string]("{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\",\"examples\":[\"Acme\"]}}}"),
	Slug:     sdk.F[string]("customer"),
	Title:    sdk.F[string]("Customer"),
	Version:  sdk.F[string]("1.2.0"),
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
schema, err := client.Schemas.Update(context.Background(), "acme", "customer", sdk.SchemaUpdateParams{})
if err != nil {
	panic(err)
}

fmt.Println(schema)
```

### Delete a shared component

Delete a schema and all related versions.

```go
schema, err := client.Schemas.Delete(context.Background(), "acme", "customer")
if err != nil {
	panic(err)
}

fmt.Println(schema)
```

### `Schemas Version`

Schemas

#### Get a shared component document

Get a specific schema version document.

| Direction | Type |
| --- | --- |
| Response | `string` |

```go
version, err := client.Schemas.Version.Get(context.Background(), "acme", "customer", "1.2.0")
if err != nil {
	panic(err)
}

fmt.Println(version)
```

#### Delete a shared component version

Delete a schema version.

```go
version, err := client.Schemas.Version.Delete(context.Background(), "acme", "customer", "1.2.0")
if err != nil {
	panic(err)
}

fmt.Println(version)
```

#### Create a shared component version

Create a schema version.

| Direction | Type |
| --- | --- |
| Request | [`SchemaVersionNewParams`](./schemaversion.go) |
| Response | [`SchemaVersionNewResponse`](./schemaversion.go) |

```go
version, err := client.Schemas.Version.New(context.Background(), "acme", "customer", sdk.SchemaVersionNewParams{
	Document: sdk.F[string]("{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\",\"examples\":[\"Acme\"]}}}"),
	Version:  sdk.F[string]("1.2.0"),
})
if err != nil {
	panic(err)
}

fmt.Println(version)
```

### `Schemas AccessGroup`

Schemas

#### Add shared component access group

Add an access group to a schema.

| Direction | Type |
| --- | --- |
| Request | [`SchemaAccessGroupNewParams`](./schemaaccessgroup.go) |

```go
accessGroup, err := client.Schemas.AccessGroup.New(context.Background(), "acme", "customer", sdk.SchemaAccessGroupNewParams{
	AccessGroup: sdk.AccessGroupParam{
		AccessGroupSlug: sdk.F[string]("acme-api"),
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
| Request | [`SchemaAccessGroupDeleteParams`](./schemaaccessgroup.go) |

```go
accessGroup, err := client.Schemas.AccessGroup.Delete(context.Background(), "acme", "customer", sdk.SchemaAccessGroupDeleteParams{
	AccessGroup: sdk.AccessGroupParam{
		AccessGroupSlug: sdk.F[string]("acme-api"),
	},
})
if err != nil {
	panic(err)
}

fmt.Println(accessGroup)
```

## `LoginPortals`

Login Portals

### Get a login portal

Get a login portal by slug.

| Direction | Type |
| --- | --- |
| Response | [`LoginPortalGetResponse`](./loginportal.go) |

```go
loginPortal, err := client.LoginPortals.Get(context.Background(), "acme-login")
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
loginPortal, err := client.LoginPortals.Update(context.Background(), "acme-login", sdk.LoginPortalUpdateParams{})
if err != nil {
	panic(err)
}

fmt.Println(loginPortal)
```

### Delete a login portal

Delete a login portal.

```go
loginPortal, err := client.LoginPortals.Delete(context.Background(), "acme-login")
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
| Response | [`UID`](./shared/shared.go) |

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
		CardColor:        sdk.F[string]("#2a2f45"),
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
	Slug:  sdk.F[string]("acme-login"),
	Title: sdk.F[string]("Acme Private Documentation"),
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

## `AccessGroups`

Access Groups

### Create an access group

Create a group for the current team. Requires docs edit permission and the access groups billing feature. Domains are exact email domains, without wildcards or implicit subdomain matching.

| Direction | Type |
| --- | --- |
| Request | [`AccessGroupNewParams`](./accessgroup.go) |
| Response | [`AccessGroupNewResponse`](./accessgroup.go) |

```go
accessGroup, err := client.AccessGroups.New(context.Background(), sdk.AccessGroupNewParams{})
if err != nil {
	panic(err)
}

fmt.Println(accessGroup)
```

### Get an access group

Get a group and its email and domain allowlists by slug.

| Direction | Type |
| --- | --- |
| Response | [`AccessGroupGetResponse`](./accessgroup.go) |

```go
accessGroup, err := client.AccessGroups.Get(context.Background(), "acme-api")
if err != nil {
	panic(err)
}

fmt.Println(accessGroup)
```

### Update an access group

Update group metadata. Requires docs edit permission. After changing the slug, use the new slug in subsequent requests.

| Direction | Type |
| --- | --- |
| Request | [`AccessGroupUpdateParams`](./accessgroup.go) |

```go
accessGroup, err := client.AccessGroups.Update(context.Background(), "acme-api", sdk.AccessGroupUpdateParams{})
if err != nil {
	panic(err)
}

fmt.Println(accessGroup)
```

### Delete an access group

Delete a group and remove its project assignments. Requires docs edit permission.

```go
accessGroup, err := client.AccessGroups.Delete(context.Background(), "acme-api")
if err != nil {
	panic(err)
}

fmt.Println(accessGroup)
```

### `AccessGroups Domains`

Access Groups

#### Add an allowed email domain

Allow an exact email domain in a group. Requires docs edit permission. A group supports up to 1000 domains.

| Direction | Type |
| --- | --- |
| Request | [`AccessGroupDomainNewParams`](./accessgroupdomain.go) |

```go
domain, err := client.AccessGroups.Domains.New(context.Background(), "acme-api", sdk.AccessGroupDomainNewParams{
	Domain: sdk.F[string]("example.com"),
})
if err != nil {
	panic(err)
}

fmt.Println(domain)
```

#### Remove an allowed email domain

Remove an exact email domain from a group. Requires docs edit permission. Other allowed domains and emails are preserved.

| Direction | Type |
| --- | --- |
| Request | [`AccessGroupDomainDeleteParams`](./accessgroupdomain.go) |

```go
domain, err := client.AccessGroups.Domains.Delete(context.Background(), "acme-api", sdk.AccessGroupDomainDeleteParams{
	Domain: sdk.F[string]("example.com"),
})
if err != nil {
	panic(err)
}

fmt.Println(domain)
```

## `Rules`

Rules

### List all rules

List all rulesets in a namespace.

| Direction | Type |
| --- | --- |
| Response | [`[]Rule`](./rule.go) |

```go
rule, err := client.Rules.ListRulesets(context.Background(), "acme")
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
| Response | [`UID`](./shared/shared.go) |

```go
rule, err := client.Rules.NewRuleset(context.Background(), "acme", sdk.RuleNewRulesetParams{
	Document: sdk.F[string]("extends: [\"spectral:oas\"]\nrules:\n  info-contact: warn\n"),
	Slug:     sdk.F[string]("acme-rules"),
	Title:    sdk.F[string]("Acme API Rules"),
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
rule, err := client.Rules.UpdateRuleset(context.Background(), "acme", "acme-rules", sdk.RuleUpdateRulesetParams{})
if err != nil {
	panic(err)
}

fmt.Println(rule)
```

### Delete a rule

Delete a rule by slug.

```go
rule, err := client.Rules.DeleteRuleset(context.Background(), "acme", "acme-rules")
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
rule, err := client.Rules.GetRulesetDocument(context.Background(), "acme", "acme-rules")
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
rule, err := client.Rules.NewRulesetAccessGroup(context.Background(), "acme", "acme-rules", sdk.RuleNewRulesetAccessGroupParams{
	AccessGroup: sdk.AccessGroupParam{
		AccessGroupSlug: sdk.F[string]("acme-api"),
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
rule, err := client.Rules.DeleteRulesetAccessGroup(context.Background(), "acme", "acme-rules", sdk.RuleDeleteRulesetAccessGroupParams{
	AccessGroup: sdk.AccessGroupParam{
		AccessGroupSlug: sdk.F[string]("acme-api"),
	},
})
if err != nil {
	panic(err)
}

fmt.Println(rule)
```

## `Themes`

Themes

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
| Response | [`UID`](./shared/shared.go) |

```go
theme, err := client.Themes.New(context.Background(), sdk.ThemeNewParams{
	Document: sdk.F[string](":root { --scalar-color-1: #1f2937; }"),
	Name:     sdk.F[string]("Acme Theme"),
	Slug:     sdk.F[string]("acme-theme"),
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
theme, err := client.Themes.Update(context.Background(), "acme-theme", sdk.ThemeUpdateParams{})
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
theme, err := client.Themes.ReplaceDocument(context.Background(), "acme-theme", sdk.ThemeReplaceDocumentParams{
	Document: sdk.F[string](":root { --scalar-color-1: #1f2937; }"),
})
if err != nil {
	panic(err)
}

fmt.Println(theme)
```

### Delete a theme

Delete a theme by slug.

```go
theme, err := client.Themes.Delete(context.Background(), "acme-theme")
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
theme, err := client.Themes.Get(context.Background(), "acme-theme")
if err != nil {
	panic(err)
}

fmt.Println(theme)
```

## `Teams`

Teams

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

### `Teams Members`

Teams

#### List team members

List the members of the current team, along with the invites still outstanding.

| Direction | Type |
| --- | --- |
| Response | [`TeamMemberListResponse`](./teammember.go) |

```go
member, err := client.Teams.Members.List(context.Background())
if err != nil {
	panic(err)
}

fmt.Println(member)
```

#### Change a member role

Change what a member of the current team is allowed to do.

| Direction | Type |
| --- | --- |
| Request | [`TeamMemberUpdateParams`](./teammember.go) |

```go
member, err := client.Teams.Members.Update(context.Background(), "UakgbKJ5m9gl0JDMbcJqL", sdk.TeamMemberUpdateParams{
	Role: sdk.F[sdk.Role](sdk.Role("owner")),
})
if err != nil {
	panic(err)
}

fmt.Println(member)
```

#### Remove a member

Remove someone from the current team.

```go
member, err := client.Teams.Members.Delete(context.Background(), "UakgbKJ5m9gl0JDMbcJqL")
if err != nil {
	panic(err)
}

fmt.Println(member)
```

### `Teams Invites`

Teams

#### Invite a member

Invite someone to the current team by email.

| Direction | Type |
| --- | --- |
| Request | [`TeamInviteMemberParams`](./teaminvite.go) |

```go
invite, err := client.Teams.Invites.Member(context.Background(), sdk.TeamInviteMemberParams{
	Email: sdk.F[string]("alex@example.com"),
	Role:  sdk.F[sdk.Role](sdk.Role("owner")),
})
if err != nil {
	panic(err)
}

fmt.Println(invite)
```

#### Resend an invite

Send the invite email again.

```go
invite, err := client.Teams.Invites.Resend(context.Background(), "UakgbKJ5m9gl0JDMbcJqL")
if err != nil {
	panic(err)
}

fmt.Println(invite)
```

#### Cancel an invite

Withdraw an invite that has not been accepted.

```go
invite, err := client.Teams.Invites.Cancel(context.Background(), "UakgbKJ5m9gl0JDMbcJqL")
if err != nil {
	panic(err)
}

fmt.Println(invite)
```

## `ScalarDocs`

Scalar Docs

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
	AllowedDomains: sdk.F[[]string]([]string{}),
	AllowedUsers:   sdk.F[[]string]([]string{}),
	IsPrivate:      sdk.F[bool](false),
	Name:           sdk.F[string]("Acme Documentation"),
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
scalarDoc, err := client.ScalarDocs.PublishGuide(context.Background(), "acme-docs")
if err != nil {
	panic(err)
}

fmt.Println(scalarDoc)
```

### List all docs projects

List every docs project on the team.

| Direction | Type |
| --- | --- |
| Request | [`ScalarDocListProjectsParams`](./scalardoc.go) |
| Response | [`ScalarDocListProjectsResponse`](./scalardoc.go) |

```go
scalarDoc, err := client.ScalarDocs.ListProjects(context.Background(), sdk.ScalarDocListProjectsParams{})
if err != nil {
	panic(err)
}

fmt.Println(scalarDoc)
```

### Create a docs project

Create a docs project. Omit `provider` to have Scalar host the repository.

| Direction | Type |
| --- | --- |
| Request | [`ScalarDocNewProjectParams`](./scalardoc.go) |
| Response | [`DocsProject`](./scalardoc.go) |

```go
scalarDoc, err := client.ScalarDocs.NewProject(context.Background(), sdk.ScalarDocNewProjectParams{
	Name:     sdk.F[string]("Acme Documentation"),
	Provider: sdk.F[sdk.ScalarDocNewProjectParamsProvider](sdk.ScalarDocNewProjectParamsProvider("forgejo")),
})
if err != nil {
	panic(err)
}

fmt.Println(scalarDoc)
```

### Get a docs project

Get a single docs project by its slug.

| Direction | Type |
| --- | --- |
| Response | [`DocsProject`](./scalardoc.go) |

```go
scalarDoc, err := client.ScalarDocs.GetProject(context.Background(), "acme-docs")
if err != nil {
	panic(err)
}

fmt.Println(scalarDoc)
```

### Update a docs project

Update project settings. Set `isPrivate` with `accessGroups` to put the site behind a login.

| Direction | Type |
| --- | --- |
| Request | [`ScalarDocUpdateProjectParams`](./scalardoc.go) |

```go
scalarDoc, err := client.ScalarDocs.UpdateProject(context.Background(), "acme-docs", sdk.ScalarDocUpdateProjectParams{})
if err != nil {
	panic(err)
}

fmt.Println(scalarDoc)
```

### Delete a docs project

Delete a docs project, its deploys, its publish records and its cached builds.

```go
scalarDoc, err := client.ScalarDocs.DeleteProject(context.Background(), "acme-docs")
if err != nil {
	panic(err)
}

fmt.Println(scalarDoc)
```

### Publish a docs project

Start a build and deploy. The returned `publishUid` identifies the publish record.

| Direction | Type |
| --- | --- |
| Request | [`ScalarDocPublishProjectParams`](./scalardoc.go) |
| Response | [`ScalarDocPublishProjectResponse`](./scalardoc.go) |

```go
scalarDoc, err := client.ScalarDocs.PublishProject(context.Background(), "acme-docs", sdk.ScalarDocPublishProjectParams{})
if err != nil {
	panic(err)
}

fmt.Println(scalarDoc)
```

### Read the site config

Read `scalar.config.json` straight from the project repository, without cloning it. `baseToken` is the compare-and-swap handle for a later write.

| Direction | Type |
| --- | --- |
| Request | [`ScalarDocListProjectConfigParams`](./scalardoc.go) |
| Response | [`ScalarDocListProjectConfigResponse`](./scalardoc.go) |

```go
scalarDoc, err := client.ScalarDocs.ListProjectConfig(context.Background(), "acme-docs", sdk.ScalarDocListProjectConfigParams{})
if err != nil {
	panic(err)
}

fmt.Println(scalarDoc)
```

### Write the site config

Commit `scalar.config.json` straight to the project repository. Pass the `baseToken` from the read this edit was based on; a conflict means the file moved underneath it.

| Direction | Type |
| --- | --- |
| Request | [`ScalarDocUpdateProjectConfigParams`](./scalardoc.go) |
| Response | [`ScalarDocUpdateProjectConfigResponse`](./scalardoc.go) |

```go
scalarDoc, err := client.ScalarDocs.UpdateProjectConfig(context.Background(), "acme-docs", sdk.ScalarDocUpdateProjectConfigParams{
	Content: sdk.F[string]("{\"name\":\"Acme Documentation\"}"),
})
if err != nil {
	panic(err)
}

fmt.Println(scalarDoc)
```

### Get the site domains

The domains the project serves on — the Scalar-hosted one and the custom one, when set.

| Direction | Type |
| --- | --- |
| Response | [`ScalarDocListProjectDomainResponse`](./scalardoc.go) |

```go
scalarDoc, err := client.ScalarDocs.ListProjectDomain(context.Background(), "acme-docs")
if err != nil {
	panic(err)
}

fmt.Println(scalarDoc)
```

### Check domain DNS

Whether the project custom domain points at Scalar yet. `expected` is the CNAME record to create; `found` is what resolves today. A project with no custom domain reports `verified` with no expected record, because Scalar serves its own subdomain directly.

| Direction | Type |
| --- | --- |
| Response | [`ScalarDocListProjectDomainStatusResponse`](./scalardoc.go) |

```go
scalarDoc, err := client.ScalarDocs.ListProjectDomainStatus(context.Background(), "acme-docs")
if err != nil {
	panic(err)
}

fmt.Println(scalarDoc)
```

## `Namespaces`

Namespaces

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

Authentication

### Exchange token

Exchange an API key for an access token.

| Direction | Type |
| --- | --- |
| Request | [`AuthenticationExchangePersonalTokenParams`](./authentication.go) |
| Response | [`AuthenticationExchangePersonalTokenResponse`](./authentication.go) |

```go
authentication, err := client.Authentication.ExchangePersonalToken(context.Background(), sdk.AuthenticationExchangePersonalTokenParams{
	PersonalToken: sdk.F[string]("scalar_example_personal_token"),
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
| Response | [`User`](./authentication.go) |

```go
authentication, err := client.Authentication.ListCurrentUser(context.Background())
if err != nil {
	panic(err)
}

fmt.Println(authentication)
```

## `Sdks`

SDKs

### List all SDKs

List every SDK on the team.

| Direction | Type |
| --- | --- |
| Request | [`SdkListParams`](./sdk.go) |
| Response | [`SdkListResponse`](./sdk.go) |

```go
sdk, err := client.Sdks.List(context.Background(), sdk.SdkListParams{})
if err != nil {
	panic(err)
}

fmt.Println(sdk)
```

### Create an SDK

Create an SDK from an API document, targeting one or more languages.

| Direction | Type |
| --- | --- |
| Request | [`SdkNewParams`](./sdk.go) |
| Response | [`UID`](./shared/shared.go) |

```go
sdk, err := client.Sdks.New(context.Background(), sdk.SdkNewParams{
	APIUID:    sdk.F[string]("UakgbKJ5m9gl0JDMbcJqL"),
	Languages: sdk.F[[]sdk.SdkNewParamsLanguage]([]sdk.SdkNewParamsLanguage{"typescript"}),
})
if err != nil {
	panic(err)
}

fmt.Println(sdk)
```

### Get an SDK

Get a single SDK by its uid.

| Direction | Type |
| --- | --- |
| Response | [`Sdk`](./sdk.go) |

```go
sdk, err := client.Sdks.Get(context.Background(), "UakgbKJ5m9gl0JDMbcJqL")
if err != nil {
	panic(err)
}

fmt.Println(sdk)
```

### Update an SDK

Update SDK metadata, its linked API, or its config.

| Direction | Type |
| --- | --- |
| Request | [`SdkUpdateParams`](./sdk.go) |

```go
sdk, err := client.Sdks.Update(context.Background(), "UakgbKJ5m9gl0JDMbcJqL", sdk.SdkUpdateParams{})
if err != nil {
	panic(err)
}

fmt.Println(sdk)
```

### Delete an SDK

Delete an SDK and every version it holds.

```go
sdk, err := client.Sdks.Delete(context.Background(), "UakgbKJ5m9gl0JDMbcJqL")
if err != nil {
	panic(err)
}

fmt.Println(sdk)
```

### Build an SDK

Start a build. Omit `version` to build the current work — the open draft, else the latest version — and the resolved version comes back in the response.

| Direction | Type |
| --- | --- |
| Request | [`SdkBuildParams`](./sdk.go) |
| Response | [`SdkBuildResponse`](./sdk.go) |

```go
sdk, err := client.Sdks.Build(context.Background(), "UakgbKJ5m9gl0JDMbcJqL", sdk.SdkBuildParams{})
if err != nil {
	panic(err)
}

fmt.Println(sdk)
```

### `Sdks Versions`

SDKs

#### Create an SDK version

Create a new SDK version against a specific API version.

| Direction | Type |
| --- | --- |
| Request | [`SdkVersionNewParams`](./sdkversion.go) |

```go
version, err := client.Sdks.Versions.New(context.Background(), "UakgbKJ5m9gl0JDMbcJqL", sdk.SdkVersionNewParams{
	APIVersion: sdk.F[string]("1.2.0"),
	Version:    sdk.F[string]("1.2.0"),
})
if err != nil {
	panic(err)
}

fmt.Println(version)
```

#### Delete an SDK version

Permanently delete one version of an SDK.

```go
version, err := client.Sdks.Versions.Delete(context.Background(), "UakgbKJ5m9gl0JDMbcJqL", "1.2.0")
if err != nil {
	panic(err)
}

fmt.Println(version)
```

### `Sdks Repositories`

SDKs

#### Link a repository

Link one language target to a GitHub repository, so builds sync there.

| Direction | Type |
| --- | --- |
| Request | [`SdkRepositoryLinkParams`](./sdkrepository.go) |
| Response | [`SdkRepositoryLinkResponse`](./sdkrepository.go) |

```go
repository, err := client.Sdks.Repositories.Link(context.Background(), "UakgbKJ5m9gl0JDMbcJqL", sdk.SdkRepositoryLinkParams{
	BaseBranch:   sdk.F[string]("main"),
	Language:     sdk.F[sdk.SdkRepositoryLinkParamsLanguage](sdk.SdkRepositoryLinkParamsLanguage("typescript")),
	RepositoryID: sdk.F[int64](123456789),
})
if err != nil {
	panic(err)
}

fmt.Println(repository)
```

#### Unlink a repository

Unlink one language target from its repository.

```go
repository, err := client.Sdks.Repositories.Unlink(context.Background(), "UakgbKJ5m9gl0JDMbcJqL", "typescript")
if err != nil {
	panic(err)
}

fmt.Println(repository)
```

#### Update publishing settings

Toggle publish-on-merge and the release settings for a linked target.

| Direction | Type |
| --- | --- |
| Request | [`SdkRepositoryUpdatePublishingParams`](./sdkrepository.go) |

```go
repository, err := client.Sdks.Repositories.UpdatePublishing(context.Background(), "UakgbKJ5m9gl0JDMbcJqL", "typescript", sdk.SdkRepositoryUpdatePublishingParams{
	PublishOnMerge: sdk.F[bool](true),
})
if err != nil {
	panic(err)
}

fmt.Println(repository)
```

## `Mcp`

### `Mcp Servers`

MCP

#### List all MCP servers

List every MCP server on the team.

| Direction | Type |
| --- | --- |
| Response | [`[]McpServer`](./mcpserver.go) |

```go
server, err := client.Mcp.Servers.List(context.Background())
if err != nil {
	panic(err)
}

fmt.Println(server)
```

#### Create an MCP server

Create an MCP server over one or more API document versions. The response carries the server and its first installation.

| Direction | Type |
| --- | --- |
| Request | [`McpServerNewParams`](./mcpserver.go) |
| Response | [`McpServerNewResponse`](./mcpserver.go) |

```go
server, err := client.Mcp.Servers.New(context.Background(), sdk.McpServerNewParams{
	Name: sdk.F[string]("Acme MCP"),
})
if err != nil {
	panic(err)
}

fmt.Println(server)
```

#### Get an MCP server

Get a single MCP server by its id.

| Direction | Type |
| --- | --- |
| Response | [`McpServer`](./mcpserver.go) |

```go
server, err := client.Mcp.Servers.Get(context.Background(), "42")
if err != nil {
	panic(err)
}

fmt.Println(server)
```

#### Update an MCP server

Update MCP server metadata and which tools it exposes.

| Direction | Type |
| --- | --- |
| Request | [`McpServerUpdateParams`](./mcpserver.go) |
| Response | [`McpServer`](./mcpserver.go) |

```go
server, err := client.Mcp.Servers.Update(context.Background(), "42", sdk.McpServerUpdateParams{})
if err != nil {
	panic(err)
}

fmt.Println(server)
```

#### Delete an MCP server

Delete an MCP server and every installation it serves.

```go
server, err := client.Mcp.Servers.Delete(context.Background(), "42")
if err != nil {
	panic(err)
}

fmt.Println(server)
```

#### `Mcp Servers Installations`

MCP

##### List installations

List the installations of an MCP server. An installation is what an MCP client connects to.

| Direction | Type |
| --- | --- |
| Response | [`[]McpInstallationListItem`](./mcpserverinstallation.go) |

```go
installation, err := client.Mcp.Servers.Installations.List(context.Background(), "42")
if err != nil {
	panic(err)
}

fmt.Println(installation)
```

##### Create an installation

Create an installation of an MCP server. `documentAuth` holds the credentials the server presents to the upstream API and is never returned.

| Direction | Type |
| --- | --- |
| Request | [`McpServerInstallationNewParams`](./mcpserverinstallation.go) |
| Response | [`McpInstallation`](./mcpserver.go) |

```go
installation, err := client.Mcp.Servers.Installations.New(context.Background(), "42", sdk.McpServerInstallationNewParams{
	DocumentAuth: sdk.F[map[string]interface{}](map[string]interface{}{}),
	Name:         sdk.F[string]("Acme MCP"),
})
if err != nil {
	panic(err)
}

fmt.Println(installation)
```

##### Get an installation

Get a single installation of an MCP server.

| Direction | Type |
| --- | --- |
| Response | [`McpInstallation`](./mcpserver.go) |

```go
installation, err := client.Mcp.Servers.Installations.Get(context.Background(), "42", "84")
if err != nil {
	panic(err)
}

fmt.Println(installation)
```

##### Update an installation

Update an installation. Set `isPrivate` and add access groups to put it behind a login.

| Direction | Type |
| --- | --- |
| Request | [`McpServerInstallationUpdateParams`](./mcpserverinstallation.go) |
| Response | [`McpInstallation`](./mcpserver.go) |

```go
installation, err := client.Mcp.Servers.Installations.Update(context.Background(), "42", "84", sdk.McpServerInstallationUpdateParams{})
if err != nil {
	panic(err)
}

fmt.Println(installation)
```

##### Delete an installation

Delete an installation of an MCP server.

```go
installation, err := client.Mcp.Servers.Installations.Delete(context.Background(), "42", "84")
if err != nil {
	panic(err)
}

fmt.Println(installation)
```

##### Add an access group

Let an access group reach a private installation.

| Direction | Type |
| --- | --- |
| Request | [`McpServerInstallationNewAccessGroupParams`](./mcpserverinstallation.go) |

```go
installation, err := client.Mcp.Servers.Installations.NewAccessGroup(context.Background(), "42", "84", sdk.McpServerInstallationNewAccessGroupParams{
	AccessGroupUID: sdk.F[string]("UakgbKJ5m9gl0JDMbcJqL"),
})
if err != nil {
	panic(err)
}

fmt.Println(installation)
```

##### Remove an access group

Stop an access group reaching a private installation.

| Direction | Type |
| --- | --- |
| Request | [`McpServerInstallationDeleteAccessGroupParams`](./mcpserverinstallation.go) |

```go
installation, err := client.Mcp.Servers.Installations.DeleteAccessGroup(context.Background(), "42", "84", sdk.McpServerInstallationDeleteAccessGroupParams{
	AccessGroupUID: sdk.F[string]("UakgbKJ5m9gl0JDMbcJqL"),
})
if err != nil {
	panic(err)
}

fmt.Println(installation)
```
