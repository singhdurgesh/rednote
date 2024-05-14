package task_register

import (
	"github.com/singhdurgesh/rednote/internal/jobs/task_server"
	"github.com/singhdurgesh/rednote/internal/jobs/tasks/notifications"
)

func RegisterTasks() {
	// Register New Tasks here
	err := task_server.Runner.Server.RegisterTasks(map[string]interface{}{
		notifications.NotificationTaskName:  notifications.ProcessNotification,
		notifications.CommunicationTaskName: notifications.ProcessCommunication,
	})

	if err != nil {
		panic(err)
	}
}
