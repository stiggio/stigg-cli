// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/stigg-cli/internal/mocktest"
	"github.com/stainless-sdks/stigg-cli/internal/requestflag"
)

func TestV1CustomersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:customers", "retrieve",
		"--id", "x",
	)
}

func TestV1CustomersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:customers", "update",
		"--id", "x",
		"--coupon-id", "couponId",
		"--email", "dev@stainless.com",
		"--integration", "{id: id, syncedEntityId: syncedEntityId, vendorIdentifier: AUTH0}",
		"--metadata", "{foo: string}",
		"--name", "name",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1CustomersUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:customers", "update",
		"--id", "x",
		"--coupon-id", "couponId",
		"--email", "dev@stainless.com",
		"--integration.id", "id",
		"--integration.synced-entity-id", "syncedEntityId",
		"--integration.vendor-identifier", "AUTH0",
		"--metadata", "{foo: string}",
		"--name", "name",
	)
}

func TestV1CustomersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:customers", "list",
		"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--created-at", "{gt: '2019-12-27T18:11:19.117Z', gte: '2019-12-27T18:11:19.117Z', lt: '2019-12-27T18:11:19.117Z', lte: '2019-12-27T18:11:19.117Z'}",
		"--email", "email",
		"--limit", "1",
		"--name", "name",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1CustomersList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:customers", "list",
		"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--created-at.gt", "2019-12-27T18:11:19.117Z",
		"--created-at.gte", "2019-12-27T18:11:19.117Z",
		"--created-at.lt", "2019-12-27T18:11:19.117Z",
		"--created-at.lte", "2019-12-27T18:11:19.117Z",
		"--email", "email",
		"--limit", "1",
		"--name", "name",
	)
}

func TestV1CustomersArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:customers", "archive",
		"--id", "x",
	)
}

func TestV1CustomersImport(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:customers", "import",
		"--customer", "{id: id, email: dev@stainless.com, name: name, metadata: {foo: string}, paymentMethodId: paymentMethodId, updatedAt: '2019-12-27T18:11:19.117Z'}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1CustomersImport)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:customers", "import",
		"--customer.id", "id",
		"--customer.email", "dev@stainless.com",
		"--customer.name", "name",
		"--customer.metadata", "{foo: string}",
		"--customer.payment-method-id", "paymentMethodId",
		"--customer.updated-at", "2019-12-27T18:11:19.117Z",
	)
}

func TestV1CustomersListResources(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:customers", "list-resources",
		"--id", "x",
		"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--limit", "1",
	)
}

func TestV1CustomersProvision(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:customers", "provision",
		"--id", "id",
		"--coupon-id", "couponId",
		"--default-payment-method", "{billingId: billingId, cardExpiryMonth: 0, cardExpiryYear: 0, cardLast4Digits: cardLast4Digits, type: CARD}",
		"--email", "dev@stainless.com",
		"--integration", "{id: id, syncedEntityId: syncedEntityId, vendorIdentifier: AUTH0}",
		"--metadata", "{foo: string}",
		"--name", "name",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(v1CustomersProvision)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:customers", "provision",
		"--id", "id",
		"--coupon-id", "couponId",
		"--default-payment-method.billing-id", "billingId",
		"--default-payment-method.card-expiry-month", "0",
		"--default-payment-method.card-expiry-year", "0",
		"--default-payment-method.card-last4-digits", "cardLast4Digits",
		"--default-payment-method.type", "CARD",
		"--email", "dev@stainless.com",
		"--integration.id", "id",
		"--integration.synced-entity-id", "syncedEntityId",
		"--integration.vendor-identifier", "AUTH0",
		"--metadata", "{foo: string}",
		"--name", "name",
	)
}

func TestV1CustomersUnarchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:customers", "unarchive",
		"--id", "x",
	)
}
