package model

import (
	"fmt"
)

type Pod struct {
	Name  string
	Image string
	Node  string
}

func MockPodList() []*Pod {
	return []*Pod{
		{Name: "one", Image: "nginx:1.18", Node: "node1"},
		{Name: "two", Image: "nginx:1.18", Node: "node1"},
		{Name: "three", Image: "nginx:1.18", Node: "node1"},
	}
}

// ParseAction 解析行为
// 从消息中解析出执行方法
func (p *Pod) ParseAction(action string) error {
	fmt.Println("xxoo", action, p.Name, p.Image, p.Node)
	return nil
}
