# Permissions

This document lists all Permissions documented in the Protobuf files via the `(buf.registry.priv.extension.v1beta1.method)` option on RPCs.

A Permission is an atomic capability that gates access to one or more RPCs.
See [permission.proto](permission.proto) for the `Permission` message definition and [extension.proto](../../priv/extension/v1beta1/extension.proto) for the method option that declares Permission requirements.

## Scopes

A grant or restriction applies within a [Scope](scope.proto), which is either the entire instance or a single resource (a User, Organization, Module, Plugin, or Policy).
A resource Scope includes the resources it owns, so a Permission granted on an Organization or User cascades down to the Modules (and their Commits, Labels, etc.) owned by that Organization or User.

A Scope is valid for a Permission when the resources the Permission operates on exist at or below that Scope:

- The Instance Scope is valid for every Permission.
- A resource Scope is valid when the Permission targets that resource type or a type owned by it.
  For example, `bsr.module.delete` is valid on a Module, and also on an Organization or User because their Modules are in scope.
  A token Permission is valid on a User but not on an Organization, because Organizations do not contain Tokens.
- Create and list Permissions target a collection, so the narrowest valid Scope is the container, not the resource itself.
  For example, `bsr.module.create` is valid on an Organization or User but not on a Module, and `bsr.organization.create` is valid only on the Instance.

The valid scopes of a Permission are exposed via the `bindable_resource_types` field on the `Permission` message.
The Instance Scope is always valid, so the field enumerates only the valid resource types; the "Bindable resource types" column below lists Instance explicitly for readability.

## Permissions by name

| Permission | Bindable resource types |
|------------|-------------------------|
| `bsr.commit.get` | Instance, Organization, User, Module |
| `bsr.commit.list` | Instance, Organization, User, Module |
| `bsr.file-descriptor-set.get` | Instance, Organization, User, Module |
| `bsr.graph.get` | Instance, Organization, User, Module |
| `bsr.label.archive` | Instance, Organization, User, Module |
| `bsr.label.create` | Instance, Organization, User, Module |
| `bsr.label.get` | Instance, Organization, User, Module |
| `bsr.label.list` | Instance, Organization, User, Module |
| `bsr.label.list-history` | Instance, Organization, User, Module |
| `bsr.label.unarchive` | Instance, Organization, User, Module |
| `bsr.label.update` | Instance, Organization, User, Module |
| `bsr.module-contributor.create` | Instance, Organization, User, Module |
| `bsr.module-contributor.delete` | Instance, Organization, User, Module |
| `bsr.module-contributor.get` | Instance, Organization, User, Module |
| `bsr.module-contributor.list` | Instance, Organization, User, Module |
| `bsr.module-contributor.update` | Instance, Organization, User, Module |
| `bsr.module.create` | Instance, Organization, User |
| `bsr.module.delete` | Instance, Organization, User, Module |
| `bsr.module.download` | Instance, Organization, User, Module |
| `bsr.module.get` | Instance, Organization, User, Module |
| `bsr.module.list` | Instance, Organization, User |
| `bsr.module.update` | Instance, Organization, User, Module |
| `bsr.module.upload` | Instance, Organization, User, Module |
| `bsr.organization-membership.create` | Instance, Organization |
| `bsr.organization-membership.delete` | Instance, Organization |
| `bsr.organization-membership.get` | Instance, Organization, User |
| `bsr.organization-membership.list` | Instance, Organization, User |
| `bsr.organization-membership.update` | Instance, Organization |
| `bsr.organization-settings.get` | Instance, Organization |
| `bsr.organization-settings.update` | Instance, Organization |
| `bsr.organization.create` | Instance |
| `bsr.organization.delete` | Instance, Organization |
| `bsr.organization.get` | Instance, Organization |
| `bsr.organization.list` | Instance |
| `bsr.organization.update` | Instance, Organization |
| `bsr.token.create` | Instance, User |
| `bsr.token.delete` | Instance, User |
| `bsr.token.get` | Instance, User |
| `bsr.token.list` | Instance, User |
| `bsr.token.update` | Instance, User |
| `bsr.user.create` | Instance |
| `bsr.user.delete` | Instance, User |
| `bsr.user.get` | Instance, User |
| `bsr.user.get-instance-role` | Instance |
| `bsr.user.list` | Instance |
| `bsr.user.update` | Instance, User |
| `bsr.user.update-instance-role` | Instance |

## RPCs with dynamic Permissions

These RPCs declare `dynamic_permission = true`.
The Permissions required are determined from the contents of the request, as documented on each method:

- `LabelService.CreateOrUpdateLabels`: a Label that does not exist requires `bsr.label.create`, and a Label that exists requires `bsr.label.update`.
- `OwnerService.GetOwners`: a User requires `bsr.user.get`, and an Organization requires `bsr.organization.get`.
- `ResourceService.GetResources`: a Module requires `bsr.module.get`, a Label requires `bsr.label.get`, and a Commit requires `bsr.commit.get`.

## RPCs with no required Permission

These RPCs declare `no_required_permission = true`.
No Permission check applies; whether the RPC requires authentication is a separate concern.

- `PermissionService.GetPermissions`
- `PermissionService.ListPermissions`
- `UserService.GetCurrentUser`
