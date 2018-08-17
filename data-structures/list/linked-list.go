package list

type node struct {
	value 	interface{}
	prev 	*node
	next 	*node
}

type LinkedList struct {
	head 	*node
	tail 	*node
	size 	int
}

func (list *LinkedList) Size() int {
	return list.size
}

func (list *LinkedList) Reverse() {
	if list.tail == nil {
		panic("Empty list.")
	}
	current_node := list.tail
	tmp_node := list.head
	list.head = list.tail
	list.tail = tmp_node
	tmp_node = nil

	var prev_node *node
	for current_node != nil {
		prev_node = current_node.prev
		current_node.prev = current_node.next
		current_node.next = prev_node
		current_node = prev_node
	}
	prev_node = nil

	current_node = nil
}

func (list *LinkedList) IndexOf(value interface{}) int {
	if list.head == nil {
		panic("Empty list")
	}
	index := 0
	current_node := list.head
	for current_node.next != nil {
		if value == current_node.value {
			current_node = nil
			return index
		} else {
			current_node = current_node.next
			index++
		}
	}
	current_node = nil
	return -1
}

func (list *LinkedList) Find(index int) interface{} {
	if index <= -1 || index >= list.size {
		panic("index out of range")
	}
	if list.head == nil {
		panic("Empty list")
	}

	target_node := list.findNodeAt(index)
	value := target_node.value
	target_node = nil
	return value
}

func (list *LinkedList) FindFirst() interface{} {
	if list.head == nil {
		panic("Empty list")
	}
	return list.head.value
}

func (list *LinkedList) FindLast() interface{} {
	if list.tail == nil {
		panic("Empty list")
	}
	return list.tail.value
}

func (list *LinkedList) findNodeAt(index int) *node {
	if index <= -1 || index >= list.size {
		panic("index out of range")
	}

	current_node := list.head
	if index <= list.size / 2 {
		for i := 0; i < index; i++ {
			current_node = current_node.next
		}
	} else {
		current_node = list.tail
		for i := list.size - 1; i > index; i-- {
			current_node = current_node.prev
		}
	}
	return current_node
}

func (list *LinkedList) Insert(value interface{}, index int) {
	if index <= -1 || index > list.size {
		panic("index out of range")
	}
	if index == 0 {
		list.InsertToFirst(value)
	} else if index == list.size {
		list.InsertToLast(value)
	} else {
		new_node := &node{value, nil, nil}
		current_node := list.findNodeAt(index)
		current_node.prev.next = new_node
		new_node.next = current_node
		new_node.prev = current_node.prev
		current_node.prev = new_node
		list.size++
		current_node = nil
		new_node = nil
	}

}

func (list *LinkedList) InsertToFirst(value interface{}) {
	new_node := &node{value, nil, list.head}
	if list.head == nil {
		list.head = new_node
		list.tail = new_node
	} else {
		list.head.prev = new_node
		list.head = new_node
	}
	list.size++
	new_node = nil
}

func (list *LinkedList) InsertToLast(value interface{}) {
	new_node := &node{value, list.tail, nil}
	if list.tail == nil {
		list.head = new_node
		list.tail = new_node
	} else {
		list.tail.next = new_node
		list.tail = new_node
	}
	list.size++
	new_node = nil
}

func (list *LinkedList) RemoveAt(index int) {
	if index <= -1 || index > list.size {
		panic("index out of range")
	}
	if index == 0 {
		list.RemoveFirst()
	} else if (index == list.size - 1) {
		list.RemoveLast()
	} else {
		target_node := list.findNodeAt(index)
		target_node.prev.next = target_node.next
		target_node.next.prev = target_node.prev
		target_node.prev = nil
		target_node.next = nil
		target_node.value = nil
		list.size--
		target_node = nil
	}
}

func (list *LinkedList) RemoveFirst() {
	if list.head == nil {
		panic("Empty list.")
	}
	first_node := list.head
	list.head = first_node.next
	first_node.next = nil
	first_node.value = nil
	list.size--
	first_node = nil
}

func (list *LinkedList) RemoveLast() {
	if list.tail == nil {
		panic("Empty list.")
	}
	last_node := list.tail
	list.tail = last_node.prev
	last_node.prev = nil
	last_node.value = nil
	list.size--
	last_node = nil
}
