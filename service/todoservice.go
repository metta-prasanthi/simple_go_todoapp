package service

import (
    "context"

    "../repositories"
    "../models"
)

func Close(ctx context.Context) {
    repositories.Close(ctx)
}

func Insert(ctx context.Context, todo *models.Todo) (int, error) {
    return repositories.Insert(ctx, todo)
}

func Delete(ctx context.Context, id int) error {
    return repositories.Delete(ctx, id)
}

func GetAll(ctx context.Context) ([]models.Todo, error) {
    return repositories.GetAll(ctx)
}