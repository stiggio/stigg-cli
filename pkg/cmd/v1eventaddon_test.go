// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/stigg-cli/internal/mocktest"
	"github.com/stainless-sdks/stigg-cli/internal/requestflag"
)

func TestV1EventsAddonsArchiveAddon(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:events:addons", "archive-addon",
		"--id", "x",
	)
}

func TestV1EventsAddonsCreateAddon(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:events:addons", "create-addon",
		"--id", "id",
		"--display-name", "displayName",
		"--product-id", "productId",
		"--billing-id", "billingId",
		"--description", "description",
		"--max-quantity", "0",
		"--metadata", "{foo: string}",
		"--pricing-type", "FREE",
		"--status", "DRAFT",
	)
}

func TestV1EventsAddonsListAddons(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:events:addons", "list-addons",
		"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--created-at", "{gt: '2019-12-27T18:11:19.117Z', gte: '2019-12-27T18:11:19.117Z', lt: '2019-12-27T18:11:19.117Z', lte: '2019-12-27T18:11:19.117Z'}",
		"--limit", "1",
		"--product-id", "productId",
		"--status", "status",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1EventsAddonsListAddons)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:events:addons", "list-addons",
		"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--created-at.gt", "2019-12-27T18:11:19.117Z",
		"--created-at.gte", "2019-12-27T18:11:19.117Z",
		"--created-at.lt", "2019-12-27T18:11:19.117Z",
		"--created-at.lte", "2019-12-27T18:11:19.117Z",
		"--limit", "1",
		"--product-id", "productId",
		"--status", "status",
	)
}

func TestV1EventsAddonsPublishAddon(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:events:addons", "publish-addon",
		"--id", "x",
		"--migration-type", "NEW_CUSTOMERS",
	)
}

func TestV1EventsAddonsRetrieveAddon(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:events:addons", "retrieve-addon",
		"--id", "x",
	)
}

func TestV1EventsAddonsUpdateAddon(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:events:addons", "update-addon",
		"--id", "x",
		"--billing-id", "billingId",
		"--dependency", "[string]",
		"--description", "description",
		"--display-name", "displayName",
		"--max-quantity", "0",
		"--metadata", "{foo: string}",
	)
}
