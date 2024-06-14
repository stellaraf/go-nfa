package flow_test

import (
	"encoding/json"
	"testing"

	"github.com/stellaraf/go-nfa/flow"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Filter(t *testing.T) {
	t.Parallel()
	f := flow.NewFilter(flow.AND)
	f = flow.AddRule(f, flow.IPVersion, flow.EQ, "4")
	b, err := json.Marshal(&f)
	require.NoError(t, err)
	exp := `{"condition":"and","rules":[{"comparisonOperator":"eq","key":"ip-version","value":"4"}]}`
	assert.Equal(t, exp, string(b))
}
