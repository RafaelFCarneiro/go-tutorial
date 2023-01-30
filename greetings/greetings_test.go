package greetings

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func TestHello(t *testing.T) {

	t.Run("Say hello to informed name", func(t *testing.T) {
		name := "Rafael"

		wantOptions := make([]string, 3)
		for i, f := range GetFormats() {
			wantOptions[i] = fmt.Sprintf(f, name)
		}

		if result, err := Hello(name); !slices.Contains(wantOptions, result) || err != nil {
			t.Errorf(`Hello() = %q, want match for any of %#q`, result, wantOptions)
		}
	})

	t.Run("Inform error if no name is informed", func(t *testing.T) {
		result, err := Hello("")
		assert.EqualError(t, err, EmptyName)
		assert.Empty(t, result)
	})

}

func TestHellos(t *testing.T) {

	t.Run("Say multiple hellos to informed names", func(t *testing.T) {
		names := []string{"Rafael", "Patty", "Livia"}

		wantMap := make(map[string]*regexp.Regexp)
		for _, name := range names {
			wantMap[name] = regexp.MustCompile(`\b` + name + `\b`)
		}

		if results, err := Hellos(names); err != nil || results == nil {
			t.Errorf(`Hellos(%v) = %q, want match for any of %#q`, names, results, wantMap)
		} else {
			for _, name := range names {
				if !wantMap[name].MatchString(results[name]) {
					t.Errorf(`Hellos(%v) = %q, want match for any of %#q`, names, results[name], wantMap[name])
				}
			}
		}
	})

	t.Run("Inform error if at least one empty name is informed", func(t *testing.T) {
		result, err := Hellos([]string{"Rafael", ""})
		assert.EqualError(t, err, EmptyName)
		assert.Nil(t, result)
	})

}
