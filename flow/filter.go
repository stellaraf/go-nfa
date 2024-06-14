package flow

type Operator string

type Condition string

func (op Operator) String() string {
	return string(op)
}

func (con Condition) String() string {
	return string(con)
}

const EQ Operator = "eq"
const NEQ Operator = "noteq"
const GE Operator = "ge"
const GT Operator = "gt"
const LT Operator = "lt"
const LE Operator = "le"

const AND Condition = "and"
const OR Condition = "or"

type Filter struct {
	Condition Condition        `json:"condition"`
	Rules     []map[string]any `json:"rules"`
}

type BaseFilter struct {
	Condition Condition `json:"condition"`
	Rules     []RuleP   `json:"rules"`
}

type SubFilter struct {
	Condition Condition        `json:"condition"`
	Rules     []map[string]any `json:"rules"`
}

func NewFilter(c Condition) *Filter {
	f := &Filter{Condition: c}
	return f
}

func NewSubFilter(c Condition) *SubFilter {
	f := &SubFilter{Condition: c}
	return f
}
