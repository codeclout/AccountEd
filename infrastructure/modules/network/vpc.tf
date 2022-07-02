locals {
  subnet_idx = [0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8]
  subnets = [
    "compute-${var.aws_region}",
    "db-${var.aws_region}",
    "public-compute-${var.aws_region}",
    "management-${var.aws_region}"
  ]
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
resource "aws_subnet" "mask_21" {
  count = var.availability_zone_count * length(local.subnets)

  availability_zone = data.aws_availability_zones.av.names[count.index % var.availability_zone_count]
  cidr_block        = cidrsubnet(aws_vpc.network.cidr_block, 5, count.index % 2 == 0 ? local.subnet_idx[count.index] : 8 + local.subnet_idx[count.index])
  vpc_id            = aws_vpc.network.id

  tags = {
    "Name" = "${local.subnets[element(slice(local.subnet_idx, 0, (length(local.subnet_idx) / 2) - 1), count.index)]}-${substr(data.aws_availability_zones.av.names[count.index % var.availability_zone_count], -1, -1)}"
  }
}
resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.network.id

  tags = merge({ Name = "igw-${var.app}-${var.environment}" }, var.tags)
}

# route table for each subnet
resource "aws_route_table" "explicit_subnet" {
  count = var.availability_zone_count * length(local.subnets)

  vpc_id = aws_vpc.network.id

  tags = merge({ Name = "route-table-${local.subnets[element(slice(local.subnet_idx, 0, (length(local.subnet_idx) / 2) - 1), count.index)]}-${substr(data.aws_availability_zones.av.names[count.index % var.availability_zone_count], -1, -1)}" }, var.tags)
}

# public route table
resource "aws_route_table" "public" {
  vpc_id = aws_vpc.network.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.igw.id
  }

  tags = {
    "Name" = "route-table-public"
  }
}

# assign each route table to a subnet
resource "aws_route_table_association" "route_table_association" {
  count = var.availability_zone_count * length(local.subnets)

  subnet_id      = aws_subnet.mask_21[count.index].id
  route_table_id = aws_route_table.explicit_subnet[count.index].id
}

# assign the public route table to the public subnet az1
resource "aws_route_table_association" "public_compute" {
  count = 2

  route_table_id = aws_route_table.public.id
  subnet_id      = aws_subnet.mask_21[4 + count.index].id
}

resource "aws_nat_gateway" "ngw" {
  count = 2

  allocation_id     = aws_eip.eip_nat_gateway[count.index].id
  connectivity_type = "public"
  subnet_id         = aws_subnet.mask_21[4 + count.index].id
}

resource "aws_eip" "eip_nat_gateway" {
  count = 2

  vpc = true
  depends_on = [
    aws_internet_gateway.igw
  ]
}

resource "aws_route" "compute_a" {
  destination_cidr_block = "0.0.0.0/0"
  nat_gateway_id         = aws_nat_gateway.ngw[0].id
  route_table_id         = aws_route_table.explicit_subnet[0].id

  depends_on = [
    aws_route_table.explicit_subnet[0]
  ]
}

resource "aws_route" "compute_b" {
  destination_cidr_block = "0.0.0.0/0"
  nat_gateway_id         = aws_nat_gateway.ngw[1].id
  route_table_id         = aws_route_table.explicit_subnet[1].id

  depends_on = [
    aws_route_table.explicit_subnet[1]
  ]
}

resource "aws_route" "management_a" {
  destination_cidr_block = "0.0.0.0/0"
  nat_gateway_id         = aws_nat_gateway.ngw[0].id
  route_table_id         = aws_route_table.explicit_subnet[6].id

  depends_on = [
    aws_route_table.explicit_subnet[6]
  ]
}

resource "aws_route" "management_b" {
  destination_cidr_block = "0.0.0.0/0"
  nat_gateway_id         = aws_nat_gateway.ngw[1].id
  route_table_id         = aws_route_table.explicit_subnet[7].id

  depends_on = [
    aws_route_table.explicit_subnet[7]
  ]
}
