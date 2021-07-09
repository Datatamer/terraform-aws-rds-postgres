output "ingress_ports" {
  value = module.sg-ports
  description = "List of ingress ports"
}

output "rds" {
  value = module.rds_postgres
}
