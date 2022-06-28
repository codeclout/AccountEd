locals {
  subnet_idx = [0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8]
  subnets = {
    public : ["compute-${var.aws_region}", "management-${var.aws_region}"],
    private : ["compute-${var.aws_region}", "db-${var.aws_region}", "management-${var.aws_region}"]
  }
  vpc_cidr = "10.0.0.0/16"
}

data "aws_availability_zones" "av" {
  state = "available"

  filter {
    name   = "opt-in-status"
    values = ["opt-in-not-required"]
  }
}

resource "aws_vpc" "network" {
  cidr_block           = local.vpc_cidr
  enable_dns_hostnames = true
  instance_tenancy     = "default"

  tags = merge({ Name = "vpc-${var.app}-${var.environment}" }, var.tags)
}

# models - 10.0.0.0/18 for each AZ 
# (10.0.{0-56}.0/21) in AZ 1
# (10.0.{64+}.0/21) in AZ 2
resource "aws_subnet" "public" {
  count = var.availability_zone_count * length(local.subnets.public)

  availability_zone = data.aws_availability_zones.av.names[count.index % var.availability_zone_count]
  cidr_block        = cidrsubnet(aws_vpc.network.cidr_block, 5, count.index % 2 == 0 ? local.subnet_idx[count.index] : 8 + local.subnet_idx[count.index])
  vpc_id            = aws_vpc.network.id

  tags = {
    "Name" = "${local.subnets.public[count.index > var.availability_zone_count / 2 ? 1 : 0]}-${substr(data.aws_availability_zones.av.names[count.index % var.availability_zone_count], -1, -1)}"
  }
}

# resource "aws_route" "public_route" {
#   route_table_id = aws_route_table.public_route_table.id

#   destination_cidr_block = "0.0.0.0/0"
#   gateway_id             = aws_internet_gateway.igw.id
# }

resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.network.id

  tags = merge({ Name = "igw-${var.app}-${var.environment}" }, var.tags)
}

resource "aws_route_table" "public_route_table" {
  count = var.availability_zone_count * length(local.subnets.public)
  route {
    cidr_block = aws_subnet.public[count.index].cidr_block
    gateway_id = aws_internet_gateway.igw.id
  }
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.igw.id
  }

  vpc_id = aws_vpc.network.id

  tags = merge({ Name = "route-table-${local.subnets.public[count.index > var.availability_zone_count / 2 ? 1 : 0]}-${substr(data.aws_availability_zones.av.names[count.index % var.availability_zone_count], -1, -1)}" }, var.tags)
}

resource "aws_route_table_association" "route_table_association" {
  count = var.availability_zone_count * length(local.subnets)

  subnet_id      = aws_subnet.public[count.index].id
  route_table_id = aws_route_table.public_route_table[count.index].id
}
