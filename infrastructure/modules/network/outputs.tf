output "compute_subnet_az_0" {
  value = aws_subnet.mask_21[0].id
}

output "compute_subnet_az_1" {
  value = aws_subnet.mask_21[1].id
}

output "public_compute_subnet_az_4" {
  value = aws_subnet.mask_21[4].id
}

output "public_compute_subnet_az_5" {
  value = aws_subnet.mask_21[5].id
}

output "public_sg_ingress_insecure_id" {
  value = aws_security_group_rule.scg_alb_rule_port80.id
}

output "public_sg_ingress_secure_id" {
  value = aws_security_group_rule.scg_alb_rule_port443.id
}

output "vpc_id" {
  value = aws_vpc.network.id
}
