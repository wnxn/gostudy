package interfaces

import (
	"bytes"
	"fmt"
	"strconv"
)

type IntSet struct {
	val int
}

func (s *IntSet) PointerReceiverString() string {
	return strconv.Itoa(s.val)
}

func (s IntSet) ValueReceiverString() string {
	return strconv.Itoa(s.val)
}

func (s *IntSet) PointerReceiverSet(i int) {
	s.val = i
}

func (s IntSet) ValueReceiverSet(i int) {
	s.val = i
}

// For pointer receiver method, we cannot call method on a non-addressable IntSet Value.
// For concrete receiver method, we can call method on a non-addressable IntSet Value.
// Imply conversion
func TestIntSet1() {
	//	_ = IntSet{}.PointerReceiverString()  // error: pointer receiver
	_ = (&IntSet{}).PointerReceiverString()

	_ = IntSet{}.ValueReceiverString()
	_ = (&IntSet{}).ValueReceiverString()
}

func TestIntSet2() {
	value := IntSet{9}
	value.PointerReceiverString()
	value.ValueReceiverString()

	pointer := &value
	pointer.PointerReceiverString()
	pointer.ValueReceiverString()

	fmt.Printf("value = %T, pointer = %T", value, pointer)

	vp := &IntSet{9}
	vp.PointerReceiverString()
	vp.ValueReceiverString()
}

// Value receiver method can only read struct variable.
// Pointer receiver method has write and read access right to struct variable.
func TestIntSet3() {
	buf := bytes.Buffer{}
	pointer := &IntSet{10}
	buf.WriteString(pointer.ValueReceiverString())

	pointer.PointerReceiverSet(20)
	buf.WriteString(pointer.ValueReceiverString())

	pointer.ValueReceiverSet(30)
	buf.WriteString(pointer.ValueReceiverString())

	fmt.Printf("%s", buf.String())
}

type Book interface {
	GetName() string
}

type Comedy struct {
	name string
}

func (c *Comedy) GetName() string {
	return c.name
}

// Can we assign a concrete type value to an interface type
func TestBook1() {
	comedy := Comedy{"Tutu"}
	//	var book Book = comedy // Cannot use type Comedy as type Book
	var book Book = &comedy

	fmt.Printf("Book name: %s, type: %T", book.GetName(), book)
}

// nil interface value vs. dysfunctional value
func TestBook2() {
	var comedy1 *Comedy
	var book1 Book = comedy1
	fmt.Printf("%t, %T, %v\n", book1 == nil, book1, book1)

	var comedy2 *Comedy = &Comedy{"Tutu"}
	var book2 Book = comedy2
	fmt.Printf("%t, %T, %v\n", book2 == nil, book2, book2)

	var book3 Book
	fmt.Printf("%t, %T, %v\n", book3 == nil, book3, book3)
}
