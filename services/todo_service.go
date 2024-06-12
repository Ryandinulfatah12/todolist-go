package services

import (
    "errors"
    "todolist-go/models"
    "todolist-go/repositories"
    "todolist-go/utils"
)

func CreateTodoService(todo *models.Todo) utils.Response {
    if len(todo.Title) == 0 {
        return utils.ErrorResponse(errors.New("title cannot be empty"))
    }
    if err := repositories.CreateTodo(todo); err != nil {
        return utils.ErrorResponse(err)
    }
    return utils.SuccessResponse(todo)
}

func GetTodosService() utils.Response {
    todos, err := repositories.GetTodos()
    if err != nil {
        return utils.ErrorResponse(err)
    }
    return utils.SuccessResponse(todos)
}

func GetTodoByIDService(id uint) utils.Response {
    todo, err := repositories.GetTodoByID(id)
    if err != nil {
        return utils.ErrorResponse(err)
    }
    return utils.SuccessResponse(todo)
}

func UpdateTodoService(todo *models.Todo) utils.Response {
    if todo.Deleted.Valid || todo.Completed {
        return utils.ErrorResponse(errors.New("cannot update a deleted or completed todo"))
    }
    if err := repositories.UpdateTodo(todo); err != nil {
        return utils.ErrorResponse(err)
    }
    return utils.SuccessResponse(todo)
}

func DeleteTodoService(todo *models.Todo) utils.Response {
    if err := repositories.DeleteTodo(todo); err != nil {
        return utils.ErrorResponse(err)
    }
    return utils.SuccessResponse(nil)
}

func CompleteTodoService(id uint) utils.Response {
    todo, err := repositories.GetTodoByID(id)
    if err != nil {
        return utils.ErrorResponse(err)
    }
    todo.Completed = true
    if err := repositories.UpdateTodo(&todo); err != nil {
        return utils.ErrorResponse(err)
    }
    return utils.SuccessResponse(todo)
}
