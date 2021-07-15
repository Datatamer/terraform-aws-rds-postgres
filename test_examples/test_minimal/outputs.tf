output "rds_db_port" {
    value = module.rds_postgres.rds_db_port
}

output "rds_dbname" {
    value = module.rds_postgres.rds_dbname
}

output "rds_hostname" {
    value = module.rds_postgres.rds_hostname
}

output "rds_postgres_id" {
    value = module.rds_postgres.rds_postgres_id
}

output "rds_postgres_pg_id" {
    value = module.rds_postgres.rds_postgres_pg_id
}

output "rds_security_group_ids" {
    value = module.rds_postgres.rds_security_group_ids
}

output "rds_username" {
    value = module.rds_postgres.rds_username
}
