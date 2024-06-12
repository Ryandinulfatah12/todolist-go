package repositories

import (
    "errors"
    "todolist-go/config"
    "todolist-go/models"
)

func CreateTodo(todo *models.Todo) error {
    return config.DB.Create(todo).Error
}

func GetTodos() ([]models.Todo, error) {
    var todos []models.Todo
    err := config.DB.Preload("Tasks").Find(&todos).Error
    return todos, err
}

func GetTodoByID(id uint) (models.Todo, error) {
    var todo models.Todo
    err := config.DB.Preload("Tasks").First(&todo, id).Error
    return todo, err
}

func UpdateTodo(todo *models.Todo) error {
    var existingTodo models.Todo
    // Check if another todo with the same title exists
    if err := config.DB.Where("title = ? AND id != ?", todo.Title, todo.ID).First(&existingTodo).Error; err == nil {
        return errors.New("a todo with the same title already exists")
    }
    return config.DB.Save(todo).Error
}

func DeleteTodo(todo *models.Todo) error {
    return config.DB.Delete(todo).Error
}
