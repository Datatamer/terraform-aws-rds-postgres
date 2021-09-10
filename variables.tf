variable "password" {
  description = "The password for the master DB user."
  type        = string
}

variable "username" {
  description = "The username for the master DB user."
  type        = string
  default     = "tamr"
}

variable "db_port" {
  description = "The port on which the database accepts connections."
  type        = number
  default     = 5432
}

variable "postgres_name" {
  description = "The name of the postgres database to create on the DB instance"
  type        = string
  default     = "tamr_rds_db"
}

variable "parameter_group_name" {
  description = "The name of the rds parameter group"
  type        = string
  default     = "rds-postgres-pg"
}

variable "identifier_prefix" {
  description = "Identifier prefix for the RDS instance"
  type        = string
  default     = "tamr-rds-"
}

variable "allocated_storage" {
  description = "Allocate storage"
  type        = number
  default     = 20
}

variable "max_allocated_storage" {
  description = "Max allocate storage"
  type        = number
  default     = 1000
}

variable "storage_type" {
  description = "Storage type (e.g. gp2, io1)"
  type        = string
  default     = "gp2"
}

variable "instance_class" {
  description = "Instance class"
  type        = string
  default     = "db.m4.large"
}

variable "maintenance_window" {
  description = "Maintenance window"
  type        = string
  default     = "sun:04:32-sun:05:02"
}

variable "backup_window" {
  description = "Backup window"
  type        = string
  default     = "03:29-03:59"
}

variable "backup_retention_period" {
  description = "Backup retention period in days"
  type        = number
  default     = 14
}

variable "skip_final_snapshot" {
  description = "Skip final snapshot"
  type        = bool
  default     = true
}

variable "apply_immediately" {
  description = "Apply immediately, do not set this to true for production"
  type        = bool
  default     = false
}

variable "subnet_group_name" {
  description = "The name of the subnet group to add the RDS instance to"
  type        = string
}

variable "rds_subnet_ids" {
  description = "VPC subnet IDs in subnet group"
  type        = list(string)
}

variable "copy_tags_to_snapshot" {
  description = "Copy tags to snapshots"
  type        = bool
  default     = true
}

variable "additional_tags" {
  type        = map(string)
  description = "[DEPRECATED: Use `tags` instead] Additional tags to set on the RDS instance."
  default     = {}
}

variable "tags" {
  type        = map(string)
  description = "A map of tags to add to all resources. Replaces `additional_tags`."
  default     = {}
}

variable "security_group_name" {
  description = "Name for the security group for the rds instance"
  type        = string
  default     = "tamr_rds_sg"
}

variable "security_group_ids" {
  description = "List of security group IDs to associate"
  type        = list(string)
}

variable "vpc_id" {
  description = "VPC ID for the rds security group"
  type        = string
}

variable "additional_cidrs" {
  description = "Additional CIDR to connect to RDS Postgres instance"
  type        = list(string)
  default     = []
}

variable "engine_version" {
  description = "Version of RDS Postgres"
  type        = string
  default     = "12"
}

variable "parameter_group_family" {
  description = "The family of the DB parameter group"
  type        = string
  default     = "postgres12"
}

variable "auto_minor_version_upgrade" {
  default     = true
  type        = bool
  description = "Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window"
}
