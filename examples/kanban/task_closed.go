package main

import (
	"context"

	"github.com/9ssi7/wfe"
)

type TaskCompletedEvent struct {
	TaskId  string
	ColId   string
	AdminId string
	Status  string
}

func createTaskCompletedFlow(completedColId string) wfe.Flow[*TaskCompletedEvent] {
	flow := wfe.New[*TaskCompletedEvent]("task_completed")
	flow.AddAction("move", func(ctx context.Context, p *TaskCompletedEvent) error {
		p.ColId = completedColId // TODO: save to db
		p.Status = "completed"
		return nil
	})
	flow.AddAction("send_notify", func(ctx context.Context, p *TaskCompletedEvent) error {
		println("New task completed by", p.AdminId, "in column", p.ColId+":", p.Status)
		return nil
	})
	flow.AddNode(wfe.NewNode("move"), wfe.NewNode("send_notify"))
	return flow
}
