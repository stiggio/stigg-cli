// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/stigg-cli/internal/mocktest"
)

func TestV1EventsAddonsDraftCreateAddonDraft(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:events:addons:draft", "create-addon-draft",
		"--id", "x",
	)
}

func TestV1EventsAddonsDraftRemoveAddonDraft(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"v1:events:addons:draft", "remove-addon-draft",
		"--id", "x",
	)
}
