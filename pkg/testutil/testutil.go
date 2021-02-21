package testutil

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

const (
	defaultMessageSize = 256
)

// P is a set of parameters we want to output when we fail a test
type P map[string]interface{}

// E non-fataly fails test
func E(t *testing.T, params *P, prefix ...string) {
	t.Error(M(params, prefix...))
}

// F fails test with fatal message
func F(t *testing.T, params *P, prefix ...string) {
	t.Fatal(M(params, prefix...))
}

// M formats a message
func M(params *P, prefix ...string) string {

	var sb strings.Builder
	sb.Grow(defaultMessageSize)

	// Add prefix if present
	if len(prefix) > 0 && prefix[0] != `` {
		sb.WriteString(`[` + prefix[0] + `]`)
	}

	if params != nil {
		m := (map[string]interface{})(*params)

		if len(m) > 0 {
			if len(prefix) > 0 {
				sb.WriteString(` `)
			}

			keys := make([]string, 0, len(m))
			for k := range m {
				keys = append(keys, k)
			}

			sort.Strings(keys)

			// Add vars
			for i, k := range keys {
				if i > 0 {
					sb.WriteString(`, `)
				}
				sb.WriteString(fmt.Sprintf("%v: `%v`", k, m[k]))
			}
		}
	}

	return sb.String()

}
