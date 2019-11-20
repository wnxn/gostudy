package kmp

import (
	"reflect"
	"testing"
)

func getNext(pattern string)[]int{
	length := len(pattern)
	if length <1{
		return nil
	}
	next := make([]int, length)
	i:=0
	j:=-1
	next[0]=-1
	for i<length-1{
		if j == -1|| pattern[i]==pattern[j]{
			i++
			j++
			if pattern[i]!= pattern[j]{
				next[i]=j
			}else{
				next[i]=next[j]
			}

		}else{
			j=next[j]
		}
	}
	return next
}

func kmp(s string, t string)int{
	next := getNext(t)
	i,j:=0,0
	for ; i< len(s) && j<len(t);{
		if j==-1|| s[i] == t[j]{
			i++
			j++
		}else{
			j=next[j]
		}
	}
	if j == len(t){
		return i-j
	}
	return -1
}

func TestGetNext(t *testing.T){
	tests := []struct{
		pattern string
		next []int
	}{
		{
			pattern: "abaabc",
			next: []int{-1,0,-1,1,0,2},
		},
		{
			pattern: "aaaab",
			next: []int{-1,-1,-1,-1,3},
		},
	}
	for _, test := range tests{
		res := getNext(test.pattern)
		if !reflect.DeepEqual(res, test.next){
			t.Errorf("expect %v, but actually %v", test.next, res)
		}
	}
}


func TestKMP(t *testing.T){
	tests := []struct{
		s string
		t string
		res int
	}{
		{
			s: "hello",
			t: "ll",
			res: 2,
		},
		{
			s: "aaaaa",
			t: "bba",
			res: -1,
		},
	}
	for _, test:=range tests{
		res := kmp(test.s, test.t)
		if res != test.res{
			t.Errorf("expect %d, but actually %d",test.res,res)
		}
	}
}
