provider "aws" {
    region = var.region
}

resource "aws_qldb_ledger" "accounted_ledger" {
    name = var.name
    tags = var.tags
}