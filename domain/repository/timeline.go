package repository

import "ddd-sample/domain/model"

type TimelineRepository interface {
	Create(timeline *model.Timeline) (*model.Timeline, error)
	FindByID(id int) (*model.Timeline, error)
	Update(timeline *model.Timeline) (*model.Timeline, error)
	Delete(timeline *model.Timeline) error
}
