# from https://registry.terraform.io/providers/hashicorp/http/latest/docs/data-sources/http
data "http" "example" {
  url = "https://checkpoint-api.hashicorp.com/v1/check/terraform"

  # Optional request headers
  request_headers = {
    Accept = "application/json"
  }
}

output "http_status" {
  value = jsondecode(data.http.example.response_body)["project_website"]
}


# get AZs list
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/availability_zones
data "aws_availability_zones" "available" {
  state = "available"
}

output "availability_zones" {
  value = data.aws_availability_zones.available.names
}

# get AMI ID
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/ami
data "aws_ami" "amazon_linux" {
  most_recent = true
  owners      = ["137112412989"]
  filter {
    name   = "name"
    values = ["amzn2-ami-kernel-5.10*"]
  }
}

output "ami" {
  value = data.aws_ami.amazon_linux.id
}


# # read data from terraform_remote_state
# #https://www.terraform.io/language/state/remote-state-data
# data "terraform_remote_state" "vpc" {
#   backend = "consul"
#   config = {
#     address = "127.0.0.1:8500"
#     scheme  = "http"
#     path    = "test/terraform.tfstate"
#   }
# }

# output "vpc_id" {
#   value = data.terraform_remote_state.vpc.outputs.vpc_id
# }

# # read data from consul
# # https://registry.terraform.io/providers/hashicorp/consul/latest/docs
# data "consul_keys" "test" {

#   key {
#     name    = "cidr"
#     path    = "test/terraform"
#     default = ""
#   }
# }

# resource "aws_vpc" "vpc" {
#   cidr_block           = jsondecode(data.consul_keys.test.var.cidr)["cidr"]
#   enable_dns_hostnames = true
#   tags = {
#     Name = "my-vpc"
#   }
# }

# read data from template_file
data "template_file" "tmp" {
  template = file("${path.module}/test.json")
  vars = {
    cidr_block = var.cidr_block
  }
}

output "tmp" {
  value = jsondecode(data.template_file.tmp.rendered)["cidr-from-tmp"]
}

output "tmp-json" {
  value = jsondecode(data.template_file.tmp.rendered)
}
