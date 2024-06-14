package flow_test

import (
	"testing"

	"github.com/stellar/go-nfa/flow"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueryParameters_Encode(t *testing.T) {
	q := &flow.QueryParameters{
		AggregateBy: flow.Octets,
	}
	f := flow.NewFilter(flow.AND)
	f = flow.AddRule(f, flow.IPVersion, flow.EQ, "4")
	q.Filters = f
	e, err := q.MarshalQuery()
	require.NoError(t, err)
	exp := map[string]string{
		"aggregate-by": "octets",
		"filters":      `{"condition":"and","rules":[{"comparisonOperator":"eq","key":"ip-version","value":"4"}]}`,
	}
	delete(e, "start-time")
	delete(e, "end-time")
	assert.Equal(t, exp, e)
}
