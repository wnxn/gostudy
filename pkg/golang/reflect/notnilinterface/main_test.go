package main

import (
	"bytes"
	"io"
	"testing"
)

func TestInterface(t *testing.T) {
	var b *bytes.Buffer
	var i io.Writer
	i = b
	if i != nil{
		t.Log("i is not nil")
	}else{
		t.Error("i is nil")
	}
	if b != nil{
		t.Errorf("b is not nil")
	}else{
		t.Log("b is nil" )
	}
	t.Logf("%T, %T", b,i)
}
