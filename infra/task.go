package infra

import (
	"ddd-sample/domain/model"
	"ddd-sample/domain/repository"

	"gorm.io/gorm"
)

type TaskRepository struct {
	Conn *gorm.DB
}

func NewTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return &TaskRepository{Conn: conn}
}

func (tr *TaskRepository) Create(task *model.Task) (*model.Task, error) {
	err := tr.Conn.Create(&task).Error
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (tr *TaskRepository) FindByID(id int) (*model.Task, error) {
	var task *model.Task
	err := tr.Conn.Where("id = ?", id).First(&task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (tr *TaskRepository) Update(task *model.Task) (*model.Task, error) {
	err := tr.Conn.Model(&task).Updates(&task).Error
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (tr *TaskRepository) Delete(task *model.Task) error {
	err := tr.Conn.Delete(&task).Error
	if err != nil {
		return err
	}

	return nil
}
