package brand

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	result := NewSource().Rand()
	fmt.Println("默认结果：", result)

	result = NewSource(Length(6)).Rand()
	fmt.Println("指定长度", result)

	result = NewSource(OnlyLetter(), Length(6)).Rand()
	fmt.Println("只有字母", result)

	result = NewSource(OnlyLowerLetter(), Length(6)).Rand()
	fmt.Println("只有小写字母", result)

	result = NewSource(OnlyUpperLetter(), Length(6)).Rand()
	fmt.Println("只有大写字母", result)

	result = NewSource(OnlyNumber(), Length(6)).Rand()
	fmt.Println("只有数字", result)

	result = NewSource(Seed("我是一个中文English123"), Length(10)).Rand()
	fmt.Println("自定义随机数种子", result)

	on := NewSource().OrderNumber(4)
	fmt.Println("创建订单号", on)
}
