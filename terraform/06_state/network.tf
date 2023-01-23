module "my_vpc" {
  source             = "keidarcy/vpc/aws"
  version            = "1.0.0"
  aws_vpc_cidr_block = var.aws_vpc_cidr_block
}

output "vpc_id" {
  value = module.my_vpc.id
}
