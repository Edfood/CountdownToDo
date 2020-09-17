package model

import (
	"fmt"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Task struct {
	ID        int    `json:"id" gorm:"praimaly_key"`
	UID       int    `json:"uid"`
	Name      string `json:"name"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type Tasks []Task

func CreateTask(task *Task) {
	db.Create(task)
}

func FindTasks(t *Task) Tasks {
	var tasks Tasks
	db.Where(t).Find(&tasks)
	return tasks
}

func DeleteTask(t *Task) error {
	if rows := db.Where(t).Delete(&Task{}).RowsAffected; rows == 0 {
		return fmt.Errorf("Could not find Task (%v) to delete", t)
	}
	return nil
}

// func UpdateTodo(t *Todo) error {
// 	rows := db.Model(t).Update(map[string]interface{}{
// 			"name": t.Name,
// 			"completed": t.Completed,
// 	}).RowsAffected
// 	if rows == 0 {
// 			return fmt.Errorf("Could not find Todo (%v) to update", t)
// 	}
// 	return nil
// }
