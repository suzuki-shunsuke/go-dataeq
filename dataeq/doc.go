// Package dataeq allows to create API to compare two values as the data format such as JSON and YAML.
// dataeq compare two values by the following way.
//
// 1. convert value to a byte string by Marshal
// 2. unmarshal a byte string to `interface{}` by Unmarshal
// 3. compare two values by reflect.DeepEqual
//
// So dataeq requires two API for the data format, Marshal and Unmarshal.
package dataeq
