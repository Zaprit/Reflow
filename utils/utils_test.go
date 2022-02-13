package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/Zaprit/Reflow/config"
)

type teststruct struct {
	Id   int
	Name string
}

type testStructWithBadJson struct {
	Id   int
	Name string
}

func (s testStructWithBadJson) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("sand")
}

func TestMarshal(t *testing.T) {

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

	config.Conf.Server.Debug = true

	jsondata2 := Marshal(&data)

	outdata = teststruct{}

	err = json.Unmarshal(jsondata2, &outdata)
	if err != nil {
		t.Error(err.Error())
	}

	if !reflect.DeepEqual(data, outdata) {
		t.Errorf("Structs NOT equal, utils.Marshal does not work")
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err, "- this was expected")
		}
	}()

	jsondata3 := Marshal(testStructWithBadJson{})
	if jsondata3 != nil {
		//		t.Error("was expecting a failure, wtf")
	}
}
