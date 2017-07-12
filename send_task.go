package main

import (
	"encoding/json"
	"time"
	"log"

	"github.com/RichardKnop/machinery/v1/tasks"

	"github.com/zhiruchen/PlayWithMachinery/client"
	"github.com/zhiruchen/PlayWithMachinery/task"
)

func main() {
	err := DelayUpdateField("Order", "1234", "price", 1000000000, 30)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	err = DelaySendMsg("Bob", []string{"Smith", "KK"}, "1231232231123312132123123123123", 60)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
}

func DelayUpdateField(objType, objID, fieldName string, fieldValue interface{}, seconds int64) error {
	eta := time.Now().Add(time.Second * time.Duration(seconds))

	args := task.UpdateFieldArgs{
		ObjID:      objID,
		ObjType:    objType,
		FieldName:  fieldName,
		FieldValue: fieldValue,
	}
	b, err := json.Marshal(&args)
	if err != nil {
		return err
	}

	taskArgs := []tasks.Arg{
		{Type: "string", Value: string(b)},
	}

	sig := task.CreateTask("updatefield", taskArgs, eta, nil, nil)
	err = task.SendTask(sig, new(client.DelayTask))
	return err
}

func DelaySendMsg(from string, to []string, content string, seconds int64) error {
	eta := time.Now().Add(time.Second * time.Duration(seconds))
	args := task.SendMsgArgs{
		From:    from,
		To:      to,
		Content: content,
	}
	b, err := json.Marshal(args)
	if err != nil {
		return err
	}
	taskArgs := []tasks.Arg{
		{Type: "string", Value: string(b)},
	}
	sig := task.CreateTask("sendmsg", taskArgs, eta, nil, nil)
	err = task.SendTask(sig, new(client.DelayTask))
	return err
}
