package stack

import "fmt"

// 模拟浏览器前进后退功能

// Browser 浏览器
type Browser struct {
	forwardStack Stack
	backStack    Stack
}

// NewBrowser 初始化
func NewBrowser() *Browser {
	return &Browser{
		forwardStack: NewArrayStack(),
		backStack:    NewLinkedListStack(),
	}
}

// CanForward 是否可以向前
func (b *Browser) CanForward() bool {
	if b.forwardStack.IsEmpty() {
		return false
	}
	return true
}

// CanBack 是否可以回退
func (b *Browser) CanBack() bool {
	if b.backStack.IsEmpty() {
		return false
	}
	return true
}

// Open 打开新的网页
func (b *Browser) Open(addr string) {
	fmt.Printf("Open new addr %+v \n", addr)
	b.forwardStack.Flush()
}

// PushBack
func (b *Browser) PushBack(addr string) {
	b.backStack.Push(addr)
}

// Forward 向前
func (b *Browser) Forward() {
	if b.forwardStack.IsEmpty() {
		return
	}
	top := b.forwardStack.Pop()
	b.backStack.Push(top)
	fmt.Printf("forward to %+v \n", top)
}

// Back 回退
func (b *Browser) Back() {
	if b.backStack.IsEmpty() {
		return
	}
	top := b.backStack.Pop()
	b.forwardStack.Push(top)
	fmt.Printf("back to %+v \n", top)
}
