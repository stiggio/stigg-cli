// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/stigg-cli/internal/mocktest"
	"github.com/stainless-sdks/stigg-cli/internal/requestflag"
)

func TestV1EventsReport(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:events", "report",
		"--api-key", "string",
		"--event", "{customerId: customerId, eventName: x, idempotencyKey: x, dimensions: {foo: string}, resourceId: resourceId, timestamp: '2019-12-27T18:11:19.117Z'}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1EventsReport)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:events", "report",
		"--event.customer-id", "customerId",
		"--event.event-name", "x",
		"--event.idempotency-key", "x",
		"--event.dimensions", "{foo: string}",
		"--event.resource-id", "resourceId",
		"--event.timestamp", "2019-12-27T18:11:19.117Z",
	)
}
