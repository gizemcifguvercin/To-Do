package service

import (
	"encoding/json"
	"fmt"
	"github.com/gizemcifguvercin/To-Do/src/todo/pkg/application/commands"
	"github.com/gizemcifguvercin/To-Do/src/todo/pkg/application/queries"
	"github.com/gizemcifguvercin/To-Do/src/todo/pkg/config"
	"github.com/gizemcifguvercin/To-Do/src/todo/pkg/domain"
	"github.com/gizemcifguvercin/To-Do/src/todo/pkg/http_client"
)

var (
	baseUrl = config.GetValueByKey("ApiURL")
)

type ToDoService struct {
	client http_client.Client
}

func NewToDoService() *ToDoService {
	return &ToDoService{
		client: http_client.New(baseUrl),
	}
}

func (a *ToDoService) Fetch(query *queries.FetchToDoByIdQuery) (todo *domain.ToDo, err error){
	fetchUrl := baseUrl + fmt.Sprintf("/%s", query.Id)
	response, err := a.client.GetWith(fetchUrl)
	if err != nil{
		return nil, err
	}
	bodyStr, err := a.client.ConvertBodyToString(response);
	if err != nil{
		return nil, err
	}
	defer response.Body.Close()

	json.Unmarshal([]byte(bodyStr), &todo)
	return todo, nil
}

func (a *ToDoService) Delete(command *commands.DeleteToDoByIdCommand) (result bool, err error) {
	deleteUrl := baseUrl + fmt.Sprintf("/%s", command.Id)

	response, err := a.client.DeleteWith(deleteUrl)
	if err != nil{
		return false, err
	}
	defer response.Body.Close()

	return true, nil
}

func (a *ToDoService) Create(todo *commands.CreateToDoCommand) (result bool, err error) {
	response, err := a.client.PostWith(baseUrl, todo)

	if err != nil{
		return false, err
	}
	defer response.Body.Close()
	return true,nil
}
