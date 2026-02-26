// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/stigg-cli/internal/mocktest"
	"github.com/stainless-sdks/stigg-cli/internal/requestflag"
)

func TestV1FeaturesArchiveFeature(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:features", "archive-feature",
		"--id", "x",
	)
}

func TestV1FeaturesCreateFeature(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:features", "create-feature",
		"--id", "id",
		"--display-name", "displayName",
		"--feature-type", "BOOLEAN",
		"--description", "description",
		"--enum-configuration", "{displayName: displayName, value: value}",
		"--feature-status", "NEW",
		"--feature-units", "featureUnits",
		"--feature-units-plural", "featureUnitsPlural",
		"--metadata", "{foo: string}",
		"--meter-type", "None",
		"--unit-transformation", "{divide: 0, featureUnits: featureUnits, featureUnitsPlural: featureUnitsPlural, round: UP}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1FeaturesCreateFeature)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:features", "create-feature",
		"--id", "id",
		"--display-name", "displayName",
		"--feature-type", "BOOLEAN",
		"--description", "description",
		"--enum-configuration.display-name", "displayName",
		"--enum-configuration.value", "value",
		"--feature-status", "NEW",
		"--feature-units", "featureUnits",
		"--feature-units-plural", "featureUnitsPlural",
		"--metadata", "{foo: string}",
		"--meter-type", "None",
		"--unit-transformation.divide", "0",
		"--unit-transformation.feature-units", "featureUnits",
		"--unit-transformation.feature-units-plural", "featureUnitsPlural",
		"--unit-transformation.round", "UP",
	)
}

func TestV1FeaturesListFeatures(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:features", "list-features",
		"--id", "id",
		"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--created-at", "{gt: '2019-12-27T18:11:19.117Z', gte: '2019-12-27T18:11:19.117Z', lt: '2019-12-27T18:11:19.117Z', lte: '2019-12-27T18:11:19.117Z'}",
		"--feature-type", "featureType",
		"--limit", "1",
		"--meter-type", "meterType",
		"--status", "status",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1FeaturesListFeatures)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:features", "list-features",
		"--id", "id",
		"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--created-at.gt", "2019-12-27T18:11:19.117Z",
		"--created-at.gte", "2019-12-27T18:11:19.117Z",
		"--created-at.lt", "2019-12-27T18:11:19.117Z",
		"--created-at.lte", "2019-12-27T18:11:19.117Z",
		"--feature-type", "featureType",
		"--limit", "1",
		"--meter-type", "meterType",
		"--status", "status",
	)
}

func TestV1FeaturesRetrieveFeature(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:features", "retrieve-feature",
		"--id", "x",
	)
}

func TestV1FeaturesUnarchiveFeature(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:features", "unarchive-feature",
		"--id", "x",
	)
}

func TestV1FeaturesUpdateFeature(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:features", "update-feature",
		"--id", "x",
		"--description", "description",
		"--display-name", "displayName",
		"--enum-configuration", "{displayName: displayName, value: value}",
		"--feature-units", "featureUnits",
		"--feature-units-plural", "featureUnitsPlural",
		"--metadata", "{foo: string}",
		"--meter", "{aggregation: {function: SUM, field: field}, filters: [{conditions: [{field: field, operation: EQUALS, value: value, values: [string]}]}]}",
		"--unit-transformation", "{divide: 0, featureUnits: featureUnits, featureUnitsPlural: featureUnitsPlural, round: UP}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1FeaturesUpdateFeature)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:features", "update-feature",
		"--id", "x",
		"--description", "description",
		"--display-name", "displayName",
		"--enum-configuration.display-name", "displayName",
		"--enum-configuration.value", "value",
		"--feature-units", "featureUnits",
		"--feature-units-plural", "featureUnitsPlural",
		"--metadata", "{foo: string}",
		"--meter.aggregation", "{function: SUM, field: field}",
		"--meter.filters", "[{conditions: [{field: field, operation: EQUALS, value: value, values: [string]}]}]",
		"--unit-transformation.divide", "0",
		"--unit-transformation.feature-units", "featureUnits",
		"--unit-transformation.feature-units-plural", "featureUnitsPlural",
		"--unit-transformation.round", "UP",
	)
}
