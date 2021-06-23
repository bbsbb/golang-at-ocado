package state

import (
	"errors"

	"github.com/dimitarkovachev/golang-at-ocado/proj/sort/gen"
)

func itemsEqual(i1 *gen.Item, i2 *gen.Item) (bool, error) {
	isItem1Empty, err1 := isEmptyItem(i1)
	isItem2Empty, err2 := isEmptyItem(i2)
	if err1 != nil || err2 != nil {
		return false, errors.New("cannot compare nil items")
	}
	if isItem1Empty || isItem2Empty {
		return false, errors.New("cannot compare empty items")
	}

	return i1.Code == i2.Code && i1.Label == i2.Label, nil
}

func isEmptyItem(item *gen.Item) (bool, error) {
	if item == nil {
		return true, errors.New("item is nil") // is it right to return true here
	}

	return item.Code == "" || item.Label == "", nil
}
