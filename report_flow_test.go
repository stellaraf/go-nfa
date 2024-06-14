package nfa_test

import (
	"encoding/json"
	"testing"

	"github.com/stellaraf/go-nfa"
	"github.com/stellaraf/go-nfa/flow"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func Test_PercentileQuery(t *testing.T) {
	t.Parallel()
	q, err := nfa.NewPercentileQuery([]string{"192.0.2.0/24"}, []string{"10.0.0.0/8"}, 95)
	require.NoError(t, err)
	b, err := json.Marshal(&q)
	require.NoError(t, err)
	var um *flow.QueryParameters
	err = json.Unmarshal(b, &um)
	require.NoError(t, err)
	assert.Equal(t, flow.Octets, um.AggregateBy)
	gj := gjson.ParseBytes(b)
	inc := gj.Get("filters.rules.0.rules.0.value").String()
	exc := gj.Get("filters.rules.0.rules.1.value").String()
	assert.Equal(t, "[\"192.0.2.0-192.0.2.255\"]", inc)
	assert.Equal(t, "[\"10.0.0.0-10.255.255.255\"]", exc)
}

func Test_CreateRange(t *testing.T) {
	t.Run("v4", func(t *testing.T) {
		r, err := nfa.CreatePrefixRange("192.0.2.0/24")
		require.NoError(t, err)
		assert.Equal(t, "192.0.2.0-192.0.2.255", r)
	})
	t.Run("v6", func(t *testing.T) {
		r, err := nfa.CreatePrefixRange("2001:db8::/32")
		require.NoError(t, err)
		assert.Equal(t, "2001:db8::-2001:db8:ffff:ffff:ffff:ffff:ffff:ffff", r)
	})
}
