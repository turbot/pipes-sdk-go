# Turbot Pipes Go SDK

## 0.12.0 [TBD]

_Enhancements_

- Add support for management of `GitLab` integrations.
- Add support for listing `GitLab` projects for a workspace integration and also list branches for a particular project.
- Add support for returning `LastActivityAt` information for `Tenant`, `Organization` and `Workspace` Member APIs.
- Add support for management of snapshot visibility settings for all workspaces in a custom tenant.

## 0.11.0 [2024-10-29]

_Breaking changes_

- Removed support for `NotificationRules` and `NotificationRule` APIs from the SDK, replaced by `Notifiers`.

_What's new?_

- Flowpipe `Mod`, `Mod Variable`, `Pipeline`, `Trigger` and `Input` APIs for `UserWorkspaces` and `OrgWorkspaces`.
- Notifier Management APIs for `Tenants`, `Users`, `Organizations`, `UserWorkspaces` and `OrgWorkspaces`.
- Billing, Subscription and Payment APIs for `Tenants`, `Users` and `Organizations`.
- Slack Integration APIs have been added to Integration Management APIs for `Tenants`, `Users` and `Organizations`.

_Enhancements_

- Add support to `GetPlan` for `Users` and `Organizations`.
- Add support for executing Process Commands for `UserWorkspaces` and `OrgWorkspaces`.
- Improved validation of some previously `string` fields by migrating to a typed enum.

## 0.10.0 [2024-08-02]

_What's new?_

- Integration Management APIs for `Tenants`, `Users`, `Organizations`, `UserWorkspaces` and `OrgWorkspaces`.
- Tenant Connection Management APIs.
- Tenant Connection Folder Management APIs.
- User Connection Folder Management APIs.
- Organization Connection Management APIs.
- Organization Connection Folder Management APIs.
- Workspace Connection Management APIs.
- Workspace Connection Folder Management APIs.
- Permission Management APIs for `Connections` and `ConnectionFolders`.

_Enhancements_

- Add support for installing mods that are associated to integrations.
- Add support for setting search path prefix for a workspace.

## 0.9.1 [2024-03-21]

_Breaking changes_

- Rename `ListIdentityUsageResponse` to `ListUsageMetricsResponse` to provide a more generic struct to be used across different methods.

_What's new?_

- New struct `UsageMetric` to query and track metrics for either an entity including `tenant`, `identity` or `workspace`.

_Enhancements_

- Added `DesiredState` when creating `Workspace`, `Pipeline`, `Datatank`, `DatatankTable` to allow extra control while setting up the respective resource.
- Added `DbVolumeSizeBytes` to `Workspace` and allow setting workspace database size during create or update.
- Added `PoweripeVersion` to `Workspace` allowing users to track the current version of `powerpipe` installed on the workspace.

## 0.9.0 [2024-01-23]

_What's new?_

- Tenant Management APIs.
- Tenant Member APIs.
- Applicable APIs made `tenant aware`.

## 0.8.1 [2023-11-03]

_Enhancements_

- Modified attribute `Frequency` for `CreateDatatankTableRequest` to use `PipelineFrequency` type.
- Modified attribute `Frequency` for `UpdateDatatankTableRequest` to use `PipelineFrequency` type.
- Modified attribute `Frequency` for `DatatankTable` to use `PipelineFrequency` type.
- Modified attribute `Frequency` for `DatatankPart` to use `PipelineFrequency` type.
- Updated `Value`, `ValueRounded` & `ValueWeighted` attributes for `IdentityUsage` struct to use `int64` type.
- Updated `UsageComputeThreshold`, `UsageStorageThreshold` & `UsageUserThreshold` attributes for `Org` struct to use `int64` type.
- Updated `UsageComputeThreshold` & `UsageStorageThreshold` attributes for `User` struct to use `int64` type.
- Updated `UsageComputeThreshold`, `UsageStorageThreshold` & `UsageUserThreshold` attributes for `UpdateOrgRequest` struct to use `int64` type.
- Updated `UsageComputeThreshold` & `UsageStorageThreshold` attributes for `UpdateUserRequest` struct to use `int64`.

