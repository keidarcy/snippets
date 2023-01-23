data "aws_ami" "amazon_linux" {
  most_recent = true
  owners      = ["137112412989"]
  filter {
    name   = "name"
    values = ["amzn2-ami-kernel-5.10*"]
  }
}

resource "aws_instance" "my_ec2" {
  ami                    = data.aws_ami.amazon_linux.id
  instance_type          = var.instance_type
  key_name               = aws_key_pair.deployer.id
  vpc_security_group_ids = [aws_security_group.my_sg.id]
  subnet_id              = aws_subnet.my_subnet.id

  tags = {
    "Name" = "terraform created"
  }

  root_block_device {
    volume_size = 10
  }

  # user_data = <<-EOT
  # 	#!/bin/bash
  # 	yum update -y
  # 	yum install -y httpd
  # 	systemctl start httpd
  # 	systemctl enable httpd
  # 	echo "Hello World from Terraform" > /var/www/html/index.html
  # 	EOT
  user_data = templatefile("${path.module}/user_data.tpl", { name : "terraform templatefile" })
}

resource "aws_eip" "my_eip" {
  instance = aws_instance.my_ec2.id
  tags = {
    "Name" = "terraform created"
  }
}

output "ec2_public_ip" {
  value = aws_eip.my_eip.public_ip
}
