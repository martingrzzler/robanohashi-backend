terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"
}

provider "aws" {
  region  = "eu-central-1"
  profile = "personal"
}

variable "key_pair_name" {
  type    = string
  default = "robanohashi_keypair"
}

variable "public_key" {
  type    = string
  default = file("~/.ssh/robanohashi_rsa.pub")
}


resource "aws_key_pair" "robanohashi_key" {
  key_name   = var.key_pair_name
  public_key = var.public_key
}

resource "aws_instance" "app_server" {
  ami                    = "ami-0c0933ae5caf0f5f9"
  instance_type          = "t2.micro"
  key_name               = aws_key_pair.robanohashi_key.key_name
  vpc_security_group_ids = [aws_security_group.vpc-ssh.id, aws_security_group.vpc-web.id]

  tags = {
    Name = "Robanohashi-Api"
  }
}

variable "domain" {
  type    = string
  default = "robanohashi.org"
}

resource "aws_route53_zone" "robanohashi_zone" {
  name = var.domain
}

resource "aws_route53_record" "robanohashi_zone_a" {
  zone_id = aws_route53_zone.robanohashi_zone.zone_id
  name    = var.domain
  type    = "A"
  ttl     = "300"
  records = [aws_instance.app_server.public_ip]
}

resource "aws_route53_record" "robanohashi_zone_api" {
  zone_id = aws_route53_zone.robanohashi_zone.zone_id
  name    = "api.${var.domain}"
  type    = "A"
  ttl     = "300"
  records = [aws_instance.app_server.public_ip]
}

resource "aws_security_group" "vpc-ssh" {
  name = "vpc-ssh"
  ingress {
    description = "Allow Port 22"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    description = "Allow all ip and ports outboun"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "vpc-web" {
  name = "vpc-web"
  ingress {
    description = "Allow Port 80"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "Allow Port 443"
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    description = "Allow all ip and ports outbound"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
