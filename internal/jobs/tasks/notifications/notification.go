package notifications

import (
	"fmt"

	"github.com/singhdurgesh/rednote/internal/jobs/tasks"
)

type NotificationTask struct {
	Phone string
	Otp   int
	Data  string
}

const NotificationTaskName = "notification"

func (n *NotificationTask) Run() error {
	fmt.Println(n)
	return nil
}

func NewNotificationTask(phone string, otp int, data string) *NotificationTask {
	return &NotificationTask{
		Phone: phone, Otp: otp, Data: data,
	}
}

func (n *NotificationTask) Name() string {
	return NotificationTaskName
}

func ProcessNotification(data string) (bool, error) {
	task := &NotificationTask{}
	return tasks.ProcessTask(task, data)
}
