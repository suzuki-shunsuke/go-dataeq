package dataeq

import (
	"reflect"
)

type (
	// Marshal converts data to a byte string.
	Marshal func(interface{}) ([]byte, error)
	// Unmarshal parses the encoded byte string and stores the result in the value pointed to by 2th argument.
	// The 2th argument must be a pointer.
	Unmarshal func([]byte, interface{}) error

	// DataFormat allows to compare values as the data format.
	// DataFormat must be created by the function `New`.
	DataFormat struct {
		marshal   func(interface{}) ([]byte, error)
		unmarshal func([]byte, interface{}) error
	}
)

// New creates DataFormat by Marshal and Unmarshal.
// DataFormat must be created by this function.
func New(marshal Marshal, unmarshal Unmarshal) DataFormat {
	return DataFormat{
		marshal:   marshal,
		unmarshal: unmarshal,
	}
}

// Convert converts value to byte string and unmarshals the byte string to dst.
// Convert can be used to normalize the value to compare with the other value.
func (df *DataFormat) Convert(x, dst interface{}) error {
	if a, ok := x.([]byte); ok {
		return df.unmarshal(a, dst)
	}
	b, err := df.marshal(x)
	if err != nil {
		return err
	}
	return df.unmarshal(b, dst)
}

// Equal returns true if two arguments are equal.
func (df *DataFormat) Equal(x, y interface{}) (bool, error) {
	if reflect.DeepEqual(x, y) {
		return true, nil
	}
	var a interface{}
	if err := df.Convert(x, &a); err != nil {
		return false, err
	}
	var b interface{}
	if err := df.Convert(y, &b); err != nil {
		return false, err
	}
	return reflect.DeepEqual(a, b), nil
}
