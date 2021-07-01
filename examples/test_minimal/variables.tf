variable "postgres_db_name" {
  type        = string
  description = "Name of the postgres db"
}

variable "parameter_group_name" {
  type        = string
  description = "Name of the parameter group"
}

variable "identifier_prefix" {
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

# variable "spark_sg_id_list" {
#   type  = list(string)
#   description = "List of IDs of Sec groups"
# }


# variable "tamr_vm_sg_id" {
#   type  = string
#   description = "ID of Tamr VM Sec groups"
# }