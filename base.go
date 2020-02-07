package gormodel

import "context"

// Base operations helpers.
type Base struct {
}

// Create action
func (b *Base) Create(ctx context.Context, i interface{}) error {
	db := ExtractContext(ctx)
	return db.Create(i).Error
}

// FirstOrCreate create if exists or return first match
func (b *Base) FirstOrCreate(ctx context.Context, i interface{}) error {
	db := ExtractContext(ctx)
	return db.Create(i).Error
}

// GetByID return records with given id.
func (b *Base) GetByID(ctx context.Context, i interface{}, id ...uint) error {
	if len(id) == 0 {
		return nil
	}
	db := ExtractContext(ctx)
	if len(id) == 1 {
		return db.Where("id == ?", id[0]).First(i).Error
	}
	return db.Where("id IN (?)", id).Find(i).Error
}

// Get return record with first match.
func (b *Base) Get(ctx context.Context, i interface{}) error {
	db := ExtractContext(ctx)
	return db.Where(i).First(i).Error
}

// UpdateOrCreate record
func (b *Base) UpdateOrCreate(ctx context.Context, i interface{}) error {
	db := ExtractContext(ctx)
	return db.Save(i).Error
}

// UpdateAttrs updates given attrs.
func (b *Base) UpdateAttrs(ctx context.Context, i interface{}, attrs map[string]interface{}) error {
	db := ExtractContext(ctx)
	return db.Model(i).Where(i).Updates(attrs).Error
}

// Remove record
func (b *Base) Remove(ctx context.Context, i interface{}) error {
	db := ExtractContext(ctx)
	return db.Delete(i).Error
}
