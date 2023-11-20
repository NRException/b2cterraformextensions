# b2cterraformextensions
An extension terraform provider to provide functionality not currently supported within the AzureRM or AzureAD Terraform providers.

## The aim
The aim of this project is to extend the functionality provided by the likes of AzureRM or AzureAD terraform providers
to leverage some newly published Graph API endpoints for the following services within Azure AD B2C:

- Authentication Policy Management (trust framework policies)
- Policy Keys (trust framework policy keys)

The scope of this provider may change in the coming months, if and when Microsoft decide to make this functionality main-stream and to publish it to the official AzureRM / AzureAD providers.
The project may become archived if this functionality is brought to the mainstream providers and this project becomes invalid.

Please note: This project is still under development, once the project is at feature parity with the intended features, a relase will be made, and proper documentation will be generated.

## How to install:

On Windows:

- Download the latest release (When available)
- Copy the .exe file to %APPDATA%/terraform.d/plugins or %APPDATA%/HashiCorp/Terraform/plugins
- Source the provider in your configuration.

On Linux:

- Download the latest release (When available)
- Copy the linux binary to $HOME/.terraform.d/plugins or terraform/plugins
- Source the provider in your configuration

## How to compile from source:

- Ensure go1.24.3 (the current build target) is installed on your machine.
- Run "make all" from the source project folder.

## How to contribute:

- Submit your PR's for review :)

## Some notes:

This project uses the Graph API to call endpoints that target your B2C tenant. This is done by using Service Principal authentication, Otherwise known as App Registrations within Azure AD / Microsoft Entra / B2C.
You MUST have a valid client ID and secret when attempting to use these providers, for more information please refer to the following docs for more info.

[https://learn.microsoft.com/en-us/entra/identity-platform/quickstart-register-app#register-an-application](https://learn.microsoft.com/en-us/entra/identity-platform/quickstart-register-app#add-a-redirect-uri)
[https://learn.microsoft.com/en-us/entra/identity-platform/quickstart-register-app#add-a-redirect-uri](https://learn.microsoft.com/en-us/entra/identity-platform/quickstart-register-app#add-credentials)

Please note that redirect URI's and scopes are not required, however the app registration will need permission to use the Microsoft Graph API, see more here:
https://learn.microsoft.com/en-us/graph/migrate-azure-ad-graph-configure-permissions?tabs=http%2Cupdatepermissions-azureadgraph-powershell#option-1-use-the-microsoft-entra-admin-center-to-find-the-apis-your-organization-uses
