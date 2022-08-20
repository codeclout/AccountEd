resource "aws_network_acl" "alb_public" {
  vpc_id = aws_vpc.network.id
}

resource "aws_network_acl_rule" "alb_public_inbound_insecure" {
  network_acl_id = aws_network_acl.alb_public

  cidr_block  = "0.0.0.0/0"
  protocol    = "tcp"
  rule_action = "allow"

  egress = false

  from_port   = 80
  rule_number = 30
  to_port     = 80
}

resource "aws_network_acl_rule" "alb_public_inbound_secure" {
  network_acl_id = aws_network_acl.alb_public

  cidr_block  = "0.0.0.0/0"
  protocol    = "tcp"
  rule_action = "allow"

  egress = false

  from_port   = 443
  rule_number = 80
  to_port     = 443
}

resource "aws_network_acl_rule" "alb_public_outbound" {
  cidr_block     = aws_vpc.network.cidr_block
  network_acl_id = aws_network_acl.alb_public

  protocol    = "tcp"
  rule_action = "allow"

  egress = false

  from_port   = 1024
  rule_number = 130
  to_port     = 65535
}

resource "aws_network_acl_rule" "alb_public_outbound" {
  network_acl_id = aws_network_acl.alb_public

  cidr_block  = "0.0.0.0/0"
  protocol    = "tcp"
  rule_action = "allow"

  egress = true

  from_port   = 1024
  rule_number = 30
  to_port     = 65535
}

resource "aws_network_acl_association" "alb_public_nacl_subnet_4" {
  network_acl_id = aws_network_acl.alb_public.id
  subnet_id      = aws_subnet.mask_21[4].id
}

resource "aws_network_acl_association" "alb_public_nacl_subnet_5" {
  network_acl_id = aws_network_acl.alb_public.id
  subnet_id      = aws_subnet.mask_21[5].id
}


