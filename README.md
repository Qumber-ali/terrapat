# Terrapat for Azure DevOps
![Qamber](https://miro.medium.com/max/600/1*wHyWnltgrOclEbAFI0-fHw.png)

![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)

This repository contains terraform code along with golang package that provides integration for azure devops pat token access from inisde terraform code. Terraform's **"external"** datasource is used to run external golang executable that contains the actual code for getting azure devops pat token. The token is then returned following external datasource's data exchange protocol between terraform and external program(Go).

Clone this repository by running:
```sh
git clone https://github.com/Qumber-ali/terrapat.git
```
After cloning navigate to scripts directory and build the golang executable by running:
```golang
go get -d ./... && go build pat.go
```
After building the executable you are ready to run the terraform code.

Initializa the terraform directory by running:
```bash
terraform init
```
And finally apply the terraform configuration as:
```bash
terraform apply
```

Note that you can integrate this complete code(terraform with external datasource and golang executable) into your terraform configuration to serve your scenario that requires azure devops pat token access inside your terraform configuration. 
## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_azurerm"></a> [azurerm](#provider\_azurerm) | n/a |
| <a name="provider_external"></a> [external](#provider\_external) | n/a |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [azurerm_client_config.current](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/data-sources/client_config) | data source |
| [external_external.pat_token](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

No inputs.

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_rendered_result"></a> [rendered\_result](#output\_rendered\_result) | n/a |

