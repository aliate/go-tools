package tools_test

import (
	"testing"
	. "github.com/aliate/go-tools"
)

type Index struct {
	Index string
	Status string
}

func PrettyJson_test(t *testing.T) {
	indices := []Index{
		{"logstash-2017-12-12", "open"},
		{"logstash-2017-12-11", "open"},
		{"logstash-2017-12-10", "open"},
		{"logstash-2017-12-14", "open"},
		{"logstash-2017-12-13", "open"},
	}

	ss, err := PrettyJson(indices)
	if err != nil {
		panic(err)
	}

	t.Log(ss)
}

