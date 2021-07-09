vpc_id              = "vpc-example"
subnet_ids          = ["subnet-az1", "subnet-az2"]
security_group_ids  = ["example-spark-service-access-sg", "example-tamr-vm-sg"]
name_prefix         = "test" # Replace me for a more specific prefix
ingress_cidr_blocks = ["1.2.3.0/24"]
