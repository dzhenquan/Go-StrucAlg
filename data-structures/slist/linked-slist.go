package slist

type node struct {
    value   interface{}
    next    *node
}

type LinkedSList struct {
    head    *node
    tail    *node
    size    int
}

type GFindFunc func (value1 interface{}, value2 interface{}) (interface{}, bool)


// get slist node length
// return slist length
func (slist *LinkedSList) Size() int {
    return slist.size
}

// get node index by value
// return index/-1
func (slist *LinkedSList) IndexOf(value interface{}) int {
    if slist.head == nil {
        panic("Empty slist.")
    }

    current_node := slist.head
    index := 0
    for current_node != nil {
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

// reverse slist
func (slist *LinkedSList) Reverse() {
    if slist.tail == nil {
        panic("Empty slist.")
    }

    prev_node := slist.tail.next
    for i := 0; i < slist.size; i++ {
        next_node := slist.head.next
        slist.head.next = prev_node

        prev_node = slist.head
        slist.head = next_node
        next_node = nil
    }

    slist.head = prev_node
    slist.tail = slist.findNodeAt(slist.size-1)

    prev_node = nil
}

// get slist node by index
func (slist *LinkedSList) Find(index int) interface{} {
    if index < 0 || (index > slist.size && slist.size > 0) {
        panic("Invalid index.")
    }
    if slist.head == nil {
        panic("Empty slist.")
    }
    target_node := slist.findNodeAt(index)
    value := target_node.value
    target_node = nil
    return value
}

// custom find func
// GFindFunc custom find function
// value will find value
// return find node/nil or true/false
func (slist *LinkedSList) FindCustom(findFunc GFindFunc, value interface{}) (interface{}, bool) {
    if slist.head == nil {
        panic("Empty slist.")
    }
    current_node := slist.head
    for current_node != nil {
        if _, ok := findFunc(current_node.value, value); ok {
            return current_node.value, true
        }
        current_node = current_node.next
    }
    current_node = nil
    return nil, false
}

// find first node from slist
// return slist node
func (slist *LinkedSList) FindFirst() interface{} {
    if slist.head == nil {
        panic("Empty slist.")
    }
    return slist.head.value
}


// find last node from slist
// return slist node
func (slist *LinkedSList) FindLast() interface{} {
    if slist.tail == nil {
        panic("Empty slist.")
    }
    return slist.tail.value
}


func (slist *LinkedSList) findNodeAt(index int) *node {
    current_node := slist.head

    for i := 0; i < index; i++ {
        current_node = current_node.next
    }
    return current_node
}

// insert node to slist by value and index
func (slist *LinkedSList) Insert(value interface{}, index int) {
    if index < 0 || (index > slist.size && slist.size > 0) {
        panic("Invalid index.")
    } else if index == 0 {
        slist.InsertToFirst(value)
    } else if index == slist.size {
        slist.InsertToLast(value)
    } else {
        if slist.head == nil {
            panic("Empty slist and invalid index.")
        } else {
            new_node := &node{value, nil}
            target_node := slist.findNodeAt(index-1)
            new_node.next = target_node.next
            target_node.next = new_node
            slist.size++
            new_node = nil
            target_node = nil
        }
    }
}

// insert node to first slist
// value will insert to slist node
func (slist *LinkedSList) InsertToFirst(value interface{}) {
    new_node := &node{value, slist.head}
    if slist.head == nil {
        slist.head = new_node
        slist.tail = new_node
    } else {
        slist.head = new_node
    }
    slist.size++
    new_node = nil
}

// insert node to last slist
// value will insert to slist node
func (slist *LinkedSList) InsertToLast(value interface{}) {
    new_node := &node{value, nil}
    if slist.tail == nil {
        slist.head = new_node
        slist.tail = new_node
    } else {
        slist.tail.next = new_node
        slist.tail = new_node
    }
    slist.size++
    new_node = nil
}

// remove node from slist by value
func (slist *LinkedSList) RemoveValue(value interface{}) {
    if slist.head == nil {
        panic("Empty slist.")
    }

    current_node := slist.head
    index := 0
    for current_node != nil {
        if value == current_node.value {
            current_node = nil
            slist.Remove(index)
            break
        }
        current_node = current_node.next
        index++
    }
    current_node = nil
}

// remove node by index from slist
func (slist *LinkedSList) Remove(index int) {
    if index < 0 || (index > slist.size && slist.size > 0) {
        panic("Invalid index.")
    } else if index == 0 {
        slist.RemoveFirst()
    } else if (index == slist.size) || (index == slist.size-1) {
        slist.RemoveLast()
    } else {
        if slist.tail == nil {
            panic("Empty slist and invalid index.")
        } else {
            target_node := slist.findNodeAt(index-1)
            remove_node := target_node.next
            target_node.next = remove_node.next

            remove_node.next = nil
            remove_node.value = nil
            remove_node = nil
            target_node = nil
            slist.size--
        }
    }
}

// remove first node from slist
func (slist *LinkedSList) RemoveFirst() {
    if slist.head == nil {
        panic("Empty slist.")
    }

    first_node := slist.head
    slist.head = first_node.next
    if slist.head == nil {
        slist.tail = nil
    }
    first_node.next = nil
    first_node.value = nil
    slist.size--
    first_node = nil
}

// remove last node from slist
func (slist *LinkedSList) RemoveLast() {
    if slist.tail == nil {
        panic("Empty slist.")
    }

    if slist.head == slist.tail {
        slist.RemoveFirst()
        return
    }

    last_node := slist.tail
    last_prev_node := slist.findNodeAt(slist.size-2)
    slist.tail = last_prev_node
    
    last_prev_node.next = nil
    last_prev_node = nil
    last_node.next = nil
    last_node.value = nil
    slist.size--
    last_node = nil
}

// remove all node from slist
func (slist *LinkedSList) RemoveAll() {
    if slist.head == nil {
        panic("Empty slist.")
    }

    current_node := slist.head
    for current_node != nil {
        target_node := current_node
        current_node = current_node.next

        target_node.next = nil
        target_node.value = nil
        target_node = nil
        slist.size--
    }
    current_node = nil
    slist.head = nil
    slist.tail = nil
}