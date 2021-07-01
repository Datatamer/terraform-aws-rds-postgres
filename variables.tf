variable "password" {
  description = "The postgres password"
  type        = string
}

variable "username" {
  description = "The postgres username"
  type        = string
  default     = "tamr"
}

variable "postgres_name" {
  description = "The name of the postgres instance"
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

variable "subnet_name" {
  description = "The name of the subnet to add the RDS instance to"
  type        = string
  default     = null
}

variable "copy_tags_to_snapshot" {
  description = "Copy tags to snapshots"
  type        = bool
  default     = true
}

variable "additional_tags" {
  description = "Additional tags to set on the RDS instance"
  type        = map
  default     = {}
}

variable "security_group_name" {
  description = "Name for the security group for the rds instance"
  type        = string
  default     = "tamr_rds_sg"
}

variable "tamr_vm_sg_id" {
  description = "Security group id attached to the tamr vm"
  type        = string
}

variable "spark_cluster_sg_ids" {
  description = "Security group is attached to the ec2 instances of EMR Spark"
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
  default     = "12.3"
}

variable "parameter_group_family" {
  description = "The family of the DB parameter group"
  type        = string
  default     = "postgres12"
}
