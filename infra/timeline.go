package infra

import (
	"ddd-sample/domain/model"
	"ddd-sample/domain/repository"

	"gorm.io/gorm"
)

type TimelineRepository struct {
	Conn *gorm.DB
}

func NewTimelineRepository(conn *gorm.DB) repository.TimelineRepository {
	return &TimelineRepository{Conn: conn}
}

func (tr *TimelineRepository) Create(timeline *model.Timeline) (*model.Timeline, error) {
	err := tr.Conn.Create(&timeline).Error
	if err != nil {
		return nil, err
	}

	return timeline, nil
}

func (tr *TimelineRepository) FindByID(id int) (*model.Timeline, error) {
	var timeline *model.Timeline
	err := tr.Conn.Where("id = ?", id).First(&timeline).Error
	if err != nil {
		return nil, err
	}

	return timeline, nil
}

func (tr *TimelineRepository) Update(timeline *model.Timeline) (*model.Timeline, error) {
	err := tr.Conn.Model(&timeline).Updates(&timeline).Error
	if err != nil {
		return nil, err
	}

	return timeline, nil
}

func (tr *TimelineRepository) Delete(timeline *model.Timeline) error {
	err := tr.Conn.Delete(&timeline).Error
	if err != nil {
		return err
	}

	return nil
}
