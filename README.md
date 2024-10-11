# Go SDK for Turbot Pipes

Turbot Pipes is an intelligence, automation & security platform built specifically for DevOps.

For help on getting started with Turbot Pipes, please visit https://turbot.com/pipes.

## Getting Started

Here's an example of listing the workspaces for your user:

```go
package main

import (
    "context"
    "fmt"
    "os"

    pipes "github.com/turbot/pipes-sdk-go"
)

func main() {
    // Create a default configuration
    configuration := pipes.NewConfiguration()

    // Add your Turbot Pipes user token as an auth header
    configuration.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("PIPES_TOKEN")))

    // Create a client
    client := pipes.NewAPIClient(configuration)

    // Find your authenticated user info
    actor, _, err := client.Actors.Get(context.Background()).Execute()

    if err != nil {
      // Do something with the error
      return
    }

    // List your workspaces
    workspaces, _, err := client.UserWorkspaces.List(context.Background(), actor.Handle).Execute()

    if err != nil {
      // Do something with the error
      return
    }
}
```

## Usages

For more detailed examples of using the SDK, please check out the following open source projects:

- https://github.com/turbot/steampipe-plugin-pipes
  - The official Steampipe plugin for accessing your Turbot Pipes resources via SQL.
- https://github.com/turbot/terraform-provider-pipes
  - The Terraform provider for managing your Turbot Pipes resources using Terraform.

## Documentation for API Endpoints

