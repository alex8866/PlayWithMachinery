package task

import (
	"time"

	"github.com/RichardKnop/machinery/v1/tasks"

	"github.com/zhiruchen/PlayWithMachinery/server"
)

type UpdateFieldArgs struct {
	ObjID      string      `json:"obj_id"`
	ObjType    string      `json:"obj_type"`
	FieldName  string      `json:"field_name"`
	FieldValue interface{} `json:"field_value"`
}

type SendMsgArgs struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Content string   `json:"content"`
}

// CreateTask 创建一个延时任务
func CreateTask(name string, args []tasks.Arg, eta time.Time, onerror, onsuccess *tasks.Signature) *tasks.Signature {
	sig := &tasks.Signature{
		Name: name,
		Args: args,
		ETA:  &eta,
	}
	if onerror != nil {
		sig.OnError = []*tasks.Signature{onerror}
	}
	if onsuccess != nil {
		sig.OnSuccess = []*tasks.Signature{onsuccess}
	}

	return sig
}

// SendTask 发送任务到server
func SendTask(sig *tasks.Signature, t server.ITask) error {
	s, err := server.StartServer(t)
	if err != nil {
		return err
	}

	_, err = s.SendTask(sig)
	if err != nil {
		return err
	}

	return nil
}
