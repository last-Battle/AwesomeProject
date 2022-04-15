package Split

import (
	"reflect"
	"testing"
)

//func TestSplit(t *testing.T) {
//	got := Split("a:b:c", ":")
//	want := []string{"a", "b", "c"}
//	if !reflect.DeepEqual(want, got) {
//		t.Errorf("expected: %v, got: %v", want, got)
//	}
//}

func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

//func TestSplit(t *testing.T) {
//	// 定义一个测试用例类型
//	type test struct {
//		input string
//		sep   string
//		want  []string
//	}
//	// 定义一个存储测试用例的切片
//	tests := []test{
//		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
//		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
//		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
//		{input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
//	}
//	for _, tc := range tests {
//		got := Split(tc.input, tc.sep)
//		if !reflect.DeepEqual(got, tc.want) {
//			t.Errorf("expect:%#v, got:%#v", tc.want, got)
//		}
//	}
//}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}

func BenchmarkSplitParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("沙河有沙又有河", "沙")
		}
	},
	)
}
