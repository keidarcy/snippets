# Description: Looping over resources
# `count` is a special argument that can be used to create multiple instances of a resource
# vpc
resource "aws_vpc" "vpc" {
  count                = 2
  cidr_block           = "10.${count.index}.0.0/16"
  enable_dns_hostnames = true
  tags = {
    Name = "terraform-vpc-${count.index}"
  }
}

resource "aws_subnet" "my_subnet" {
  vpc_id     = aws_vpc.vpc[0].id
  cidr_block = "10.0.1.0/24"
  tags = {
    Name = "terraform-subnet"
  }
}

# output all vpc ids
output "vpc_id" {
  value = aws_vpc.vpc[*].id
}

# for_each
resource "aws_vpc" "vpc2" {
  for_each = {
    private = "10.1.0.0/16"
    public  = "192.168.0.0/16"
  }
  cidr_block           = each.value
  enable_dns_hostnames = true
  tags = {
    Name = "terraform-vpc-${each.key}"
    # Name = "terraform-vpc-${each["private"]}"
  }
}

locals {
  admin_users = {
    for name, user in merge(var.users) : name => user
    if user.is_admin
  }
  regular_users = {
    for name, user in var.users : name => user
    if !user.is_admin
  }
}


