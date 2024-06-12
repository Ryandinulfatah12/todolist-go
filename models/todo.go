package models

import (
    "gorm.io/gorm"
)

type Todo struct {
    ID        uint       `gorm:"primarykey" json:"id"`
    Title     string     `gorm:"unique;not null" json:"title"`
    Tasks     []Task     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"tasks"`
    Completed bool       `json:"completed"`
    Deleted   gorm.DeletedAt `gorm:"index" json:"-"`
}

type Task struct {
    ID     uint   `gorm:"primarykey" json:"id"`
    TodoID uint   `json:"todo_id"`
    Title  string `json:"title"`
    Done   bool   `json:"done"`
}
