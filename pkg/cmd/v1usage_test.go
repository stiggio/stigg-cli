// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/stigg-cli/internal/mocktest"
	"github.com/stainless-sdks/stigg-cli/internal/requestflag"
)

func TestV1UsageHistory(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:usage", "history",
			"--api-key", "string",
			"--customer-id", "customerId",
			"--feature-id", "featureId",
			"--start-date", "'2019-12-27T18:11:19.117Z'",
			"--end-date", "'2019-12-27T18:11:19.117Z'",
			"--group-by", "groupBy",
			"--include-historical-usage=true",
			"--resource-id", "resourceId",
		)
	})
}

func TestV1UsageReport(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:usage", "report",
			"--api-key", "string",
			"--usage", "{customerId: customerId, featureId: featureId, value: -9007199254740991, createdAt: '2019-12-27T18:11:19.117Z', dimensions: {foo: string}, resourceId: resourceId, updateBehavior: DELTA}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1UsageReport)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "v1:usage", "report",
			"--api-key", "string",
			"--usage.customer-id", "customerId",
			"--usage.feature-id", "featureId",
			"--usage.value", "-9007199254740991",
			"--usage.created-at", "2019-12-27T18:11:19.117Z",
			"--usage.dimensions", "{foo: string}",
			"--usage.resource-id", "resourceId",
			"--usage.update-behavior", "DELTA",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"usages:\n" +
			"  - customerId: customerId\n" +
			"    featureId: featureId\n" +
			"    value: -9007199254740991\n" +
			"    createdAt: '2019-12-27T18:11:19.117Z'\n" +
			"    dimensions:\n" +
			"      foo: string\n" +
			"    resourceId: resourceId\n" +
			"    updateBehavior: DELTA\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:usage", "report",
			"--api-key", "string",
		)
	})
}
