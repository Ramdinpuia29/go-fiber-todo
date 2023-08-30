package repositories

import (
	"rdp/gotodos/app/models"
	"rdp/gotodos/config/database"

	"gorm.io/gorm/clause"
)

func GetAllTodos(todos *[]models.Todo) (err error) {
	if err = database.DB.Find(todos).Error; err != nil {
		return err
	}
	return nil
}

func GetTodoById(todo *models.Todo, id string) (err error) {
	if err = database.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateTodo(todo *models.Todo) (err error) {
	if err = database.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTodo(todo *models.Todo, id string) (err error) {
	if err = database.DB.Clauses(clause.Returning{}).Where("id = ?", id).Updates(todo).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTodo(todo *models.Todo, id string) (err error) {
	if err = database.DB.Where("id = ?", id).Delete(todo).Error; err != nil {
		return err
	}
	return nil
}
