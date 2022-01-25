package kinesis

import (
	"testing"

	"github.com/aquasecurity/defsec/provider/aws/kinesis"
	"github.com/aquasecurity/defsec/rules"
        "github.com/aquasecurity/defsec/state"
	"github.com/stretchr/testify/assert"
)

func TestCheckEnableInTransitEncryption(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		name     string
		input    kinesis.Kinesis
		expected bool
	}{
		{
			name:     "positive result",
			input:    kinesis.Kinesis{},
			expected: true,
		},
		{
			name:     "negative result",
			input:    kinesis.Kinesis{},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.AWS.Kinesis = test.input
			results := CheckEnableInTransitEncryption.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() != rules.StatusPassed && result.Rule().LongID() == CheckEnableInTransitEncryption.Rule().LongID() {
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
