package tests

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/retry"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTerraformCreateRDS(t *testing.T) {
	t.Parallel()

	// randomId := "qurwrs"
	randomId := strings.ToLower(random.UniqueId())
	namePrefix := fmt.Sprintf("%s-%s", "terratest-rds", randomId)
	expectedDBName := strings.ReplaceAll(namePrefix, "-", "_") // db names don't accept hyphens

	// Getting a random region between the US ones
	awsRegion := aws.GetRandomRegion(t, []string{"us-east-1", "us-east-2", "us-west-1", "us-west-2"}, nil)
	// awsRegion := "us-east-1"

	expectedUser := "tamruser"
	pw := "tamrpassword"
	expectedPort := int64(5432)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../test_examples/test_minimal",

		Vars: map[string]interface{}{
			"postgres_db_name":     expectedDBName,
			"parameter_group_name": fmt.Sprintf("%s-rds-postgres-pg", namePrefix),
			"name_prefix":          fmt.Sprintf("%s-", namePrefix),
			"pg_username":          expectedUser,
			"pg_password":          pw,
		},
		// Environment variables to set when running Terraform
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	ords := terraform.OutputAll(t, terraformOptions)

	oRDSInstanceID := ords["rds_postgres_id"].(string)
	oRDSSGIDs := ords["rds_security_group_ids"].([]interface{})
	oRDShostname := ords["rds_hostname"].(string)
	oRDSport := ords["rds_db_port"].(float64)
	oRDSuser := ords["rds_username"].(string)
	oDBName := ords["rds_dbname"].(string)

	// Fails test if Instance ID is nil
	require.NotNil(t, oRDSInstanceID)

	// Information in RDS API can take more than 20 mins to be available. We retry for 40mins before proceeding
	rdsObj := retry.DoWithRetryInterface(t, "Waiting RDS API to be available", 20, 2*time.Minute, func() (interface{}, error) {
		return aws.GetRdsInstanceDetailsE(t, oRDSInstanceID, awsRegion)
	}).(*rds.DBInstance)

	// Verify that the address is not null and equal to output
	address := aws.GetAddressOfRdsInstance(t, oRDSInstanceID, awsRegion)
	assert.NotNil(t, address)
	assert.Equal(t, oRDShostname, address)

	// Verify that the DB instance is listening on the expected port and equal to output
	port := aws.GetPortOfRdsInstance(t, oRDSInstanceID, awsRegion)
	assert.Equal(t, expectedPort, port)
	assert.Equal(t, int64(oRDSport), port)

	// Verify Sec Group IDs output is not nil
	assert.NotNil(t, oRDSSGIDs)

	// Verify that user is the same as expected and equal to output
	assert.Equal(t, expectedUser, *rdsObj.MasterUsername)
	assert.NotNil(t, oRDSuser)

	// Verify that user is the same as expected and equal to output
	assert.Equal(t, expectedDBName, *rdsObj.DBName)
	assert.Equal(t, oDBName, *rdsObj.DBName)

}
