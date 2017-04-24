package goset

import (
	"testing"
	"encoding/json"
)

func TestSet_MarshalJSON(t *testing.T) {
	set := make(Set)
	set.Add("element1")
	set.Add([]string{"element2"}...)
	data1, _ := set.MarshalJSON()
	data2, err := json.Marshal(&set)
	t.Logf("set:%v set.MarshalJSON():%s json:%s", set, data1, data2)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestSet_UnmarshalJSON(t *testing.T) {
	set := make(Set)
	err := json.Unmarshal([]byte(`["element1", "element2"]`), &set)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("set:%v", set)
}