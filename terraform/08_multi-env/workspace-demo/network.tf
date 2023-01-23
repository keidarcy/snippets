locals {
  common_tags = {
    Environment = "${terraform.workspace}"
  }
}

# create vpc
resource "aws_vpc" "my_vpc" {
  cidr_block           = var.vpc_cidr_block[terraform.workspace]
  instance_tenancy     = "default"
  enable_dns_hostnames = true
  tags                 = merge(local.common_tags, { Name = "my-vpc-${terraform.workspace}" })
}

# Create a Subnet
resource "aws_subnet" "my_subnet" {
  count      = var.subnet_count[terraform.workspace]
  vpc_id     = aws_vpc.my_vpc.id
  cidr_block = cidrsubnet(var.vpc_cidr_block[terraform.workspace], 8, count.index)
  tags       = merge(local.common_tags, { Name = "my-subnet-${terraform.workspace}-${count.index}" })
}
