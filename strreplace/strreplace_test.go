package strreplace

import (
	"fmt"
	"testing"
)

func TestHidden(t *testing.T) {
	// 测试隐藏手机号（默认参数，中间隐藏）
	result := NewSource([]rune("13800138000")).Hidden()
	fmt.Println(result)

	// 测试隐藏手机号（两端隐藏）
	result = NewSource([]rune("13800138000"), IsHiddenCenter(true)).Hidden()
	fmt.Println(result)

	// 测试隐藏手机号（隐藏首部）
	result = NewSource([]rune("13800138000"), Start(0)).Hidden()
	fmt.Println(result)

	// 测试隐藏手机号（隐藏尾部）
	result = NewSource([]rune("13800138000"), End(0)).Hidden()
	fmt.Println(result)

	// 测试隐藏中文
	result = NewSource([]rune("一二三四五六七八九十一")).Hidden()
	fmt.Println(result)

	// 测试隐藏中文， 替换符号
	result = NewSource([]rune("一二三四五六七八九十一"), Symbol('-')).Hidden()
	fmt.Println(result)

	// 测试隐藏中文， 替换符号（字体图标）
	result = NewSource([]rune("一二三四五六七八九十一"), Symbol('🍀')).Hidden()
	fmt.Println(result)
}
