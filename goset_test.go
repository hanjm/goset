package goset

import (
	"testing"
	"encoding/json"
)

func TestSet_MarshalJSON(t *testing.T) {
	set := make(Set)
	set.Add("element")
	data, err := json.Marshal(&set)
	t.Logf("json:%s", data)
	if err != nil {
		t.Fatalf(err.Error())
	}
}