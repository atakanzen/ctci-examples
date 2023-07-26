package linkedlists_test

import (
	"testing"

	"github.com/atakanzen/ctci-examples/linkedlists"
	"github.com/atakanzen/go-dsa/linkedlist"
	"github.com/stretchr/testify/assert"
)

// 2.1 --------------------------------------------------------------------------------------------------------------

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

func BenchmarkRemoveDups(b *testing.B) {
	ll := linkedlist.NewDoublyLinkedList(15, 25, 23, 11, 25, 23)
	for i := 0; i < b.N; i++ {
		linkedlists.RemoveDups(ll)
	}
}

func TestRemoveDupsWithoutBuffer(t *testing.T) {
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
			actual := linkedlists.RemoveDupsWithoutBuffer(tC.input)
			assert.Equal(t, tC.want, actual)
		})
	}
}

func BenchmarkRemoveDupsWithoutBuffer(b *testing.B) {
	ll := linkedlist.NewDoublyLinkedList(15, 25, 23, 11, 25, 23)
	for i := 0; i < b.N; i++ {
		linkedlists.RemoveDupsWithoutBuffer(ll)
	}
}

// 2.2 --------------------------------------------------------------------------------------------------------------

func TestKthToLastElement(t *testing.T) {
	testCases := []struct {
		desc       string
		linkedList *linkedlist.DoublyLinkedList
		k          int
		want       interface{}
	}{
		{
			desc:       "should return 2nd to last element",
			linkedList: linkedlist.NewDoublyLinkedList(15, 25, 23, 11, 14),
			k:          2,
			want:       11,
		},
		{
			desc:       "should return 4th to last element",
			linkedList: linkedlist.NewDoublyLinkedList(15, 25, 23, 11, 14),
			k:          4,
			want:       25,
		},
		{
			desc:       "should return 1st to last element",
			linkedList: linkedlist.NewDoublyLinkedList(15, 25, 23, 11, 14),
			k:          1,
			want:       14,
		},
		{
			desc:       "should return nil for k out of bounds",
			linkedList: linkedlist.NewDoublyLinkedList(15, 25, 23, 11, 14),
			k:          5,
			want:       nil,
		},
		{
			desc:       "should return nil for k negative",
			linkedList: linkedlist.NewDoublyLinkedList(15, 25, 23, 11, 14),
			k:          -1,
			want:       nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := linkedlists.KthToLastElement(tC.linkedList.Head, tC.k)
			assert.Equal(t, tC.want, actual)
		})
	}
}

func TestKthToLastElementRecursive(t *testing.T) {
	testCases := []struct {
		desc       string
		linkedList *linkedlist.DoublyLinkedList
		k          int
		want       interface{}
	}{
		{
			desc:       "should return 2nd to last element",
			linkedList: linkedlist.NewDoublyLinkedList(15, 25, 23, 11, 14),
			k:          2,
			want:       11,
		},
		{
			desc:       "should return 4th to last element",
			linkedList: linkedlist.NewDoublyLinkedList(15, 25, 23, 11, 14),
			k:          4,
			want:       25,
		},
		{
			desc:       "should return 1st to last element",
			linkedList: linkedlist.NewDoublyLinkedList(15, 25, 23, 11, 14),
			k:          1,
			want:       14,
		},
		{
			desc:       "should return nil for k out of bounds",
			linkedList: linkedlist.NewDoublyLinkedList(15, 25, 23, 11, 14),
			k:          6,
			want:       nil,
		},
		{
			desc:       "should return nil for k negative",
			linkedList: linkedlist.NewDoublyLinkedList(15, 25, 23, 11, 14),
			k:          -1,
			want:       nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			i := 0
			actual := linkedlists.KthToLastElementRecursive(tC.linkedList.Head, tC.k, &i)
			assert.Equal(t, tC.want, actual)
		})
	}
}

func TestKthToLastElementRunner(t *testing.T) {
	testCases := []struct {
		desc       string
		linkedList *linkedlist.DoublyLinkedList
		k          int
		want       interface{}
	}{
		{
			desc:       "should return 2nd to last element",
			linkedList: linkedlist.NewDoublyLinkedList(15, 25, 23, 11, 14),
			k:          2,
			want:       11,
		},
		{
			desc:       "should return 4th to last element",
			linkedList: linkedlist.NewDoublyLinkedList(15, 25, 23, 11, 14),
			k:          4,
			want:       25,
		},
		{
			desc:       "should return 1st to last element",
			linkedList: linkedlist.NewDoublyLinkedList(15, 25, 23, 11, 14),
			k:          1,
			want:       14,
		},
		{
			desc:       "should return nil for k out of bounds",
			linkedList: linkedlist.NewDoublyLinkedList(15, 25, 23, 11, 14),
			k:          6,
			want:       nil,
		},
		{
			desc:       "should return nil for k negative",
			linkedList: linkedlist.NewDoublyLinkedList(15, 25, 23, 11, 14),
			k:          -1,
			want:       nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := linkedlists.KthToLastElementRunner(tC.linkedList.Head, tC.k)
			assert.Equal(t, tC.want, actual)
		})
	}
}

