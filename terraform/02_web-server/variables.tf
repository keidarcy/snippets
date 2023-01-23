variable "ami_amazon_linux" {
	type = string
	description = "ec2 ami id"
	default = "ami-0bba69335379e17f8"
}

variable "instance_type" {
	type = string
	description = "ec2 instance type"
	default = "t2.micro"
}