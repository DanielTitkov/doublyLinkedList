package main


import (
    "fmt"
)


// List

type List struct {
    first *Item
    last *Item
}


func (self *List) First() *Item {
    return self.first
}


func (self *List) Last() *Item {
    return self.last
}


func (self *List) Len() int {
    i, len := self.first, 0
    for i != nil {
        len++
        i = i.next
    }
    return len
}


func (self *List) PushFront(value interface{}) {
    newFirst := &Item{value: value}
    if oldFirst := self.first; oldFirst != nil {
        newFirst.next = oldFirst
        oldFirst.prev = newFirst
    } else {
        self.last = newFirst
    }
    self.first = newFirst
}


func (self *List) PushBack(value interface{}) {
    newLast := &Item{value: value}
    if oldLast := self.last; oldLast != nil {
        newLast.prev = oldLast
        oldLast.next = newLast
    } else {
        self.first = newLast
    }
    self.last = newLast
}


// Item

type Item struct {
    value interface{}
    prev *Item
    next *Item
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


func main() {
    i1 := Item{value: "foo"}
    i2 := Item{value: "bar", prev:  &i1}
    i3 := Item{value: "spam", prev: &i2}
    i1.next = &i2
    i2.next = &i3

    l := List{}

    // fmt.Println(l.first.value, l.last.value)
    l.PushBack("FOO")
    fmt.Println(l.first.value, l.last.value)
    fmt.Println(l.Len())
    l.PushBack("BAR")
    fmt.Println(l.first.value, l.last.value)
    fmt.Println(l.Len())

    // fmt.Println(i1, i2, i3)
    // fmt.Println(l)
    // fmt.Println(i1.Next())
    // fmt.Println(i2.Prev())
    // fmt.Println(l.Len())
}
