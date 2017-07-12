package client

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/zhiruchen/PlayWithMachinery/task"
)

type DelayTask struct{}

func (d *DelayTask) UpdateField(args string) error {
	arg := task.UpdateFieldArgs{}
	err := json.Unmarshal([]byte(args), &arg)
	if err != nil {
		return err
	}

	return updateField(arg.ObjID, arg.ObjType, arg.FieldName, arg.FieldValue)
}

func (d *DelayTask) SendMsg(args string) error {
	arg := task.SendMsgArgs{}
	err := json.Unmarshal([]byte(args), &arg)
	if err != nil {
		return err
	}
	return sendMsg(arg.From, arg.To, arg.Content)
}

func updateField(objid, objtype, fieldname string, fieldvalue interface{}) error {
	result := fmt.Sprintf("objId: %s, objType: %s, fieldname: %s, fieldValue: %v", objid, objtype, fieldname, fieldvalue)
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(result)

	return nil
}

func sendMsg(from string, to []string, content string) error {
	result := fmt.Sprintf("from: %s, to: %v, content: %s", from, to, content)
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(result)
	return nil
}
