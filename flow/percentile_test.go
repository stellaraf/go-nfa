package flow_test

import (
	"encoding/json"
	"testing"

	"github.com/stellar/go-nfa/flow"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPercentileData_UnmarshalJSON(t *testing.T) {
	t.Parallel()
	type S struct {
		Data []flow.PercentileData `json:"data"`
	}
	var s *S
	r := []byte(`{"data": [["2024-06-04T00:00:00Z",4638717.514563107],["2024-06-04T00:10:00Z",4239642.642533937]]}`)
	err := json.Unmarshal(r, &s)
	require.NoError(t, err)
	assert.Equal(t, s.Data[0].Data, 4638717.514563107)
}
