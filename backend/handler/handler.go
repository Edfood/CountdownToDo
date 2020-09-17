package handler

import (
	"net/http"
	"strconv"

	"github.com/CountdownToDo/backend/model"
	"github.com/labstack/echo"
)

func AddTask(c echo.Context) error {
	task := new(model.Task)
	if err := c.Bind(task); err != nil {
		return err
	}

	if task.Name == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid to or message fields",
		}
	}

	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	task.UID = uid
	model.CreateTask(task)

	return c.JSON(http.StatusCreated, task)
}

func GetTasks(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	tasks := model.FindTasks(&model.Task{UID: uid})
	return c.JSON(http.StatusOK, tasks)
}

func DeleteTask(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	if err := model.DeleteTask(&model.Task{ID: taskID, UID: uid}); err != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}
