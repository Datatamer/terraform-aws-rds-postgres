package tests

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	terratest_testing "github.com/gruntwork-io/terratest/modules/testing"
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

	awsRegion := "us-east-2" // remove this later

	user := "tamruser"
	pw := "tamrpassword"
	expectedPort := 5432

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
	logger.Log(t, fmt.Sprintf("RDS object: %+v", rds))

	// test connectivity
	schemaExistsInRdsInstance := GetWhetherSchemaExistsInRdsPostgresInstance(t, rds["rds_hostname"].(string), int64(expectedPort), user, pw, expectedDBName)

	// Look up the endpoint address and port of the RDS instance
	address := aws.GetAddressOfRdsInstance(t, dbInstanceID, awsRegion)
	port := aws.GetPortOfRdsInstance(t, dbInstanceID, awsRegion)

	// describe rds -  SG 5432 port

	// describe rds - param group

	// rds_postgres_id := rds["rds_postgres_id"].(string)
	// assert.Equal(t, expectedName)

	// Lookup parameter values. All defined values are strings in the API call response
	// generalLogParameterValue := aws.GetParameterValueForParameterOfRdsInstance(t, "general_log", dbInstanceID, awsRegion)
	// allowSuspiciousUdfsParameterValue := aws.GetParameterValueForParameterOfRdsInstance(t, "allow-suspicious-udfs", dbInstanceID, awsRegion)

	// Lookup option values. All defined values are strings in the API call response
	// mariadbAuditPluginServerAuditEventsOptionValue := aws.GetOptionSettingForOfRdsInstance(t, "MARIADB_AUDIT_PLUGIN", "SERVER_AUDIT_EVENTS", dbInstanceID, awsRegion)

	// Verify that the address is not null
	assert.NotNil(t, address)
	// Verify that the DB instance is listening on the port mentioned
	assert.Equal(t, expectedPort, port)
	// Verify that the table/schema requested for creation is actually present in the database
	assert.True(t, schemaExistsInRdsInstance)
	// Booleans are (string) "0", "1"
	// assert.Equal(t, "0", generalLogParameterValue)
	// Values not set are "". This is custom behavior defined.
	// assert.Equal(t, "", allowSuspiciousUdfsParameterValue)
	// assert.Equal(t, "", mariadbAuditPluginServerAuditEventsOptionValue)
	// assert.Equal(t, "CONNECT", mariadbAuditPluginServerAuditEventsOptionValue)

}

// GetWhetherSchemaExistsInRdsPostgresInstance checks whether the specified schema/table name exists in the RDS instance
func GetWhetherSchemaExistsInRdsPostgresInstance(t terratest_testing.TestingT, dbUrl string, dbPort int64, dbUsername string, dbPassword string, expectedSchemaName string) bool {
	output, err := GetWhetherSchemaExistsInRdsPostgresInstanceE(t, dbUrl, dbPort, dbUsername, dbPassword, expectedSchemaName)
	if err != nil {
		t.Fatal(err)
	}
	return output
}

// GetWhetherSchemaExistsInRdsPostgresInstanceE checks whether the specified schema/table name exists in the RDS instance
func GetWhetherSchemaExistsInRdsPostgresInstanceE(t terratest_testing.TestingT, dbUrl string, dbPort int64, dbUsername string, dbPassword string, expectedSchemaName string) (bool, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", dbUrl, dbPort, dbUsername, dbPassword, expectedSchemaName)

	db, connErr := sql.Open("postgres", connectionString)
	if connErr != nil {
		return false, connErr
	}
	defer db.Close()
	var (
		schemaName string
	)
	sqlStatement := `SELECT "catalog_name" FROM "information_schema"."schemata" where catalog_name=$1`
	row := db.QueryRow(sqlStatement, expectedSchemaName)
	scanErr := row.Scan(&schemaName)
	if scanErr != nil {
		return false, scanErr
	}
	return true, nil
}
