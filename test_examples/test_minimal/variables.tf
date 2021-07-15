variable "postgres_db_name" {
  type        = string
  description = "Name of the postgres db"
}

variable "parameter_group_name" {
  type        = string
  description = "Name of the parameter group"
}

variable "name_prefix" {
  type        = string
  description = "Identifier prefix for the resources"
}

variable "pg_username" {
  type        = string
  description = "Username for postgres"
}

variable "pg_password" {
  type        = string
  description = "Password for postgres"
}

variable "egress_cidr_blocks" {
  description = "CIDR blocks to attach to security groups for egress"
  type = list(string)
  default = ["0.0.0.0/0"]
}