// 2.3 --------------------------------------------------------------------------------------------------------------

func TestDeleteMiddleNode(t *testing.T) {
	testCases := []struct {
		desc        string
		ll          linkedlist.DoublyLinkedList
		valueToFind interface{}
		valuesAfter []interface{}
		want        bool
	}{
		{
			desc:        "should delete B",
			ll:          *linkedlist.NewDoublyLinkedList("A", "B", "C", "D", "E"),
			valueToFind: "A",
			valuesAfter: []interface{}{"A", "C", "D", "E"},
			want:        true,
		},
		{
			desc:        "should delete C",
			ll:          *linkedlist.NewDoublyLinkedList("A", "B", "C", "D", "E"),
			valueToFind: "B",
			valuesAfter: []interface{}{"A", "B", "D", "E"},
			want:        true,
		},
		{
			desc:        "should not delete with last item",
			ll:          *linkedlist.NewDoublyLinkedList("A", "B", "C", "D", "E"),
			valueToFind: "E",
			valuesAfter: []interface{}{"A", "B", "C", "D", "E"},
			want:        false,
		},
		{
			desc:        "should not delete with not existing item",
			ll:          *linkedlist.NewDoublyLinkedList("A", "B", "C", "D", "E"),
			valueToFind: 15,
			valuesAfter: []interface{}{"A", "B", "C", "D", "E"},
			want:        false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			node := tC.ll.Get(tC.valueToFind)

			actual := linkedlists.DeleteMiddleNode(node)
			actualValuesAfter := tC.ll.Values()

			assert.Equal(t, tC.want, actual)
			assert.Equal(t, tC.valuesAfter, actualValuesAfter)
		})
	}
}

// 2.4 --------------------------------------------------------------------------------------------------------------

func TestPartition(t *testing.T) {
	testCases := []struct {
		desc  string
		ll    *linkedlist.DoublyLinkedList
		pivot interface{}
		want  []interface{}
	}{
		{
			desc:  "should partition around 5",
			ll:    linkedlist.NewDoublyLinkedList(3, 5, 8, 5, 10, 2, 1),
			pivot: 5,
			want:  []interface{}{3, 2, 1, 5, 8, 5, 10},
		},
		{
			desc:  "should partition around 10",
			ll:    linkedlist.NewDoublyLinkedList(3, 25, 8, 10, 10, 13, 1),
			pivot: 10,
			want:  []interface{}{3, 8, 1, 25, 10, 10, 13},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			linkedlists.Partition(tC.ll, tC.pivot)
			assert.Equal(t, tC.want, tC.ll.Values())
		})
	}
}

// 2.5 --------------------------------------------------------------------------------------------------------------

func TestSumByLinkedList(t *testing.T) {
	testCases := []struct {
		desc                 string
		numberOne, numberTwo *linkedlist.DoublyLinkedList
		want                 *linkedlist.DoublyLinkedList
	}{
		{
			desc:      "should return 219",
			numberOne: linkedlist.NewDoublyLinkedList(7, 1, 6),
			numberTwo: linkedlist.NewDoublyLinkedList(5, 9, 2),
			want:      linkedlist.NewDoublyLinkedList(2, 1, 9),
		},
		{
			desc:      "should return 5911",
			numberOne: linkedlist.NewDoublyLinkedList(2, 7, 6),
			numberTwo: linkedlist.NewDoublyLinkedList(3, 2, 5),
			want:      linkedlist.NewDoublyLinkedList(5, 9, 1, 1),
		},
		{
			desc:      "should return 9591",
			numberOne: linkedlist.NewDoublyLinkedList(5, 7, 9),
			numberTwo: linkedlist.NewDoublyLinkedList(4, 8, 9),
			want:      linkedlist.NewDoublyLinkedList(9, 5, 9, 1),
		},
		{
			desc:      "should return 0004",
			numberOne: linkedlist.NewDoublyLinkedList(0, 0, 0, 2),
			numberTwo: linkedlist.NewDoublyLinkedList(0, 0, 0, 2),
			want:      linkedlist.NewDoublyLinkedList(0, 0, 0, 4),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := linkedlists.SumByLinkedList(tC.numberOne, tC.numberTwo)
			assert.Equal(t, tC.want.Values(), actual.Values())
		})
	}
}
