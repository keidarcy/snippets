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

variable "enable_dns_hostnames" {
  type        = bool
  description = "Enable DNS hostnames"
  default     = true
}

variable "volumn_size" {
  default = 10
}

variable "aws_instance_sizes" {
  type        = map(string)
  description = "aws instance sizes"
  default = {
    small  = "t2.micro"
    medium = "t2.small"
    large  = "t2.medium"
  }
}

variable "students" {
  type        = map(string)
  description = "student information"
  default = {
    name = "xxx"
    age  = 20
  }
}

variable "tuple_test" {
  type        = tuple([string, number, bool])
  description = "tuple test"
  default     = ["xxx", 20, true]
}

variable "db_port" {
  type = object({
    external = number
    internal = number
    protocol = string
  })
  default = {
    external = 5432
    internal = 5432
    protocol = "tcp"
  }
}

variable "aws_regions" {
  type    = list(string)
  default = ["us-east-1", "us-east-2", "ap-northeast-1"]
}


