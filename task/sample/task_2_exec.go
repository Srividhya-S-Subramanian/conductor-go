package sample

import (
	"github.com/netflix/conductor/client/go/model"
	log "github.com/sirupsen/logrus"
)

func Task_2_Execution_Function(t *model.Task) (taskResult *model.TaskResult, err error) {
	log.Println("Executing Task_2_Execution_Function for", t.TaskType)

	//Do some logic
	taskResult = model.NewTaskResult(t)

	output := map[string]interface{}{
		"task": "task_2",
		"key2": "value2",
		"key3": 3,
		"key4": false,
	}
	taskResult.OutputData = output
	taskResult.Status = "COMPLETED"
	err = nil

	return taskResult, err
}