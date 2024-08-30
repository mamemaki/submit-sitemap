package flagutil

import (
	"fmt"
	"strings"
)

// ChoiceSet is a flag value type that allows multiple values to be selected from multiple options.
// ChoiceSet implements flag.Value
// see https://github.com/spf13/pflag/blob/d5e0c0615acee7028e1e2740a11102313be88de1/flag.go#L187
type ChoiceSet struct {
	Options []string
	Value   *[]string
}

func NewChoiceSet(options, value []string, p *[]string) *ChoiceSet {
	cs := new(ChoiceSet)
	cs.Options = options
	cs.Value = p
	*cs.Value = value
	return cs
}

func (c ChoiceSet) String() string {
	return fmt.Sprintf("[%s]", strings.Join(*c.Value, ","))
}

func (c *ChoiceSet) isValueAllowed(opts []string, val string) bool {
	for _, opt := range opts {
		if val == opt {
			return true
		}
	}
	return false
}

func (c *ChoiceSet) Set(val string) error {
	values := strings.Split(val, ",")

	for _, v := range values {
		if !c.isValueAllowed(c.Options, v) {
			return fmt.Errorf("%s is not included in [%s]", v, strings.Join(c.Options, ","))
		}
	}

	*c.Value = values
	return nil
}

func (c *ChoiceSet) Type() string {
	return "stringSlice"
}
