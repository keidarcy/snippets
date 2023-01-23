locals {
  common_tags = {
    Environment = "${var.environment}"
  }
}

# create vpc
resource "aws_vpc" "my_vpc" {
  cidr_block           = var.vpc_cidr_block
  instance_tenancy     = "default"
  enable_dns_hostnames = true
  tags                 = merge(local.common_tags, { Name = "my-vpc-${var.environment}" })
}

# Create a Subnet
resource "aws_subnet" "my_subnet" {
  count      = var.subnet_count
  vpc_id     = aws_vpc.my_vpc.id
  cidr_block = cidrsubnet(var.vpc_cidr_block, 8, count.index)
  tags       = merge(local.common_tags, { Name = "my-subnet-${var.environment}-${count.index}" })
}
