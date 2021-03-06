package sqlgen

import (
	"fmt"
	"strings"
)

type Values []Value

type Value struct {
	Value interface{}
}

func (self Value) Compile(layout *Template) string {
	if raw, ok := self.Value.(Raw); ok {
		return raw.Raw
	}
	return mustParse(layout.ValueQuote, Raw{fmt.Sprintf(`%v`, self.Value)})
}

func (self Values) Compile(layout *Template) string {
	l := len(self)

	if l > 0 {
		chunks := make([]string, 0, l)

		for i := 0; i < l; i++ {
			chunks = append(chunks, self[i].Compile(layout))
		}

		return strings.Join(chunks, layout.ValueSeparator)
	}

	return ""
}