## 0.8.0 [2023-10-20]

_What's new?_

- Identity Datatank APIs.
- Identity Datatank Table APIs.
- Identity Datatank Part APIs.
- Identity Usage APIs.
- Identity Workspace Usage APIs.

_Enhancements_

- Added attribute `InstanceType` to `CreateWorkspaceRequest` for passing the instance type information of the workspace to be created.
- Added attribute `InstanceType` to `Workspace` for storing the instance type information of a workspace.
- Added attribute `InstanceType` to `Pipeline` for storing the instance type information of a pipeline.
- Added various usage attributes to `user` and `org` structs.

## 0.7.0 [2023-05-22]

_What's new?_

- Identity Workspace Schema APIs.
- Identity Constraint APIs.

## 0.6.0 [2023-04-20]

_Breaking changes_

- Removed attribute `AvatarUrl` from CreateOrgRequest struct.
- Removed attribute `AvatarUrl` from UpdateOrgRequest struct.
- Removed attribute `AvatarUrl` from UpdateUserRequest struct.

_What's new?_

- Identity Workspace Aggregator APIs.
- Identity APIs to get Avatar URLs
- Org APIs to create / delete custom avatar.
- User APIs to create / delete custom avatar.
- Add `TokenMinIssuedAt` attribute to Org struct to determine the time after which tokens will be accepted for this identity.
- Add `TokenMinIssuedAt` attribute to User struct to determine the time after which tokens will be accepted for this identity.

## 0.5.0 [2023-02-06]

_Breaking changes_

- Modified attribute `Layout` of the WorkspaceSnapshotData struct to have defined type `WorkspaceSnapshotDataLayout` instead of `map[string]interface{}`.

_What's new?_

- Identity Workspace Pipeline APIs
- Identity Process APIs
- Identity Workspace Process APIs

_Enhancements_

- Added attribute `PluginVersion` in the `Connection` struct to return the version of the plugin.
- Added ability to pass in `where` clauses for org workspace snapshot filtering capabilities.

## 0.4.0 [2023-01-09]

_Breaking changes_

- Removed attribute `id` in the `ErrorModel` struct.
- Renamed attribute `errors` to `validation_errors` in the `ErrorModel` struct.
- Attribute `instance` in the `ErrorModel` struct will now return the instance of the turbot pipes error instead of the `id` attribute where it was previously accessed.

## 0.3.0 [2022-12-06]

- Remove SearchPath attribute from WorkspaceSnapshotData model
- Add ExpiresAt attribute to WorkspaceSnapshot model

## 0.2.0 [2022-11-03]

- User Preferences APIs
- User Emails APIs
- Pipes Login APIs
- Add option to set snapshot titles

## 0.1.3 [2022-08-12]

- Add Snapshot APIs.
- Remove email from all responses where users are returned.
- Various enhancements related to APIs returning redundant data.

## 0.1.2 [2022-07-19]

- Allow users to be added to workspaces directly, bypassing the invite process.
- Add search capability for org members.

## 0.1.1 [2022-07-13]

- Move Orgs for a user under actor orgs.
- Modify response structure for workspaces of an actor.
- Org and Workspace members to be listed via a unified method rather than separate methods for invites and members.

## 0.1.0 [2022-07-01]

- Add Org Workspace Member APIs.
- Modify various existing APIs to include metadata information such as created_by etc. 

## 0.0.4 [2022-05-31]

- Add Mod, Mod Variables APIs

## 0.0.3 [2021-12-15]

- Update the types docs.
- Update the description of input parameters for API.

## 0.0.2 [2021-12-14]

- Updated: Removed the suffix types from the struct.
- Updated: Removed the suffix sperr from the error struct.
- Fixed: Naming of workspace data APIs.

## 0.0.1 [2021-12-13]

- Initial version
