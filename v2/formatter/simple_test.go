package print

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func TestLogMe(t *testing.T) {
	var buf bytes.Buffer
	log.SetFlags(0)
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stdout)
	}()
	expected := "test"
	LogMe2(expected)
	expected = fmt.Sprintln(expected) //platform adaptive println

	got := strings.Trim(buf.String(), "")
	if got != expected {
		t.Fatalf("Log failed - Expect: \"%s\" Got: \"%s\"", expected, got)
	}
	t.Log("Success in log output")
	t.Log("Success in printline works!")
}
