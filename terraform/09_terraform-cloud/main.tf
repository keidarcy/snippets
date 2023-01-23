

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"
}

terraform {
  cloud {
    organization = "xiaopeng163"

    workspaces {
      name = "terraform-cli-demo"
    }
  }
}


provider "aws" {
  region = var.aws_region
}

resource "aws_vpc" "vpc" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  tags = {
    Name = "my-vpc"
  }
}

variable "aws_region" {
  type    = string
  default = "eu-central-1"
}
