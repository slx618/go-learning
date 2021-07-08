package split_string

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	//在单元测试的开始最先运行
	fmt.Println("start...")
	m.Run()
	//在单元测试的结束后运行
	fmt.Println("end...")

}

func TestSplit(t *testing.T) {
	got := Split("aaaa", "a")
	want := []string{"", "", "", "", ""}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%#v got: %#v", want, got)
	}
}

func TestSplit2(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%#v got: %#v", want, got)
	}
}

func TestSplit3(t *testing.T) {
	got := Split("abcabcab", "a")
	want := []string{"", "bc", "bc", "b"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%#v got: %#v", want, got)
	}
}

func TestSplit4(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroup := []testCase{
		{"aaaa", "a", []string{"", "", "", "", ""}},
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{"abcabcab", "a", []string{"", "bc", "bc", "b"}},
	}

	for _, v := range testGroup {
		got := Split(v.str, v.sep)
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("want:%#v got: %#v", v.want, got)
		}
	}
}

func TestSplit5(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	//单独运行某个 go test -v -run="TestSplit5/a"
	//测试覆盖率 go test -cover
	//测试覆盖率输出文件 go test -cover -coverprofile=1.out
	//html查看覆盖率 go tool cover -html=1.out

	testGroup := map[string]testCase{
		"a": testCase{"aaaa", "a", []string{"", "", "", "", ""}},
		"b": testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		"c": testCase{"abcabcab", "a", []string{"", "bc", "bc", "b"}},
	}

	for k, v := range testGroup {
		t.Run(k, func(t *testing.T) {
			got := Split(v.str, v.sep)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("want:%#v got: %#v", v.want, got)
			}
		})
	}

}

//go test -bench=. 运行所有
//go test -bench=Split
//go test -bench=Split -test.benchtime=2 强制运行 2 秒

func BenchmarkSplit(b *testing.B) {
	//b.ResetTimer() 重置时间
	//b.SetParallelism(1)
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}

// ExampleSplit 示例函数
func ExampleSplit() {
	got := Split("a:b:c", ":")
	fmt.Println(got)

	// Output 会自动生成文档
	// xxx
	// xxx
}
