package model

import (
	"fmt"
)

type Pod struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	Node  string `json:"node"`
}

func MockPodList() []*Pod {
	return []*Pod{
		{Name: "one", Image: "nginx:1.18", Node: "node1"},
		{Name: "two", Image: "nginx:1.18", Node: "node2"},
		{Name: "three", Image: "nginx:1.18", Node: "node3"},
	}
}

// ParseAction 解析行为
// 从消息中解析出执行方法
func (p *Pod) ParseAction(action string) (*Response, error) {
	fmt.Println(action, p.Name, p.Image, p.Node)
	return NewResponse("pod-list", MockPodList()), nil
}
