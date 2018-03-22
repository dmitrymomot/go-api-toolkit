package repository

import (
	"github.com/jinzhu/gorm"
)

// BaseRepository structure
type BaseRepository struct {
	DB *gorm.DB
}

// GetByID function fetches model by ID
func (r *BaseRepository) GetByID(out interface{}, id string, relations ...string) error {
	q := r.DB
	for _, relation := range relations {
		q = q.Preload(relation)
	}
	return q.Find(out, "id = ?", id).Error
}

// All function fetches all models list
func (r *BaseRepository) All(out interface{}, relations []string, query interface{}, args ...interface{}) error {
	q := r.DB
	for _, relation := range relations {
		q = q.Preload(relation)
	}
	if query != nil && query != "" {
		q = q.Where(query, args...)
	}
	return q.Unscoped().Find(out).Error
}

// List function fetches models list
// If the limit parameter equal 0, it be set as 18
// Why 18? Because this number can be divided by 2 and 3 without residue :)
func (r *BaseRepository) List(out interface{}, relations []string, limit, offset int, query interface{}, args ...interface{}) error {
	q := r.DB
	if limit == 0 {
		limit = 18
	}
	for _, relation := range relations {
		q = q.Preload(relation)
	}
	if query != nil && query != "" {
		q = q.Where(query, args...)
	}
	return q.Limit(limit).Offset(offset).Find(out).Error
}

// First function fetches the first model from list by given conditions
func (r *BaseRepository) First(out interface{}, relations []string, where ...interface{}) error {
	q := r.DB
	for _, relation := range relations {
		q = q.Preload(relation)
	}
	return q.First(out, where).Error
}

// Last function fetches the last model from list by given conditions
func (r *BaseRepository) Last(out interface{}, relations []string, where ...interface{}) error {
	q := r.DB
	for _, relation := range relations {
		q = q.Preload(relation)
	}
	return q.Last(out, where).Error
}

// Create a new model
func (r *BaseRepository) Create(model interface{}) error {
	return r.DB.Create(model).Error
}

// Save a model
func (r *BaseRepository) Save(model interface{}) error {
	return r.DB.Save(model).Error
}

// Update a model
func (r *BaseRepository) Update(model interface{}, data map[string]interface{}) error {
	return r.DB.Model(model).Updates(data).Error
}

// Delete a model
func (r *BaseRepository) Delete(model interface{}) error {
	return r.DB.Delete(model).Error
}

// ForceDelete models
func (r *BaseRepository) ForceDelete(model interface{}) error {
	return r.DB.Unscoped().Delete(model).Error
}

// Related function fetches related models
func (r *BaseRepository) Related(model interface{}, out interface{}, foreignKeys ...string) error {
	return r.DB.Model(model).Related(out, foreignKeys...).Error
}
