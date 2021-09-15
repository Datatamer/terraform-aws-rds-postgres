package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

const expectedPw = "tamrpassword"

func initTestCases() []RdsTestCase {
	return []RdsTestCase{
		{
			testName:         "test1",
			expectApplyError: false,
			vars: map[string]interface{}{
				"pg_password":          expectedPw,
				"pg_username":          "",
				"postgres_db_name":     "",
				"parameter_group_name": "",
				"name_prefix":          "",
			},
		},
	}
}
func TestTerraformCreateRDS(t *testing.T) {

	testCases := initTestCases()

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.testName, func(t *testing.T) {
			t.Parallel()

			// These will create a tempTestFolder for each bucketTestCase.
			tempTestFolder := test_structure.CopyTerraformFolderToTemp(t, "..", "test_examples/minimal")

			// this stage will generate a random `awsRegion` and a `uniqueId` to be used in tests.
			test_structure.RunTestStage(t, "pick_new_randoms", func() {
				usRegions := []string{"us-east-1", "us-east-2", "us-west-1", "us-west-2"}
				// This function will first check for the Env Var TERRATEST_REGION and return its value if != ""
				awsRegion := aws.GetRandomStableRegion(t, usRegions, nil)

				test_structure.SaveString(t, tempTestFolder, "region", awsRegion)
				test_structure.SaveString(t, tempTestFolder, "unique_id", strings.ToLower(random.UniqueId()))
			})

			defer test_structure.RunTestStage(t, "teardown", func() {
				teraformOptions := test_structure.LoadTerraformOptions(t, tempTestFolder)
				terraform.Destroy(t, teraformOptions)
			})

			test_structure.RunTestStage(t, "setup_options", func() {
				awsRegion := test_structure.LoadString(t, tempTestFolder, "region")
				uniqueID := test_structure.LoadString(t, tempTestFolder, "unique_id")

				testCase.vars["parameter_group_name"] = fmt.Sprintf("terratest-pg-%s", uniqueID)
				testCase.vars["name_prefix"] = fmt.Sprintf("terratest-%s", uniqueID)
				testCase.vars["pg_username"] = uniqueID
				testCase.vars["postgres_db_name"] = fmt.Sprintf("db_%s", uniqueID)

				terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
					TerraformDir: tempTestFolder,
					Vars:         testCase.vars,
					EnvVars: map[string]string{
						"AWS_REGION": awsRegion,
					},
				})

				test_structure.SaveTerraformOptions(t, tempTestFolder, terraformOptions)
			})

			test_structure.RunTestStage(t, "create_rds", func() {
				terraformOptions := test_structure.LoadTerraformOptions(t, tempTestFolder)
				terraform.InitAndApply(t, terraformOptions)
			})

			test_structure.RunTestStage(t, "validate", func() {
				awsRegion := test_structure.LoadString(t, tempTestFolder, "region")
				terraformOptions := test_structure.LoadTerraformOptions(t, tempTestFolder)
				validateModuleOutputs(t,
					terraformOptions,
					awsRegion,
					int64(5432),
					testCase.vars["pg_username"].(string),
					testCase.vars["postgres_db_name"].(string),
				)
			})
		})
	}
}
