package model

import "errors"

type Task struct {
	ID      int
	UserID  int
	Title   string
	Content string
	User    User
}

func NewTask(user_id int, title string, content string) (*Task, error) {
	if title == "" {
		return nil, errors.New("please enter the title")
	}

	task := &Task{
		UserID:  user_id,
		Title:   title,
		Content: content,
	}

	return task, nil
}

func (t *Task) Set(title string, content string) error {
	if title == "" {
		return errors.New("please enter the title")
	}

	t.Title = title
	t.Content = content
	return nil
}
