package tree

import "testing"

var root = CreateTree()

var testcase = []struct{
	valuep int
	valueq int
	expect int
	}{
	{7,7,7},
	{4,6,1},
	{6,7,3},
	{7,6,3},
}

func TestLowestCommonAncestorBinaryTree1(t *testing.T) {
	for _, test:=range testcase{
		actual := LowestCommonAncestorBinaryTree1(root, test.valuep,test.valueq)
		if actual.Value != test.expect{
			t.Errorf("LCA1 input %d,%d, expect is %d, but actual is %d",
				test.valuep,test.valueq,test.expect,actual)
		}
	}
}

func TestLowestCommonAncestorBinaryTree2(t *testing.T) {
	for _, test:=range testcase{
		actual := LowestCommonAncestorBinaryTree1(root, test.valuep,test.valueq)
		if actual.Value != test.expect{
			t.Errorf("LCA2 input %d,%d, expect is %d, but actual is %d",
				test.valuep,test.valueq,test.expect,actual)
		}
	}
}

func TestLowestCommonAncestorBinaryTree3(t *testing.T) {
	for _, test:=range testcase{
		actual := LowestCommonAncestorBinaryTree1(root, test.valuep,test.valueq)
		if actual.Value != test.expect{
			t.Errorf("LCA3 input %d,%d, expect is %d, but actual is %d",
				test.valuep,test.valueq,test.expect,actual)
		}
	}
}
//
//func BenchmarkLowestCommonAncestorBinaryTree1(b *testing.B) {
//	test := testcase[1]
//	b.ResetTimer()
//	for i:=0;i<b.N;i++{
//		actual := LowestCommonAncestorBinaryTree1(root,test.valuep,test.valueq)
//		if actual.Value != test.expect{
//			b.Errorf("LCA1 input %d,%d, expect is %d, but actual is %d",
//				test.valuep,test.valueq,test.expect,actual)
//		}
//	}
//}

//func BenchmarkLowestCommonAncestorBinaryTree2(b *testing.B) {
//	test := testcase[1]
//	b.ResetTimer()
//	for i:=0;i<b.N;i++{
//		actual := LowestCommonAncestorBinaryTree2(root,test.valuep,test.valueq)
//		if actual.Value != test.expect{
//			b.Errorf("LCA2 input %d,%d, expect is %d, but actual is %d",
//				test.valuep,test.valueq,test.expect,actual)
//		}
//	}
//}
//
func BenchmarkLowestCommonAncestorBinaryTree3(b *testing.B) {
	test := testcase[1]
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		actual := LowestCommonAncestorBinaryTree3(root,test.valuep,test.valueq)
		if actual.Value != test.expect{
			b.Errorf("LCA3 input %d,%d, expect is %d, but actual is %d",
				test.valuep,test.valueq,test.expect,actual)
		}
	}
}