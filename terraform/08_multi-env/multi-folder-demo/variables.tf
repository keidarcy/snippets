variable "aws_region" {
  type    = string
}

variable "aws_access_key" {

  type        = string
  description = "aws access key"
  sensitive   = true
}

variable "aws_secret_key" {

  type        = string
  description = "aws secret key"
  sensitive   = true
}

variable "vpc_cidr_block" {
  type        = string
  description = "base cidr block for vpc"

}
variable "subnet_count" {
  type        = number
  description = "number of subnets to create"
}

variable "environment" {
  type        = string
  description = "environment name"
}