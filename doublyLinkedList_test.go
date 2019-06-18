package main


import "testing"


func TestItemValue(t *testing.T) {
    v := "the biggest cat"
    i := Item{value: v}
    if res := i.Value(); res != v {
        t.Errorf("expected %v, got %v", v, res)
    }
}


func TestListLen(t *testing.T) {
    l := List{}
    for _, v := range([]string{"foo", "bar", "spam", "bangs", "jam"}) {
        l.PushBack(v)
    }
    expected := 5
    if len := l.Len(); len != expected {
        t.Errorf("expected %v, got %v", expected, len)
    }
}


func TestListPushBack(t *testing.T) {
    l := List{}
    l.PushBack("1")
    l.PushBack("2")
    l.PushBack("3")
    if l.First().Value() != "1" || l.Last().Value() != "3" {
        t.Errorf("expected fst == %v, last == %v, got %v and %v", "1", "3", l.First().Value(), l.Last().Value())
    }
    if l.First().Next().Value() != "2" || l.Last().Prev().Value() != "2" {
        t.Errorf("expected %v and %v, got %v and %v", "2", "2", l.First().Next().Value(), l.last.Prev().Value())
    }
    if l.First().Prev() != nil || l.Last().Next() != nil {
        t.Errorf("border items incorrect: expected nil, got %v and %v", l.First().Prev(), l.Last().Next())
    }
}


func TestListPushFront(t *testing.T) {
    l := List{}
    l.PushFront("1")
    l.PushFront("2")
    l.PushFront("3")
    if l.First().Value() != "3" || l.Last().Value() != "1" {
        t.Errorf("expected fst == %v, last == %v, got %v and %v", "3", "1", l.First().Value(), l.Last().Value())
    }
    if l.First().Next().Value() != "2" || l.Last().Prev().Value() != "2" {
        t.Errorf("expected %v and %v, got %v and %v", "2", "2", l.First().Next().Value(), l.last.Prev().Value())
    }
    if l.First().Prev() != nil || l.Last().Next() != nil {
        t.Errorf("border items incorrect: expected nil, got %v and %v", l.First().Prev(), l.Last().Next())
    }
}


func TestItemRemove(t *testing.T) {
    l := List{}
    l.PushBack("1")
    l.PushBack("2")
    l.PushBack("3")
    i2 := *l.First().Next()
    i2.Remove()
    if l.First().Next() != l.Last() || l.Last().Prev() != l.First() || l.Len() != 2 {
        t.Errorf("incorrect resulting list structure, got len %v", l.Len())
    }
}


func TestLastItemRemove(t *testing.T) {
    l := List{}
    l.PushBack("1")
    l.First().Remove()
    if l.First() != nil || l.Last() != nil || l.Len() != 0 {
        t.Errorf("expected nil, nil, 0, got %v, %v, %v", l.First(), l.Last(), l.Len())
    }
}


func TestRemovePenultItem(t *testing.T) {
    l := List{}
    l.PushBack("1")
    l.PushBack("2")
    l.Last().Remove()
    if l.First() != l.Last() || l.Len() != 1 {
        t.Errorf("expected one item in list, got %v, %v, len %v", l.First(), l.Last(), l.Len())
    }
}
