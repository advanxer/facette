package set

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSet_New(t *testing.T) {
	s := New()

	if s.Size() != 0 {
		t.Error("New: calling without any parameters should create a set with zero size")
	}

	u := s.New()
	if u.Size() != 0 {
		t.Error("New: creating a new set via s.New() should create a set with zero size")
	}
}

func TestSet_New_parameters(t *testing.T) {
	s := New("string", "another_string", 1, 3.14)

	if s.Size() != 4 {
		t.Error("New: calling with parameters should create a set with size of four")
	}
}

func TestSet_Add(t *testing.T) {
	s := New()
	s.Add(1)
	s.Add(2)
	s.Add(2) // duplicate
	s.Add("fatih")
	s.Add("zeynep")
	s.Add("zeynep") // another duplicate

	if s.Size() != 4 {
		t.Error("Add: items are not unique. The set size should be four")
	}

	if !s.Has(1, 2, "fatih", "zeynep") {
		t.Error("Add: added items are not availabile in the set.")
	}
}

func TestSet_Add_multiple(t *testing.T) {
	s := New()
	s.Add("ankara", "san francisco", 3.14)

	if s.Size() != 3 {
		t.Error("Add: items are not unique. The set size should be three")
	}

	if !s.Has("ankara", "san francisco", 3.14) {
		t.Error("Add: added items are not availabile in the set.")
	}
}

func TestSet_Remove(t *testing.T) {
	s := New()
	s.Add(1)
	s.Add(2)
	s.Add("fatih")

	s.Remove(1)
	if s.Size() != 2 {
		t.Error("Remove: set size should be two after removing")
	}

	s.Remove(1)
	if s.Size() != 2 {
		t.Error("Remove: set size should be not change after trying to remove a non-existing item")
	}

	s.Remove(2)
	s.Remove("fatih")
	if s.Size() != 0 {
		t.Error("Remove: set size should be zero")
	}

	s.Remove("fatih") // try to remove something from a zero length set
}

func TestSet_Remove_multiple(t *testing.T) {
	s := New()
	s.Add("ankara", "san francisco", 3.14, "istanbul")
	s.Remove("ankara", "san francisco", 3.14)

	if s.Size() != 1 {
		t.Error("Remove: items are not unique. The set size should be four")
	}

	if !s.Has("istanbul") {
		t.Error("Add: added items are not availabile in the set.")
	}
}

func TestSet_Pop(t *testing.T) {
	s := New()
	s.Add(1)
	s.Add(2)
	s.Add("fatih")

	a := s.Pop()
	if s.Size() != 2 {
		t.Error("Pop: set size should be two after popping out")
	}

	if s.Has(a) {
		t.Error("Pop: returned item should not exist")
	}

	s.Pop()
	s.Pop()
	b := s.Pop()
	if b != nil {
		t.Error("Pop: should return nil because set is empty")
	}

	s.Pop() // try to remove something from a zero length set
}

func TestSet_Has(t *testing.T) {
	s := New("1", "2", "3", "4")

	if !s.Has("1") {
		t.Error("Has: the item 1 exist, but 'Has' is returning false")
	}

	if !s.Has("1", "2", "3", "4") {
		t.Error("Has: the items all exist, but 'Has' is returning false")
	}
}

func TestSet_Clear(t *testing.T) {
	s := New()
	s.Add(1)
	s.Add("istanbul")
	s.Add("san francisco")

	s.Clear()
	if s.Size() != 0 {
		t.Error("Clear: set size should be zero")
	}
}

func TestSet_IsEmpty(t *testing.T) {
	s := New()

	empty := s.IsEmpty()
	if !empty {
		t.Error("IsEmpty: set is empty, it should be true")
	}

	s.Add(2)
	s.Add(3)
	notEmpty := s.IsEmpty()

	if notEmpty {
		t.Error("IsEmpty: set is filled, it should be false")
	}
}

func TestSet_IsEqual(t *testing.T) {
	s := New("1", "2", "3")
	u := New("1", "2", "3")

	ok := s.IsEqual(u)
	if !ok {
		t.Error("IsEqual: set s and t are equal. However it returns false")
	}
}

func TestSet_IsSubset(t *testing.T) {
	s := New("1", "2", "3", "4")
	u := New("1", "2", "3")

	ok := s.IsSubset(u)
	if !ok {
		t.Error("IsSubset: u is a subset of s. However it returns false")
	}

	ok = u.IsSubset(s)
	if ok {
		t.Error("IsSubset: s is not a subset of u. However it returns true")
	}

}

func TestSet_IsSuperset(t *testing.T) {
	s := New("1", "2", "3", "4")
	u := New("1", "2", "3")

	ok := u.IsSuperset(s)
	if !ok {
		t.Error("IsSuperset: s is a superset of u. However it returns false")
	}

	ok = s.IsSuperset(u)
	if ok {
		t.Error("IsSuperset: u is not a superset of u. However it returns true")
	}

}

