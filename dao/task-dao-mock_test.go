package dao_test // /!\ don't change this! create a new test file instead.

import (
	"github.com/Sfeir/golang-200/dao"
	"github.com/Sfeir/golang-200/model"
	"github.com/satori/go.uuid"
	"testing"
	"time"
)

func TestDAOMock(t *testing.T) {

	daoMock, err := dao.GetTaskDAO("", "", dao.DAOMock)
	if err != nil {
		t.Error(err)
	}

	// TODO create a new task called "toSave" for testing purpose
	toSave := model.Task{
		ID:           uuid.NewV4().String(),
		Title:        "Use Go",
		Description:  "Let's use the Go programming language in a real project.",
		Status:       model.StatusTodo,
		Priority:     model.PriorityMedium,
		CreationDate: time.Date(2017, 02, 01, 0, 0, 0, 0, time.UTC),
		DueDate:      time.Date(2017, 02, 02, 0, 0, 0, 0, time.UTC),
	}

	// TODO save the "toSave" task
	err = daoMock.Save(&toSave)

	// TODO check the error
	if err != nil {
		t.Error(err)
	}

	t.Log("initial task saved", toSave)


	tasks, err := daoMock.GetAll(dao.NoPaging, dao.NoPaging)
	if err != nil {
		t.Error(err)
	}
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
	toSave2 := model.Task{
		ID:           uuid.NewV4().String(),
		Title:        "Use Go Again",
		Description:  "Let's use the Go programming language in a real project.",
		Status:       model.StatusDone,
		Priority:     model.PriorityHigh,
		CreationDate: time.Date(2017, 02, 01, 0, 0, 0, 0, time.UTC),
		DueDate:      time.Date(2017, 02, 02, 0, 0, 0, 0, time.UTC),
	}
	err = daoMock.Save(&toSave2)
	if err != nil {
		t.Error(err)
	}

	// check indexes search
	tasks, err = daoMock.GetAll(0, 0)
	if err != nil {
		t.Error(err)
	}

	if len(tasks) != 1 {
		t.Errorf("Expected 1 tasks, got %d", len(tasks))
	}

	oneTask, err := daoMock.GetByID(toSave.ID)
	if err != nil {
		t.Error(err)
	}
	if toSave != *oneTask {
		t.Error("Got wrong task by ID")
	}

	err = daoMock.Delete(oneTask.ID)
	if err != nil {
		t.Error(err)
	}

	oneTask, err = daoMock.GetByID(oneTask.ID)
	if err == nil {
		t.Error("Task should have been deleted", oneTask)
	}

	// TODO get "toSave" task by ID and verify that it is successfully retrieved


	// TODO delete the "toSave" task and verify with a get by ID that it is removed

}
