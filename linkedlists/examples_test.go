package linkedlists_test

import (
	"testing"

	"github.com/atakanzen/ctci-examples/linkedlists"
	"github.com/atakanzen/go-dsa/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestRemoveDups(t *testing.T) {
	testCases := []struct {
		desc  string
		input *linkedlist.DoublyLinkedList
		want  []interface{}
	}{
		{
			desc:  "should remove 2 duplicates",
			input: linkedlist.NewDoublyLinkedList(15, 25, 23, 11, 25, 23),
			want:  []interface{}{15, 25, 23, 11},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := linkedlists.RemoveDups(tC.input)
			assert.Equal(t, tC.want, actual)
		})
	}
}
