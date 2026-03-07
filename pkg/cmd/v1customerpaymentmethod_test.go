// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/stigg-cli/internal/mocktest"
)

func TestV1CustomersPaymentMethodAttach(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:customers:payment-method", "attach",
			"--api-key", "string",
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
			t, pipeData, "v1:customers:payment-method", "attach",
			"--api-key", "string",
			"--id", "x",
		)
	})
}

func TestV1CustomersPaymentMethodDetach(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:customers:payment-method", "detach",
			"--api-key", "string",
			"--id", "x",
		)
	})
}
