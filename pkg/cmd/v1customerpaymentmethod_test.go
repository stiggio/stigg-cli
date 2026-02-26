// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/stigg-cli/internal/mocktest"
)

func TestV1CustomersPaymentMethodAttach(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:customers:payment-method", "attach",
		"--id", "x",
		"--integration-id", "integrationId",
		"--payment-method-id", "paymentMethodId",
		"--vendor-identifier", "AUTH0",
		"--billing-currency", "usd",
	)
}

func TestV1CustomersPaymentMethodDetach(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:customers:payment-method", "detach",
		"--id", "x",
	)
}
