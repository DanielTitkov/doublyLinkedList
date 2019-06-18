package main


import (
    "fmt"
)


// List

type List struct {
    first, last *Item
}


func (self *List) First() *Item {
    return self.first
}


func (self *List) Last() *Item {
    return self.last
}


func (self *List) Map(f func(i *Item)) {
    i := self.first
    for i != nil {
        f(i)
        i = i.next
    }
}


func (self *List) Len() int {
    len := 0
    self.Map(func(_ *Item) {len++})
    return len
}


func (self *List) PushFront(value interface{}) {
    newFirst := &Item{value: value, list: self}
    if oldFirst := self.first; oldFirst != nil {
        newFirst.next = oldFirst
        oldFirst.prev = newFirst
    } else {
        self.last = newFirst
    }
    self.first = newFirst
}


func (self *List) PushBack(value interface{}) {
    newLast := &Item{value: value, list: self}
    if oldLast := self.last; oldLast != nil {
        newLast.prev = oldLast
        oldLast.next = newLast
    } else {
        self.first = newLast
    }
    self.last = newLast
}


func (self *List) PrintValues() {
    self.Map(func (i *Item) {fmt.Println(i.Value())})
}


// Item

type Item struct {
    value interface{}
    list *List
    prev, next *Item
}


func (self *Item) Value() interface{} {
    return self.value
}


func (self *Item) Next() *Item {
    return self.next
}


func (self *Item) Prev() *Item {
    return self.prev
}


func (self *Item) Remove() {
    if self.prev == nil {
        self.list.first = self.next
    } else {
        self.prev.next = self.next
    }
    if self.next == nil {
        self.list.last = self.prev
    } else {
        self.next.prev = self.prev
    }
}


func main() {
    l := List{}

    for _, v := range([]string{"Who", "is", "John", "Galt", "?"}) {
        l.PushBack(v)
    }

    l.PrintValues()
}
