# resource "aws_vpc" "my_vpc" {
#   cidr_block         = var.aws_vpc_cidr_block
#   instance_tenancy   = "default"
#   enable_dns_support = true
# }

# resource "aws_subnet" "my_subnet" {
#   vpc_id     = aws_vpc.my_vpc.id
#   cidr_block = cidrsubnet(var.aws_vpc_cidr_block, 8, 1)
# }

# resource "aws_internet_gateway" "my_igw" {
#   vpc_id = aws_vpc.my_vpc.id
# }

# module "my_vpc" {
#   source             = "./modules/vpc"
#   aws_vpc_cidr_block = var.aws_vpc_cidr_block
# }

module "my_vpc" {
  source             = "keidarcy/vpc/aws"
  version            = "1.0.0"
  aws_vpc_cidr_block = var.aws_vpc_cidr_block
}

output "vpc_id" {
  value = module.my_vpc.id
}

# module "my_vpc_2" {
#   source             = "./modules/vpc"
#   aws_vpc_cidr_block = "10.1.0.0/16"
# }
