package main

import (
	"encoding/json"
	"os/exec"
)

type Taskwarrior struct {
}

func NewTaskwarrior() *Taskwarrior {
	return &Taskwarrior{}
}

type Task struct {
	Description string  `json:"description"`
	Project     string  `json:"project"`
	Priority    string  `json:"priority"`
	Urgency     float64 `json:"urgency"`
	Due         string  `json:"due"`
	Scheduled   string  `json:"scheduled"`
	// Tags        string `json:"tags"`
	Uuid string `json:"uuid"`
}

func (t *Taskwarrior) GetTodayTasks() ([]Task, error) {
	cmd := exec.Command("task", "(scheduled.before=tod or due.before=eod) status:pending", "export")
	var tasks []Task
	output, err := cmd.Output()
	if err != nil {
		return tasks, err
	}
	err = json.Unmarshal(output, &tasks)
	return tasks, err
}

// func (t *Taskwarrior) GetContext() (string, error) {
// }
