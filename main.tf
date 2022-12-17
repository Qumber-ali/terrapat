data "azurerm_client_config" "current" {}

data "external" "pat_token" {
  program = ["${path.module}/scripts/pat"]
  query = {
    tenant = data.azurerm_client_config.current.tenant_id
  }
}


output "rendered_result" {
 value = data.external.pat_token.result.output
}
