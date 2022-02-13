package ws

import (
	"encoding/json"
	"fmt"
	"gin-chat/model"
	"reflect"
)

const (
	ClientPing = 0 // 心跳包
	CreatePod  = 1 // 新增商品
)

var (
	CommandModel = map[int]model.IModel{}
)

func init() {
	CommandModel[CreatePod] = (*model.Pod)(nil) // 反射，类型是pod，值是nil
	CommandModel[ClientPing] = (*model.Ping)(nil)
}

// Command 消息命令
type Command struct {
	Type   int
	Action string
	Data   map[string]interface{}
}

// Parse 执行解析
func (command *Command) Parse() (*model.Response, error) {
	if v, ok := CommandModel[command.Type]; ok {
		// 通过反射初始化对象
		obj := reflect.New(reflect.TypeOf(v).Elem()).Interface()
		// 通过json方式，将map转成struct
		bytes, err := json.Marshal(command.Data)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(bytes, obj); err != nil {
			return nil, err
		}

		return obj.(model.IModel).ParseAction(command.Action)
	}

	return nil, fmt.Errorf("error command")
}
