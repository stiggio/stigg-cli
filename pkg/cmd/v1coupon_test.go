// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/stigg-cli/internal/mocktest"
	"github.com/stainless-sdks/stigg-cli/internal/requestflag"
)

func TestV1CouponsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:coupons", "create",
		"--api-key", "string",
		"--id", "id",
		"--amounts-off", "[{amount: 0, currency: usd}]",
		"--description", "description",
		"--duration-in-months", "1",
		"--metadata", "{foo: string}",
		"--name", "name",
		"--percent-off", "1",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1CouponsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:coupons", "create",
		"--api-key", "string",
		"--id", "id",
		"--amounts-off.amount", "0",
		"--amounts-off.currency", "usd",
		"--description", "description",
		"--duration-in-months", "1",
		"--metadata", "{foo: string}",
		"--name", "name",
		"--percent-off", "1",
	)
}

func TestV1CouponsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:coupons", "retrieve",
		"--api-key", "string",
		"--id", "x",
	)
}

func TestV1CouponsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:coupons", "list",
		"--api-key", "string",
		"--id", "id",
		"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--created-at", "{gt: '2019-12-27T18:11:19.117Z', gte: '2019-12-27T18:11:19.117Z', lt: '2019-12-27T18:11:19.117Z', lte: '2019-12-27T18:11:19.117Z'}",
		"--limit", "1",
		"--status", "status",
		"--type", "FIXED",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1CouponsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:coupons", "list",
		"--api-key", "string",
		"--id", "id",
		"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--created-at.gt", "2019-12-27T18:11:19.117Z",
		"--created-at.gte", "2019-12-27T18:11:19.117Z",
		"--created-at.lt", "2019-12-27T18:11:19.117Z",
		"--created-at.lte", "2019-12-27T18:11:19.117Z",
		"--limit", "1",
		"--status", "status",
		"--type", "FIXED",
	)
}

func TestV1CouponsArchiveCoupon(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:coupons", "archive-coupon",
		"--api-key", "string",
		"--id", "x",
	)
}

func TestV1CouponsUpdateCoupon(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:coupons", "update-coupon",
		"--api-key", "string",
		"--id", "x",
		"--description", "description",
		"--metadata", "{foo: string}",
		"--name", "name",
	)
}
