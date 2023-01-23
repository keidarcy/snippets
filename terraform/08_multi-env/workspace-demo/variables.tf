# variable "aws_region" {
#   type    = string
#   default = "eu-central-1"
# }

# variable "aws_access_key" {

#   type        = string
#   description = "aws access key"
#   sensitive   = true
# }

# variable "aws_secret_key" {

#   type        = string
#   description = "aws secret key"
#   sensitive   = true
# }

variable "vpc_cidr_block" {
  type        = map(string)
  description = "base cidr block for vpc"

}
variable "subnet_count" {
  type        = map(number)
  description = "number of subnets to create"
}