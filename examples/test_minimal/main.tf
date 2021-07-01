module "rds_postgres" {
  #source               = "git::https://github.com/Datatamer/terraform-rds-postgres.git?ref=0.1.0"
  source               = "../../"
  postgres_name        = var.postgres_db_name
  parameter_group_name = var.parameter_group_name
  identifier_prefix    = var.identifier_prefix
  instance_class = "db.t3.medium"

  username = var.pg_username
  password = var.pg_password

  subnet_name          = module.vpc.database_subnet_group_name
  vpc_id               = module.vpc.vpc_id

  spark_cluster_sg_ids = [aws_security_group.example[0].id] #mock
  tamr_vm_sg_id        = aws_security_group.example[1].id #mock
}

data "aws_region" "current" {

}
module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "3.1.0"

  name = "hugo-test-vpc"
  cidr = "172.18.0.0/18"

  azs             = [for i in ["a","b","c"] : "${data.aws_region.current.name}${i}"]
  private_subnets = ["172.18.0.0/24", "172.18.1.0/24", "172.18.2.0/24"]
  public_subnets  = ["172.18.3.0/24", "172.18.4.0/24", "172.18.5.0/24"]
  
  database_subnets    = ["172.18.6.0/24", "172.18.7.0/24", "172.18.8.0/24"]

  create_database_subnet_group           = true
  create_database_subnet_route_table     = true
  create_database_internet_gateway_route = true

  enable_nat_gateway = false
  enable_vpn_gateway = false

  tags = {
    Terraform = "true"
    Environment = "dev"
  }
}

resource "aws_security_group" "example" {
  count = 2
  # ... other configuration ...

  vpc_id = module.vpc.vpc_id
  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }
}
