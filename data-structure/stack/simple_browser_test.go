package stack

import (
	"github.com/zepeng-jiang/go-basic-demo/data-structure/stack"
	"testing"
)

func TestBrowser_Forward(t *testing.T) {
	b := stack.NewBrowser()
	b.PushBack("www.qq.com")
	b.PushBack("www.baidu.com")
	b.PushBack("www.sina.com")
	if b.CanBack() {
		b.Back()
	}
	if b.CanForward() {
		b.Forward()
	}
	if b.CanBack() {
		b.Back()
	}
	if b.CanBack() {
		b.Back()
	}
	if b.CanBack() {
		b.Back()
	}
	b.Open("www.taobao.com")
	if b.CanForward() {
		b.Forward()
	}
}