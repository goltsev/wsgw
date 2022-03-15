package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type WriterJSONStub struct {
	v *Notification
}

func (w *WriterJSONStub) WriteJSON(v interface{}) error {
	w.v = v.(*Notification)
	return nil
}

func TestUpdate(t *testing.T) {
	type input struct {
		name         string
		list         []string
		notification *Notification
		expected     *Notification
	}
	tests := []input{
		func() input {
			n := &Notification{
				Symbol: "one",
				Price:  1,
			}
			return input{
				name:         "empty list",
				list:         []string{},
				notification: n,
				expected:     n,
			}
		}(),
		func() input {
			n := &Notification{
				Symbol: "one",
				Price:  1,
			}
			return input{
				name:         "yes",
				list:         []string{"one"},
				notification: n,
				expected:     n,
			}
		}(),
		func() input {
			return input{
				name: "no",
				list: []string{"two"},
				notification: &Notification{
					Symbol: "one",
					Price:  1,
				},
				expected: nil,
			}
		}(),
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			stub := &WriterJSONStub{}
			sub := &Subscriber{
				List: test.list,
				conn: stub,
			}
			assert.NoError(t, sub.write(test.notification))
			assert.Equal(t, test.expected, stub.v)
		})
	}
}
