package main

import "context"

func main() {
	defaultAdminId := "admin"
	defaultColId := "todo"
	completedColId := "done"
	taskCreatedFlow := createTaskCreatedFlow(defaultAdminId, defaultColId)
	taskCompletedFlow := createTaskCompletedFlow(completedColId)

	// on task created
	taskCreatedFlow.Run(context.Background(), &TaskCreatedEvent{
		TaskId:    "1",
		TaskTitle: "Task 1",
	})

	// on task completed
	taskCompletedFlow.Run(context.Background(), &TaskCompletedEvent{
		TaskId:  "1",
		AdminId: "admin",
	})

}
