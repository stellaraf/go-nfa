package flow

type RuleInterface interface {
	Key() string
	Value() any
	Operator() Operator
}

type RuleP struct {
	Key      string   `json:"key"`
	Operator Operator `json:"comparisonOperator"`
	Value    any      `json:"value"`
}

type CountryR struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type RuleVal interface {
	string | int | uint8 | uint16 | uint32 | uint64 | []string | []uint32 | CountryR
}

func AddRule[T RuleVal](f *Filter, key string, op Operator, value T) *Filter {
	// r := RuleP{key, op, value}
	// f.rules = append(f.rules, r)
	f.Rules = append(f.Rules, map[string]any{"key": key, "comparisonOperator": op.String(), "value": value})
	return f
}

func AddSubFilterRule(f *Filter, value *SubFilter) *Filter {
	sr := map[string]any{
		"condition": value.Condition.String(),
		"rules":     value.Rules,
	}
	f.Rules = append(f.Rules, sr)
	return f
}

func AddSubRule[T RuleVal](sf *SubFilter, key string, op Operator, value T) *SubFilter {
	sf.Rules = append(sf.Rules, map[string]any{"key": key, "comparisonOperator": op.String(), "value": value})
	return sf
}
