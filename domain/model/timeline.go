package model

import "errors"

type Timeline struct {
	ID      int
	UserID  int
	Content string
}

func NewTimeline(user_id int, content string) (*Timeline, error) {
	if content == "" {
		return nil, errors.New("please enter the content")
	}

	timeline := &Timeline{
		UserID:  user_id,
		Content: content,
	}

	return timeline, nil
}

func (t *Timeline) Set(content string) error {
	if content == "" {
		return errors.New("please enter the content")
	}

	t.Content = content
	return nil
}
