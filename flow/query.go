package flow

import (
	"encoding/json"
	"time"
)

type QueryParameters struct {
	// required
	StartTime time.Time `json:"start-time"`

	// required
	EndTime time.Time `json:"end-time"`

	Filters *Filter `json:"filters"`

	// ts, ip-version, src-addr, dst-addr, tos, protocol, src-port, dst-port, input-interface,
	// output-interface, src-mask, dst-mask, src-as, dst-as, next-hop, src-as-path, dst-as-path,
	// src-country, dst-country, src-city, dst-city, exporter-ip, exporter-id, exporter-as,
	// bgp-localpref, bgp-med, l2-mac-src-addr, l2-mac-dst-addr, l2-ethernet-type,
	// mpls-top-label-stack-section, mpls-top-label-stack-section1, mpls-top-label-stack-section2,
	// mpls-top-label-stack-section3, mpls-top-label-stack-section4, mpls-top-label-stack-section5,
	// mpls-top-label-stack-section6, mpls-top-label-stack-section7, mpls-top-label-stack-section8,
	// mpls-top-label-stack-section9, mpls-top-label-type, mpls-top-label-ipv4,
	// mpls-vpn-route-distinguisher, mpls-top-label-prefix-length, mpls-top-label-ipv6,
	// mpls-top-label-ttl, mpls-label-stack-length, mpls-label-stack-depth, mpls-top-label-exp,
	// mpls-post-top-label-exp, pseudo-wire-id, pseudo-wire-type, pseudo-wire-control-word,
	// community, l3-ip-min-total-length, l3-ip-max-total-length, l3-ip-total-length,
	// l3-ip-min-ttl, l3-ip-max-ttl, l3-ip-ttl, l4-tcp-flag, flow-role, application-name
	GroupBy string `json:"groupby,omitempty"`

	// ts, ip-version, src-addr, dst-addr, tos, protocol, src-port, dst-port,
	// input-interface, output-interface, src-mask, dst-mask, src-as, dst-as,
	// next-hop, octets, packets, flows, src-vlan, dst-vlan, dst-as-path, src-as-path,
	// src-country, dst-country, src-city, dst-city, bgp-localpref, bgp-med, exporter-ip,
	// exporter-id, exporter-as, l2-mac-src-addr, l2-mac-dst-addr, l2-ethernet-type,
	// mpls-top-label-stack-section, mpls-top-label-stack-section1, mpls-top-label-stack-section2,
	// mpls-top-label-stack-section3, mpls-top-label-stack-section4, mpls-top-label-stack-section5,
	// mpls-top-label-stack-section6, mpls-top-label-stack-section7, mpls-top-label-stack-section8,
	// mpls-top-label-stack-section9, mpls-top-label-type, mpls-top-label-ipv4,
	// mpls-vpn-route-distinguisher, mpls-top-label-prefix-length, mpls-top-label-ipv6,
	// mpls-top-label-ttl, mpls-label-stack-length, mpls-label-stack-depth, mpls-top-label-exp,
	// mpls-post-top-label-exp, pseudo-wire-id, pseudo-wire-type, pseudo-wire-control-word,
	// community, l3-ip-min-total-length, l3-ip-max-total-length, l3-ip-total-length, l3-ip-min-ttl,
	// l3-ip-max-ttl, l3-ip-ttl, l4-tcp-syn-count, l4-tcp-fin-count, l4-tcp-rst-count,
	// l4-tcp-psh-count, l4-tcp-ack-count, l4-tcp-urg-count, flow-role, application-name
	TableColumns []string `json:"table-columns,omitempty"`

	// octets, packets, flows, l4-tcp-flag-count
	AggregateBy string `json:"aggregate-by,omitempty"`

	// count, min, max, avg, uniq, sum
	AggregateFunction string `json:"aggregate-function,omitempty"`

	// ts, ip-version, src-addr, dst-addr, tos, protocol, src-port, dst-port, input-interface,
	// output-interface, src-mask, dst-mask, src-as, dst-as, next-hop, src-as-path, dst-as-path,
	// octets, packets, exporter-ip, exporter-id, exporter-as, flows, bgp-localpref, bgp-med,
	// l2-mac-src-addr, l2-mac-dst-addr, l2-ethernet-type, mpls-top-label-stack-section,
	// mpls-top-label-stack-section1, mpls-top-label-stack-section2, mpls-top-label-stack-section3,
	// mpls-top-label-stack-section4, mpls-top-label-stack-section5, mpls-top-label-stack-section6,
	// mpls-top-label-stack-section7, mpls-top-label-stack-section8, mpls-top-label-stack-section9,
	// mpls-top-label-type, mpls-top-label-ipv4, mpls-vpn-route-distinguisher,
	// mpls-top-label-prefix-length, mpls-top-label-ipv6, mpls-top-label-ttl,
	// mpls-label-stack-length, mpls-label-stack-depth, mpls-top-label-exp,
	// mpls-post-top-label-exp, pseudo-wire-id, pseudo-wire-type, pseudo-wire-control-word,
	// community, l3-ip-min-total-length, l3-ip-max-total-length, l3-ip-total-length,
	// l3-ip-min-ttl, l3-ip-max-ttl, l3-ip-ttl, l4-tcp-flag-count, flow-role, application-name
	OrderBy string `json:"order-by,omitempty"`

	// descending, ascending
	Order string `json:"order,omitempty"`

	Top      uint32 `json:"top,omitempty"`
	Page     uint32 `json:"page,omitempty"`
	PageSize uint16 `json:"page-size,omitempty"`

	// location, device, interface
	FiltersNarrowBy []string `json:"filters-narrow-by,omitempty"`

	// seconds
	RateUnit string `json:"rate-unit,omitempty"`

	IsolateEmptySpaces  bool  `json:"isolate-empty-spaces,omitempty"`
	GroupBySrcPrefix    bool  `json:"group-by-src-prefix,omitempty"`
	GroupBySrcPrefixLen bool  `json:"group-by-src-prefix-length,omitempty"`
	GroupByDstPrefix    bool  `json:"group-by-dst-prefix,omitempty"`
	GroupByDstPrefixLen bool  `json:"group-by-dst-prefix-length,omitempty"`
	ASPathPrependRemove bool  `json:"as-path-prepend-remove,omitempty"`
	Percentile          uint8 `json:"percentile,omitempty"`
	Dictionary          bool  `json:"dictionary,omitempty"`
	TableMinMax         bool  `json:"table-min-max,omitempty"`

	// addr, port, as, country, city, l2-mac-addr
	GroupByBidirectional string `json:"groupby-bidirectional,omitempty"`

	GroupByTSBidirectional bool `json:"groupby-ts-bidirectional,omitempty"`

	// ts, octets, flows, packets, l4-tcp-flag-count, bidirectional-left, bidirectional-right
	OrderByBidirectional string `json:"orderby-bidirectional,omitempty"`

	// as-ingress, as-egress
	TrafficDirection string `json:"traffic-direction,omitempty"`

	// minimum 60
	Step uint32 `json:"step,omitempty"`
}

func (q *QueryParameters) MarshalQuery() (map[string]string, error) {
	var m map[string]any
	b, err := json.Marshal(&q)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	ms := make(map[string]string, len(m))
	for k, v := range m {
		s, ok := v.(string)
		if ok {
			ms[k] = s
			continue
		}
		b, err := json.Marshal(&v)
		if err != nil {
			return nil, err
		}
		ms[k] = string(b)
	}
	// ms["start-time"] = q.StartTime.Format("2006-01-02T15:04:05-0700")
	// ms["end-time"] = q.EndTime.Format("2006-01-02T15:04:05-0700")
	ms["start-time"] = q.StartTime.Format("2006-01-02T15:04:05Z0700")
	ms["end-time"] = q.EndTime.Format("2006-01-02T15:04:05Z0700")
	return ms, nil
}
