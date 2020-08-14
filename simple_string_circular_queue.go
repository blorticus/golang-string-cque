package stringcque

// SimpleStringCircularQueue is a simple circular buffer of strings
type SimpleStringCircularQueue struct {
	stringSlice                       []string
	capacity                          int
	indexOfLastElementAfterWrapAround int
	indexInSliceForNextInsert         int
	indexOfLastSliceElement           int
	countOfItemsInQueue               int
}

// NewSimpleStringCircularBuffer creates a simple circular buffer of strings, with
// the queue holding up to 'capacity' number of items.  If the capacity is less than
// 1, the nil pointer is returned.
func NewSimpleStringCircularBuffer(capacity int) *SimpleStringCircularQueue {
	if int(capacity) < 1 {
		return nil
	}

	backingSlice := make([]string, capacity)
	return &SimpleStringCircularQueue{
		stringSlice:                       backingSlice,
		capacity:                          len(backingSlice),
		indexOfLastElementAfterWrapAround: -1,
		indexInSliceForNextInsert:         0,
		indexOfLastSliceElement:           capacity - 1,
		countOfItemsInQueue:               0,
	}
}

// PutItemAtEnd places an item at the end of the circular queue
func (queue *SimpleStringCircularQueue) PutItemAtEnd(item string) {
	queue.stringSlice[queue.indexInSliceForNextInsert] = item
	queue.indexInSliceForNextInsert = (queue.indexInSliceForNextInsert + 1) % queue.capacity

	if queue.indexOfLastElementAfterWrapAround == -1 {
		if queue.backingSliceIsFull() && queue.indexInSliceForNextInsert > 0 {
			queue.indexOfLastElementAfterWrapAround = 0
		} else {
			queue.countOfItemsInQueue++
		}
	} else {
		queue.indexOfLastElementAfterWrapAround = (queue.indexOfLastElementAfterWrapAround + 1) % queue.capacity
	}
}

// IsEmpty returns true if the queue has no items in it; false otherwise
func (queue *SimpleStringCircularQueue) IsEmpty() bool {
	return queue.countOfItemsInQueue == 0
}

// IsNotEmpty returns true if the queue has at least one item in it; false otherwise
func (queue *SimpleStringCircularQueue) IsNotEmpty() bool {
	return queue.countOfItemsInQueue != 0
}

// NumberOfItemsInTheQueue returns a count of the number of items in the queue
func (queue *SimpleStringCircularQueue) NumberOfItemsInTheQueue() uint {
	return uint(queue.countOfItemsInQueue)
}

// GetItemAtIndex retrieves the string at the specified index (0 is the first item)
func (queue *SimpleStringCircularQueue) GetItemAtIndex(index uint) (item string, thereIsAnItemAtThatIndex bool) {
	if int(index) > queue.countOfItemsInQueue-1 {
		return "", false
	}

	indexOfItemInBackingSlice := (queue.indexOfLastElementAfterWrapAround + int(index) + 1) % queue.capacity
	return queue.stringSlice[indexOfItemInBackingSlice], true
}

func (queue *SimpleStringCircularQueue) backingSliceIsFull() bool {
	return queue.countOfItemsInQueue == len(queue.stringSlice)
}
