package main

import (
	"context"

	"github.com/9ssi7/wfe"
)

type TaskCreatedEvent struct {
	TaskId    string
	ColId     string
	AdminId   string
	TaskTitle string
}

func createTaskCreatedFlow(defaultUserId string, defaultColId string) wfe.Flow[*TaskCreatedEvent] {
	flow := wfe.New[*TaskCreatedEvent]("task_created")
	flow.AddAction("assigne_defaults", func(ctx context.Context, p *TaskCreatedEvent) error {
		if p.AdminId == "" {
			p.AdminId = defaultUserId
		}
		if p.ColId == "" {
			p.ColId = defaultColId
		}
		return nil
	})
	flow.AddAction("send_notify", func(ctx context.Context, p *TaskCreatedEvent) error {
		println("New task created by", p.AdminId, "in column", p.ColId+":", p.TaskTitle)
		return nil
	})
	flow.AddNode(wfe.NewNode("assigne_defaults"), wfe.NewNode("send_notify"))
	return flow
}
