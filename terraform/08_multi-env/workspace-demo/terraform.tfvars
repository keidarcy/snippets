vpc_cidr_block = {
  dev  = "10.0.0.0/16"
  qa   = "10.1.0.0/16"
  prod = "10.2.0.0/16"
}

subnet_count = {
  dev  = 1
  qa   = 1
  prod = 2
}