variable "vpc_cidr_block" {
  type = map(string)
  default = {
    private = "10.1.0.0/16"
    public  = "192.168.0.0/16"
  }
}

variable "users" {
  type = map(object({
    is_admin = bool
    name     = string
  }))

  default = {
    alice = {
      is_admin = true
      name     = "Alice"
    }
    bob = {
      is_admin = false
      name     = "Bob"
    }
  }
}

variable "aws_regions" {
  type    = list(string)
  default = ["eu-central-1", "us-east-1", "us-east-2"]
}
