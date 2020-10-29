// https://blog.gruntwork.io/an-introduction-to-terraform-f17df9c6d180#3fd2

provider "aws" {
  region = "eu-central-1"
}

data "aws_availability_zones" "all" {}

resource "aws_instance" "example" {
  instance_type = "t2.micro"
  ami = "ami-03d85bfa79ad10274" // Ubuntu 20.04 Focal
  vpc_security_group_ids = [aws_security_group.instance.id]

# this is deeply despicable because it runs docker as root and it doesn't even work
  user_data = <<-EOF
              sudo apt update
              sudo apt install docker.io -y
              sudo docker run -p ${var.server_port}:${var.server_port} loefesto/wiki-go:latest
              EOF

  tags = {
    "Name" = "terraform-example"
  }

  key_name = aws_key_pair.example.key_name

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_key_pair" "example" {
  key_name   = "example_instance_key"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCuSxAYFYaUiYI2AasLdYD0sgynRkn3USM1XC8p3tPKSh4jfrIcu7kCgpJd3Ma1PTRxVAy5Gm4jq6f9lLsjH9wOtQyqs4/8GgLfEats4iVoPYzPa0URngWtgvnN08k9dK64hf/P5POkRlc3OmGIh7bEMaYBFj/HobAjIKwy19uJefPtV/TA+HMMkaEIcq0Y5kThJEg4C2EYDzff1OoObs9VlpM+yqEUxyiqscWRzNaWwF0vzfoPeKZJa50tr2iswpSujnhpQZJMyY16VWAsic3+op3sHbLMJcl6eAGcQany/bOnmH+rfXC8XnktSvA0Z4bf7Bb5lNqY74v80OV8bMEDK/zQkf4yP5NzwPgEn+mcirBST5gYcOXkeiABXm8ngLw9symMlpRe+pfou4Qe7yzrRL1tKDC4GwYJWql1MbKE2AGkd0Msi7lNyBXKxcwvJLeHQRAUtm5V+atxHAy6eNyQM0ZIbN8MnELDKTNzd8w+40cV3kEFjaXAUjJNtf9SbpM= marcopolita@SBVPN308"
}

resource "aws_security_group" "instance" {
  name = "terraform-example-instance"

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = var.server_port
    to_port = var.server_port
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "elb" {
  name = "terraform-example-elb"

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_elb" "example" {
  name               = "terraform-example"
  security_groups = [aws_security_group.elb.id]
  availability_zones = data.aws_availability_zones.all.names
  instances = [aws_instance.example.id]

  health_check {
    target              = "HTTP:${var.server_port}/"
    interval            = 30
    timeout             = 3
    healthy_threshold   = 2
    unhealthy_threshold = 2
  }

  listener {
    lb_port           = 80
    lb_protocol       = "http"
    instance_port     = var.server_port
    instance_protocol = "http"
  }
}