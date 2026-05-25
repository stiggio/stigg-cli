// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stiggio/stigg-cli/internal/mocktest"
	"github.com/stiggio/stigg-cli/internal/requestflag"
)

func TestV1BetaEntitiesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1-beta:entities", "retrieve",
			"--id", "id",
			"--entity-id", "x",
		)
	})
}

func TestV1BetaEntitiesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1-beta:entities", "list",
			"--max-items", "10",
			"--id", "id",
			"--after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--include-archived", "true",
			"--limit", "1",
			"--type-ref-id", "typeRefId",
		)
	})
}

func TestV1BetaEntitiesArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1-beta:entities", "archive",
			"--id", "id",
			"--id", "user-7f3a0c1d",
			"--id", "user-c4d1b2e9",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"ids:\n" +
			"  - user-7f3a0c1d\n" +
			"  - user-c4d1b2e9\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1-beta:entities", "archive",
			"--id", "id",
		)
	})
}

func TestV1BetaEntitiesUnarchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1-beta:entities", "unarchive",
			"--id", "id",
			"--id", "user-7f3a0c1d",
			"--id", "user-c4d1b2e9",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"ids:\n" +
			"  - user-7f3a0c1d\n" +
			"  - user-c4d1b2e9\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1-beta:entities", "unarchive",
			"--id", "id",
		)
	})
}

func TestV1BetaEntitiesUpsert(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1-beta:entities", "upsert",
			"--id", "id",
			"--entity", "{id: user-7f3a0c1d, metadata: {email: jane@acme.com, role: admin}, typeRefId: user}",
			"--entity", "{id: user-c4d1b2e9, metadata: {email: john@acme.com}, typeRefId: user}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1BetaEntitiesUpsert)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1-beta:entities", "upsert",
			"--id", "id",
			"--entity.id", "user-7f3a0c1d",
			"--entity.metadata", "{email: jane@acme.com, role: admin}",
			"--entity.type-ref-id", "user",
			"--entity.id", "user-c4d1b2e9",
			"--entity.metadata", "{email: john@acme.com}",
			"--entity.type-ref-id", "user",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"entities:\n" +
			"  - id: user-7f3a0c1d\n" +
			"    metadata:\n" +
			"      email: jane@acme.com\n" +
			"      role: admin\n" +
			"    typeRefId: user\n" +
			"  - id: user-c4d1b2e9\n" +
			"    metadata:\n" +
			"      email: john@acme.com\n" +
			"    typeRefId: user\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1-beta:entities", "upsert",
			"--id", "id",
		)
	})
}
