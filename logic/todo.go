package logic

import (
	"bubble_todolist/dao"
	"bubble_todolist/models"
)

func CreateTask(todo *models.Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetTask(id string) (todo *models.Todo, err error) {
	todo = new(models.Todo)
	if err = dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return
}

func GetTodoList() (todoList []*models.Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateTask(todo *models.Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}

func DeleteTask(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&models.Todo{}).Error
	return
}
