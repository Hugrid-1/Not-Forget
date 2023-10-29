package entity

import "time"

type TaskType string

var (
	TASK_TYPE_JOB   = TaskType("Job")
	TASK_TYPE_WATCH = TaskType("Watch")
	TASK_TYPE_READ  = TaskType("Read")
)

type Task struct {
	id          int
	name        string
	taskType    TaskType
	description string
	createDate  time.Time
	endDate     time.Time
}

// TODO: Обработка ошибок при создании задачи
func NewTask(
	id int,
	name string,
	taskType string,
	description string,
	createDate time.Time,
	endDate time.Time,
) (Task, error) {
	return Task{
		id:          id,
		name:        name,
		taskType:    TaskType(taskType),
		description: description,
		createDate:  createDate,
		endDate:     endDate,
	}, nil
}

func (t *Task) GetID() int {
	return t.id
}

func (t *Task) GetName() string {
	return t.name
}

func (t *Task) GetTaskType() TaskType {
	return t.taskType
}

func (t *Task) GetDescription() string {
	return t.description
}

func (t *Task) GetCreateDate() time.Time {
	return t.createDate
}

func (t *Task) GetEndDate() time.Time {
	return t.endDate
}
