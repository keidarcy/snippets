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
  /* access_key = var.aws_access_key */
  /* secret_key = var.aws_secret_key */
  region = var.aws_regions[2]
}


locals {
  common_tags = {
    Company = "MyCompany"
    Owner   = "test"
  }
}
resource "aws_vpc" "vpc" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = var.enable_dns_hostnames
  /* tags = { */
  /*   Name = "my-vpc-1" */
  /* } */
  tags = merge(
    local.common_tags,
  { Name = "my-vpc-1" })
}
