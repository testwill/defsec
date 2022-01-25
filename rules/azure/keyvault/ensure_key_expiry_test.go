package keyvault

import (
	"testing"

	"github.com/aquasecurity/defsec/provider/azure/keyvault"
	"github.com/aquasecurity/defsec/rules"
        "github.com/aquasecurity/defsec/state"
	"github.com/stretchr/testify/assert"
)

func TestCheckEnsureKeyExpiry(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		name     string
		input    keyvault.KeyVault
		expected bool
	}{
		{
			name:     "positive result",
			input:    keyvault.KeyVault{},
			expected: true,
		},
		{
			name:     "negative result",
			input:    keyvault.KeyVault{},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.Azure.KeyVault = test.input
			results := CheckEnsureKeyExpiry.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() != rules.StatusPassed && result.Rule().LongID() == CheckEnsureKeyExpiry.Rule().LongID() {
					found = true
				}
			}
			if test.expected {
				assert.True(t, found, "Rule should have been found")
			} else {
				assert.False(t, found, "Rule should not have been found")
			}
		})
	}
}
