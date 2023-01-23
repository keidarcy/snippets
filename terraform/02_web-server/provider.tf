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
    bucket = "terraform-remote-state-xyh-20230114"
    key    = "terraform.tfstate"
    region = "ap-northeast-1"
    # need to create dynamodb table first
    # https://developer.hashicorp.com/terraform/language/settings/backends/s3#dynamodb-state-locking
    /* dynamodb_table = "terraform-remote-state-lock" */
  }
}
