module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = ">=3.14.2"

  name           = var.aws_vpc_name
  cidr           = var.aws_vpc_cidr
  azs            = var.aws_availability_zones
  public_subnets = [cidrsubnet(var.aws_vpc_cidr, 8, 1)]

  create_igw = true

  manage_default_security_group  = true
  default_security_group_egress  = var.aws_sg_egress_rules
  default_security_group_ingress = var.aws_sg_ingress_rules
}
