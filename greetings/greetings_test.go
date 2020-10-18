package greetings

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, _ := Hello("Gladys")

	assert.Regexp(t, want, msg)
}

func TestHelloEmpty(t *testing.T) {
	_, err := Hello("")
	assert.EqualError(t, err, "Empty name")
}
