// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stiggio/stigg-cli/internal/mocktest"
)

func TestInternalBetaEventQueuesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"internal:beta:event-queues", "retrieve",
			"--queue-name", "x",
		)
	})
}

func TestInternalBetaEventQueuesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"internal:beta:event-queues", "update",
			"--queue-name", "x",
			"--allowed-assume-role-arn", "string",
			"--create-low-priority-queues=true",
			"--event-type", "MEMBER_INVITED",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"allowedAssumeRoleArns:\n" +
			"  - string\n" +
			"createLowPriorityQueues: true\n" +
			"eventTypes:\n" +
			"  - MEMBER_INVITED\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"internal:beta:event-queues", "update",
			"--queue-name", "x",
		)
	})
}

func TestInternalBetaEventQueuesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"internal:beta:event-queues", "list",
		)
	})
}

func TestInternalBetaEventQueuesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"internal:beta:event-queues", "delete",
			"--queue-name", "x",
		)
	})
}

func TestInternalBetaEventQueuesProvision(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"internal:beta:event-queues", "provision",
			"--region", "us-east-1",
			"--allowed-assume-role-arn", "string",
			"--create-low-priority-queues=true",
			"--event-type", "MEMBER_INVITED",
			"--suffix", "suffix",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"region: us-east-1\n" +
			"allowedAssumeRoleArns:\n" +
			"  - string\n" +
			"createLowPriorityQueues: true\n" +
			"eventTypes:\n" +
			"  - MEMBER_INVITED\n" +
			"suffix: suffix\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"internal:beta:event-queues", "provision",
		)
	})
}
