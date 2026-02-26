// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/stigg-cli/internal/mocktest"
	"github.com/stainless-sdks/stigg-cli/internal/requestflag"
)

func TestV1EventsPlansCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:events:plans", "create",
		"--id", "id",
		"--display-name", "displayName",
		"--product-id", "productId",
		"--billing-id", "billingId",
		"--default-trial-config", "{duration: 0, units: DAY, budget: {hasSoftLimit: true, limit: 0}, trialEndBehavior: CONVERT_TO_PAID}",
		"--description", "description",
		"--metadata", "{foo: string}",
		"--parent-plan-id", "parentPlanId",
		"--pricing-type", "FREE",
		"--status", "DRAFT",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1EventsPlansCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:events:plans", "create",
		"--id", "id",
		"--display-name", "displayName",
		"--product-id", "productId",
		"--billing-id", "billingId",
		"--default-trial-config.duration", "0",
		"--default-trial-config.units", "DAY",
		"--default-trial-config.budget", "{hasSoftLimit: true, limit: 0}",
		"--default-trial-config.trial-end-behavior", "CONVERT_TO_PAID",
		"--description", "description",
		"--metadata", "{foo: string}",
		"--parent-plan-id", "parentPlanId",
		"--pricing-type", "FREE",
		"--status", "DRAFT",
	)
}

func TestV1EventsPlansRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:events:plans", "retrieve",
		"--id", "x",
	)
}

func TestV1EventsPlansList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:events:plans", "list",
		"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--created-at", "{gt: '2019-12-27T18:11:19.117Z', gte: '2019-12-27T18:11:19.117Z', lt: '2019-12-27T18:11:19.117Z', lte: '2019-12-27T18:11:19.117Z'}",
		"--limit", "1",
		"--product-id", "productId",
		"--status", "status",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1EventsPlansList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:events:plans", "list",
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
