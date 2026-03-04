// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/stigg-cli/internal/mocktest"
)

func TestV1SubscriptionsInvoiceMarkAsPaid(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:subscriptions:invoice", "mark-as-paid",
		"--api-key", "string",
		"--id", "x",
	)
}
