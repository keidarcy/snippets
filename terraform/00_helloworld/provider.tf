terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
    http = {
      source  = "hashicorp/http"
      version = "~> 3.1"
    }
  }

  required_version = ">= 1.2.0"
}

provider "aws" {
  profile = "terraform"
  region  = "ap-northeast-1"
}

