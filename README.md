# Tamr AWS RDS Terraform Module
This terraform module creates an AWS RDS postgres instance that will be used by TAMR.
This repo follows the [terraform standard module structure](https://www.terraform.io/docs/modules/index.html#standard-module-structure).

# Examples
## Basic
Inline example implementation of the module.  This is the most basic example of what it would look like to use this module.
```
module "rds_postgres" {
  source               = "git::https://github.com/Datatamer/terraform-aws-rds-postgres.git?ref=x.y.z"
  postgres_name        = "example_rds_postgres"
  parameter_group_name = "example-rds-postgres-pg"
  identifier_prefix    = "example-rds-"
  username             = "exampleUsername"
  password             = "examplePassword"

  subnet_group_name    = "example_subnet"
  rds_subnet_ids       = ["example-subnet-1", "example-subnet-2"]
  ingress_sg_ids       = ["sg-sparksecuritygroup1", "sg-sparksecuritygroup2", "sg-tamrvmsecuritygroup"]
  vpc_id               = "vpc-examplevpcnetworkid"
}
```
## Minimal
Smallest complete fully working example. This example might require extra resources to run the example.
- [Minimal](https://github.com/Datatamer/terraform-aws-rds-postgres/tree/master/examples/minimal)

# Resources Created
This terraform module will create:
* an AWS RDS Postgres instance
* a database parameter group
* a database subnet group

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| terraform | >= 0.13 |
| aws | >= 3.36, !=4.0.0, !=4.1.0, !=4.2.0, !=4.3.0, !=4.4.0, !=4.5.0, !=4.6.0, !=4.7.0, !=4.8.0 |

## Providers

| Name | Version |
|------|---------|
| aws | >= 3.36, !=4.0.0, !=4.1.0, !=4.2.0, !=4.3.0, !=4.4.0, !=4.5.0, !=4.6.0, !=4.7.0, !=4.8.0 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| password | The password for the master DB user. | `string` | n/a | yes |
| rds\_subnet\_ids | VPC subnet IDs in subnet group | `list(string)` | n/a | yes |
| security\_group\_ids | List of security group IDs to associate | `list(string)` | n/a | yes |
| subnet\_group\_name | The name of the subnet group to add the RDS instance to | `string` | n/a | yes |
| additional\_tags | [DEPRECATED: Use `tags` instead] Additional tags to set on the RDS instance. | `map(string)` | `{}` | no |
| allocated\_storage | Allocate storage | `number` | `20` | no |
| apply\_immediately | Apply immediately, do not set this to true for production | `bool` | `false` | no |
| auto\_minor\_version\_upgrade | Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window | `bool` | `true` | no |
| backup\_retention\_period | Backup retention period in days | `number` | `14` | no |
| backup\_window | Backup window | `string` | `"03:29-03:59"` | no |
| copy\_tags\_to\_snapshot | Copy tags to snapshots | `bool` | `true` | no |
| db\_port | The port on which the database accepts connections. | `number` | `5432` | no |
| enabled\_cloudwatch\_logs\_exports | Indicates that postgresql logs will be configured to be sent automatically to Cloudwatch | `bool` | `true` | no |
| engine\_version | Version of RDS Postgres | `string` | `"12"` | no |
| identifier\_prefix | Identifier prefix for the RDS instance | `string` | `"tamr-rds-"` | no |
| instance\_class | Instance class | `string` | `"db.m4.large"` | no |
| maintenance\_window | Maintenance window | `string` | `"sun:04:32-sun:05:02"` | no |
| max\_allocated\_storage | Max allocate storage | `number` | `1000` | no |
| multi\_az | Specifies if the RDS instance is multi-AZ. | `bool` | `true` | no |
| param\_log\_min\_duration\_statement | (ms) Sets the minimum execution time above which statements will be logged. | `string` | `"-1"` | no |
| param\_log\_statement | Sets the type of statements logged. Valid values are none, ddl, mod, all | `string` | `"none"` | no |
| parameter\_group\_family | The family of the DB parameter group | `string` | `"postgres12"` | no |
| parameter\_group\_name | The name of the rds parameter group | `string` | `"rds-postgres-pg"` | no |
| performance\_insights\_enabled | Specifies whether Performance Insights are enabled. | `bool` | `false` | no |
| performance\_insights\_retention\_period | The amount of time in days to retain Performance Insights data. Either 7 (7 days) or 731 (2 years). | `number` | `7` | no |
| postgres\_name | The name of the postgres database to create on the DB instance | `string` | `"tamr_rds_db"` | no |
| skip\_final\_snapshot | Skip final snapshot | `bool` | `true` | no |
| storage\_type | Storage type (e.g. gp2, io1) | `string` | `"gp2"` | no |
| tags | A map of tags to add to all resources. Replaces `additional_tags`. | `map(string)` | `{}` | no |
| username | The username for the master DB user. | `string` | `"tamr"` | no |

## Outputs

| Name | Description |
|------|-------------|
| cloudwatch\_logs\_path | n/a |
| rds\_db\_port | n/a |
| rds\_dbname | n/a |
| rds\_hostname | n/a |
| rds\_postgres\_id | ID of the of the RDS instance |
| rds\_postgres\_pg\_id | ID of the RDS postgres parameter group |
| rds\_security\_group\_ids | List of security group ids attached to the rds instance |
| rds\_username | n/a |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

# References
* AWS RDS: https://aws.amazon.com/rds/features/
* Terraform module structure: https://www.terraform.io/docs/modules/index.html#standard-module-structure

# Development
## Releasing new versions
* Updated version contained in `VERSION`
* Documented changes in `CHANGELOG.md`
* Create a tag in github for the commit associated with the version

# License
Apache 2 Licensed. See LICENSE for full details.
