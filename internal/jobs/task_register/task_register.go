package task_register

import (
	"github.com/singhdurgesh/rednote/internal/jobs/task_server"
	"github.com/singhdurgesh/rednote/internal/jobs/tasks"
	"github.com/singhdurgesh/rednote/internal/jobs/tasks/notification"
)

func RegisterTasks() {
	// Register New Tasks here
	err := task_server.Runner.Server.RegisterTasks(map[string]interface{}{
		"send_otp":     notification.SendMail,
		"notification": tasks.ProcessNotification,
	})

	if err != nil {
		panic(err)
	}
}
