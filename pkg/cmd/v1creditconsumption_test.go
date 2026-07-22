// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stiggio/stigg-cli/internal/mocktest"
	"github.com/stiggio/stigg-cli/internal/requestflag"
)

func TestV1CreditsConsumptionConsume(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:credits:consumption", "consume",
			"--amount", "1",
			"--currency-id", "currencyId",
			"--customer-id", "customerId",
			"--idempotency-key", "x",
			"--created-at", "'2019-12-27T18:11:19.117Z'",
			"--dimensions", "{foo: string}",
			"--resource-id", "resourceId",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"amount: 1\n" +
			"currencyId: currencyId\n" +
			"customerId: customerId\n" +
			"idempotencyKey: x\n" +
			"createdAt: '2019-12-27T18:11:19.117Z'\n" +
			"dimensions:\n" +
			"  foo: string\n" +
			"resourceId: resourceId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:credits:consumption", "consume",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})
}

func TestV1CreditsConsumptionConsumeAsync(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:credits:consumption", "consume-async",
			"--consumption", "{amount: 1, currencyId: currencyId, customerId: customerId, idempotencyKey: x, createdAt: '2019-12-27T18:11:19.117Z', dimensions: {foo: string}, resourceId: resourceId}",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1CreditsConsumptionConsumeAsync)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:credits:consumption", "consume-async",
			"--consumption.amount", "1",
			"--consumption.currency-id", "currencyId",
			"--consumption.customer-id", "customerId",
			"--consumption.idempotency-key", "x",
			"--consumption.created-at", "2019-12-27T18:11:19.117Z",
			"--consumption.dimensions", "{foo: string}",
			"--consumption.resource-id", "resourceId",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"consumptions:\n" +
			"  - amount: 1\n" +
			"    currencyId: currencyId\n" +
			"    customerId: customerId\n" +
			"    idempotencyKey: x\n" +
			"    createdAt: '2019-12-27T18:11:19.117Z'\n" +
			"    dimensions:\n" +
			"      foo: string\n" +
			"    resourceId: resourceId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:credits:consumption", "consume-async",
			"--x-account-id", "X-ACCOUNT-ID",
			"--x-environment-id", "X-ENVIRONMENT-ID",
		)
	})
}
