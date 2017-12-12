package tools_test

import (
	"testing"
	. "github.com/aliate/go-tools"
)

type Index struct {
	Index string
	Status string
}

func TestPrettyJSON(t *testing.T) {
	indices := []Index{
		{"logstash-2017-12-12", "open"},
		{"logstash-2017-12-11", "open"},
		{"logstash-2017-12-10", "open"},
		{"logstash-2017-12-14", "open"},
		{"logstash-2017-12-13", "open"},
	}
	ss, err := PrettyJSON(indices)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ss)
}

func TestPrettyJSON_string(t *testing.T) {
	str := "{\"test\":[{\"test\":\"test\",\"s\":1},{\"test\":\"ttt\",\"s\":3}]}"
	ss, err := PrettyJSON(str)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ss)
}

