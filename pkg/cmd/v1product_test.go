// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stiggio/stigg-cli/internal/mocktest"
	"github.com/stiggio/stigg-cli/internal/requestflag"
)

func TestV1ProductsArchiveProduct(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:products", "archive-product",
			"--api-key", "string",
			"--id", "x",
		)
	})
}

func TestV1ProductsCreateProduct(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:products", "create-product",
			"--api-key", "string",
			"--id", "id",
			"--description", "description",
			"--display-name", "displayName",
			"--metadata", "{foo: string}",
			"--multiple-subscriptions=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: id\n" +
			"description: description\n" +
			"displayName: displayName\n" +
			"metadata:\n" +
			"  foo: string\n" +
			"multipleSubscriptions: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:products", "create-product",
			"--api-key", "string",
		)
	})
}

func TestV1ProductsDuplicateProduct(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:products", "duplicate-product",
			"--api-key", "string",
			"--id", "x",
			"--id", "id",
			"--description", "description",
			"--display-name", "displayName",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: id\n" +
			"description: description\n" +
			"displayName: displayName\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:products", "duplicate-product",
			"--api-key", "string",
			"--id", "x",
		)
	})
}

func TestV1ProductsListProducts(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:products", "list-products",
			"--api-key", "string",
			"--max-items", "10",
			"--id", "id",
			"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--created-at", "{gt: '2019-12-27T18:11:19.117Z', gte: '2019-12-27T18:11:19.117Z', lt: '2019-12-27T18:11:19.117Z', lte: '2019-12-27T18:11:19.117Z'}",
			"--limit", "1",
			"--status", "status",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1ProductsListProducts)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "v1:products", "list-products",
			"--api-key", "string",
			"--max-items", "10",
			"--id", "id",
			"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--created-at.gt", "2019-12-27T18:11:19.117Z",
			"--created-at.gte", "2019-12-27T18:11:19.117Z",
			"--created-at.lt", "2019-12-27T18:11:19.117Z",
			"--created-at.lte", "2019-12-27T18:11:19.117Z",
			"--limit", "1",
			"--status", "status",
		)
	})
}

func TestV1ProductsUnarchiveProduct(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:products", "unarchive-product",
			"--api-key", "string",
			"--id", "x",
		)
	})
}

func TestV1ProductsUpdateProduct(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:products", "update-product",
			"--api-key", "string",
			"--id", "x",
			"--description", "description",
			"--display-name", "displayName",
			"--metadata", "{foo: string}",
			"--multiple-subscriptions=true",
			"--product-settings", "{subscriptionCancellationTime: END_OF_BILLING_PERIOD, subscriptionEndSetup: DOWNGRADE_TO_FREE, subscriptionStartSetup: PLAN_SELECTION, downgradePlanId: downgradePlanId, prorateAtEndOfBillingPeriod: true, subscriptionStartPlanId: subscriptionStartPlanId}",
			"--usage-reset-cutoff-rule", "{behavior: NEVER_RESET}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1ProductsUpdateProduct)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "v1:products", "update-product",
			"--api-key", "string",
			"--id", "x",
			"--description", "description",
			"--display-name", "displayName",
			"--metadata", "{foo: string}",
			"--multiple-subscriptions=true",
			"--product-settings.subscription-cancellation-time", "END_OF_BILLING_PERIOD",
			"--product-settings.subscription-end-setup", "DOWNGRADE_TO_FREE",
			"--product-settings.subscription-start-setup", "PLAN_SELECTION",
			"--product-settings.downgrade-plan-id", "downgradePlanId",
			"--product-settings.prorate-at-end-of-billing-period=true",
			"--product-settings.subscription-start-plan-id", "subscriptionStartPlanId",
			"--usage-reset-cutoff-rule.behavior", "NEVER_RESET",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"displayName: displayName\n" +
			"metadata:\n" +
			"  foo: string\n" +
			"multipleSubscriptions: true\n" +
			"productSettings:\n" +
			"  subscriptionCancellationTime: END_OF_BILLING_PERIOD\n" +
			"  subscriptionEndSetup: DOWNGRADE_TO_FREE\n" +
			"  subscriptionStartSetup: PLAN_SELECTION\n" +
			"  downgradePlanId: downgradePlanId\n" +
			"  prorateAtEndOfBillingPeriod: true\n" +
			"  subscriptionStartPlanId: subscriptionStartPlanId\n" +
			"usageResetCutoffRule:\n" +
			"  behavior: NEVER_RESET\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:products", "update-product",
			"--api-key", "string",
			"--id", "x",
		)
	})
}
