package greetings

import (
	"regexp"
	"testing"
)

// Go 内单元测试，测试文件约定以 `_test.go` 作为后缀
// 在指定目录下，执行 `go test -v` 启动测试脚本并查看输出

// 单元测试离不开输入、输出和预期值。这里我们输出 name，输出为 `Hello(name)` 返回值，预期值是一个包含 name 的字符串
// 一方面，我们测试其正常值的返回情况
// test 方法接受一个 testing.T 的指针作为参数，我们用它来进行日志打印
func TestHello(t *testing.T) {
	name := "abc"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello(%v) = %q, %v, want match for %#q, nil`, name, msg, err, want)
	}
}

// 另一方面，我们反向测试其异常
// 不是所有测试都需要正则匹配，这里如果返回不为空，则直接抛出异常
func TestHelloEmptyName(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want ""`, msg, err)
	}
}
