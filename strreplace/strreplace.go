package strreplace

import (
	"bytes"
)

type Option func(origin *Source)

type Source struct {
	// 要处理的数据
	Data []rune
	// 开始的位置
	Start int
	// 结束的位置
	End int
	// 隐藏内容的位置： true: 中间隐藏； false: 两端隐藏
	IsHiddenCenter bool
	// 要替换的内容
	Symbol rune
}

// NewSource 初始化
func NewSource(data []rune, options ...Option) (s *Source) {
	s = &Source{
		Data:           data,
		Start:          3,
		End:            4,
		IsHiddenCenter: true,
		Symbol:         '*',
	}

	for _, option := range options {
		option(s)
	}

	return s
}

// Start 截取开始的位置
func Start(start int) Option {
	return func(o *Source) {
		if start >= 0 && start <= len(o.Data) && o.IsHiddenCenter {
			o.Start = start
		}
	}
}

// End 截取结束的位置
func End(end int) Option {
	return func(o *Source) {
		if end >= 0 && end <= len(o.Data) {
			o.End = end
		}
	}
}

// IsHiddenCenter 截取的方式
func IsHiddenCenter(mode bool) Option {
	return func(o *Source) {
		o.IsHiddenCenter = mode
	}
}

// Symbol 设置需要替换的符号
func Symbol(symbol rune) Option {
	return func(o *Source) {
		o.Symbol = symbol
	}
}

// Hidden 开始隐藏内容
func (o *Source) Hidden() string {

	buf := bytes.Buffer{}

	for k, v := range o.Data {
		// 中间隐藏
		if k >= o.Start && k < len(o.Data)-o.End {
			if o.IsHiddenCenter {
				buf.WriteRune(o.Symbol)
			} else {
				buf.WriteRune(v)
			}
		} else {
			if o.IsHiddenCenter {
				buf.WriteRune(v)
			} else {
				buf.WriteRune(o.Symbol)
			}
		}
	}

	return buf.String()
}
