package tasks

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/singhdurgesh/rednote/internal/jobs/task_server"
)

// import (
// 	"encoding/json"
// 	"fmt"

// 	"github.com/RichardKnop/machinery/v2/tasks"
// 	"github.com/singhdurgesh/rednote/internal/jobs/task_server"
// )

// type Task[T any] struct {
// 	name          string
// 	encryptedArgs bool
// }

// type taskMessage[T any] struct {
// 	Arg T
// }

// func (task *Task[T]) Run(arg T) {

// }

// func (task *Task[T]) RunAsync(arg T) {
// 	data, err := json.Marshal(taskMessage[T]{
// 		Arg: arg,
// 	})
// 	fmt.Printf("Running task: %s\n", task.name)
// 	if err != nil {
// 		panic(err)
// 	}
// 	task_server.Runner.Server.SendTask(&tasks.Signature{
// 		Name: task.name,
// 		Args: []tasks.Arg{
// 			{
// 				Name:  "arg",
// 				Type:  "string",
// 				Value: string(data),
// 			},
// 		},
// 	})
// }

// func CreateTask[T any](taskName string, arg T, encryptedArgs bool) Task[T] {
// 	task := Task[T]{
// 		name:          taskName,
// 		encryptedArgs: encryptedArgs,
// 	}
// 	fmt.Printf("Creating task: %s\n", taskName)
// 	task_server.Runner.Server.RegisterTask(taskName, func(arg string) error {
// 		t := new(taskMessage[T])
// 		err := json.Unmarshal([]byte(arg), &t)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// 	return task
// }

// type AsyncTask[T any] struct {
// 	Name string
// }

// func (a *AsyncTask[T]) Run(arg T) {
// 	fmt.Errorf("Method to be defined to in the Implementation Class")
// }

// func (a *AsyncTask[T]) RunAsync(arg T) {
// 	data, err := json.Marshal()
// }

// NewNotificationTask(name: string, phone: string, otp: int, data: string).Run
// - Will process the task specific to the task
// NewNotificationTask(name: string, phone: string, otp: int, data: string).RunAsync
// - create a new notificationTask object that will have the task nae and build the json from the json from argument
// - Push it to the queue for the record

type NotificationTask struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Otp   int    `json:"otp"`
	Data  string `json:"data"`
}

func (n *NotificationTask) Run() {
	fmt.Println(n.Name, n.Phone, n.Phone, n.Data)
}

func NewNotificationTask(phone string, otp int, data string) *NotificationTask {
	return &NotificationTask{Name: "notification", Phone: phone, Otp: otp, Data: data}
}

func (n *NotificationTask) RunAsync() error {
	payload, err := json.Marshal(n)

	if err != nil {
		return err
	}

	// if n.encryptedArgs {
	// 	encryptedData, err := base64.StdEncoding.EncodeToString(payload)
	// 	payload = json.Marshal()
	// }

	task := tasks.Signature{
		Name: n.Name,
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: payload,
			},
		},
	}

	_, err = task_server.Runner.Server.SendTask(&task)

	if err != nil {
		return err
	}

	return nil
}

func ProcessNotification(data string) (bool, error) {
	decodedstg, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Println(err)
		return false, nil
	}

	task := NotificationTask{}
	err = json.Unmarshal([]byte(decodedstg), &task)

	if err != nil {
		fmt.Println("JSON decoding error", err)
		return false, err
	}

	task.Run()
	fmt.Println(task, "Message Received")

	return true, nil
}
