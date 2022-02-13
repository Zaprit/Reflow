package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/Zaprit/Reflow/config"
)

type testStruct struct {
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

	data := testStruct{
		Id:   1,
		Name: "Test Data",
	}

	jsonData, err := Marshal(&data)

	outData := testStruct{}

	err2 := json.Unmarshal(jsonData, &outData)
	if err2 != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(data, outData) {
		t.Errorf("Structs NOT equal, utils.Marshal does not work")
	}

	config.Conf.Server.Debug = true

	jsonData2, err := Marshal(&data)

	outData = testStruct{}

	err2 = json.Unmarshal(jsonData2, &outData)
	if err2 != nil {
		t.Error(err.Error())
	}

	if !reflect.DeepEqual(data, outData) {
		t.Errorf("Structs NOT equal, utils.Marshal does not work")
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err, "- this was expected")
		}
	}()

	jsonData3, err2 := Marshal(testStructWithBadJson{})
	if err2 == nil {
		t.Error("was expecting a failure, hmm. Anyway, here's what json.marshal spat out\n", jsonData3)
	}

}
