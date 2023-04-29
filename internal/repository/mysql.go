package repository

import (
	"context"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Create(ctx context.Context, model interface{}) error {
	result := r.db.WithContext(ctx).Create(model)
	return result.Error
}

func (r *Repository) FindAll(ctx context.Context, models interface{}) error {
	result := r.db.WithContext(ctx).Find(models)
	return result.Error
}

func (r *Repository) FindByID(ctx context.Context, model interface{}, id uint) error {
	result := r.db.WithContext(ctx).First(model, id)
	return result.Error
}

func (r *Repository) Update(ctx context.Context, model interface{}) error {
	result := r.db.WithContext(ctx).Save(model)
	return result.Error
}

func (r *Repository) Delete(ctx context.Context, model interface{}, id uint) error {
	result := r.db.WithContext(ctx).Delete(model, id)
	return result.Error
}

func (r *Repository) GetByAdvanceQuery(models interface{}, query AdvanceQuery) error {
	var err error
	db := r.db

	for _, join := range query.Joins {
		db = db.Joins(join)
	}

	for _, where := range query.Wheres {
		db = db.Where(where.Column+" "+string(where.Operator)+" ?", where.Value)
	}

	for _, preload := range query.Preloads {
		db = db.Preload(preload)
	}

	if query.Select != "" {
		db = db.Select(query.Select)
	}

	if query.FindType == FindAll {
		err = db.Find(models).Error
	} else {
		err = db.First(models).Error
	}

	return err
}
