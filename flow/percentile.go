package flow

import (
	"encoding/json"
	"fmt"
	"time"
)

type QueryResults struct {
	Code string `json:"code"`
}

type PercentileData struct {
	Timestamp time.Time
	Data      float64
}

func (p *PercentileData) UnmarshalJSON(b []byte) error {
	sl := make([]any, 2)
	err := json.Unmarshal(b, &sl)
	if err != nil {
		return err
	}
	tr := sl[0].(string)
	d, ok := sl[1].(float64)
	if !ok {
		return fmt.Errorf("failed to parse value '%v' as float64", sl[1])
	}
	t, err := time.Parse(time.RFC3339Nano, tr)
	if err != nil {
		return err
	}
	p.Timestamp = t
	p.Data = d
	return nil
}

type ResPercentileQuery struct {
	Results         []QueryResults   `json:"results"`
	PercentileValue float64          `json:"percentileValue"`
	QueryParams     map[string]any   `json:"queryParams"`
	Title           string           `json:"title"`
	Header          []string         `json:"header"`
	Step            int              `json:"step"`
	Data            []PercentileData `json:"data"`
}
