package dataeq

import (
	"reflect"
)

type (
	Marshal   func(interface{}) ([]byte, error)
	Unmarshal func([]byte, interface{}) error

	DataFormat struct {
		marshal   func(interface{}) ([]byte, error)
		unmarshal func([]byte, interface{}) error
	}
)

func New(marshal Marshal, unmarshal Unmarshal) DataFormat {
	return DataFormat{
		marshal:   marshal,
		unmarshal: unmarshal,
	}
}

func (df DataFormat) ConvertByte(b []byte) (interface{}, error) {
	var d interface{}
	err := df.unmarshal(b, &d)
	if err == nil {
		return d, nil
	}
	return nil, err
}

func (df DataFormat) Convert(x interface{}) (interface{}, error) {
	if a, ok := x.([]byte); ok {
		return df.ConvertByte(a)
	}
	b, err := df.marshal(x)
	if err != nil {
		return nil, err
	}
	return df.ConvertByte(b)
}

func (df DataFormat) Equal(x, y interface{}) (bool, error) {
	if reflect.DeepEqual(x, y) {
		return true, nil
	}
	a, err := df.Convert(x)
	if err != nil {
		return false, err
	}
	b, err := df.Convert(y)
	if err != nil {
		return false, err
	}
	return reflect.DeepEqual(a, b), nil
}
