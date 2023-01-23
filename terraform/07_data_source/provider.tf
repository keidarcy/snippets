terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
    http = {
      source  = "hashicorp/http"
      version = "3.2.1"
    }
    template = {
      source  = "hashicorp/template"
      version = "2.2.0"
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
    bucket = "terraform-remote-state-xyh-20230114"
    key    = "terraform.tfstate"
    region = "ap-northeast-1"
  }
}
