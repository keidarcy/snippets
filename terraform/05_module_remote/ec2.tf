resource "aws_instance" "my_ec2" {
  ami                    = var.ami_amazon_linux
  instance_type          = var.instance_type
  key_name               = aws_key_pair.deployer.id
  vpc_security_group_ids = [module.vpc.default_security_group_id]
  subnet_id              = module.vpc.public_subnets[0]

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


  # user_data = data.template_file.cloud-init.rendered
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
