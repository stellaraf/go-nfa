package nfa_test

import (
	"testing"

	"github.com/stellaraf/go-nfa"
	"github.com/stellaraf/go-nfa/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	t.Parallel()
	u, err := nfa.ParseURL(test.Env.URL)
	require.NoError(t, err)
	client, err := nfa.New(nfa.Username(test.Env.Username), nfa.Password(test.Env.Password), nfa.URL(u))
	require.NoError(t, err)
	require.NotNil(t, client)
}

func TestClient_PercentileQuery(t *testing.T) {
	t.Parallel()
	u, err := nfa.ParseURL(test.Env.URL)
	require.NoError(t, err)
	client, err := nfa.New(nfa.Username(test.Env.Username), nfa.Password(test.Env.Password), nfa.URL(u))
	require.NoError(t, err)
	require.NotNil(t, client)

	res, err := client.PercentileQuery(test.Env.Test.Prefixes, test.Env.Test.Exclude, 95)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, res.PercentileValue, uint64(100))
}
