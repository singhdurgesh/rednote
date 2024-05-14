package tasks

import (
	"encoding/base64"
	"encoding/json"

	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/singhdurgesh/rednote/internal/jobs/task_server"
	"github.com/singhdurgesh/rednote/internal/pkg/logger"
)

type AsyncTask interface {
	Name() string
	Run() error
}

func RunAsync(t AsyncTask) error {
	payload, err := json.Marshal(t)

	if err != nil {
		return err
	}

	logger.LogrusLogger.Println("Pushing Job with payload: ", string(payload))
	task := tasks.Signature{
		Name: t.Name(),
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: payload,
			},
		},
	}

	_, err = task_server.Runner.Server.SendTask(&task)

	return err
}

func ProcessTask(t AsyncTask, payload string) (bool, error) {
	logger := logger.LogrusLogger
	logger.Printf("Processing the Task: %v", payload)

	decodedstg, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		logger.Println("Base64 Decoding Error: ", err)
		return false, nil
	}

	err = json.Unmarshal(decodedstg, &t)

	if err != nil {
		logger.Println("JSON Decoding Error: ", err)
		return false, err
	}

	err = t.Run()

	if err != nil {
		return false, err
	}

	return true, nil
}
