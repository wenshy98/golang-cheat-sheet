package option

import (
	"errors"
	"fmt"
)

// 在一些场景下，我们可能并不想对外暴露具体的配置结构体，而是仅仅对外提供一个功能函数。这时我们会将对应的结构体定义为小写字母开头，将其限制只在包内部使用。

type sthOptions struct {
	value int
	name  string
}

type Option interface {
	apply(*sthOptions) error
}

// applyOption 实现 option 接口
type applyOption struct {
	f func(*sthOptions) error
}

func (a *applyOption) apply(opts *sthOptions) error {
	return a.f(opts)
}

func newApplyOption(f func(opts *sthOptions) error) *applyOption {
	return &applyOption{f: f}
}

func WithName(name string) Option {
	return newApplyOption(func(o *sthOptions) error {
		o.name = name
		return nil
	})
}

func WithValue(value int) Option {
	return newApplyOption(func(o *sthOptions) error {
		if value == 0 {
			return errors.New("value is zero")
		}
		o.value = value
		return nil
	})
}

// DoSomething 包对外提供的函数
func DoSomething(name string, opts ...Option) error {
	o := &sthOptions{name: name}
	for _, opt := range opts {
		err := opt.apply(o)
		if err != nil {
			return err
		}
	}
	// 在包内部基于o实现逻辑...
	fmt.Printf("o:%#v\n", o)
	return nil
}
