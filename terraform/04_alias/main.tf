# https://developer.hashicorp.com/terraform/language/providers/configuration#alias-multiple-provider-configurations
provider "aws" {
  region = "us-east-1"
}

provider "aws" {
  alias  = "west"
  region = "us-west-2"
}

resource "aws_instance" "foo" {
  provider = aws.west

  # ...
}
