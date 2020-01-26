package dataeq

import (
	"reflect"
)

type (
	Marshal   func(interface{}) ([]byte, error)
	Unmarshal func([]byte, interface{}) error

	DataFormat interface {
		ConvertByte([]byte) (interface{}, error)
		Convert(interface{}) (interface{}, error)
		Equal(x, y interface{}) (bool, error)
	}

	converter struct {
		marshal   func(interface{}) ([]byte, error)
		unmarshal func([]byte, interface{}) error
	}
)

func New(marshal Marshal, unmarshal Unmarshal) DataFormat {
	return &converter{
		marshal:   marshal,
		unmarshal: unmarshal,
	}
}

func (c *converter) ConvertByte(b []byte) (interface{}, error) {
	var d interface{}
	err := c.unmarshal(b, &d)
	if err == nil {
		return d, nil
	}
	return nil, err
}

func (c *converter) Convert(x interface{}) (interface{}, error) {
	if a, ok := x.([]byte); ok {
		return c.ConvertByte(a)
	}
	b, err := c.marshal(x)
	if err != nil {
		return nil, err
	}
	return c.ConvertByte(b)
}

func (c *converter) Equal(x, y interface{}) (bool, error) {
	if reflect.DeepEqual(x, y) {
		return true, nil
	}
	a, err := c.Convert(x)
	if err != nil {
		return false, err
	}
	b, err := c.Convert(y)
	if err != nil {
		return false, err
	}
	return reflect.DeepEqual(a, b), nil
}
