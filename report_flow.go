package nfa

import (
	"fmt"
	"net"
	"net/netip"

	"github.com/c-robinson/iplib/v2"
	"github.com/golang-module/carbon/v2"
	"github.com/stellaraf/go-nfa/flow"
)

func CreatePrefixRange(prefix string) (string, error) {
	n, err := netip.ParsePrefix(prefix)
	if err != nil {
		return "", err
	}
	if n.Addr().Is6() {
		ip := net.IP(n.Addr().AsSlice())
		pfx := iplib.NewNet6(ip, n.Bits(), 0)
		start := pfx.IP().String()
		end := pfx.LastAddress().String()
		return fmt.Sprintf("%s-%s", start, end), nil
	}
	ip := net.IP(n.Addr().AsSlice())
	pfx := iplib.NewNet4(ip, n.Bits())
	start := pfx.IP().String()
	end := pfx.BroadcastAddress().String()
	return fmt.Sprintf("%s-%s", start, end), nil
}

func NewPercentileQuery(prefixes []string, exclusions []string, p uint8) (*flow.QueryParameters, error) {
	now := carbon.Now().SetTimezone("Etc/UTC").EndOfDay()
	lastMonth := now.SubMonth().StartOfDay()

	include := make([]string, 0)
	exclude := make([]string, 0)

	for _, prefix := range prefixes {
		r, err := CreatePrefixRange(prefix)
		if err != nil {
			return nil, err
		}
		include = append(include, r)
	}

	for _, prefix := range exclusions {
		r, err := CreatePrefixRange(prefix)
		if err != nil {
			return nil, err
		}
		exclude = append(exclude, r)
	}

	q := &flow.QueryParameters{
		StartTime:         lastMonth.StdTime(),
		EndTime:           now.StdTime(),
		Percentile:        p,
		AggregateBy:       flow.Octets,
		AggregateFunction: flow.Avg,
		GroupBy:           flow.TS,
	}
	srcFilter := flow.NewSubFilter(flow.AND)
	srcFilter = flow.AddSubRule(srcFilter, flow.SrcAddr, flow.EQ, include)
	srcFilter = flow.AddSubRule(srcFilter, flow.DstAddr, flow.NEQ, exclude)

	dstFilter := flow.NewSubFilter(flow.AND)
	dstFilter = flow.AddSubRule(dstFilter, flow.DstAddr, flow.EQ, include)
	dstFilter = flow.AddSubRule(dstFilter, flow.SrcAddr, flow.NEQ, exclude)

	filter := flow.NewFilter(flow.OR)
	filter = flow.AddSubFilterRule(filter, srcFilter)
	filter = flow.AddSubFilterRule(filter, dstFilter)
	q.Filters = filter
	return q, nil
}
