package dataeq_test

import (
	"fmt"
	"log"

	"github.com/suzuki-shunsuke/go-dataeq/v2/dataeq"
)

type (
	Foo struct {
		Foo string `json:"foo"`
	}
)

func checkResult(b bool, err error) {
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)
}

func Example() {
	b, err := dataeq.JSON.DeepEqual(
		Foo{Foo: "bar"},
		map[string]string{"foo": "bar"},
	)
	checkResult(b, err)

	// when a type of value is []byte, it is treated as JSON string
	b, err = dataeq.JSON.DeepEqual(
		Foo{Foo: "bar"},
		[]byte(`{"foo": "bar"}`),
	)
	checkResult(b, err)

	// Output:
	// true
	// true
}
