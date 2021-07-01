package tests

import (
	"database/sql"
	"fmt"

	terratest_testing "github.com/gruntwork-io/terratest/modules/testing"
)

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
