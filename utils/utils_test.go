package utils

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMarshal(t *testing.T) {

	type teststruct struct {
		Id   int
		Name string
	}

	data := teststruct{
		Id:   1,
		Name: "Test Data",
	}

	jsondata := Marshal(&data)

	outdata := teststruct{}

	err := json.Unmarshal(jsondata, &outdata)
	if err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(data, outdata) {
		t.Errorf("Structs NOT equal, utils.Marshal does not work")
	}
}
