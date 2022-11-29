// Package brand 生成随机数，防止和系统包名冲突，前面加了个b
package brand

import (
	"math/rand"
	"strings"
	"time"
)

const (
	defaultNumber      = "1234567890"
	defaultLowerLetter = "abcdefghijklmnopqrstuvwxyz"
	defaultUpperLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type Source struct {
	Seed   []rune
	length uint
}

type Result []rune

type Option func(source *Source)

func NewSource(options ...Option) *Source {
	s := &Source{
		Seed:   []rune(strings.Join([]string{defaultNumber, defaultLowerLetter, defaultUpperLetter}, "")),
		length: 8,
	}

	for _, option := range options {
		option(s)
	}

	return s
}

// OnlyNumber 只有纯数字，包含 1234567890
func OnlyNumber() Option {
	return func(source *Source) {
		source.Seed = []rune("1234567890")
	}
}

// OnlyLetter 只有字母，包含26个字母的小写
func OnlyLetter() Option {
	return func(source *Source) {
		source.Seed = []rune(strings.Join([]string{defaultLowerLetter, defaultUpperLetter}, ""))
	}
}

// OnlyLowerLetter 只有小写字母
func OnlyLowerLetter() Option {
	return func(source *Source) {
		source.Seed = []rune(defaultLowerLetter)
	}
}

// OnlyUpperLetter 只有大写字母
func OnlyUpperLetter() Option {
	return func(source *Source) {
		source.Seed = []rune(defaultUpperLetter)
	}
}

// Seed 自定义随机字符串种子
func Seed(s string) Option {
	return func(source *Source) {
		source.Seed = []rune(s)
	}
}

func Length(l uint) Option {
	return func(source *Source) {
		source.length = l
	}
}

func (s *Source) Rand() (rs Result) {
	if s.length == 0 || len(s.Seed) == 0 {
		return make(Result, 0)
	}

	n := int(s.length)
	// 随机数的种子
	src := rand.NewSource(time.Now().UnixNano())

	var letterIdxBits uint = 6
	var letterIdxMask int64 = 1<<letterIdxBits - 1
	var letterIdxMax = 63 / letterIdxBits

	rs = make([]rune, n)

	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(s.Seed) {
			rs[i] = s.Seed[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return
}

// OrderNumber 创建子订单号
func (s *Source) OrderNumber(l uint, pres ...string) (orderNumber string) {
	// 后面随机数的长度
	suffix := ""
	if l > 0 {
		s.length = l
		s.Seed = []rune(defaultNumber)

		suffix = s.Rand().String()
	}

	// 根据当前时间的年月日时分秒创建订单的前缀
	ts := time.Now().Format("20060102150405")

	prefix := ""
	if len(pres) > 0 {
		prefix = pres[0]
	}

	orderNumber = strings.Join([]string{prefix, ts, suffix}, "")

	return
}

// 将结果转换成 string 类型
func (r Result) String() string {
	return string(r)
}
