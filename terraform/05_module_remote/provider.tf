terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"
}

provider "aws" {
  profile = "terraform"
  region  = "ap-northeast-1"
}

terraform {
  backend "s3" {
    # need to create bucket first
    bucket = "terraform-remote-state-keidarcy"
    key    = "terraform.tfstate"
    region = "ap-northeast-1"
    # need to create dynamodb table first
    # https://developer.hashicorp.com/terraform/language/settings/backends/s3#dynamodb-state-locking
    # dynamodb_table = "terraform-remote-state-lock"
  }
}


# terraform {
#   backend "consul" {
#     address = "127.0.0.1:8500"
#     scheme = "http"
#     path = "terraform.tfstate"
#   }
# }
