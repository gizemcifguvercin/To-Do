package main

import (
	"github.com/gizemcifguvercin/To-Do/src/todo/pkg/application/commands"
	"github.com/gizemcifguvercin/To-Do/src/todo/pkg/application/queries"
	error_pkg "github.com/gizemcifguvercin/To-Do/src/todo/pkg/error"
	"github.com/gizemcifguvercin/To-Do/src/todo/pkg/service"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ToDoId string

func TestCreateToDo_ShouldBeFail_WhenUnCorrectParameters(t *testing.T) {
	command := new(commands.CreateToDoCommand)
	uuid , _ := uuid.NewV4()
	command.Id =  uuid.String()
	command.Type = "Work"

	todoService := service.NewToDoService()
	_, err := todoService.Create(command)
	if err != nil {
		custom_error, _ := err.(*error_pkg.CustomError)
		assert.Equal(t, custom_error.Code, 400, "Name is required")
	}
}

func TestCreateToDo_ShouldBeSuccessful_WhenCorrectParameters(t *testing.T) {
	command := new(commands.CreateToDoCommand)
	uuid , _ := uuid.NewV4()
	command.Id =  uuid.String()
	command.Type = "Work"
	command.Name = "Write Unit Tests"

	todoService := service.NewToDoService()
	data, _ := todoService.Create(command)
	ToDoId = command.Id
	assert.Equal(t, data, true)
}

func TestFetchToDo_ShouldBeFail_WhenTypeOfToDoIdIsntCorrect(t *testing.T) {
	id := "3"
	query := new(queries.FetchToDoByIdQuery)
	query.Id = id

	ToDoService := service.NewToDoService()
	_, err := ToDoService.Fetch(query)
	if err != nil {
		custom_error, _ := err.(*error_pkg.CustomError)
		assert.Equal(t, custom_error.Code, 400, "ToDo Id validation error")
	}
}

func TestFetchToDo_ShouldReturnRelatedToDo_WhenToDoExists(t *testing.T) {
	query := new(queries.FetchToDoByIdQuery)
	query.Id = ToDoId

	ToDoService := service.NewToDoService()
	data, _ := ToDoService.Fetch(query)
	assert.Equal(t, data.Id, ToDoId)
}

func TestDeleteToDo_ShouldBeSuccessful_WhenToDoExists(t *testing.T) {
	command := new(commands.DeleteToDoByIdCommand)
	command.Id = ToDoId

	ToDoService := service.NewToDoService()
	data, _ := ToDoService.Delete(command)
	assert.Equal(t, data, true)
}

func TestDeleteToDo_ShouldBeFail_WhenToDoDoesntExist(t *testing.T) {
	command := new(commands.DeleteToDoByIdCommand)
	uuid , _ := uuid.NewV4()
	command.Id = uuid.String()

	ToDoService := service.NewToDoService()
	data, _ := ToDoService.Delete(command)
	assert.Equal(t, data, false, "ToDo doesn't exists")
}