All URIs are relative to *https://pipes.turbot.com/api/v0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*Actors* | [**Get**](docs/Actors.md#get) | **Get** /actor | Actor information
*Actors* | [**ListActivity**](docs/Actors.md#listactivity) | **Get** /actor/activity | List actor activity
*Actors* | [**ListConnections**](docs/Actors.md#listconnections) | **Get** /actor/conn | List actor connections
*Actors* | [**ListOrgInvites**](docs/Actors.md#listorginvites) | **Get** /actor/org/invite | List org invites for actor
*Actors* | [**ListOrgs**](docs/Actors.md#listorgs) | **Get** /actor/org | List actor orgs
*Actors* | [**ListWorkspaces**](docs/Actors.md#listworkspaces) | **Get** /actor/workspace | List actor workspaces
*Auth* | [**ConfirmLogin**](docs/Auth.md#confirmlogin) | **Get** /login/confirm | Confirm user login
*Auth* | [**ConfirmSignup**](docs/Auth.md#confirmsignup) | **Get** /signup/confirm | Confirm user signup
*Auth* | [**Login**](docs/Auth.md#login) | **Post** /login | User login
*Auth* | [**LoginTokenCreate**](docs/Auth.md#logintokencreate) | **Post** /login/token | Generate temporary token request
*Auth* | [**LoginTokenDelete**](docs/Auth.md#logintokendelete) | **Delete** /login/token/{temporary_token_request_id} | Delete temporary token request
*Auth* | [**LoginTokenGet**](docs/Auth.md#logintokenget) | **Get** /login/token/{temporary_token_request_id} | Get temporary token request
*Auth* | [**LoginTokenUpdate**](docs/Auth.md#logintokenupdate) | **Patch** /login/token/{temporary_token_request_id} | Update temporary token request
*Auth* | [**Logout**](docs/Auth.md#logout) | **Get** /logout/{provider} | User logout
*Auth* | [**PostProviderCallback**](docs/Auth.md#postprovidercallback) | **Post** /auth/{provider}/callback | Post auth provider callback
*Auth* | [**Provider**](docs/Auth.md#provider) | **Get** /auth/{provider} | Auth Provider
*Auth* | [**ProviderCallback**](docs/Auth.md#providercallback) | **Get** /auth/{provider}/callback | Auth provider callback
*Auth* | [**Signup**](docs/Auth.md#signup) | **Post** /signup | User signup
*Identities* | [**Get**](docs/Identities.md#get) | **Get** /identity/{identity_handle} | Get identity
*Identities* | [**GetAvatar**](docs/Identities.md#getavatar) | **Get** /identity/{identity_handle}/avatar | Get identity avatar
*Identities* | [**List**](docs/Identities.md#list) | **Get** /identity | List identities
*OrgBilling* | [**CreateBillingIntent**](docs/OrgBilling.md#createbillingintent) | **Post** /org/{org_handle}/billing/payment_method/intent | Create Stripe billing intent
*OrgBilling* | [**CreateOrgSubscription**](docs/OrgBilling.md#createorgsubscription) | **Post** /org/{org_handle}/billing/subscription | Create a new subscription
*OrgBilling* | [**DeleteOrgBillingPaymentMethod**](docs/OrgBilling.md#deleteorgbillingpaymentmethod) | **Delete** /org/{org_handle}/billing/payment_method/{payment_method_id} | Delete org billing payment method.
*OrgBilling* | [**DeleteOrgBillingSubscription**](docs/OrgBilling.md#deleteorgbillingsubscription) | **Delete** /org/{org_handle}/billing/subscription/{subscription_id} | Delete org subscription
*OrgBilling* | [**GetOrgBillingInvoice**](docs/OrgBilling.md#getorgbillinginvoice) | **Get** /org/{org_handle}/billing/invoice/{invoice_id} | Get an organization invoice
*OrgBilling* | [**GetOrgBillingPaymentMethod**](docs/OrgBilling.md#getorgbillingpaymentmethod) | **Get** /org/{org_handle}/billing/payment_method/{payment_method_id} | Get org billing payment method.
*OrgBilling* | [**GetOrgBillingSubscription**](docs/OrgBilling.md#getorgbillingsubscription) | **Get** /org/{org_handle}/billing/subscription/{subscription_id} | Get org subscription
*OrgBilling* | [**GetOrgBillingUpcomingInvoice**](docs/OrgBilling.md#getorgbillingupcominginvoice) | **Get** /org/{org_handle}/billing/stripe/invoice/upcoming | Get upcoming invoice
*OrgBilling* | [**ListOrgBillingInvoices**](docs/OrgBilling.md#listorgbillinginvoices) | **Get** /org/{org_handle}/billing/invoice | List org invoices
*OrgBilling* | [**ListOrgBillingPaymentMethod**](docs/OrgBilling.md#listorgbillingpaymentmethod) | **Get** /org/{org_handle}/billing/payment_method | List org billing payment methods.
*OrgBilling* | [**ListOrgSubscription**](docs/OrgBilling.md#listorgsubscription) | **Get** /org/{org_handle}/billing/subscription | List org subscriptions
*OrgBilling* | [**UpdateOrgBillingPaymentMethod**](docs/OrgBilling.md#updateorgbillingpaymentmethod) | **Patch** /org/{org_handle}/billing/payment_method/{payment_method_id} | Update org billing payment method.
*OrgBilling* | [**UpdateOrgBillingSubscription**](docs/OrgBilling.md#updateorgbillingsubscription) | **Patch** /org/{org_handle}/billing/subscription/{subscription_id} | Update org subscription
*OrgConnectionFolders* | [**Create**](docs/OrgConnectionFolders.md#create) | **Post** /org/{org_handle}/connection_folder | Create org connection folder
*OrgConnectionFolders* | [**CreatePermission**](docs/OrgConnectionFolders.md#createpermission) | **Post** /org/{org_handle}/connection_folder/{folder_id}/permission | Create org connection folder permission
*OrgConnectionFolders* | [**Delete**](docs/OrgConnectionFolders.md#delete) | **Delete** /org/{org_handle}/connection_folder/{folder_id} | Delete org connection folder
*OrgConnectionFolders* | [**DeletePermission**](docs/OrgConnectionFolders.md#deletepermission) | **Delete** /org/{org_handle}/connection_folder/{folder_id}/permission/{permission_id} | Delete org connection folder permission
*OrgConnectionFolders* | [**Get**](docs/OrgConnectionFolders.md#get) | **Get** /org/{org_handle}/connection_folder/{folder_id} | Get org connection folder
*OrgConnectionFolders* | [**GetPermission**](docs/OrgConnectionFolders.md#getpermission) | **Get** /org/{org_handle}/connection_folder/{folder_id}/permission/{permission_id} | Get org connection folder permission
*OrgConnectionFolders* | [**List**](docs/OrgConnectionFolders.md#list) | **Get** /org/{org_handle}/connection_folder | List org connection folders
*OrgConnectionFolders* | [**ListPermission**](docs/OrgConnectionFolders.md#listpermission) | **Get** /org/{org_handle}/connection_folder/{folder_id}/permission | List org connection folder permissions
*OrgConnectionFolders* | [**ListWorkspaces**](docs/OrgConnectionFolders.md#listworkspaces) | **Get** /org/{org_handle}/connection_folder/{folder_id}/workspace | List org connection folder workspaces
*OrgConnectionFolders* | [**Update**](docs/OrgConnectionFolders.md#update) | **Patch** /org/{org_handle}/connection_folder/{folder_id} | Update org connection folder
*OrgConnectionFolders* | [**UpdatePermission**](docs/OrgConnectionFolders.md#updatepermission) | **Patch** /org/{org_handle}/connection_folder/{folder_id}/permission/{permission_id} | Update org connection folder permission
*OrgConnectionTree* | [**List**](docs/OrgConnectionTree.md#list) | **Get** /org/{org_handle}/connection_tree | List tenant connection tree
*OrgConnections* | [**Create**](docs/OrgConnections.md#create) | **Post** /org/{org_handle}/connection | Create org connection
*OrgConnections* | [**CreateDeprecated**](docs/OrgConnections.md#createdeprecated) | **Post** /org/{org_handle}/conn | Create org connection
*OrgConnections* | [**CreatePermission**](docs/OrgConnections.md#createpermission) | **Post** /org/{org_handle}/connection/{connection_handle}/permission | Create org connection permission
*OrgConnections* | [**Delete**](docs/OrgConnections.md#delete) | **Delete** /org/{org_handle}/connection/{connection_handle} | Delete org connection
*OrgConnections* | [**DeleteDeprecated**](docs/OrgConnections.md#deletedeprecated) | **Delete** /org/{org_handle}/conn/{conn_handle} | Delete org connection
*OrgConnections* | [**DeletePermission**](docs/OrgConnections.md#deletepermission) | **Delete** /org/{org_handle}/connection/{connection_handle}/permission/{permission_id} | Delete org connection permission
*OrgConnections* | [**Get**](docs/OrgConnections.md#get) | **Get** /org/{org_handle}/connection/{connection_handle} | Get org connection
*OrgConnections* | [**GetDeprecated**](docs/OrgConnections.md#getdeprecated) | **Get** /org/{org_handle}/conn/{conn_handle} | Get org connection
*OrgConnections* | [**GetPermission**](docs/OrgConnections.md#getpermission) | **Get** /org/{org_handle}/connection/{connection_handle}/permission/{permission_id} | Get org connection permission
*OrgConnections* | [**List**](docs/OrgConnections.md#list) | **Get** /org/{org_handle}/connection | List org connections
*OrgConnections* | [**ListDeprecated**](docs/OrgConnections.md#listdeprecated) | **Get** /org/{org_handle}/conn | List org connections
*OrgConnections* | [**ListPermission**](docs/OrgConnections.md#listpermission) | **Get** /org/{org_handle}/connection/{connection_handle}/permission | List org connection permissions
*OrgConnections* | [**ListWorkspaces**](docs/OrgConnections.md#listworkspaces) | **Get** /org/{org_handle}/connection/{connection_handle}/workspace | List org connection workspaces
*OrgConnections* | [**ListWorkspacesDeprecated**](docs/OrgConnections.md#listworkspacesdeprecated) | **Get** /org/{org_handle}/conn/{conn_handle}/workspace | List org connection workspaces
*OrgConnections* | [**Test**](docs/OrgConnections.md#test) | **Post** /org/{org_handle}/connection/{connection_handle}/test | Test org connection
*OrgConnections* | [**TestDeprecated**](docs/OrgConnections.md#testdeprecated) | **Post** /org/{org_handle}/conn/{conn_handle}/test | Test org connection
*OrgConnections* | [**Update**](docs/OrgConnections.md#update) | **Patch** /org/{org_handle}/connection/{connection_handle} | Update org connection
*OrgConnections* | [**UpdateDeprecated**](docs/OrgConnections.md#updatedeprecated) | **Patch** /org/{org_handle}/conn/{conn_handle} | Update org connection
*OrgConnections* | [**UpdatePermission**](docs/OrgConnections.md#updatepermission) | **Patch** /org/{org_handle}/connection/{connection_handle}/permission/{permission_id} | Update org connection permission
*OrgIntegrations* | [**Command**](docs/OrgIntegrations.md#command) | **Post** /org/{org_handle}/integration/{integration_handle}/command | Run org integration command
*OrgIntegrations* | [**Create**](docs/OrgIntegrations.md#create) | **Post** /org/{org_handle}/integration | Create org integration
*OrgIntegrations* | [**Delete**](docs/OrgIntegrations.md#delete) | **Delete** /org/{org_handle}/integration/{integration_handle} | Delete org integration
*OrgIntegrations* | [**Get**](docs/OrgIntegrations.md#get) | **Get** /org/{org_handle}/integration/{integration_handle} | Get org integration
*OrgIntegrations* | [**InstallGithubIntegration**](docs/OrgIntegrations.md#installgithubintegration) | **Get** /org/{org_handle}/integration/{integration_handle}/github/install | Install GitHub integration for an org
*OrgIntegrations* | [**InstallSlackIntegration**](docs/OrgIntegrations.md#installslackintegration) | **Get** /org/{org_handle}/integration/{integration_handle}/slack/install | Install Slack integration for an org
*OrgIntegrations* | [**List**](docs/OrgIntegrations.md#list) | **Get** /org/{org_handle}/integration | List org integrations
*OrgIntegrations* | [**ListSlackChannels**](docs/OrgIntegrations.md#listslackchannels) | **Get** /org/{org_handle}/integration/{integration_handle}/slack/channel | List org integration Slack channels
*OrgIntegrations* | [**Test**](docs/OrgIntegrations.md#test) | **Post** /org/{org_handle}/integration/{integration_handle}/test | Test org integration
*OrgIntegrations* | [**Update**](docs/OrgIntegrations.md#update) | **Patch** /org/{org_handle}/integration/{integration_handle} | Update org integration
*OrgMembers* | [**ConfirmInvite**](docs/OrgMembers.md#confirminvite) | **Get** /org/{org_handle}/member/invite/confirm | Confirm org member invite
*OrgMembers* | [**Create**](docs/OrgMembers.md#create) | **Post** /org/{org_handle}/member | Create Org Member
*OrgMembers* | [**Delete**](docs/OrgMembers.md#delete) | **Delete** /org/{org_handle}/member/{user_handle} | Delete org member
*OrgMembers* | [**DeleteInvite**](docs/OrgMembers.md#deleteinvite) | **Delete** /org/{org_handle}/member/invite | Delete org member invite
*OrgMembers* | [**Get**](docs/OrgMembers.md#get) | **Get** /org/{org_handle}/member/{user_handle} | Get org member
*OrgMembers* | [**Invite**](docs/OrgMembers.md#invite) | **Post** /org/{org_handle}/member/invite | Invite org member
*OrgMembers* | [**List**](docs/OrgMembers.md#list) | **Get** /org/{org_handle}/member | List Organization Members
*OrgMembers* | [**Update**](docs/OrgMembers.md#update) | **Patch** /org/{org_handle}/member/{user_handle} | Update org member
*OrgProcesses* | [**Get**](docs/OrgProcesses.md#get) | **Get** /org/{org_handle}/process/{process_id} | Get Org process
*OrgProcesses* | [**List**](docs/OrgProcesses.md#list) | **Get** /org/{org_handle}/process | List Org processes
*OrgProcesses* | [**Log**](docs/OrgProcesses.md#log) | **Get** /org/{org_handle}/process/{process_id}/log/{log_file}.{content_type} | List Org process logs
*OrgWorkspaceAggregators* | [**Create**](docs/OrgWorkspaceAggregators.md#create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/aggregator | Create an aggregator for an org workspace
*OrgWorkspaceAggregators* | [**Delete**](docs/OrgWorkspaceAggregators.md#delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle} | Delete an aggregator for a org workspace
*OrgWorkspaceAggregators* | [**Get**](docs/OrgWorkspaceAggregators.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle} | Get an aggregator for a org workspace
*OrgWorkspaceAggregators* | [**GetConnection**](docs/OrgWorkspaceAggregators.md#getconnection) | **Get** /org/{org_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle}/connection/{connection_handle} | Get a connection in the scope of an aggregator for a org workspace
*OrgWorkspaceAggregators* | [**List**](docs/OrgWorkspaceAggregators.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/aggregator | List aggregators for an org workspace
*OrgWorkspaceAggregators* | [**ListConnections**](docs/OrgWorkspaceAggregators.md#listconnections) | **Get** /org/{org_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle}/connection | List connections in the scope of an aggregator for a org workspace
*OrgWorkspaceAggregators* | [**Update**](docs/OrgWorkspaceAggregators.md#update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle} | Update an aggregator for a org workspace
*OrgWorkspaceConnectionAssociations* | [**Create**](docs/OrgWorkspaceConnectionAssociations.md#create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/conn | Create org workspace connection association
*OrgWorkspaceConnectionAssociations* | [**Delete**](docs/OrgWorkspaceConnectionAssociations.md#delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/conn/{conn_handle} | Delete org workspace connection association
*OrgWorkspaceConnectionAssociations* | [**Get**](docs/OrgWorkspaceConnectionAssociations.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/conn/{conn_handle} | Get org workspace connection association
*OrgWorkspaceConnectionAssociations* | [**List**](docs/OrgWorkspaceConnectionAssociations.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/conn | List org workspace connection associations
*OrgWorkspaceConnectionFolders* | [**Create**](docs/OrgWorkspaceConnectionFolders.md#create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/connection_folder | Create org workspace connection folder
*OrgWorkspaceConnectionFolders* | [**Delete**](docs/OrgWorkspaceConnectionFolders.md#delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/connection_folder/{folder_id} | Delete org workspace connection folder
*OrgWorkspaceConnectionFolders* | [**Get**](docs/OrgWorkspaceConnectionFolders.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/connection_folder/{folder_id} | Get org workspace connection folder
*OrgWorkspaceConnectionFolders* | [**List**](docs/OrgWorkspaceConnectionFolders.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/connection_folder | List org workspace connection folders
*OrgWorkspaceConnectionFolders* | [**Update**](docs/OrgWorkspaceConnectionFolders.md#update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/connection_folder/{folder_id} | Update org workspace connection folder
*OrgWorkspaceConnectionTree* | [**List**](docs/OrgWorkspaceConnectionTree.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/connection_tree | List org workspace connection tree
*OrgWorkspaceConnections* | [**Create**](docs/OrgWorkspaceConnections.md#create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/connection | Create a connection for an org workspace or associate an org connection to the workspace
*OrgWorkspaceConnections* | [**Delete**](docs/OrgWorkspaceConnections.md#delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/connection/{connection_handle} | Delete org workspace connection association
*OrgWorkspaceConnections* | [**Get**](docs/OrgWorkspaceConnections.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/connection/{connection_handle} | Get org workspace connection association
*OrgWorkspaceConnections* | [**List**](docs/OrgWorkspaceConnections.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/connection | List org workspace connection associations
*OrgWorkspaceConnections* | [**Test**](docs/OrgWorkspaceConnections.md#test) | **Post** /org/{org_handle}/workspace/{workspace_handle}/connection/{connection_handle}/test | Test org workspace connection
*OrgWorkspaceConnections* | [**Update**](docs/OrgWorkspaceConnections.md#update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/connection/{connection_handle} | Update org workspace connection
*OrgWorkspaceDatatankParts* | [**Command**](docs/OrgWorkspaceDatatankParts.md#command) | **Post** /org/{org_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name}/part/{datatank_part_id}/command | Run org workspace Datatank table partition command
*OrgWorkspaceDatatankParts* | [**Get**](docs/OrgWorkspaceDatatankParts.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name}/table/{datatank_part_id} | Get org workspace Datatank table partition
*OrgWorkspaceDatatankParts* | [**List**](docs/OrgWorkspaceDatatankParts.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name}/part | List org workspace Datatank partitions
*OrgWorkspaceDatatankParts* | [**Update**](docs/OrgWorkspaceDatatankParts.md#update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name}/part/{datatank_part_id} | Update org workspace Datatank table partition
*OrgWorkspaceDatatankTables* | [**Create**](docs/OrgWorkspaceDatatankTables.md#create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table | Create org workspace Datatank table
*OrgWorkspaceDatatankTables* | [**Delete**](docs/OrgWorkspaceDatatankTables.md#delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name} | Delete org workspace Datatank table
*OrgWorkspaceDatatankTables* | [**Get**](docs/OrgWorkspaceDatatankTables.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name} | Get org workspace Datatank table
*OrgWorkspaceDatatankTables* | [**List**](docs/OrgWorkspaceDatatankTables.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table | List org workspace Datatank tables
*OrgWorkspaceDatatankTables* | [**Update**](docs/OrgWorkspaceDatatankTables.md#update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name} | Update org workspace Datatank table
*OrgWorkspaceDatatanks* | [**Create**](docs/OrgWorkspaceDatatanks.md#create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/datatank | Create org workspace Datatank
*OrgWorkspaceDatatanks* | [**Delete**](docs/OrgWorkspaceDatatanks.md#delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/datatank/{datatank_handle} | Delete org workspace Datatank
*OrgWorkspaceDatatanks* | [**Get**](docs/OrgWorkspaceDatatanks.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/datatank/{datatank_handle} | Get org workspace Datatank
*OrgWorkspaceDatatanks* | [**List**](docs/OrgWorkspaceDatatanks.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/datatank | List org workspace Datatank
*OrgWorkspaceDatatanks* | [**Update**](docs/OrgWorkspaceDatatanks.md#update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/datatank/{datatank_handle} | Update org workspace Datatank
*OrgWorkspaceFlowpipeInputs* | [**Create**](docs/OrgWorkspaceFlowpipeInputs.md#create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/input | Create org workspace flowpipe input
*OrgWorkspaceFlowpipeInputs* | [**Get**](docs/OrgWorkspaceFlowpipeInputs.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/form/{step_execution_id}/{random_id} | Get org workspace flowpipe form
*OrgWorkspaceFlowpipeInputs* | [**Get_0**](docs/OrgWorkspaceFlowpipeInputs.md#get_0) | **Get** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/input/{input_id} | Get org workspace flowpipe input
*OrgWorkspaceFlowpipeInputs* | [**List**](docs/OrgWorkspaceFlowpipeInputs.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/input | List org workspace flowpipe inputs
*OrgWorkspaceFlowpipeInputs* | [**Post**](docs/OrgWorkspaceFlowpipeInputs.md#post) | **Post** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/form/{step_execution_id}/{random_id}/submit | Post org workspace flowpipe form response
*OrgWorkspaceFlowpipeModVariables* | [**CreateSetting**](docs/OrgWorkspaceFlowpipeModVariables.md#createsetting) | **Post** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/mod/{mod_alias}/variable | Create a setting for a flowpipe mod variable in an organization workspace
*OrgWorkspaceFlowpipeModVariables* | [**DeleteSetting**](docs/OrgWorkspaceFlowpipeModVariables.md#deletesetting) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/mod/{mod_alias}/variable/{variable_name} | Delete setting for a flowpipe mod variable in an organization workspace
*OrgWorkspaceFlowpipeModVariables* | [**GetSetting**](docs/OrgWorkspaceFlowpipeModVariables.md#getsetting) | **Get** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/mod/{mod_alias}/variable/{variable_name} | Get setting for a flowpipe mod variable in an organization workspace
*OrgWorkspaceFlowpipeModVariables* | [**List**](docs/OrgWorkspaceFlowpipeModVariables.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/mod/{mod_alias}/variable | List variables in an organization flowpipe workspace mod
*OrgWorkspaceFlowpipeModVariables* | [**UpdateSetting**](docs/OrgWorkspaceFlowpipeModVariables.md#updatesetting) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/mod/{mod_alias}/variable/{variable_name} | Update setting for a flowpipe mod variable in an organization workspace
*OrgWorkspaceFlowpipeMods* | [**Get**](docs/OrgWorkspaceFlowpipeMods.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/mod/{mod_alias} | Get organization workspace installed flowpipe mod
*OrgWorkspaceFlowpipeMods* | [**Install**](docs/OrgWorkspaceFlowpipeMods.md#install) | **Post** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/mod | Install flowpipe mod to organization workspace.
*OrgWorkspaceFlowpipeMods* | [**Uninstall**](docs/OrgWorkspaceFlowpipeMods.md#uninstall) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/mod/{mod_alias} | Uninstall flowpipe mod from organization workspace.
*OrgWorkspaceFlowpipeMods* | [**Update**](docs/OrgWorkspaceFlowpipeMods.md#update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/mod/{mod_alias} | Update a flowpipe mod in an organization workspace
*OrgWorkspaceFlowpipePipelines* | [**Command**](docs/OrgWorkspaceFlowpipePipelines.md#command) | **Post** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/pipeline/{pipeline_name}/command | Run organization workspace Flowpipe pipeline command
*OrgWorkspaceFlowpipePipelines* | [**Get**](docs/OrgWorkspaceFlowpipePipelines.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/pipeline/{pipeline_id} | Get org workspace pipeline
*OrgWorkspaceFlowpipePipelines* | [**ListTriggers**](docs/OrgWorkspaceFlowpipePipelines.md#listtriggers) | **Get** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/pipeline/{pipeline_id}/trigger | Get org workspace pipeline triggers
*OrgWorkspaceFlowpipeTriggers* | [**Command**](docs/OrgWorkspaceFlowpipeTriggers.md#command) | **Post** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/trigger/{trigger_name}/command | Run organization workspace Flowpipe trigger command
*OrgWorkspaceFlowpipeTriggers* | [**Create**](docs/OrgWorkspaceFlowpipeTriggers.md#create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/trigger | Create org workspace trigger
*OrgWorkspaceFlowpipeTriggers* | [**Delete**](docs/OrgWorkspaceFlowpipeTriggers.md#delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/trigger/{trigger_name} | Delete org workspace pipeline
*OrgWorkspaceFlowpipeTriggers* | [**Get**](docs/OrgWorkspaceFlowpipeTriggers.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/trigger/{trigger_name} | Get org workspace flowpipe trigger
*OrgWorkspaceFlowpipeTriggers* | [**Update**](docs/OrgWorkspaceFlowpipeTriggers.md#update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/trigger/{trigger_name} | Update org workspace trigger
*OrgWorkspaceIntegrations* | [**Get**](docs/OrgWorkspaceIntegrations.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/integration/{integration_handle} | Get org workspace integration
*OrgWorkspaceIntegrations* | [**List**](docs/OrgWorkspaceIntegrations.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/integration | List org workspace integrations
*OrgWorkspaceIntegrations* | [**ListGithubRepositories**](docs/OrgWorkspaceIntegrations.md#listgithubrepositories) | **Get** /org/{org_handle}/workspace/{workspace_handle}/integration/{integration_handle}/github/repository | List org workspace integration github respositories
*OrgWorkspaceIntegrations* | [**ListGithubRepositoryHeads**](docs/OrgWorkspaceIntegrations.md#listgithubrepositoryheads) | **Get** /org/{org_handle}/workspace/{workspace_handle}/integration/{integration_handle}/github/repository/{repository_owner}/{repository_name}/head | List org workspace integration github respository heads
*OrgWorkspaceMembers* | [**Create**](docs/OrgWorkspaceMembers.md#create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/member | Create Org Workspace Member
*OrgWorkspaceMembers* | [**Delete**](docs/OrgWorkspaceMembers.md#delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/member/{user_handle} | Delete Org Workspace Member
*OrgWorkspaceMembers* | [**Get**](docs/OrgWorkspaceMembers.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/member/{user_handle} | Get Org Workspace Member
*OrgWorkspaceMembers* | [**List**](docs/OrgWorkspaceMembers.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/member | List Organization Workspace Members
*OrgWorkspaceMembers* | [**Update**](docs/OrgWorkspaceMembers.md#update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/member/{user_handle} | Update Org Workspace Member
*OrgWorkspaceModVariables* | [**CreateSetting**](docs/OrgWorkspaceModVariables.md#createsetting) | **Post** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable | Create a setting for a mod variable in an organization workspace
*OrgWorkspaceModVariables* | [**DeleteSetting**](docs/OrgWorkspaceModVariables.md#deletesetting) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable/{variable_name} | Delete setting for a mod variable in an organization workspace
*OrgWorkspaceModVariables* | [**GetSetting**](docs/OrgWorkspaceModVariables.md#getsetting) | **Get** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable/{variable_name} | Get setting for a mod variable in an organization workspace
*OrgWorkspaceModVariables* | [**List**](docs/OrgWorkspaceModVariables.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable | List variables in an organization workspace mod
*OrgWorkspaceModVariables* | [**UpdateSetting**](docs/OrgWorkspaceModVariables.md#updatesetting) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable/{variable_name} | Update setting for a mod variable in an organization workspace
*OrgWorkspaceMods* | [**Get**](docs/OrgWorkspaceMods.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias} | Get organization workspace installed mod
*OrgWorkspaceMods* | [**Install**](docs/OrgWorkspaceMods.md#install) | **Post** /org/{org_handle}/workspace/{workspace_handle}/mod | Install a mod to an organization workspace
*OrgWorkspaceMods* | [**List**](docs/OrgWorkspaceMods.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/mod | List organization workspace installed mods
*OrgWorkspaceMods* | [**Uninstall**](docs/OrgWorkspaceMods.md#uninstall) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias} | Uninstall mod from organization workspace.
*OrgWorkspaceMods* | [**Update**](docs/OrgWorkspaceMods.md#update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias} | Update a mod in an organization workspace
*OrgWorkspacePipelines* | [**Command**](docs/OrgWorkspacePipelines.md#command) | **Post** /org/{org_handle}/workspace/{workspace_handle}/pipeline/{pipeline_id}/command | Run org workspace pipeline command
*OrgWorkspacePipelines* | [**Create**](docs/OrgWorkspacePipelines.md#create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/pipeline | Create org workspace pipeline
*OrgWorkspacePipelines* | [**Delete**](docs/OrgWorkspacePipelines.md#delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/pipeline/{pipeline_id} | Delete org workspace pipeline
*OrgWorkspacePipelines* | [**Get**](docs/OrgWorkspacePipelines.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/pipeline/{pipeline_id} | Get org workspace pipeline
*OrgWorkspacePipelines* | [**List**](docs/OrgWorkspacePipelines.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/pipeline | List org workspace pipelines
*OrgWorkspacePipelines* | [**Update**](docs/OrgWorkspacePipelines.md#update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/pipeline/{pipeline_id} | Update org workspace pipeline
*OrgWorkspaceProcesses* | [**Get**](docs/OrgWorkspaceProcesses.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/process/{process_id} | Get org workspace process
*OrgWorkspaceProcesses* | [**List**](docs/OrgWorkspaceProcesses.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/process | List org workspace processes
*OrgWorkspaceProcesses* | [**Log**](docs/OrgWorkspaceProcesses.md#log) | **Get** /org/{org_handle}/workspace/{workspace_handle}/process/{process_id}/log/{log_file}.{content_type} | List org workspace process logs
*OrgWorkspaceSchemas* | [**Attach**](docs/OrgWorkspaceSchemas.md#attach) | **Post** /org/{org_handle}/workspace/{workspace_handle}/schema | Attach a schema to an org workspace
*OrgWorkspaceSchemas* | [**Detach**](docs/OrgWorkspaceSchemas.md#detach) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/schema/{schema_name} | Detach a schema from an org workspace
*OrgWorkspaceSchemas* | [**Get**](docs/OrgWorkspaceSchemas.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/schema/{schema_name} | Get org workspace schema
*OrgWorkspaceSchemas* | [**Get_0**](docs/OrgWorkspaceSchemas.md#get_0) | **Get** /org/{org_handle}/workspace/{workspace_handle}/schema/{schema_name}/table | List org workspace schema tables
*OrgWorkspaceSchemas* | [**List**](docs/OrgWorkspaceSchemas.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/schema | List org workspace schemas
*OrgWorkspaceSnapshots* | [**Create**](docs/OrgWorkspaceSnapshots.md#create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/snapshot | Create org workspace snapshot
*OrgWorkspaceSnapshots* | [**Delete**](docs/OrgWorkspaceSnapshots.md#delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id} | Delete org workspace snapshot
*OrgWorkspaceSnapshots* | [**Download**](docs/OrgWorkspaceSnapshots.md#download) | **Get** /org/{org_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id}.{content_type} | Download org workspace snapshot
*OrgWorkspaceSnapshots* | [**Get**](docs/OrgWorkspaceSnapshots.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id} | Get org workspace snapshot
*OrgWorkspaceSnapshots* | [**List**](docs/OrgWorkspaceSnapshots.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/snapshot | List org workspace snapshots
*OrgWorkspaceSnapshots* | [**Update**](docs/OrgWorkspaceSnapshots.md#update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id} | Update org workspace snapshot
*OrgWorkspaceUsages* | [**List**](docs/OrgWorkspaceUsages.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/usage | List org workspace usage
*OrgWorkspaces* | [**Command**](docs/OrgWorkspaces.md#command) | **Post** /org/{org_handle}/workspace/{workspace_handle}/command | Run org workspace command
*OrgWorkspaces* | [**Create**](docs/OrgWorkspaces.md#create) | **Post** /org/{org_handle}/workspace | Create org workspace
*OrgWorkspaces* | [**Delete**](docs/OrgWorkspaces.md#delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle} | Delete org workspace
*OrgWorkspaces* | [**Get**](docs/OrgWorkspaces.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle} | Get org workspace
*OrgWorkspaces* | [**GetQuery**](docs/OrgWorkspaces.md#getquery) | **Get** /org/{org_handle}/workspace/{workspace_handle}/query | Query org workspace
*OrgWorkspaces* | [**GetQueryWithExtensions**](docs/OrgWorkspaces.md#getquerywithextensions) | **Get** /org/{org_handle}/workspace/{workspace_handle}/query/data.{extensions} | Query org workspace with extensions
*OrgWorkspaces* | [**List**](docs/OrgWorkspaces.md#list) | **Get** /org/{org_handle}/workspace | List org workspaces
*OrgWorkspaces* | [**ListAuditLogs**](docs/OrgWorkspaces.md#listauditlogs) | **Get** /org/{org_handle}/workspace/{workspace_handle}/audit_log | Org workspace audit logs
*OrgWorkspaces* | [**ListDBLogs**](docs/OrgWorkspaces.md#listdblogs) | **Get** /org/{org_handle}/workspace/{workspace_handle}/db_log | Org workspace logs
*OrgWorkspaces* | [**PostQuery**](docs/OrgWorkspaces.md#postquery) | **Post** /org/{org_handle}/workspace/{workspace_handle}/query | Query org workspace
*OrgWorkspaces* | [**PostQueryWithExtensions**](docs/OrgWorkspaces.md#postquerywithextensions) | **Post** /org/{org_handle}/workspace/{workspace_handle}/query/data.{extensions} | Query org workspace with extensions
*OrgWorkspaces* | [**Update**](docs/OrgWorkspaces.md#update) | **Patch** /org/{org_handle}/workspace/{workspace_handle} | Update org workspace
*Orgs* | [**Create**](docs/Orgs.md#create) | **Post** /org | Create org
*Orgs* | [**CreateAvatar**](docs/Orgs.md#createavatar) | **Post** /org/{org_handle}/avatar | Create org avatar
*Orgs* | [**Delete**](docs/Orgs.md#delete) | **Delete** /org/{org_handle} | Delete org
*Orgs* | [**DeleteAvatar**](docs/Orgs.md#deleteavatar) | **Delete** /org/{org_handle}/avatar | Delete org avatar
*Orgs* | [**Get**](docs/Orgs.md#get) | **Get** /org/{org_handle} | Get org
*Orgs* | [**GetPlan**](docs/Orgs.md#getplan) | **Get** /org/{org_handle}/billing/plan | Get org billing plan.
*Orgs* | [**List**](docs/Orgs.md#list) | **Get** /org | List orgs
*Orgs* | [**ListAuditLogs**](docs/Orgs.md#listauditlogs) | **Get** /org/{org_handle}/audit_log | Org audit logs
*Orgs* | [**ListConstraints**](docs/Orgs.md#listconstraints) | **Get** /org/{org_handle}/constraint | List org constraints
*Orgs* | [**ListUsage**](docs/Orgs.md#listusage) | **Get** /org/{org_handle}/usage | List org usage
*Orgs* | [**Update**](docs/Orgs.md#update) | **Patch** /org/{org_handle} | Update org
*TenantBilling* | [**CreateBillingIntent**](docs/TenantBilling.md#createbillingintent) | **Post** /tenant/{tenant_handle}/billing/payment_method/intent | Create tenant Stripe billing intent
*TenantBilling* | [**CreateTenantSubscription**](docs/TenantBilling.md#createtenantsubscription) | **Post** /tenant/{tenant_handle}/billing/subscription | Create a new subscription
*TenantBilling* | [**DeleteTenantBillingPaymentMethod**](docs/TenantBilling.md#deletetenantbillingpaymentmethod) | **Delete** /tenant/{tenant_handle}/billing/payment_method/{payment_method_id} | Delete tenant billing payment method.
*TenantBilling* | [**DeleteTenantBillingSubscription**](docs/TenantBilling.md#deletetenantbillingsubscription) | **Delete** /tenant/{tenant_handle}/billing/subscription/{subscription_id} | Delete tenant subscription
*TenantBilling* | [**GetPlan**](docs/TenantBilling.md#getplan) | **Get** /tenant/{tenant_handle}/billing/plan | Get tenant billing plan.
*TenantBilling* | [**GetTenantBillingInvoice**](docs/TenantBilling.md#gettenantbillinginvoice) | **Get** /tenant/{tenant_handle}/billing/invoice/{invoice_id} | Get a tenant invoice
*TenantBilling* | [**GetTenantBillingPaymentMethod**](docs/TenantBilling.md#gettenantbillingpaymentmethod) | **Get** /tenant/{tenant_handle}/billing/payment_method/{payment_method_id} | Get tenant billing payment method.
*TenantBilling* | [**GetTenantBillingSubscription**](docs/TenantBilling.md#gettenantbillingsubscription) | **Get** /tenant/{tenant_handle}/billing/subscription/{subscription_id} | Get tenant subscription
*TenantBilling* | [**GetTenantBillingUpcomingInvoice**](docs/TenantBilling.md#gettenantbillingupcominginvoice) | **Get** /tenant/{tenant_handle}/billing/stripe/invoice/upcoming | Get upcoming invoice
*TenantBilling* | [**ListTenantBillingInvoices**](docs/TenantBilling.md#listtenantbillinginvoices) | **Get** /tenant/{tenant_handle}/billing/invoice | List tenant invoices
*TenantBilling* | [**ListTenantBillingPaymentMethod**](docs/TenantBilling.md#listtenantbillingpaymentmethod) | **Get** /tenant/{tenant_handle}/billing/payment_method | List tenant billing payment methods.
*TenantBilling* | [**ListTenantSubscription**](docs/TenantBilling.md#listtenantsubscription) | **Get** /tenant/{tenant_handle}/billing/subscription | List tenant subscriptions
*TenantBilling* | [**UpdateTenantBillingPaymentMethod**](docs/TenantBilling.md#updatetenantbillingpaymentmethod) | **Patch** /tenant/{tenant_handle}/billing/payment_method/{payment_method_id} | Update tenant billing payment method.
*TenantBilling* | [**UpdateTenantBillingSubscription**](docs/TenantBilling.md#updatetenantbillingsubscription) | **Patch** /tenant/{tenant_handle}/billing/subscription/{subscription_id} | Update tenant subscription
*TenantConnectionFolders* | [**Create**](docs/TenantConnectionFolders.md#create) | **Post** /connection_folder | Create tenant connection folder
*TenantConnectionFolders* | [**CreatePermission**](docs/TenantConnectionFolders.md#createpermission) | **Post** /connection_folder/{folder_id}/permission | Create tenant connection folder permission
*TenantConnectionFolders* | [**Delete**](docs/TenantConnectionFolders.md#delete) | **Delete** /connection_folder/{folder_id} | Delete tenant connection folder
*TenantConnectionFolders* | [**DeletePermission**](docs/TenantConnectionFolders.md#deletepermission) | **Delete** /connection_folder/{folder_id}/permission/{permission_id} | Delete tenant connection folder permission
*TenantConnectionFolders* | [**Get**](docs/TenantConnectionFolders.md#get) | **Get** /connection_folder/{folder_id} | Get tenant connection folder
*TenantConnectionFolders* | [**GetPermission**](docs/TenantConnectionFolders.md#getpermission) | **Get** /connection_folder/{folder_id}/permission/{permission_id} | Get tenant connection folder permission
*TenantConnectionFolders* | [**List**](docs/TenantConnectionFolders.md#list) | **Get** /connection_folder | List tenant connection folders
*TenantConnectionFolders* | [**ListPermission**](docs/TenantConnectionFolders.md#listpermission) | **Get** /connection_folder/{folder_id}/permission | List tenant connection folder permissions
*TenantConnectionFolders* | [**ListWorkspaces**](docs/TenantConnectionFolders.md#listworkspaces) | **Get** /connection_folder/{folder_id}/workspace | List tenant connection folder workspaces
*TenantConnectionFolders* | [**Update**](docs/TenantConnectionFolders.md#update) | **Patch** /connection_folder/{folder_id} | Update tenant connection folder
*TenantConnectionFolders* | [**UpdatePermission**](docs/TenantConnectionFolders.md#updatepermission) | **Patch** /connection_folder/{folder_id}/permission/{permission_id} | Update tenant connection folder permission
*TenantConnectionTree* | [**List**](docs/TenantConnectionTree.md#list) | **Get** /connection_tree | List tenant connection tree
*TenantConnections* | [**Create**](docs/TenantConnections.md#create) | **Post** /connection | Create tenant connection
*TenantConnections* | [**CreatePermission**](docs/TenantConnections.md#createpermission) | **Post** /connection/{connection_handle}/permission | Create tenant connection permission
*TenantConnections* | [**Delete**](docs/TenantConnections.md#delete) | **Delete** /connection/{connection_handle} | Delete tenant connection
*TenantConnections* | [**DeletePermission**](docs/TenantConnections.md#deletepermission) | **Delete** /connection/{connection_handle}/permission/{permission_id} | Delete tenant connection permission
*TenantConnections* | [**Get**](docs/TenantConnections.md#get) | **Get** /connection/{connection_handle} | Get tenant connection
*TenantConnections* | [**GetPermission**](docs/TenantConnections.md#getpermission) | **Get** /connection/{connection_handle}/permission/{permission_id} | Get tenant connection permission
*TenantConnections* | [**List**](docs/TenantConnections.md#list) | **Get** /connection | List tenant connections
*TenantConnections* | [**ListPermission**](docs/TenantConnections.md#listpermission) | **Get** /connection/{connection_handle}/permission | List tenant connection permissions
*TenantConnections* | [**ListWorkspaces**](docs/TenantConnections.md#listworkspaces) | **Get** /connection/{connection_handle}/workspace | List tenant connection workspaces
*TenantConnections* | [**Test**](docs/TenantConnections.md#test) | **Post** /connection/{connection_handle}/test | Test tenant connection
*TenantConnections* | [**Update**](docs/TenantConnections.md#update) | **Patch** /connection/{connection_handle} | Update tenant connection
*TenantConnections* | [**UpdatePermission**](docs/TenantConnections.md#updatepermission) | **Patch** /connection/{connection_handle}/permission/{permission_id} | Update tenant connection permission
*TenantIntegrations* | [**Command**](docs/TenantIntegrations.md#command) | **Post** /integration/{integration_handle}/command | Run tenant integration command
*TenantIntegrations* | [**Create**](docs/TenantIntegrations.md#create) | **Post** /integration | Create tenant integration
*TenantIntegrations* | [**Delete**](docs/TenantIntegrations.md#delete) | **Delete** /integration/{integration_handle} | Delete tenant integration
*TenantIntegrations* | [**Get**](docs/TenantIntegrations.md#get) | **Get** /integration/{integration_handle} | Get tenant integration
*TenantIntegrations* | [**InstallGithubIntegration**](docs/TenantIntegrations.md#installgithubintegration) | **Get** /integration/{integration_handle}/github/install | Install GitHub integration on a custom tenant
*TenantIntegrations* | [**InstallSlackIntegration**](docs/TenantIntegrations.md#installslackintegration) | **Get** /integration/{integration_handle}/slack/install | Install Slack integration on a custom tenant
*TenantIntegrations* | [**List**](docs/TenantIntegrations.md#list) | **Get** /integration | List tenant integrations
*TenantIntegrations* | [**ListSlackChannels**](docs/TenantIntegrations.md#listslackchannels) | **Get** /integration/{integration_handle}/slack/channel | List tenant integration Slack channels
*TenantIntegrations* | [**Test**](docs/TenantIntegrations.md#test) | **Post** /integration/{integration_handle}/test | Test custom tenant integration
*TenantIntegrations* | [**Update**](docs/TenantIntegrations.md#update) | **Patch** /integration/{integration_handle} | Update tenant integration
*TenantMembers* | [**ConfirmInvite**](docs/TenantMembers.md#confirminvite) | **Get** /tenant/{tenant_handle}/member/invite/confirm | Confirm tenant member invite
*TenantMembers* | [**Delete**](docs/TenantMembers.md#delete) | **Delete** /tenant/{tenant_handle}/member/{user_handle} | Delete tenant member
*TenantMembers* | [**DeleteInvite**](docs/TenantMembers.md#deleteinvite) | **Delete** /tenant/{tenant_handle}/member/invite | Delete tenant member invite
*TenantMembers* | [**Get**](docs/TenantMembers.md#get) | **Get** /tenant/{tenant_handle}/member/{user_handle} | Get tenant member
*TenantMembers* | [**Invite**](docs/TenantMembers.md#invite) | **Post** /tenant/{tenant_handle}/member/invite | Invite tenant member
*TenantMembers* | [**List**](docs/TenantMembers.md#list) | **Get** /tenant/{tenant_handle}/member | List Tenant Members
*TenantMembers* | [**Update**](docs/TenantMembers.md#update) | **Patch** /tenant/{tenant_handle}/member/{user_handle} | Update tenant member
*Tenants* | [**Create**](docs/Tenants.md#create) | **Post** /tenant | Create tenant
*Tenants* | [**CreateAvatar**](docs/Tenants.md#createavatar) | **Post** /tenant/{tenant_handle}/avatar | Create tenant avatar
*Tenants* | [**Delete**](docs/Tenants.md#delete) | **Delete** /tenant/{tenant_handle} | Delete tenant
*Tenants* | [**DeleteAvatar**](docs/Tenants.md#deleteavatar) | **Delete** /tenant/{tenant_handle}/avatar | Delete tenant avatar
*Tenants* | [**Get**](docs/Tenants.md#get) | **Get** /tenant/{tenant_handle} | Get tenant
*Tenants* | [**GetAvatar**](docs/Tenants.md#getavatar) | **Get** /tenant/{tenant_handle}/avatar | Get tenant avatar
*Tenants* | [**GetSettings**](docs/Tenants.md#getsettings) | **Get** /settings | Get tenant settings
*Tenants* | [**List**](docs/Tenants.md#list) | **Get** /tenant | List tenants
*Tenants* | [**ListAuditLogs**](docs/Tenants.md#listauditlogs) | **Get** /audit_log | Tenant audit logs
*Tenants* | [**ListConstraints**](docs/Tenants.md#listconstraints) | **Get** /tenant/{tenant_handle}/constraint | List tenant constraints
*Tenants* | [**ListUsage**](docs/Tenants.md#listusage) | **Get** /tenant/{tenant_handle}/usage | List tenant usage
*Tenants* | [**Update**](docs/Tenants.md#update) | **Patch** /tenant/{tenant_handle} | Update tenant
*Tenants* | [**UpdateSettings**](docs/Tenants.md#updatesettings) | **Patch** /settings | Update tenant settings
*UserBilling* | [**CreateBillingIntent**](docs/UserBilling.md#createbillingintent) | **Post** /user/{user_handle}/billing/payment_method/intent | Create Stripe billing intent
*UserBilling* | [**CreateUserSubscription**](docs/UserBilling.md#createusersubscription) | **Post** /user/{user_handle}/billing/subscription | Create a new subscription
*UserBilling* | [**DeleteUserBillingPaymentMethod**](docs/UserBilling.md#deleteuserbillingpaymentmethod) | **Delete** /user/{user_handle}/billing/payment_method/{payment_method_id} | Delete user billing payment method.
*UserBilling* | [**DeleteUserBillingSubscription**](docs/UserBilling.md#deleteuserbillingsubscription) | **Delete** /user/{user_handle}/billing/subscription/{subscription_id} | Delete user subscription
*UserBilling* | [**GetUserBillingInvoice**](docs/UserBilling.md#getuserbillinginvoice) | **Get** /user/{user_handle}/billing/invoice/{invoice_id} | Get a user invoice
*UserBilling* | [**GetUserBillingPaymentMethod**](docs/UserBilling.md#getuserbillingpaymentmethod) | **Get** /user/{user_handle}/billing/payment_method/{payment_method_id} | Get user billing payment method.
*UserBilling* | [**GetUserBillingSubscription**](docs/UserBilling.md#getuserbillingsubscription) | **Get** /user/{user_handle}/billing/subscription/{subscription_id} | Get user subscription
*UserBilling* | [**GetUserBillingUpcomingInvoice**](docs/UserBilling.md#getuserbillingupcominginvoice) | **Get** /user/{user_handle}/billing/stripe/invoice/upcoming | Get upcoming user invoice
*UserBilling* | [**ListUserBillingInvoices**](docs/UserBilling.md#listuserbillinginvoices) | **Get** /user/{user_handle}/billing/invoice | List user invoices
*UserBilling* | [**ListUserBillingPaymentMethod**](docs/UserBilling.md#listuserbillingpaymentmethod) | **Get** /user/{user_handle}/billing/payment_method | List user billing payment methods.
*UserBilling* | [**ListUserSubscription**](docs/UserBilling.md#listusersubscription) | **Get** /user/{user_handle}/billing/subscription | List user subscriptions
*UserBilling* | [**UpdateUserBillingPaymentMethod**](docs/UserBilling.md#updateuserbillingpaymentmethod) | **Patch** /user/{user_handle}/billing/payment_method/{payment_method_id} | Update user billing payment method.
*UserConnections* | [**Create**](docs/UserConnections.md#create) | **Post** /user/{user_handle}/connection | Create user connection
*UserConnections* | [**CreateDeprecated**](docs/UserConnections.md#createdeprecated) | **Post** /user/{user_handle}/conn | Create user connection
*UserConnections* | [**Delete**](docs/UserConnections.md#delete) | **Delete** /user/{user_handle}/connection/{connection_handle} | Delete user connection
*UserConnections* | [**DeleteDeprecated**](docs/UserConnections.md#deletedeprecated) | **Delete** /user/{user_handle}/conn/{conn_handle} | Delete user connection
*UserConnections* | [**Get**](docs/UserConnections.md#get) | **Get** /user/{user_handle}/connection/{connection_handle} | Get user connection
*UserConnections* | [**GetDeprecated**](docs/UserConnections.md#getdeprecated) | **Get** /user/{user_handle}/conn/{conn_handle} | Get user connection
*UserConnections* | [**List**](docs/UserConnections.md#list) | **Get** /user/{user_handle}/connection | List user connections
*UserConnections* | [**ListDeprecated**](docs/UserConnections.md#listdeprecated) | **Get** /user/{user_handle}/conn | List user connections
*UserConnections* | [**ListWorkspaces**](docs/UserConnections.md#listworkspaces) | **Get** /user/{user_handle}/connection/{connection_handle}/workspace | List user connection workspaces
*UserConnections* | [**ListWorkspacesDeprecated**](docs/UserConnections.md#listworkspacesdeprecated) | **Get** /user/{user_handle}/conn/{conn_handle}/workspace | List user connection workspaces
*UserConnections* | [**Test**](docs/UserConnections.md#test) | **Post** /user/{user_handle}/connection/{connection_handle}/test | Test user connection
*UserConnections* | [**TestDeprecated**](docs/UserConnections.md#testdeprecated) | **Post** /user/{user_handle}/conn/{conn_handle}/test | Test user connection
*UserConnections* | [**Update**](docs/UserConnections.md#update) | **Patch** /user/{user_handle}/connection/{connection_handle} | Update user connection
*UserConnections* | [**UpdateDeprecated**](docs/UserConnections.md#updatedeprecated) | **Patch** /user/{user_handle}/conn/{conn_handle} | Update user connection
*UserIntegrations* | [**Command**](docs/UserIntegrations.md#command) | **Post** /user/{user_handle}/integration/{integration_handle}/command | Run user integration command
*UserIntegrations* | [**Create**](docs/UserIntegrations.md#create) | **Post** /user/{user_handle}/integration | Create user integration
*UserIntegrations* | [**Delete**](docs/UserIntegrations.md#delete) | **Delete** /user/{user_handle}/integration/{integration_handle} | Delete user integration
*UserIntegrations* | [**Get**](docs/UserIntegrations.md#get) | **Get** /user/{user_handle}/integration/{integration_handle} | Get user integration
*UserIntegrations* | [**InstallGithubIntegration**](docs/UserIntegrations.md#installgithubintegration) | **Get** /user/{user_handle}/integration/{integration_handle}/github/install | Install GitHub integration for a user
*UserIntegrations* | [**InstallSlackIntegration**](docs/UserIntegrations.md#installslackintegration) | **Get** /user/{user_handle}/integration/{integration_handle}/slack/install | Install Slack integration for a user
*UserIntegrations* | [**List**](docs/UserIntegrations.md#list) | **Get** /user/{user_handle}/integration | List user integrations
*UserIntegrations* | [**ListSlackChannels**](docs/UserIntegrations.md#listslackchannels) | **Get** /user/{user_handle}/integration/{integration_handle}/slack/channel | List user integration Slack channels
*UserIntegrations* | [**Test**](docs/UserIntegrations.md#test) | **Post** /user/{user_handle}/integration/{integration_handle}/test | Test user integration
*UserIntegrations* | [**Update**](docs/UserIntegrations.md#update) | **Patch** /user/{user_handle}/integration/{integration_handle} | Update user integration
*UserProcesses* | [**Get**](docs/UserProcesses.md#get) | **Get** /user/{user_handle}/process/{process_id} | Get User process
*UserProcesses* | [**List**](docs/UserProcesses.md#list) | **Get** /user/{user_handle}/process | List User processes
*UserProcesses* | [**Log**](docs/UserProcesses.md#log) | **Get** /user/{user_handle}/process/{process_id}/log/{log_file}.{content_type} | List user process logs
*UserTokens* | [**Create**](docs/UserTokens.md#create) | **Post** /user/{user_handle}/token | Create token
*UserTokens* | [**Delete**](docs/UserTokens.md#delete) | **Delete** /user/{user_handle}/token/{token_id} | Delete token
*UserTokens* | [**Get**](docs/UserTokens.md#get) | **Get** /user/{user_handle}/token/{token_id} | Get token
*UserTokens* | [**List**](docs/UserTokens.md#list) | **Get** /user/{user_handle}/token | List tokens
*UserTokens* | [**Update**](docs/UserTokens.md#update) | **Patch** /user/{user_handle}/token/{token_id} | Update token
*UserWorkspaceAggregators* | [**Create**](docs/UserWorkspaceAggregators.md#create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/aggregator | Create an aggregator for a user workspace
*UserWorkspaceAggregators* | [**Delete**](docs/UserWorkspaceAggregators.md#delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle} | Delete an aggregator for a user workspace
*UserWorkspaceAggregators* | [**Get**](docs/UserWorkspaceAggregators.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle} | Get an aggregator for a user workspace
*UserWorkspaceAggregators* | [**GetConnection**](docs/UserWorkspaceAggregators.md#getconnection) | **Get** /user/{user_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle}/connection/{connection_handle} | Get a connection in the scope of an aggregator for a user workspace
*UserWorkspaceAggregators* | [**List**](docs/UserWorkspaceAggregators.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/aggregator | List aggregators for a user workspace
*UserWorkspaceAggregators* | [**ListConnections**](docs/UserWorkspaceAggregators.md#listconnections) | **Get** /user/{user_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle}/connection | List connections in the scope of an aggregator for a user workspace
*UserWorkspaceAggregators* | [**Update**](docs/UserWorkspaceAggregators.md#update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle} | Update an aggregator for a user workspace
*UserWorkspaceConnectionAssociations* | [**Create**](docs/UserWorkspaceConnectionAssociations.md#create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/conn | Create user workspace connection association
*UserWorkspaceConnectionAssociations* | [**Delete**](docs/UserWorkspaceConnectionAssociations.md#delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/conn/{conn_handle} | Delete user workspace connection association
*UserWorkspaceConnectionAssociations* | [**Get**](docs/UserWorkspaceConnectionAssociations.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/conn/{conn_handle} | Get user workspace connection association
*UserWorkspaceConnectionAssociations* | [**List**](docs/UserWorkspaceConnectionAssociations.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/conn | List user workspace connection associations
*UserWorkspaceConnectionFolders* | [**Create**](docs/UserWorkspaceConnectionFolders.md#create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/connection_folder | Create user workspace connection folder
*UserWorkspaceConnectionFolders* | [**Delete**](docs/UserWorkspaceConnectionFolders.md#delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/connection_folder/{folder_id} | Delete user workspace connection folder
*UserWorkspaceConnectionFolders* | [**Get**](docs/UserWorkspaceConnectionFolders.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/connection_folder/{folder_id} | Get user workspace connection folder
*UserWorkspaceConnectionFolders* | [**List**](docs/UserWorkspaceConnectionFolders.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/connection_folder | List user workspace connections folders
*UserWorkspaceConnectionFolders* | [**Update**](docs/UserWorkspaceConnectionFolders.md#update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/connection_folder/{folder_id} | Update user workspace connection folder
*UserWorkspaceConnectionTree* | [**List**](docs/UserWorkspaceConnectionTree.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/connection_tree | List user workspace connection tree
*UserWorkspaceConnections* | [**Create**](docs/UserWorkspaceConnections.md#create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/connection | Create a connection for a user workspace
*UserWorkspaceConnections* | [**Delete**](docs/UserWorkspaceConnections.md#delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/connection/{connection_handle} | Delete user workspace connection association
*UserWorkspaceConnections* | [**Get**](docs/UserWorkspaceConnections.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/connection/{connection_handle} | Get user workspace connection association
*UserWorkspaceConnections* | [**List**](docs/UserWorkspaceConnections.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/connection | List connections for a user workspace
*UserWorkspaceConnections* | [**Test**](docs/UserWorkspaceConnections.md#test) | **Post** /user/{user_handle}/workspace/{workspace_handle}/connection/{connection_handle}/test | Test user workspace connection
*UserWorkspaceConnections* | [**Update**](docs/UserWorkspaceConnections.md#update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/connection/{connection_handle} | Update user workspace connection
*UserWorkspaceDatatankParts* | [**Command**](docs/UserWorkspaceDatatankParts.md#command) | **Post** /user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name}/part/{datatank_part_id}/command | Run user workspace Datatank table partition command
*UserWorkspaceDatatankParts* | [**Get**](docs/UserWorkspaceDatatankParts.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name}/table/{datatank_part_id} | Get user workspace Datatank table partition
*UserWorkspaceDatatankParts* | [**List**](docs/UserWorkspaceDatatankParts.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name}/part | List user workspace Datatank partitions
*UserWorkspaceDatatankParts* | [**Update**](docs/UserWorkspaceDatatankParts.md#update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name}/part/{datatank_part_id} | Update user workspace Datatank table partition
*UserWorkspaceDatatankTables* | [**Create**](docs/UserWorkspaceDatatankTables.md#create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table | Create user workspace Datatank table
*UserWorkspaceDatatankTables* | [**Delete**](docs/UserWorkspaceDatatankTables.md#delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name} | Delete user workspace Datatank table
*UserWorkspaceDatatankTables* | [**Get**](docs/UserWorkspaceDatatankTables.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name} | Get user workspace Datatank table
*UserWorkspaceDatatankTables* | [**List**](docs/UserWorkspaceDatatankTables.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table | List user workspace Datatank tables
*UserWorkspaceDatatankTables* | [**Update**](docs/UserWorkspaceDatatankTables.md#update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name} | Update user workspace Datatank table
*UserWorkspaceDatatanks* | [**Create**](docs/UserWorkspaceDatatanks.md#create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/datatank | Create user workspace Datatank
*UserWorkspaceDatatanks* | [**Delete**](docs/UserWorkspaceDatatanks.md#delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle} | Delete user workspace Datatank table
*UserWorkspaceDatatanks* | [**Get**](docs/UserWorkspaceDatatanks.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle} | Get user workspace Datatank
*UserWorkspaceDatatanks* | [**List**](docs/UserWorkspaceDatatanks.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/datatank | List user workspace Datatank
*UserWorkspaceDatatanks* | [**Update**](docs/UserWorkspaceDatatanks.md#update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle} | Update user workspace Datatank
*UserWorkspaceFlowpipeInputs* | [**Create**](docs/UserWorkspaceFlowpipeInputs.md#create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/input | Create user workspace flowpipe input
*UserWorkspaceFlowpipeInputs* | [**Get**](docs/UserWorkspaceFlowpipeInputs.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/form/{step_execution_id}/{random_id} | Get user workspace flowpipe form
*UserWorkspaceFlowpipeInputs* | [**Get_0**](docs/UserWorkspaceFlowpipeInputs.md#get_0) | **Get** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/input/{input_id} | Get user workspace flowpipe input
*UserWorkspaceFlowpipeInputs* | [**List**](docs/UserWorkspaceFlowpipeInputs.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/input | List user workspace flowpipe inputs
*UserWorkspaceFlowpipeInputs* | [**Post**](docs/UserWorkspaceFlowpipeInputs.md#post) | **Post** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/form/{step_execution_id}/{random_id}/submit | Post user workspace flowpipe form response
*UserWorkspaceFlowpipeModVariables* | [**CreateSetting**](docs/UserWorkspaceFlowpipeModVariables.md#createsetting) | **Post** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/mod/{mod_alias}/variable | Create a setting for a flowpipe mod variable in a user workspace
*UserWorkspaceFlowpipeModVariables* | [**DeleteSetting**](docs/UserWorkspaceFlowpipeModVariables.md#deletesetting) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/mod/{mod_alias}/variable/{variable_name} | Delete setting for a flowpipe mod variable in a user workspace
*UserWorkspaceFlowpipeModVariables* | [**GetSetting**](docs/UserWorkspaceFlowpipeModVariables.md#getsetting) | **Get** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/mod/{mod_alias}/variable/{variable_name} | Get setting for a flowpipe mod variable in a user workspace
*UserWorkspaceFlowpipeModVariables* | [**List**](docs/UserWorkspaceFlowpipeModVariables.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/mod/{mod_alias}/variable | List variables for a user workspace flowpipe mod
*UserWorkspaceFlowpipeModVariables* | [**UpdateSetting**](docs/UserWorkspaceFlowpipeModVariables.md#updatesetting) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/mod/{mod_alias}/variable/{variable_name} | Update setting for a flowpipe mod variable in a user workspace
*UserWorkspaceFlowpipeMods* | [**Get**](docs/UserWorkspaceFlowpipeMods.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/mod/{mod_alias} | Get user workspace installed flowpipe mod
*UserWorkspaceFlowpipeMods* | [**Install**](docs/UserWorkspaceFlowpipeMods.md#install) | **Post** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/mod | Install flowpipe mod to a user&#39;s workspace
*UserWorkspaceFlowpipeMods* | [**Uninstall**](docs/UserWorkspaceFlowpipeMods.md#uninstall) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/mod/{mod_alias} | Uninstall flowpipe mod from a user&#39;s workspace.
*UserWorkspaceFlowpipeMods* | [**Update**](docs/UserWorkspaceFlowpipeMods.md#update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/mod/{mod_alias} | Update a flowpipe mod in a user&#39;s workspace
*UserWorkspaceFlowpipePipelines* | [**Command**](docs/UserWorkspaceFlowpipePipelines.md#command) | **Post** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/pipeline/{pipeline_name}/command | Run user workspace Flowpipe pipeline command
*UserWorkspaceFlowpipePipelines* | [**Get**](docs/UserWorkspaceFlowpipePipelines.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/pipeline/{pipeline_id} | Get user workspace pipeline
*UserWorkspaceFlowpipePipelines* | [**ListTriggers**](docs/UserWorkspaceFlowpipePipelines.md#listtriggers) | **Get** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/pipeline/{pipeline_id}/trigger | Get user workspace pipeline triggers
*UserWorkspaceFlowpipeTriggers* | [**Command**](docs/UserWorkspaceFlowpipeTriggers.md#command) | **Post** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/trigger/{trigger_name}/command | Run user workspace Flowpipe trigger command
*UserWorkspaceFlowpipeTriggers* | [**Create**](docs/UserWorkspaceFlowpipeTriggers.md#create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/trigger | Create user workspace trigger
*UserWorkspaceFlowpipeTriggers* | [**Delete**](docs/UserWorkspaceFlowpipeTriggers.md#delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/trigger/{trigger_name} | Delete user workspace pipeline
*UserWorkspaceFlowpipeTriggers* | [**Get**](docs/UserWorkspaceFlowpipeTriggers.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/trigger/{trigger_name} | Get user workspace flowpipe trigger
*UserWorkspaceFlowpipeTriggers* | [**Update**](docs/UserWorkspaceFlowpipeTriggers.md#update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/trigger/{trigger_name} | Update user workspace trigger
*UserWorkspaceIntegrations* | [**Get**](docs/UserWorkspaceIntegrations.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/integration/{integration_handle} | Get user workspace integration
*UserWorkspaceIntegrations* | [**List**](docs/UserWorkspaceIntegrations.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/integration | List user workspace integrations
*UserWorkspaceIntegrations* | [**ListGithubRepositories**](docs/UserWorkspaceIntegrations.md#listgithubrepositories) | **Get** /user/{user_handle}/workspace/{workspace_handle}/integration/{integration_handle}/github/repository | List user workspace integration github respositories
*UserWorkspaceIntegrations* | [**ListGithubRepositoryHeads**](docs/UserWorkspaceIntegrations.md#listgithubrepositoryheads) | **Get** /user/{user_handle}/workspace/{workspace_handle}/integration/{integration_handle}/github/repository/{repository_owner}/{repository_name}/head | List user workspace integration github respository heads
*UserWorkspaceModVariables* | [**CreateSetting**](docs/UserWorkspaceModVariables.md#createsetting) | **Post** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable | Create a setting for a mod variable in a user workspace
*UserWorkspaceModVariables* | [**DeleteSetting**](docs/UserWorkspaceModVariables.md#deletesetting) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable/{variable_name} | Delete setting for a mod variable in a user workspace
*UserWorkspaceModVariables* | [**GetSetting**](docs/UserWorkspaceModVariables.md#getsetting) | **Get** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable/{variable_name} | Get setting for a mod variable in a user workspace
*UserWorkspaceModVariables* | [**List**](docs/UserWorkspaceModVariables.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable | List variables for a user workspace mod
*UserWorkspaceModVariables* | [**UpdateSetting**](docs/UserWorkspaceModVariables.md#updatesetting) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable/{variable_name} | Update setting for a mod variable in a user workspace
*UserWorkspaceMods* | [**Get**](docs/UserWorkspaceMods.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias} | Get user workspace installed mod
*UserWorkspaceMods* | [**Install**](docs/UserWorkspaceMods.md#install) | **Post** /user/{user_handle}/workspace/{workspace_handle}/mod | Install a mod to a user&#39;s workspace
*UserWorkspaceMods* | [**List**](docs/UserWorkspaceMods.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/mod | List user workspace installed mods
*UserWorkspaceMods* | [**Uninstall**](docs/UserWorkspaceMods.md#uninstall) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias} | Uninstall mod from a user&#39;s workspace.
*UserWorkspaceMods* | [**Update**](docs/UserWorkspaceMods.md#update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias} | Update a mod in a user&#39;s workspace
*UserWorkspacePipelines* | [**Command**](docs/UserWorkspacePipelines.md#command) | **Post** /user/{user_handle}/workspace/{workspace_handle}/pipeline/{pipeline_id}/command | Run user workspace pipeline command
*UserWorkspacePipelines* | [**Create**](docs/UserWorkspacePipelines.md#create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/pipeline | Create user workspace pipeline
*UserWorkspacePipelines* | [**Delete**](docs/UserWorkspacePipelines.md#delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/pipeline/{pipeline_id} | Delete user workspace pipeline
*UserWorkspacePipelines* | [**Get**](docs/UserWorkspacePipelines.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/pipeline/{pipeline_id} | Get user workspace pipeline
*UserWorkspacePipelines* | [**List**](docs/UserWorkspacePipelines.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/pipeline | List user workspace pipelines
*UserWorkspacePipelines* | [**Update**](docs/UserWorkspacePipelines.md#update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/pipeline/{pipeline_id} | Update user workspace pipeline
*UserWorkspaceProcesses* | [**Get**](docs/UserWorkspaceProcesses.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/process/{process_id} | Get user workspace process
*UserWorkspaceProcesses* | [**List**](docs/UserWorkspaceProcesses.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/process | List user workspace processes
*UserWorkspaceProcesses* | [**Log**](docs/UserWorkspaceProcesses.md#log) | **Get** /user/{user_handle}/workspace/{workspace_handle}/process/{process_id}/log/{log_file}.{content_type} | List user workspace process logs
*UserWorkspaceSchemas* | [**Attach**](docs/UserWorkspaceSchemas.md#attach) | **Post** /user/{user_handle}/workspace/{workspace_handle}/schema | Attach a schema to a user workspace
*UserWorkspaceSchemas* | [**Detach**](docs/UserWorkspaceSchemas.md#detach) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/schema/{schema_name} | Detach a schema from a user workspace
*UserWorkspaceSchemas* | [**Get**](docs/UserWorkspaceSchemas.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/schema/{schema_name} | Get user workspace schema
*UserWorkspaceSchemas* | [**Get_0**](docs/UserWorkspaceSchemas.md#get_0) | **Get** /user/{user_handle}/workspace/{workspace_handle}/schema/{schema_name}/table | List user workspace schema tables
*UserWorkspaceSchemas* | [**List**](docs/UserWorkspaceSchemas.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/schema | List user workspace schemas
*UserWorkspaceSnapshots* | [**Create**](docs/UserWorkspaceSnapshots.md#create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/snapshot | Create user workspace snapshot
*UserWorkspaceSnapshots* | [**Delete**](docs/UserWorkspaceSnapshots.md#delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id} | Delete user workspace snapshot
*UserWorkspaceSnapshots* | [**Download**](docs/UserWorkspaceSnapshots.md#download) | **Get** /user/{user_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id}.{content_type} | Download user workspace snapshot
*UserWorkspaceSnapshots* | [**Get**](docs/UserWorkspaceSnapshots.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id} | Get user workspace snapshot
*UserWorkspaceSnapshots* | [**List**](docs/UserWorkspaceSnapshots.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/snapshot | List user workspace snapshots
*UserWorkspaceSnapshots* | [**Update**](docs/UserWorkspaceSnapshots.md#update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id} | Update user workspace snapshot
*UserWorkspaceUsages* | [**List**](docs/UserWorkspaceUsages.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/usage | List user workspace usage
*UserWorkspaces* | [**Command**](docs/UserWorkspaces.md#command) | **Post** /user/{user_handle}/workspace/{workspace_handle}/command | Run user workspace command
*UserWorkspaces* | [**Create**](docs/UserWorkspaces.md#create) | **Post** /user/{user_handle}/workspace | Create user workspace
*UserWorkspaces* | [**Delete**](docs/UserWorkspaces.md#delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle} | Delete user workspace
*UserWorkspaces* | [**Get**](docs/UserWorkspaces.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle} | Get user workspace
*UserWorkspaces* | [**GetQuery**](docs/UserWorkspaces.md#getquery) | **Get** /user/{user_handle}/workspace/{workspace_handle}/query | Query user workspace
*UserWorkspaces* | [**GetQueryWithExtensions**](docs/UserWorkspaces.md#getquerywithextensions) | **Get** /user/{user_handle}/workspace/{workspace_handle}/query/data.{extensions} | Query user workspace with extensions
*UserWorkspaces* | [**List**](docs/UserWorkspaces.md#list) | **Get** /user/{user_handle}/workspace | List user workspaces
*UserWorkspaces* | [**ListAuditLogs**](docs/UserWorkspaces.md#listauditlogs) | **Get** /user/{user_handle}/workspace/{workspace_handle}/audit_log | User workspace audit logs
*UserWorkspaces* | [**ListDBLogs**](docs/UserWorkspaces.md#listdblogs) | **Get** /user/{user_handle}/workspace/{workspace_handle}/db_log | User workspace logs
*UserWorkspaces* | [**PostQuery**](docs/UserWorkspaces.md#postquery) | **Post** /user/{user_handle}/workspace/{workspace_handle}/query | Query user workspace
*UserWorkspaces* | [**PostQueryWithExtensions**](docs/UserWorkspaces.md#postquerywithextensions) | **Post** /user/{user_handle}/workspace/{workspace_handle}/query/data.{extensions} | Query user workspace with extensions
*UserWorkspaces* | [**Update**](docs/UserWorkspaces.md#update) | **Patch** /user/{user_handle}/workspace/{workspace_handle} | Update user workspace
*Users* | [**Create**](docs/Users.md#create) | **Post** /user | Create user
*Users* | [**CreateAvatar**](docs/Users.md#createavatar) | **Post** /user/{user_handle}/avatar | Create user avatar
*Users* | [**CreateDBPassword**](docs/Users.md#createdbpassword) | **Post** /user/{user_handle}/password | Create user password
*Users* | [**Delete**](docs/Users.md#delete) | **Delete** /user/{user_handle} | Delete user
*Users* | [**DeleteAvatar**](docs/Users.md#deleteavatar) | **Delete** /user/{user_handle}/avatar | Delete user avatar
*Users* | [**Get**](docs/Users.md#get) | **Get** /user/{user_handle} | Get user
*Users* | [**GetDBPassword**](docs/Users.md#getdbpassword) | **Get** /user/{user_handle}/password | Get user password
*Users* | [**GetEmail**](docs/Users.md#getemail) | **Get** /user/{user_handle}/email/{email_id} | Get user email
*Users* | [**GetPlan**](docs/Users.md#getplan) | **Get** /user/{user_handle}/billing/plan | Get user billing plan.
*Users* | [**GetPreferences**](docs/Users.md#getpreferences) | **Get** /user/{user_handle}/preferences | Get user preferences
*Users* | [**List**](docs/Users.md#list) | **Get** /user | List users
*Users* | [**ListAuditLogs**](docs/Users.md#listauditlogs) | **Get** /user/{user_handle}/audit_log | User audit logs
*Users* | [**ListConstraints**](docs/Users.md#listconstraints) | **Get** /user/{user_handle}/constraint | List user constraints
*Users* | [**ListEmails**](docs/Users.md#listemails) | **Get** /user/{user_handle}/email | List user emails
*Users* | [**ListUsage**](docs/Users.md#listusage) | **Get** /user/{user_handle}/usage | List user usage
*Users* | [**Update**](docs/Users.md#update) | **Patch** /user/{user_handle} | Update user
*Users* | [**UpdatePreferences**](docs/Users.md#updatepreferences) | **Patch** /user/{user_handle}/preferences | Update user preferences


## Documentation For Models

 - [ActorWorkspace](docs/ActorWorkspace.md)
 - [Aggregator](docs/Aggregator.md)
 - [AttachWorkspaceSchemaRequest](docs/AttachWorkspaceSchemaRequest.md)
 - [AuditRecord](docs/AuditRecord.md)
 - [BillingInfo](docs/BillingInfo.md)
 - [BillingInvoice](docs/BillingInvoice.md)
 - [BillingSetupIntent](docs/BillingSetupIntent.md)
 - [BillingSubscription](docs/BillingSubscription.md)
 - [BillingSubscriptionItem](docs/BillingSubscriptionItem.md)
 - [Card](docs/Card.md)
 - [CmdDatatankPartRequest](docs/CmdDatatankPartRequest.md)
 - [Connection](docs/Connection.md)
 - [ConnectionConfigSource](docs/ConnectionConfigSource.md)
 - [ConnectionCredentialSource](docs/ConnectionCredentialSource.md)
 - [ConnectionHandleMode](docs/ConnectionHandleMode.md)
 - [ConnectionTestResult](docs/ConnectionTestResult.md)
 - [ConnectionTrunkItem](docs/ConnectionTrunkItem.md)
 - [Constraint](docs/Constraint.md)
 - [ConstraintOverrideRequest](docs/ConstraintOverrideRequest.md)
 - [ConstraintVisibility](docs/ConstraintVisibility.md)
 - [CreateAggregatorRequest](docs/CreateAggregatorRequest.md)
 - [CreateAvatarResponse](docs/CreateAvatarResponse.md)
 - [CreateConnectionFolderRequest](docs/CreateConnectionFolderRequest.md)
 - [CreateConnectionRequest](docs/CreateConnectionRequest.md)
 - [CreateDatatankRequest](docs/CreateDatatankRequest.md)
 - [CreateDatatankTableRequest](docs/CreateDatatankTableRequest.md)
 - [CreateFlowpipeInputRequest](docs/CreateFlowpipeInputRequest.md)
 - [CreateIntegrationRequest](docs/CreateIntegrationRequest.md)
 - [CreateOrgRequest](docs/CreateOrgRequest.md)
 - [CreateOrgUserRequest](docs/CreateOrgUserRequest.md)
 - [CreateOrgWorkspaceUserRequest](docs/CreateOrgWorkspaceUserRequest.md)
 - [CreatePermissionRequest](docs/CreatePermissionRequest.md)
 - [CreatePipelineRequest](docs/CreatePipelineRequest.md)
 - [CreateTenantRequest](docs/CreateTenantRequest.md)
 - [CreateTriggerRequest](docs/CreateTriggerRequest.md)
 - [CreateUserPasswordRequest](docs/CreateUserPasswordRequest.md)
 - [CreateUserRequest](docs/CreateUserRequest.md)
 - [CreateWorkspaceConnRequest](docs/CreateWorkspaceConnRequest.md)
 - [CreateWorkspaceModRequest](docs/CreateWorkspaceModRequest.md)
 - [CreateWorkspaceModVariableSettingRequest](docs/CreateWorkspaceModVariableSettingRequest.md)
 - [CreateWorkspaceRequest](docs/CreateWorkspaceRequest.md)
 - [CreateWorkspaceSnapshotRequest](docs/CreateWorkspaceSnapshotRequest.md)
 - [Datatank](docs/Datatank.md)
 - [DatatankPart](docs/DatatankPart.md)
 - [DatatankPartCmd](docs/DatatankPartCmd.md)
 - [DatatankPartFreshnessState](docs/DatatankPartFreshnessState.md)
 - [DatatankPartState](docs/DatatankPartState.md)
 - [DatatankState](docs/DatatankState.md)
 - [DatatankTable](docs/DatatankTable.md)
 - [DatatankTableFreshness](docs/DatatankTableFreshness.md)
 - [DatatankTableState](docs/DatatankTableState.md)
 - [DeleteAvatarResponse](docs/DeleteAvatarResponse.md)
 - [DesiredState](docs/DesiredState.md)
 - [ErrorDetailModel](docs/ErrorDetailModel.md)
 - [ErrorModel](docs/ErrorModel.md)
 - [FlowpipeInput](docs/FlowpipeInput.md)
 - [FlowpipeInputState](docs/FlowpipeInputState.md)
 - [FlowpipeInputStepType](docs/FlowpipeInputStepType.md)
 - [GithubRepository](docs/GithubRepository.md)
 - [GithubRepositoryBranch](docs/GithubRepositoryBranch.md)
 - [Identity](docs/Identity.md)
 - [IdentityPlan](docs/IdentityPlan.md)
 - [IdentityUsageThresholdAction](docs/IdentityUsageThresholdAction.md)
 - [Integration](docs/Integration.md)
 - [IntegrationCommandAction](docs/IntegrationCommandAction.md)
 - [IntegrationCommandRequest](docs/IntegrationCommandRequest.md)
 - [IntegrationState](docs/IntegrationState.md)
 - [IntegrationType](docs/IntegrationType.md)
 - [InviteOrgUserRequest](docs/InviteOrgUserRequest.md)
 - [InviteTenantUserRequest](docs/InviteTenantUserRequest.md)
 - [ListActorWorkspacesResponse](docs/ListActorWorkspacesResponse.md)
 - [ListAuditLogsResponse](docs/ListAuditLogsResponse.md)
 - [ListBillingInvoicesResponse](docs/ListBillingInvoicesResponse.md)
 - [ListBillingPaymentMethodsResponse](docs/ListBillingPaymentMethodsResponse.md)
 - [ListConnectionsResponse](docs/ListConnectionsResponse.md)
 - [ListConstraintsResponse](docs/ListConstraintsResponse.md)
 - [ListDatatankPartResponse](docs/ListDatatankPartResponse.md)
 - [ListDatatankResponse](docs/ListDatatankResponse.md)
 - [ListDatatankTableResponse](docs/ListDatatankTableResponse.md)
 - [ListFlowpipeInputsResponse](docs/ListFlowpipeInputsResponse.md)
 - [ListFlowpipeTriggersResponse](docs/ListFlowpipeTriggersResponse.md)
 - [ListGithubRepositoriesResponse](docs/ListGithubRepositoriesResponse.md)
 - [ListGithubRepositoryBranchesResponse](docs/ListGithubRepositoryBranchesResponse.md)
 - [ListIdentitiesResponse](docs/ListIdentitiesResponse.md)
 - [ListIntegrationsResponse](docs/ListIntegrationsResponse.md)
 - [ListLogsResponse](docs/ListLogsResponse.md)
 - [ListOrgUsersResponse](docs/ListOrgUsersResponse.md)
 - [ListOrgWorkspaceUsersResponse](docs/ListOrgWorkspaceUsersResponse.md)
 - [ListOrgsResponse](docs/ListOrgsResponse.md)
 - [ListPermissionsResponse](docs/ListPermissionsResponse.md)
 - [ListPipelinesResponse](docs/ListPipelinesResponse.md)
 - [ListProcessesResponse](docs/ListProcessesResponse.md)
 - [ListSlackChannelsResponse](docs/ListSlackChannelsResponse.md)
 - [ListSubscriptionsResponse](docs/ListSubscriptionsResponse.md)
 - [ListTenantUsersResponse](docs/ListTenantUsersResponse.md)
 - [ListTenantsResponse](docs/ListTenantsResponse.md)
 - [ListTokensResponse](docs/ListTokensResponse.md)
 - [ListUsageMetricsResponse](docs/ListUsageMetricsResponse.md)
 - [ListUserEmailsResponse](docs/ListUserEmailsResponse.md)
 - [ListUserOrgsResponse](docs/ListUserOrgsResponse.md)
 - [ListUsersResponse](docs/ListUsersResponse.md)
 - [ListWorkspaceAggregatorsResponse](docs/ListWorkspaceAggregatorsResponse.md)
 - [ListWorkspaceConnResponse](docs/ListWorkspaceConnResponse.md)
 - [ListWorkspaceConnectionAssociationsResponse](docs/ListWorkspaceConnectionAssociationsResponse.md)
 - [ListWorkspaceConnectionsResponse](docs/ListWorkspaceConnectionsResponse.md)
 - [ListWorkspaceModVariablesResponse](docs/ListWorkspaceModVariablesResponse.md)
 - [ListWorkspaceModsResponse](docs/ListWorkspaceModsResponse.md)
 - [ListWorkspaceSchemaResponse](docs/ListWorkspaceSchemaResponse.md)
 - [ListWorkspaceSchemaTableResponse](docs/ListWorkspaceSchemaTableResponse.md)
 - [ListWorkspaceSnapshotsResponse](docs/ListWorkspaceSnapshotsResponse.md)
 - [ListWorkspacesResponse](docs/ListWorkspacesResponse.md)
 - [LogRecord](docs/LogRecord.md)
 - [ModPipeling](docs/ModPipeling.md)
 - [Notifier](docs/Notifier.md)
 - [NotifierPrecedence](docs/NotifierPrecedence.md)
 - [NotifierState](docs/NotifierState.md)
 - [Org](docs/Org.md)
 - [OrgUser](docs/OrgUser.md)
 - [OrgWorkspaceUser](docs/OrgWorkspaceUser.md)
 - [PaymentMethod](docs/PaymentMethod.md)
 - [Permission](docs/Permission.md)
 - [Pipeline](docs/Pipeline.md)
 - [PipelineCommandAction](docs/PipelineCommandAction.md)
 - [PipelineCommandRequest](docs/PipelineCommandRequest.md)
 - [PipelineCommandResponse](docs/PipelineCommandResponse.md)
 - [PipelineFrequency](docs/PipelineFrequency.md)
 - [PipelineInstanceType](docs/PipelineInstanceType.md)
 - [PipelineState](docs/PipelineState.md)
 - [Plan](docs/Plan.md)
 - [ProcessState](docs/ProcessState.md)
 - [SlackChannel](docs/SlackChannel.md)
 - [SnapshotState](docs/SnapshotState.md)
 - [SnapshotVisibility](docs/SnapshotVisibility.md)
 - [SpProcess](docs/SpProcess.md)
 - [TemporaryTokenRequest](docs/TemporaryTokenRequest.md)
 - [TemporaryTokenRequestState](docs/TemporaryTokenRequestState.md)
 - [Tenant](docs/Tenant.md)
 - [TenantLoginSettings](docs/TenantLoginSettings.md)
 - [TenantPlan](docs/TenantPlan.md)
 - [TenantSamlLoginSettings](docs/TenantSamlLoginSettings.md)
 - [TenantSettings](docs/TenantSettings.md)
 - [TenantUser](docs/TenantUser.md)
 - [TestConnectionRequest](docs/TestConnectionRequest.md)
 - [TestIntegrationRequest](docs/TestIntegrationRequest.md)
 - [Token](docs/Token.md)
 - [TriggerCommandAction](docs/TriggerCommandAction.md)
 - [TriggerCommandRequest](docs/TriggerCommandRequest.md)
 - [TriggerCommandResponse](docs/TriggerCommandResponse.md)
 - [TriggerState](docs/TriggerState.md)
 - [USBankAccount](docs/USBankAccount.md)
 - [UpdateAggregatorRequest](docs/UpdateAggregatorRequest.md)
 - [UpdateBillingPaymentMethodRequest](docs/UpdateBillingPaymentMethodRequest.md)
 - [UpdateBillingSubscriptionRequest](docs/UpdateBillingSubscriptionRequest.md)
 - [UpdateConnectionFolderRequest](docs/UpdateConnectionFolderRequest.md)
 - [UpdateConnectionRequest](docs/UpdateConnectionRequest.md)
 - [UpdateDatatankPartRequest](docs/UpdateDatatankPartRequest.md)
 - [UpdateDatatankRequest](docs/UpdateDatatankRequest.md)
 - [UpdateDatatankTableRequest](docs/UpdateDatatankTableRequest.md)
 - [UpdateIntegrationRequest](docs/UpdateIntegrationRequest.md)
 - [UpdateOrgRequest](docs/UpdateOrgRequest.md)
 - [UpdateOrgUserRequest](docs/UpdateOrgUserRequest.md)
 - [UpdateOrgWorkspaceUserRequest](docs/UpdateOrgWorkspaceUserRequest.md)
 - [UpdatePermissionRequest](docs/UpdatePermissionRequest.md)
 - [UpdatePipelineRequest](docs/UpdatePipelineRequest.md)
 - [UpdateTemporaryTokenRequest](docs/UpdateTemporaryTokenRequest.md)
 - [UpdateTenantLoginSettings](docs/UpdateTenantLoginSettings.md)
 - [UpdateTenantRequest](docs/UpdateTenantRequest.md)
 - [UpdateTenantSamlLoginSettings](docs/UpdateTenantSamlLoginSettings.md)
 - [UpdateTenantSettingsRequest](docs/UpdateTenantSettingsRequest.md)
 - [UpdateTenantUserRequest](docs/UpdateTenantUserRequest.md)
 - [UpdateTokenRequest](docs/UpdateTokenRequest.md)
 - [UpdateTriggerRequest](docs/UpdateTriggerRequest.md)
 - [UpdateUserPreferencesRequest](docs/UpdateUserPreferencesRequest.md)
 - [UpdateUserRequest](docs/UpdateUserRequest.md)
 - [UpdateWorkspaceModRequest](docs/UpdateWorkspaceModRequest.md)
 - [UpdateWorkspaceModVariableSettingRequest](docs/UpdateWorkspaceModVariableSettingRequest.md)
 - [UpdateWorkspaceRequest](docs/UpdateWorkspaceRequest.md)
 - [UpdateWorkspaceSnapshotRequest](docs/UpdateWorkspaceSnapshotRequest.md)
 - [UsageDimensionType](docs/UsageDimensionType.md)
 - [UsageMetric](docs/UsageMetric.md)
 - [UsageMetricType](docs/UsageMetricType.md)
 - [UsageUnitType](docs/UsageUnitType.md)
 - [User](docs/User.md)
 - [UserDatabasePassword](docs/UserDatabasePassword.md)
 - [UserEmail](docs/UserEmail.md)
 - [UserLoginRequest](docs/UserLoginRequest.md)
 - [UserOrg](docs/UserOrg.md)
 - [UserPreferences](docs/UserPreferences.md)
 - [UserSignupRequest](docs/UserSignupRequest.md)
 - [Workspace](docs/Workspace.md)
 - [WorkspaceAggregator](docs/WorkspaceAggregator.md)
 - [WorkspaceAggregatorAssociation](docs/WorkspaceAggregatorAssociation.md)
 - [WorkspaceCommandAction](docs/WorkspaceCommandAction.md)
 - [WorkspaceCommandRequest](docs/WorkspaceCommandRequest.md)
 - [WorkspaceCommandResponse](docs/WorkspaceCommandResponse.md)
 - [WorkspaceConn](docs/WorkspaceConn.md)
 - [WorkspaceConnection](docs/WorkspaceConnection.md)
 - [WorkspaceConnectionAssociation](docs/WorkspaceConnectionAssociation.md)
 - [WorkspaceInstanceType](docs/WorkspaceInstanceType.md)
 - [WorkspaceMod](docs/WorkspaceMod.md)
 - [WorkspaceModState](docs/WorkspaceModState.md)
 - [WorkspaceModTrigger](docs/WorkspaceModTrigger.md)
 - [WorkspaceModVariable](docs/WorkspaceModVariable.md)
 - [WorkspaceQueryResult](docs/WorkspaceQueryResult.md)
 - [WorkspaceQueryResultColumn](docs/WorkspaceQueryResultColumn.md)
 - [WorkspaceSchema](docs/WorkspaceSchema.md)
 - [WorkspaceSchemaState](docs/WorkspaceSchemaState.md)
 - [WorkspaceSchemaTable](docs/WorkspaceSchemaTable.md)
 - [WorkspaceSchemaTableColumn](docs/WorkspaceSchemaTableColumn.md)
 - [WorkspaceSnapshot](docs/WorkspaceSnapshot.md)
 - [WorkspaceSnapshotData](docs/WorkspaceSnapshotData.md)
 - [WorkspaceSnapshotDataLayout](docs/WorkspaceSnapshotDataLayout.md)
 - [WorkspaceState](docs/WorkspaceState.md)


## Documentation For Authorization



### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



