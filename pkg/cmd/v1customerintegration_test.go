// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stiggio/stigg-cli/internal/mocktest"
)

func TestV1CustomersIntegrationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:customers:integrations", "retrieve",
			"--id", "id",
			"--integration-id", "integrationId",
		)
	})
}

func TestV1CustomersIntegrationsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:customers:integrations", "update",
			"--id", "id",
			"--integration-id", "integrationId",
			"--synced-entity-id", "syncedEntityId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("syncedEntityId: syncedEntityId")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:customers:integrations", "update",
			"--id", "id",
			"--integration-id", "integrationId",
		)
	})
}

func TestV1CustomersIntegrationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:customers:integrations", "list",
			"--max-items", "10",
			"--id", "x",
			"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
			"--vendor-identifier", "AUTH0",
		)
	})
}

func TestV1CustomersIntegrationsLink(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:customers:integrations", "link",
			"--id", "x",
			"--id", "id",
			"--synced-entity-id", "syncedEntityId",
			"--vendor-identifier", "AUTH0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: id\n" +
			"syncedEntityId: syncedEntityId\n" +
			"vendorIdentifier: AUTH0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:customers:integrations", "link",
			"--id", "x",
		)
	})
}

func TestV1CustomersIntegrationsUnlink(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:customers:integrations", "unlink",
			"--id", "id",
			"--integration-id", "integrationId",
		)
	})
}
