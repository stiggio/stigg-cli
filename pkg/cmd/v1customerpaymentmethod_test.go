// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stiggio/stigg-cli/internal/mocktest"
)

func TestV1CustomersPaymentMethodAttach(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:customers:payment-method", "attach",
			"--id", "x",
			"--integration-id", "integrationId",
			"--payment-method-id", "paymentMethodId",
			"--vendor-identifier", "AUTH0",
			"--billing-currency", "usd",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"integrationId: integrationId\n" +
			"paymentMethodId: paymentMethodId\n" +
			"vendorIdentifier: AUTH0\n" +
			"billingCurrency: usd\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:customers:payment-method", "attach",
			"--id", "x",
		)
	})
}

func TestV1CustomersPaymentMethodDetach(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:customers:payment-method", "detach",
			"--id", "x",
		)
	})
}
