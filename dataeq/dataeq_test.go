package dataeq_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/go-dataeq/dataeq"
)

type (
	invalidMarshaler struct{}
)

func (m *invalidMarshaler) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("failed to marshal")
}

func TestJSON_Convert(t *testing.T) {
	t.Parallel()
	data := []struct {
		title   string
		x       interface{}
		isError bool
		exp     interface{}
	}{
		{
			title: "simple []byte map",
			x:     []byte(`{"foo": "bar"}`),
			exp: map[string]interface{}{
				"foo": "bar",
			},
		},
		{
			title: "simple array",
			x:     []string{"foo", "bar"},
			exp:   []interface{}{"foo", "bar"},
		},
		{
			title: "simple int",
			x:     5,
			exp:   float64(5),
		},
		{
			title: "simple nil",
			x:     nil,
			exp:   nil,
		},
		{
			title:   "failed to marshal",
			x:       &invalidMarshaler{},
			isError: true,
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			var a interface{}
			err := dataeq.JSON.Convert(d.x, &a)
			if d.isError {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)
			require.Equal(t, d.exp, a)
		})
	}
}

func TestJSON_Equal(t *testing.T) { //nolint:funlen
	t.Parallel()
	data := []struct {
		title   string
		x       interface{}
		y       interface{}
		isError bool
		exp     bool
	}{
		{
			title: "compare equal map",
			x: map[string]string{
				"foo": "bar",
			},
			y: map[string]string{
				"foo": "bar",
			},
			exp: true,
		},
		{
			title: "compare []byte and map",
			x:     []byte(`{"foo": "bar"}`),
			y: map[string]string{
				"foo": "bar",
			},
			exp: true,
		},
		{
			title: "compare empty map and nil",
			x:     nil,
			y:     map[string]string{},
			exp:   false,
		},
		{
			title:   "failed to marshal x",
			x:       &invalidMarshaler{},
			y:       map[string]string{},
			isError: true,
		},
		{
			title:   "failed to marshal y",
			x:       map[string]string{},
			y:       &invalidMarshaler{},
			isError: true,
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			f, err := dataeq.JSON.Equal(d.x, d.y)
			if d.isError {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)
			if d.exp {
				require.True(t, f)
			} else {
				require.False(t, f)
			}
		})
	}
}
