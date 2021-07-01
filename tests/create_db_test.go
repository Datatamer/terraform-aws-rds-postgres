package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestTerraformCreateRDS(t *testing.T) {
	t.Parallel()

	namePrefix := "terratest-aws-rds-example"
	// expectedName := fmt.Sprintf("%s-%s", namePrefix, strings.ToLower(random.UniqueId()))
	expectedName := namePrefix // remove this later
	expectedDBName := strings.ReplaceAll(expectedName, "-", "_")

	//awsRegion := aws.GetRandomStableRegion(t, nil, nil)
	awsRegion := "us-east-1" // remove this later

	user := "tamruser"
	pw := "tamrpassword"
	expectedPort := int64(5432)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/test_minimal",

		Vars: map[string]interface{}{
			"postgres_db_name":     expectedDBName,
			"parameter_group_name": fmt.Sprintf("%s-rds-postgres-pg", namePrefix),
			"identifier_prefix":    fmt.Sprintf("%s-", namePrefix),
			"pg_username":          user,
			"pg_password":          pw,
			// "spark_sg_id_list":     []string{"sg-0b103ea834dfa1234"},
			// "tamr_vm_sg_id":        "sg-0b103ea834dfac123",
		},
		// Environment variables to set when running Terraform
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	})

	//
	// defer terraform.Destroy(t, terraformOptions)
	// terraform.InitAndApply(t, terraformOptions)

	rds := terraform.OutputMapOfObjects(t, terraformOptions, "module_rds")
	//   output "rds_postgres_pg_id"
	//   output "rds_postgres_id"
	//   output "rds_sg_id"
	//   output "rds_hostname"

	dbInstanceID := rds["rds_postgres_id"].(string)
	// dbHostname := rds["rds_hostname"].(string)
	// dbSecGroupID := rds["rds_sg_id"].(string)
	// dbParamGroupID := rds["rds_postgres_pg_id"].(string)

	logger.Log(t, fmt.Sprintf("RDS object: %+v", rds))

	// Look up the endpoint address and port of the RDS instance
	address := aws.GetAddressOfRdsInstance(t, dbInstanceID, awsRegion)
	port := aws.GetPortOfRdsInstance(t, dbInstanceID, awsRegion)

	// describe rds -  SG 5432 port
	// aws.GetAllParametersOfRdsInstance()

	// describe rds - param group

	// test connectivity
	// schemaExistsInRdsInstance := GetWhetherSchemaExistsInRdsPostgresInstance(t, rds["rds_hostname"].(string), int64(expectedPort), user, pw, expectedDBName)

	// Verify that the table/schema requested for creation is actually present in the database
	// assert.True(t, schemaExistsInRdsInstance)

	// assert.Equal(t, expectedName, )

	// Verify that the address is not null
	assert.NotNil(t, address)
	// Verify that the DB instance is listening on the port mentioned
	assert.Equal(t, expectedPort, port)

}