func TestSet_String(t *testing.T) {
	s := New()
	if s.String() != "[]" {
		t.Error("String: output is not what is excepted", s.String())
	}

	s.Add("1", "2", "3", "4")
	if s.String() != "[1, 2, 3, 4]" {
		t.Error("String: output is not what is excepted")
	}
}

func TestSet_List(t *testing.T) {
	s := New("1", "2", "3", "4")

	// this returns a slice of interface{}
	if len(s.List()) != 4 {
		t.Error("List: slice size should be four.")
	}

	for _, item := range s.List() {
		r := reflect.TypeOf(item)
		if r.Kind().String() != "string" {
			t.Error("List: slice item should be a string")
		}
	}
}

func TestSet_Copy(t *testing.T) {
	s := New("1", "2", "3", "4")
	r := s.Copy()

	if !s.IsEqual(r) {
		t.Error("Copy: set s and r are not equal")
	}
}

func TestSet_Union(t *testing.T) {
	s := New("1", "2", "3")
	r := New("3", "4", "5")
	u := s.Union(r)

	if u.Size() != 5 {
		t.Error("Union: the merged set doesn't have all items in it.")
	}

	if !u.Has("1", "2", "3", "4", "5") {
		t.Error("Union: merged items are not availabile in the set.")
	}
}

func TestSet_Merge(t *testing.T) {
	s := New("1", "2", "3")
	r := New("3", "4", "5")
	s.Merge(r)

	if s.Size() != 5 {
		t.Error("Merge: the set doesn't have all items in it.")
	}

	if !s.Has("1", "2", "3", "4", "5") {
		t.Error("Merge: merged items are not availabile in the set.")
	}
}

func TestSet_Separate(t *testing.T) {
	s := New("1", "2", "3")
	r := New("3", "5")
	s.Separate(r)

	if s.Size() != 2 {
		t.Error("Separate: the set doesn't have all items in it.")
	}

	if !s.Has("1", "2") {
		t.Error("Separate: items after separation are not availabile in the set.")
	}
}

func TestSet_Intersection(t *testing.T) {
	s := New("1", "2", "3")
	r := New("3", "5")
	u := s.Intersection(r)

	if u.Size() != 1 {
		t.Error("Intersection: the set doesn't have all items in it.")
	}

	if !u.Has("3") {
		t.Error("Intersection: items after intersection are not availabile in the set.")
	}
}

func TestSet_Difference(t *testing.T) {
	s := New("1", "2", "3")
	r := New("2", "3", "5")
	u := s.Difference(r)

	if u.Size() != 1 {
		t.Error("Difference: the set doesn't have all items in it.")
	}

	if !u.Has("1") {
		t.Error("Difference: items are not availabile in the set.")
	}
}

func TestSet_SymmetricDifference(t *testing.T) {
	s := New("1", "2", "3")
	r := New("3", "4", "5")
	u := s.SymmetricDifference(r)

	if u.Size() != 4 {
		t.Error("SymmetricDifference: the set doesn't have all items in it.")
	}

	if !u.Has("1", "2", "4", "5") {
		t.Error("SymmetricDifference: items are not availabile in the set.")
	}
}

func TestSet_StringSlice(t *testing.T) {
	s := New("san francisco", "istanbul", 3.14, 1321, "ankara")
	u := s.StringSlice()

	if len(u) != 3 {
		t.Error("StringSlice: slice should only have three items")
	}

	for _, item := range u {
		r := reflect.TypeOf(item)
		if r.Kind().String() != "string" {
			t.Error("StringSlice: slice item should be a string")
		}
	}
}

func TestSet_IntSlice(t *testing.T) {
	s := New("san francisco", "istanbul", 3.14, 1321, "ankara", 8876)
	u := s.IntSlice()

	if len(u) != 2 {
		t.Error("IntSlice: slice should only have two items")
	}

	for _, item := range u {
		r := reflect.TypeOf(item)
		if r.Kind().String() != "int" {
			t.Error("Intslice: slice item should be a int")
		}
	}
}

func TestSet_RaceAdd(t *testing.T) {
	// Create two sets. Add concurrently items to each of them. Remove from the
	// other one.
	// "go test -race" should detect this if the library is not thread-safe.
	s := New()
	u := New()

	go func() {
		for i := 0; i < 1000; i++ {
			item := "item" + strconv.Itoa(i)
			go func(i int) {
				s.Add(item)
				u.Add(item)
			}(i)
		}
	}()

	for i := 0; i < 1000; i++ {
		item := "item" + strconv.Itoa(i)
		go func(i int) {
			s.Add(item)
			u.Add(item)
		}(i)
	}
}