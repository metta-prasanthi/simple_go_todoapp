package repositories

import (
	"time"
	"../models"
)

type Sample struct{}

func (s *Sample) Close() {}

func (s *Sample) Insert(todo *models.Todo) (int, error) {
    return 0, nil
}

func (s *Sample) Delete(id int) error {
    return nil
}

func (s *Sample) GetAll() ([]models.Todo, error) {
    todoList := []models.Todo{
        {
            ID:      1,
            Title:   "Practice GoLang",
            Note:    "Calculator",
            DueDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
        },
        {
            ID:      2,
            Title:   "Do homework",
            Note:    "ToDoApp",
            DueDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
        },
        {
            ID:      3,
            Title:   "Start Final Project",
            Note:    "Finalize Requirements",
            DueDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
        },
    }

    return todoList, nil
}