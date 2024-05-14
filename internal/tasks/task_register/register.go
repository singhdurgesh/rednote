package task_register

import (
	"github.com/singhdurgesh/rednote/cmd/app"
	"github.com/singhdurgesh/rednote/internal/tasks/notifications"
)

func RegisterTasks() error {
	err := app.Broker.RegisterTasks(map[string]interface{}{
		notifications.NotificationTaskName:  notifications.ProcessNotification,
		notifications.CommunicationTaskName: notifications.ProcessCommunication,
	})

	if err != nil {
		panic(err)
	}

	return nil
}
