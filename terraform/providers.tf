terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "3.58.0"
    }
  }

  backend "azurerm" {
    resource_group_name  = "terraform_backends"
    storage_account_name = "msprarosaje"
    container_name       = "tfstate"
    key                  = "backend.terraform.tfstate"
  }
}

provider "azurerm" {
  # Configuration options$
  features {

  }

  client_id       = "90e09199-b541-45d0-a415-a772edf9a745"
  tenant_id       = "5369cc18-884b-45b0-80a5-7b66171f60cf"
  subscription_id = var.subscription_id
  client_secret   = var.client_secret
}
