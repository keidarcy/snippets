
resource "aws_vpc" "vpc" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = "true"
  tags = {
    Name = jsondecode(data.http.example.response_body)["project_website"]
  }
}

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
