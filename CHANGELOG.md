# Changelog

## [0.3.3](https://github.com/scalar/scalar-go/compare/v0.3.2...v0.3.3) (2026-08-07)


### Chores

* **api:** regenerate SDK ([e42fe57](https://github.com/scalar/scalar-go/commit/e42fe5753f11219fc6120edca090cce4bd353f0c))
* **api:** update generated SDK content ([dfc277c](https://github.com/scalar/scalar-go/commit/dfc277ce6ad36e2758c24f2bb1069c9ffe630f34))
* **api:** update generated SDK content ([3e11d54](https://github.com/scalar/scalar-go/commit/3e11d54ec66bba80b4307da1dfc06e3ab844308f))

## [0.3.2](https://github.com/scalar/scalar-go/compare/v0.3.1...v0.3.2) (2026-07-28)


### Chores

* **api:** regenerate SDK ([0de1cd0](https://github.com/scalar/scalar-go/commit/0de1cd02992c79f52b3edde36e10a96406243f63))

## [0.3.1](https://github.com/scalar/scalar-go/compare/v0.3.0...v0.3.1) (2026-07-27)


### Chores

* **api:** regenerate SDK ([6d74d45](https://github.com/scalar/scalar-go/commit/6d74d455978650aaf8d63b3d47f830e5422c90b6))

## [0.3.0](https://github.com/scalar/scalar-go/compare/v0.2.0...v0.3.0) (2026-07-23)


### ⚠ BREAKING CHANGES

* **api:** 316 breaking changes to the SDK surface.
    - `400` error response of `registry.listAllApiDocuments` changed from `_400` to `400`.
    - `401` error response of `registry.listAllApiDocuments` changed from `_401` to `401`.
    - `403` error response of `registry.listAllApiDocuments` changed from `_403` to `403`.
    - `404` error response of `registry.listAllApiDocuments` changed from `_404` to `404`.
    - `422` error response of `registry.listAllApiDocuments` changed from `_422` to `422`.
    - `500` error response of `registry.listAllApiDocuments` changed from `_500` to `500`.
    - `400` error response of `registry.listApiDocuments` changed from `_400` to `400`.
    - `401` error response of `registry.listApiDocuments` changed from `_401` to `401`.
    - `403` error response of `registry.listApiDocuments` changed from `_403` to `403`.
    - `404` error response of `registry.listApiDocuments` changed from `_404` to `404`.
    - `422` error response of `registry.listApiDocuments` changed from `_422` to `422`.
    - `500` error response of `registry.listApiDocuments` changed from `_500` to `500`.
    - `400` error response of `registry.createApiDocument` changed from `_400` to `400`.
    - `401` error response of `registry.createApiDocument` changed from `_401` to `401`.
    - `403` error response of `registry.createApiDocument` changed from `_403` to `403`.
    - `404` error response of `registry.createApiDocument` changed from `_404` to `404`.
    - `422` error response of `registry.createApiDocument` changed from `_422` to `422`.
    - `500` error response of `registry.createApiDocument` changed from `_500` to `500`.
    - `400` error response of `registry.updateApiDocument` changed from `_400` to `400`.
    - `401` error response of `registry.updateApiDocument` changed from `_401` to `401`.
    - `403` error response of `registry.updateApiDocument` changed from `_403` to `403`.
    - `404` error response of `registry.updateApiDocument` changed from `_404` to `404`.
    - `422` error response of `registry.updateApiDocument` changed from `_422` to `422`.
    - `500` error response of `registry.updateApiDocument` changed from `_500` to `500`.
    - `400` error response of `registry.deleteApiDocument` changed from `_400` to `400`.
    - `401` error response of `registry.deleteApiDocument` changed from `_401` to `401`.
    - `403` error response of `registry.deleteApiDocument` changed from `_403` to `403`.
    - `404` error response of `registry.deleteApiDocument` changed from `_404` to `404`.
    - `422` error response of `registry.deleteApiDocument` changed from `_422` to `422`.
    - `500` error response of `registry.deleteApiDocument` changed from `_500` to `500`.
    - `400` error response of `registry.retrieveApiDocumentVersion` changed from `_400` to `400`.
    - `401` error response of `registry.retrieveApiDocumentVersion` changed from `_401` to `401`.
    - `403` error response of `registry.retrieveApiDocumentVersion` changed from `_403` to `403`.
    - `404` error response of `registry.retrieveApiDocumentVersion` changed from `_404` to `404`.
    - `422` error response of `registry.retrieveApiDocumentVersion` changed from `_422` to `422`.
    - `500` error response of `registry.retrieveApiDocumentVersion` changed from `_500` to `500`.
    - `400` error response of `registry.updateApiDocumentVersion` changed from `_400` to `400`.
    - `401` error response of `registry.updateApiDocumentVersion` changed from `_401` to `401`.
    - `403` error response of `registry.updateApiDocumentVersion` changed from `_403` to `403`.
    - `404` error response of `registry.updateApiDocumentVersion` changed from `_404` to `404`.
    - `422` error response of `registry.updateApiDocumentVersion` changed from `_422` to `422`.
    - `500` error response of `registry.updateApiDocumentVersion` changed from `_500` to `500`.
    - `400` error response of `registry.deleteApiDocumentVersion` changed from `_400` to `400`.
    - `401` error response of `registry.deleteApiDocumentVersion` changed from `_401` to `401`.
    - `403` error response of `registry.deleteApiDocumentVersion` changed from `_403` to `403`.
    - `404` error response of `registry.deleteApiDocumentVersion` changed from `_404` to `404`.
    - `422` error response of `registry.deleteApiDocumentVersion` changed from `_422` to `422`.
    - `500` error response of `registry.deleteApiDocumentVersion` changed from `_500` to `500`.
    - Response of `registry.listApiDocumentVersionMetadata` changed from `ManagedDocVersion` to `managed_doc_version`.
    - `400` error response of `registry.listApiDocumentVersionMetadata` changed from `_400` to `400`.
    - `401` error response of `registry.listApiDocumentVersionMetadata` changed from `_401` to `401`.
    - `403` error response of `registry.listApiDocumentVersionMetadata` changed from `_403` to `403`.
    - `404` error response of `registry.listApiDocumentVersionMetadata` changed from `_404` to `404`.
    - `422` error response of `registry.listApiDocumentVersionMetadata` changed from `_422` to `422`.
    - `500` error response of `registry.listApiDocumentVersionMetadata` changed from `_500` to `500`.
    - Response of `registry.createApiDocumentVersion` changed from `ManagedDocVersion` to `managed_doc_version`.
    - `400` error response of `registry.createApiDocumentVersion` changed from `_400` to `400`.
    - `401` error response of `registry.createApiDocumentVersion` changed from `_401` to `401`.
    - `403` error response of `registry.createApiDocumentVersion` changed from `_403` to `403`.
    - `404` error response of `registry.createApiDocumentVersion` changed from `_404` to `404`.
    - `422` error response of `registry.createApiDocumentVersion` changed from `_422` to `422`.
    - `500` error response of `registry.createApiDocumentVersion` changed from `_500` to `500`.
    - `400` error response of `registry.createApiDocumentAccessGroup` changed from `_400` to `400`.
    - `401` error response of `registry.createApiDocumentAccessGroup` changed from `_401` to `401`.
    - `403` error response of `registry.createApiDocumentAccessGroup` changed from `_403` to `403`.
    - `404` error response of `registry.createApiDocumentAccessGroup` changed from `_404` to `404`.
    - `422` error response of `registry.createApiDocumentAccessGroup` changed from `_422` to `422`.
    - `500` error response of `registry.createApiDocumentAccessGroup` changed from `_500` to `500`.
    - `400` error response of `registry.deleteApiDocumentAccessGroup` changed from `_400` to `400`.
    - `401` error response of `registry.deleteApiDocumentAccessGroup` changed from `_401` to `401`.
    - `403` error response of `registry.deleteApiDocumentAccessGroup` changed from `_403` to `403`.
    - `404` error response of `registry.deleteApiDocumentAccessGroup` changed from `_404` to `404`.
    - `422` error response of `registry.deleteApiDocumentAccessGroup` changed from `_422` to `422`.
    - `500` error response of `registry.deleteApiDocumentAccessGroup` changed from `_500` to `500`.
    - `400` error response of `schemas.list` changed from `_400` to `400`.
    - `401` error response of `schemas.list` changed from `_401` to `401`.
    - `403` error response of `schemas.list` changed from `_403` to `403`.
    - `404` error response of `schemas.list` changed from `_404` to `404`.
    - `422` error response of `schemas.list` changed from `_422` to `422`.
    - `500` error response of `schemas.list` changed from `_500` to `500`.
    - Response of `schemas.create` changed from `Uid` to `uid`.
    - `400` error response of `schemas.create` changed from `_400` to `400`.
    - `401` error response of `schemas.create` changed from `_401` to `401`.
    - `403` error response of `schemas.create` changed from `_403` to `403`.
    - `404` error response of `schemas.create` changed from `_404` to `404`.
    - `422` error response of `schemas.create` changed from `_422` to `422`.
    - `500` error response of `schemas.create` changed from `_500` to `500`.
    - `400` error response of `schemas.update` changed from `_400` to `400`.
    - `401` error response of `schemas.update` changed from `_401` to `401`.
    - `403` error response of `schemas.update` changed from `_403` to `403`.
    - `404` error response of `schemas.update` changed from `_404` to `404`.
    - `422` error response of `schemas.update` changed from `_422` to `422`.
    - `500` error response of `schemas.update` changed from `_500` to `500`.
    - `400` error response of `schemas.delete` changed from `_400` to `400`.
    - `401` error response of `schemas.delete` changed from `_401` to `401`.
    - `403` error response of `schemas.delete` changed from `_403` to `403`.
    - `404` error response of `schemas.delete` changed from `_404` to `404`.
    - `422` error response of `schemas.delete` changed from `_422` to `422`.
    - `500` error response of `schemas.delete` changed from `_500` to `500`.
    - `400` error response of `schemas.version.retrieveSchema` changed from `_400` to `400`.
    - `401` error response of `schemas.version.retrieveSchema` changed from `_401` to `401`.
    - `403` error response of `schemas.version.retrieveSchema` changed from `_403` to `403`.
    - `404` error response of `schemas.version.retrieveSchema` changed from `_404` to `404`.
    - `422` error response of `schemas.version.retrieveSchema` changed from `_422` to `422`.
    - `500` error response of `schemas.version.retrieveSchema` changed from `_500` to `500`.
    - `400` error response of `schemas.version.deleteSchema` changed from `_400` to `400`.
    - `401` error response of `schemas.version.deleteSchema` changed from `_401` to `401`.
    - `403` error response of `schemas.version.deleteSchema` changed from `_403` to `403`.
    - `404` error response of `schemas.version.deleteSchema` changed from `_404` to `404`.
    - `422` error response of `schemas.version.deleteSchema` changed from `_422` to `422`.
    - `500` error response of `schemas.version.deleteSchema` changed from `_500` to `500`.
    - Response of `schemas.version.createSchema` changed from `Uid` to `uid`.
    - `400` error response of `schemas.version.createSchema` changed from `_400` to `400`.
    - `401` error response of `schemas.version.createSchema` changed from `_401` to `401`.
    - `403` error response of `schemas.version.createSchema` changed from `_403` to `403`.
    - `404` error response of `schemas.version.createSchema` changed from `_404` to `404`.
    - `422` error response of `schemas.version.createSchema` changed from `_422` to `422`.
    - `500` error response of `schemas.version.createSchema` changed from `_500` to `500`.
    - `400` error response of `schemas.accessGroup.createSchema` changed from `_400` to `400`.
    - `401` error response of `schemas.accessGroup.createSchema` changed from `_401` to `401`.
    - `403` error response of `schemas.accessGroup.createSchema` changed from `_403` to `403`.
    - `404` error response of `schemas.accessGroup.createSchema` changed from `_404` to `404`.
    - `422` error response of `schemas.accessGroup.createSchema` changed from `_422` to `422`.
    - `500` error response of `schemas.accessGroup.createSchema` changed from `_500` to `500`.
    - `400` error response of `schemas.accessGroup.deleteSchema` changed from `_400` to `400`.
    - `401` error response of `schemas.accessGroup.deleteSchema` changed from `_401` to `401`.
    - `403` error response of `schemas.accessGroup.deleteSchema` changed from `_403` to `403`.
    - `404` error response of `schemas.accessGroup.deleteSchema` changed from `_404` to `404`.
    - `422` error response of `schemas.accessGroup.deleteSchema` changed from `_422` to `422`.
    - `500` error response of `schemas.accessGroup.deleteSchema` changed from `_500` to `500`.
    - `400` error response of `loginPortals.retrieve` changed from `_400` to `400`.
    - `401` error response of `loginPortals.retrieve` changed from `_401` to `401`.
    - `403` error response of `loginPortals.retrieve` changed from `_403` to `403`.
    - `404` error response of `loginPortals.retrieve` changed from `_404` to `404`.
    - `422` error response of `loginPortals.retrieve` changed from `_422` to `422`.
    - `500` error response of `loginPortals.retrieve` changed from `_500` to `500`.
    - `400` error response of `loginPortals.update` changed from `_400` to `400`.
    - `401` error response of `loginPortals.update` changed from `_401` to `401`.
    - `403` error response of `loginPortals.update` changed from `_403` to `403`.
    - `404` error response of `loginPortals.update` changed from `_404` to `404`.
    - `422` error response of `loginPortals.update` changed from `_422` to `422`.
    - `500` error response of `loginPortals.update` changed from `_500` to `500`.
    - `400` error response of `loginPortals.delete` changed from `_400` to `400`.
    - `401` error response of `loginPortals.delete` changed from `_401` to `401`.
    - `403` error response of `loginPortals.delete` changed from `_403` to `403`.
    - `404` error response of `loginPortals.delete` changed from `_404` to `404`.
    - `422` error response of `loginPortals.delete` changed from `_422` to `422`.
    - `500` error response of `loginPortals.delete` changed from `_500` to `500`.
    - Response of `loginPortals.create` changed from `Uid` to `uid`.
    - `400` error response of `loginPortals.create` changed from `_400` to `400`.
    - `401` error response of `loginPortals.create` changed from `_401` to `401`.
    - `403` error response of `loginPortals.create` changed from `_403` to `403`.
    - `404` error response of `loginPortals.create` changed from `_404` to `404`.
    - `422` error response of `loginPortals.create` changed from `_422` to `422`.
    - `500` error response of `loginPortals.create` changed from `_500` to `500`.
    - `400` error response of `loginPortals.list` changed from `_400` to `400`.
    - `401` error response of `loginPortals.list` changed from `_401` to `401`.
    - `403` error response of `loginPortals.list` changed from `_403` to `403`.
    - `404` error response of `loginPortals.list` changed from `_404` to `404`.
    - `422` error response of `loginPortals.list` changed from `_422` to `422`.
    - `500` error response of `loginPortals.list` changed from `_500` to `500`.
    - `400` error response of `rules.listRulesets` changed from `_400` to `400`.
    - `401` error response of `rules.listRulesets` changed from `_401` to `401`.
    - `403` error response of `rules.listRulesets` changed from `_403` to `403`.
    - `404` error response of `rules.listRulesets` changed from `_404` to `404`.
    - `422` error response of `rules.listRulesets` changed from `_422` to `422`.
    - `500` error response of `rules.listRulesets` changed from `_500` to `500`.
    - Response of `rules.createRuleset` changed from `Uid` to `uid`.
    - `400` error response of `rules.createRuleset` changed from `_400` to `400`.
    - `401` error response of `rules.createRuleset` changed from `_401` to `401`.
    - `403` error response of `rules.createRuleset` changed from `_403` to `403`.
    - `404` error response of `rules.createRuleset` changed from `_404` to `404`.
    - `422` error response of `rules.createRuleset` changed from `_422` to `422`.
    - `500` error response of `rules.createRuleset` changed from `_500` to `500`.
    - `400` error response of `rules.updateRuleset` changed from `_400` to `400`.
    - `401` error response of `rules.updateRuleset` changed from `_401` to `401`.
    - `403` error response of `rules.updateRuleset` changed from `_403` to `403`.
    - `404` error response of `rules.updateRuleset` changed from `_404` to `404`.
    - `422` error response of `rules.updateRuleset` changed from `_422` to `422`.
    - `500` error response of `rules.updateRuleset` changed from `_500` to `500`.
    - `400` error response of `rules.deleteRuleset` changed from `_400` to `400`.
    - `401` error response of `rules.deleteRuleset` changed from `_401` to `401`.
    - `403` error response of `rules.deleteRuleset` changed from `_403` to `403`.
    - `404` error response of `rules.deleteRuleset` changed from `_404` to `404`.
    - `422` error response of `rules.deleteRuleset` changed from `_422` to `422`.
    - `500` error response of `rules.deleteRuleset` changed from `_500` to `500`.
    - `400` error response of `rules.retrieveRulesetDocument` changed from `_400` to `400`.
    - `401` error response of `rules.retrieveRulesetDocument` changed from `_401` to `401`.
    - `403` error response of `rules.retrieveRulesetDocument` changed from `_403` to `403`.
    - `404` error response of `rules.retrieveRulesetDocument` changed from `_404` to `404`.
    - `422` error response of `rules.retrieveRulesetDocument` changed from `_422` to `422`.
    - `500` error response of `rules.retrieveRulesetDocument` changed from `_500` to `500`.
    - `400` error response of `rules.createRulesetAccessGroup` changed from `_400` to `400`.
    - `401` error response of `rules.createRulesetAccessGroup` changed from `_401` to `401`.
    - `403` error response of `rules.createRulesetAccessGroup` changed from `_403` to `403`.
    - `404` error response of `rules.createRulesetAccessGroup` changed from `_404` to `404`.
    - `422` error response of `rules.createRulesetAccessGroup` changed from `_422` to `422`.
    - `500` error response of `rules.createRulesetAccessGroup` changed from `_500` to `500`.
    - `400` error response of `rules.deleteRulesetAccessGroup` changed from `_400` to `400`.
    - `401` error response of `rules.deleteRulesetAccessGroup` changed from `_401` to `401`.
    - `403` error response of `rules.deleteRulesetAccessGroup` changed from `_403` to `403`.
    - `404` error response of `rules.deleteRulesetAccessGroup` changed from `_404` to `404`.
    - `422` error response of `rules.deleteRulesetAccessGroup` changed from `_422` to `422`.
    - `500` error response of `rules.deleteRulesetAccessGroup` changed from `_500` to `500`.
    - `400` error response of `themes.list` changed from `_400` to `400`.
    - `401` error response of `themes.list` changed from `_401` to `401`.
    - `403` error response of `themes.list` changed from `_403` to `403`.
    - `404` error response of `themes.list` changed from `_404` to `404`.
    - `422` error response of `themes.list` changed from `_422` to `422`.
    - `500` error response of `themes.list` changed from `_500` to `500`.
    - Response of `themes.create` changed from `Uid` to `uid`.
    - `400` error response of `themes.create` changed from `_400` to `400`.
    - `401` error response of `themes.create` changed from `_401` to `401`.
    - `403` error response of `themes.create` changed from `_403` to `403`.
    - `404` error response of `themes.create` changed from `_404` to `404`.
    - `422` error response of `themes.create` changed from `_422` to `422`.
    - `500` error response of `themes.create` changed from `_500` to `500`.
    - `400` error response of `themes.update` changed from `_400` to `400`.
    - `401` error response of `themes.update` changed from `_401` to `401`.
    - `403` error response of `themes.update` changed from `_403` to `403`.
    - `404` error response of `themes.update` changed from `_404` to `404`.
    - `422` error response of `themes.update` changed from `_422` to `422`.
    - `500` error response of `themes.update` changed from `_500` to `500`.
    - `400` error response of `themes.replaceDocument` changed from `_400` to `400`.
    - `401` error response of `themes.replaceDocument` changed from `_401` to `401`.
    - `403` error response of `themes.replaceDocument` changed from `_403` to `403`.
    - `404` error response of `themes.replaceDocument` changed from `_404` to `404`.
    - `422` error response of `themes.replaceDocument` changed from `_422` to `422`.
    - `500` error response of `themes.replaceDocument` changed from `_500` to `500`.
    - `400` error response of `themes.delete` changed from `_400` to `400`.
    - `401` error response of `themes.delete` changed from `_401` to `401`.
    - `403` error response of `themes.delete` changed from `_403` to `403`.
    - `404` error response of `themes.delete` changed from `_404` to `404`.
    - `422` error response of `themes.delete` changed from `_422` to `422`.
    - `500` error response of `themes.delete` changed from `_500` to `500`.
    - `400` error response of `themes.retrieve` changed from `_400` to `400`.
    - `401` error response of `themes.retrieve` changed from `_401` to `401`.
    - `403` error response of `themes.retrieve` changed from `_403` to `403`.
    - `404` error response of `themes.retrieve` changed from `_404` to `404`.
    - `422` error response of `themes.retrieve` changed from `_422` to `422`.
    - `500` error response of `themes.retrieve` changed from `_500` to `500`.
    - `400` error response of `teams.list` changed from `_400` to `400`.
    - `401` error response of `teams.list` changed from `_401` to `401`.
    - `403` error response of `teams.list` changed from `_403` to `403`.
    - `404` error response of `teams.list` changed from `_404` to `404`.
    - `422` error response of `teams.list` changed from `_422` to `422`.
    - `500` error response of `teams.list` changed from `_500` to `500`.
    - `400` error response of `scalarDocs.listGuides` changed from `_400` to `400`.
    - `401` error response of `scalarDocs.listGuides` changed from `_401` to `401`.
    - `403` error response of `scalarDocs.listGuides` changed from `_403` to `403`.
    - `404` error response of `scalarDocs.listGuides` changed from `_404` to `404`.
    - `422` error response of `scalarDocs.listGuides` changed from `_422` to `422`.
    - `500` error response of `scalarDocs.listGuides` changed from `_500` to `500`.
    - `400` error response of `scalarDocs.createGuide` changed from `_400` to `400`.
    - `401` error response of `scalarDocs.createGuide` changed from `_401` to `401`.
    - `403` error response of `scalarDocs.createGuide` changed from `_403` to `403`.
    - `404` error response of `scalarDocs.createGuide` changed from `_404` to `404`.
    - `422` error response of `scalarDocs.createGuide` changed from `_422` to `422`.
    - `500` error response of `scalarDocs.createGuide` changed from `_500` to `500`.
    - `400` error response of `scalarDocs.publishGuide` changed from `_400` to `400`.
    - `401` error response of `scalarDocs.publishGuide` changed from `_401` to `401`.
    - `403` error response of `scalarDocs.publishGuide` changed from `_403` to `403`.
    - `404` error response of `scalarDocs.publishGuide` changed from `_404` to `404`.
    - `422` error response of `scalarDocs.publishGuide` changed from `_422` to `422`.
    - `500` error response of `scalarDocs.publishGuide` changed from `_500` to `500`.
    - `400` error response of `namespaces.list` changed from `_400` to `400`.
    - `401` error response of `namespaces.list` changed from `_401` to `401`.
    - `403` error response of `namespaces.list` changed from `_403` to `403`.
    - `404` error response of `namespaces.list` changed from `_404` to `404`.
    - `422` error response of `namespaces.list` changed from `_422` to `422`.
    - `500` error response of `namespaces.list` changed from `_500` to `500`.
    - `400` error response of `authentication.exchangePersonalToken` changed from `_400` to `400`.
    - `401` error response of `authentication.exchangePersonalToken` changed from `_401` to `401`.
    - `403` error response of `authentication.exchangePersonalToken` changed from `_403` to `403`.
    - `404` error response of `authentication.exchangePersonalToken` changed from `_404` to `404`.
    - `422` error response of `authentication.exchangePersonalToken` changed from `_422` to `422`.
    - `500` error response of `authentication.exchangePersonalToken` changed from `_500` to `500`.
    - Response of `authentication.listCurrentUser` changed from `User` to `user`.
    - `400` error response of `authentication.listCurrentUser` changed from `_400` to `400`.
    - `401` error response of `authentication.listCurrentUser` changed from `_401` to `401`.
    - `403` error response of `authentication.listCurrentUser` changed from `_403` to `403`.
    - `404` error response of `authentication.listCurrentUser` changed from `_404` to `404`.
    - `422` error response of `authentication.listCurrentUser` changed from `_422` to `422`.
    - `500` error response of `authentication.listCurrentUser` changed from `_500` to `500`.
    - Removed schema `_400`.
    - Removed schema `_401`.
    - Removed schema `_403`.
    - Removed schema `_404`.
    - Removed schema `_422`.
    - Removed schema `_500`.
    - Removed schema `ApiDocument`.
    - Removed schema `Nanoid`.
    - Removed schema `Version`.
    - Removed schema `Slug`.
    - Removed schema `Namespace`.
    - Removed schema `ManagedDocVersion`.
    - Removed schema `Method`.
    - Removed schema `AccessGroup`.
    - Removed schema `Schema`.
    - Removed schema `ManagedSchemaVersion`.
    - Removed schema `Timestamp`.
    - Removed schema `Uid`.
    - Removed schema `LoginPortalEmail`.
    - Removed schema `LoginPortalPage`.
    - Removed schema `LoginPortal`.
    - Removed schema `Rule`.
    - Removed schema `Theme`.
    - Removed schema `Team`.
    - Removed schema `TeamName`.
    - Removed schema `TeamImage`.
    - Removed schema `GithubProject`.
    - Removed schema `ActiveDeployment`.
    - Removed schema `GithubProjectRepository`.
    - Removed schema `Email`.
    - Removed schema `TeamSummary`.
    - Removed schema `User`.

### Features

* **api:** update SDK surface (349 changes) ([362a5b3](https://github.com/scalar/scalar-go/commit/362a5b3b87b2fa0a97006143614881cea5a8b662))

## [0.2.0](https://github.com/scalar/scalar-go/compare/v0.1.0...v0.2.0) (2026-07-20)


### Features

* **api:** initial SDK generation ([a74adc7](https://github.com/scalar/scalar-go/commit/a74adc70bc83ace7f96538f6b7197b56a9774d2e))


### Chores

* **api:** update generated SDK content ([9fa10b8](https://github.com/scalar/scalar-go/commit/9fa10b812e871152c627ba33cca581b7b0d2cf71))
