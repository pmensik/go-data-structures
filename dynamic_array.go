package main

var capacity, length = 10, 0
var array = make([]int, capacity)

func resizeArray() bool {
	tempArray := array
	array = make([]int, capacity)
	for index := range tempArray {
		array[index] = tempArray[index]
	}
	return true
}

func checkArray() bool {
	if len(array) <= length+1 {
		capacity = capacity * 2
		if !resizeArray() {
			return false
		}
	}
	return true
}

// Add will insert new element into an array and resize it if needed
func Add(element int) bool {
	if !checkArray() {
		return false
	}
	array[length], length = element, length+1
	return true
}

// Contains check whether array contains an element
func Contains(element int) bool {
	for _, value := range array {
		if value == element {
			return true
		}
	}
	return false
}

// Clear will empty the array
func Clear() {
	array = nil
	capacity = 10
	length = 0
}

// Insert element into an index
func Insert(element int, index int) bool {
	checkArray()
	oldArray := array
	array = make([]int, capacity)
	for i := 0; i < index; i++ {
		array[i] = oldArray[i]
	}
	array[index] = element
	for i := index + 1; i < len(oldArray); i++ {
		array[i] = oldArray[i-1]
	}
	length++
	return true
}

// Remove element from an index
func Remove(index int) int {
	if index > length || index < 0 {
		return 0
	}
	oldArray := array
	array = make([]int, capacity)
	for i := 0; i < index; i++ {
		array[i] = oldArray[i]
	}
	for i := index + 1; i < len(oldArray); i++ {
		array[i-1] = oldArray[i]
	}
	length--
	return oldArray[index]
}
