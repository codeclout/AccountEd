resource "aws_security_group" "scg_ecs" {
  name   = "${var.environment}-scg-ecs"
  vpc_id = aws_vpc.network.id
  tags   = var.tags
}

resource "aws_security_group_rule" "scg_ecs_rule" {
  protocol = "tcp"
  type     = "ingress"

  from_port = 32768
  to_port   = 65535

  cidr_blocks              = [aws_vpc.network.cidr_block]
  security_group_id        = aws_security_group.scg_ecs.id
  source_security_group_id = aws_security_group.scg_alb.id
}

resource "aws_security_group" "scg_alb" {
  name   = "${var.environment}-scg-alb"
  vpc_id = aws_vpc.network.id
  tags   = var.tags
}

resource "aws_security_group_rule" "scg_alb_rule_port80" {
  protocol = "tcp"
  type     = "ingress"

  from_port = 80
  to_port   = 80

  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.scg_alb.id
}

resource "aws_security_group_rule" "scg_alb_rule_port443" {
  protocol = "tcp"
  type     = "ingress"

  from_port = 443
  to_port   = 443

  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.scg_alb.id
}
