package main


import "testing"



func TestItemValue(t *testing.T) {
    v := "the biggest cat"
    i := Item{value: v}
    if res := i.Value(); res != v {
        t.Errorf("Expected %v, got %v", v, res)
    }
}


func TestListLen(t *testing.T) {
    l := List{}
    for _, v := range([]string{"foo", "bar", "spam", "bangs", "jam"}) {
        l.PushBack(v)
    }
    expected := 5
    if len := l.Len(); len != expected {
        t.Errorf("Expected %v, got %v", expected, len)
    }
}


func TestListPushBack(t *testing.T) {
    l := List{}
    l.PushBack("1")
    l.PushBack("2")
    l.PushBack("3")
    if l.first.Value() != "1" || l.last.Value() != "3" {
        t.Errorf("Expected fst == %v, last == %v, got %v and %v", "1", "3", l.first.Value(), l.last.Value())
    }
    if l.first.Next().Value() != "2" || l.last.Prev().Value() != "2" {
        t.Errorf("Expected %v and %v, got %v and %v", "2", "2", l.first.Next().Value(), l.last.Prev().Value())
    }
    if l.first.Prev() != nil || l.last.Next() != nil {
        t.Errorf("Border items incorrect: expected nil, got %v and %v", l.first.Prev(), l.last.Next())
    }
}


func TestListPushFront(t *testing.T) {
    l := List{}
    l.PushFront("1")
    l.PushFront("2")
    l.PushFront("3")
    if l.first.Value() != "3" || l.last.Value() != "1" {
        t.Errorf("Expected fst == %v, last == %v, got %v and %v", "3", "1", l.first.Value(), l.last.Value())
    }
    if l.first.Next().Value() != "2" || l.last.Prev().Value() != "2" {
        t.Errorf("Expected %v and %v, got %v and %v", "2", "2", l.first.Next().Value(), l.last.Prev().Value())
    }
    if l.first.Prev() != nil || l.last.Next() != nil {
        t.Errorf("Border items incorrect: expected nil, got %v and %v", l.first.Prev(), l.last.Next())
    }
}
