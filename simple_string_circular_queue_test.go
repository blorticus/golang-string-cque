package stringcque_test

import (
	"fmt"
	"testing"

	"github.com/blorticus/stringcque"
)

func TestInvalidStringCque(t *testing.T) {
	q := stringcque.NewSimpleStringCircularBuffer(0)

	if q != nil {
		t.Errorf("On NewSimpleStringCircularBuffer of length 0, should return nil but does not")
	}

	q = stringcque.NewSimpleStringCircularBuffer(-1)

	if q != nil {
		t.Errorf("On NewSimpleStringCircularBuffer of length -1, should return nil but does not")
	}

}

func TestStringCque(t *testing.T) {
	q := stringcque.NewSimpleStringCircularBuffer(4)

	if !q.IsEmpty() {
		t.Errorf("No items in queue, IsEmpty should be true, but is not")
	}

	if q.IsNotEmpty() {
		t.Errorf("No items in queue, IsNotEmpty should be false, but is not")
	}

	if q.NumberOfItemsInTheQueue() != 0 {
		t.Errorf("No items in queue, NumberOfItemsInQueue should be 0, but is %d", q.NumberOfItemsInTheQueue())
	}

	for i := 0; i < 10; i++ {
		item, thereIsAnItemAtThisIndex := q.GetItemAtIndex(uint(i))

		if item != "" {
			t.Errorf("No items in queue, item at index (%d) should be empty string, is (%s)", i, item)
		}

		if thereIsAnItemAtThisIndex {
			t.Errorf("No items in queue, GetItemAtIndex at index (%d) should indicate there is no item, but does not", i)
		}
	}

	testStrings := []string{"first", "second 2", "3 third here", "alkfsdjaslfkj", "this\nthat\nther\t", "", " ", "\n", "\n ", " \n ", "eleven", "twelve"}
	for i, insertString := range testStrings {
		q.PutItemAtEnd(insertString)
		if err := performPostInsertTest(q, i, insertString, testStrings, 4); err != nil {
			t.Error(err.Error())
		}
	}
}

func performPostInsertTest(q *stringcque.SimpleStringCircularQueue, indexInSetOfInsertedItem int, stringThatWasJustInserted string, setOfTestStrings []string, queueCapacity int) error {
	if q.IsEmpty() {
		return fmt.Errorf("After insert number (%d) queue should not be empty but IsEmpty returns true", indexInSetOfInsertedItem)
	}
	if !q.IsNotEmpty() {
		return fmt.Errorf("After insert number (%d) queue should be not empty, but IsNotEmpty returns false", indexInSetOfInsertedItem)
	}

	indexInSetOfFirstQueueItem := 0
	if indexInSetOfInsertedItem+1 > queueCapacity {
		indexInSetOfFirstQueueItem = (indexInSetOfInsertedItem + 1) - queueCapacity
	}

	indexInSetOfNextExpectedItemInQueue := indexInSetOfFirstQueueItem
	for i := 0; i < queueCapacity; i++ {
		item, thereIsAnItemAtThisIndex := q.GetItemAtIndex(uint(i))

		if indexInSetOfInsertedItem < 10 {
			if i > indexInSetOfInsertedItem {
				if thereIsAnItemAtThisIndex {
					return fmt.Errorf("On insert of item (%d) expect that there is no item in queue at index (%d), but GetItemAtIndex reports that there is", indexInSetOfInsertedItem, i)
				}
				continue
			}
		}

		if !thereIsAnItemAtThisIndex {
			return fmt.Errorf("On insert of item (%d) into queue, element (%d) of queue should have element, but GetItemAtIndex says it does not", indexInSetOfInsertedItem, i)
		}
		if item != setOfTestStrings[indexInSetOfNextExpectedItemInQueue] {
			return fmt.Errorf("On insert of item (%d) into queue, element (%d) expected to be (%s), is (%s)", indexInSetOfInsertedItem, i, setOfTestStrings[indexInSetOfNextExpectedItemInQueue], item)
		}
		indexInSetOfNextExpectedItemInQueue++
	}

	return nil
}
