package usecase

import (
	"ddd-sample/domain/model"
	"ddd-sample/domain/repository"
)

type TaskUsecase interface {
	Create(user_id int, title, content string) (*model.Task, error)
	FindByID(id int) (*model.Task, error)
	Update(id, user_id int, title, content string) (*model.Task, error)
	Delete(id int) error
}

type taskUsecase struct {
	taskRepository repository.TaskRepository
}

func NewTaskUsecase(taskRepo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepository: taskRepo}
}

func (tu *taskUsecase) Create(user_id int, title, content string) (*model.Task, error) {
	task, err := model.NewTask(user_id, title, content)
	if err != nil {
		return nil, err
	}

	createdTask, err := tu.taskRepository.Create(task)
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

func (tu *taskUsecase) FindByID(id int) (*model.Task, error) {
	foundTask, err := tu.taskRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundTask, nil
}

func (tu *taskUsecase) Update(id, user_id int, title, content string) (*model.Task, error) {

}

func (tu *taskUsecase) Delete(id int) error {
	task, err := tu.taskRepository.FindByID(id)
	if err != nil {
		return err
	}

	err = tu.taskRepository.Delete(task)
	if err != nil {
		return err
	}

	return nil
}
