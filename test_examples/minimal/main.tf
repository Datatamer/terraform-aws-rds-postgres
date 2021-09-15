module "rds_postgres" {
  #source               = "git::https://github.com/Datatamer/terraform-rds-postgres.git?ref=x.y.z"
  source               = "../../"
  postgres_name        = var.postgres_db_name
  parameter_group_name = var.parameter_group_name
  identifier_prefix    = var.name_prefix
  instance_class       = "db.t3.medium"
  engine_version       = "12.5"
  username             = var.pg_username
  password             = var.pg_password

  subnet_group_name = "${var.name_prefix}-subnet-group"
  vpc_id            = module.vpc.vpc_id

  # Network requirement: DB subnet group needs a subnet in at least two Availability Zones
  rds_subnet_ids     = module.vpc.database_subnets
  security_group_ids = module.rds-postgres-sg.security_group_ids
}

data "aws_region" "current" {}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "3.1.0"

  name = "${var.name_prefix}vpc"
  cidr = "172.18.0.0/18"

  azs             = [for i in ["a", "b"] : "${data.aws_region.current.name}${i}"]
  private_subnets = ["172.18.0.0/24", "172.18.1.0/24"]
  public_subnets  = ["172.18.3.0/24", "172.18.4.0/24"]

  database_subnets = ["172.18.6.0/24", "172.18.7.0/24"]

  create_database_subnet_group           = false
  create_database_subnet_route_table     = true
  create_database_internet_gateway_route = true

  enable_nat_gateway = false
  enable_vpn_gateway = false

  tags = {
    Terraform   = "true"
    Terratest   = "true"
    Environment = "dev"
  }
}

module "sg-ports" {
  # source               = "git::https://github.com/Datatamer/terraform-aws-rds-postgres.git//modules/rds-postgres-ports?ref=2.0.0"
  source = "../../modules/rds-postgres-ports"
}

module "rds-postgres-sg" {
  source              = "git::git@github.com:Datatamer/terraform-aws-security-groups.git?ref=1.0.0"
  vpc_id              = module.vpc.vpc_id
  ingress_cidr_blocks = module.vpc.database_subnets_cidr_blocks
  egress_cidr_blocks  = var.egress_cidr_blocks
  ingress_ports       = module.sg-ports.ingress_ports
  sg_name_prefix      = var.name_prefix
  egress_protocol     = "all"
  ingress_protocol    = "tcp"
}
