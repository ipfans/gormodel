package gormodel

import (
	"context"
	"sync"

	"github.com/jinzhu/gorm"
)

// IBase interface
type IBase interface {
	Create(context.Context, interface{}) error
	GetByID(context.Context, interface{}, uint) error
	GetByMultiID(context.Context, interface{}, []uint) error
	Update(context.Context, interface{}, map[string]interface{}) error
	Remove(context.Context, interface{}) error
}

// HookFunc for executions
type HookFunc func(context.Context, string, ...interface{})

// Base operations helpers.
type Base struct {
	lock       sync.RWMutex
	defaultOps map[string]interface{}

	beforeHook func(context.Context, string, ...interface{})
	afterHook  func(context.Context, string, ...interface{})
}

// Option set options to db query.
func (b *Base) Option(k string, v interface{}) {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.defaultOps == nil {
		b.defaultOps = make(map[string]interface{})
	}
	b.defaultOps[k] = v
}

// BeforeHook set hook before execution.
func (b *Base) BeforeHook(fn HookFunc) {
	b.beforeHook = fn
}

// AfterHook set hook after execution.
func (b *Base) AfterHook(fn HookFunc) {
	b.afterHook = fn
}

func (b *Base) perparedDB(ctx context.Context) (db *gorm.DB) {
	db = ExtractContext(ctx)
	b.lock.RLock()
	opts := b.defaultOps
	for k := range opts {
		db = db.Set(k, opts[k])
	}
	b.lock.RUnlock()
	return
}

// Create new record.
func (b *Base) Create(ctx context.Context, i interface{}) (err error) {
	db := b.perparedDB(ctx)
	err = db.Create(i).Error
	return
}

// GetByID returns record by primary key.
func (b *Base) GetByID(ctx context.Context, i interface{}, id uint) (err error) {
	db := b.perparedDB(ctx)
	if b.beforeHook != nil {
		b.beforeHook(ctx, "GetByID", id)
	}
	err = db.Where("id = ?", id).First(i).Error
	if b.afterHook != nil {
		b.beforeHook(ctx, "GetByID", id, i, err)
	}
	return
}

// GetByMultiID returns record by primary key.
func (b *Base) GetByMultiID(ctx context.Context, i interface{}, id []uint) (err error) {
	db := b.perparedDB(ctx)
	if b.beforeHook != nil {
		b.beforeHook(ctx, "GetByMultiID", i, id)
	}
	err = db.Where("id IN (?)", id).Find(i).Error
	if b.afterHook != nil {
		b.beforeHook(ctx, "GetByMultiID", i, id, err)
	}
	return
}

// Update record with given attrs.
func (b *Base) Update(ctx context.Context, i interface{}, attrs map[string]interface{}) (err error) {
	db := b.perparedDB(ctx)
	if b.beforeHook != nil {
		b.beforeHook(ctx, "Update", i, attrs)
	}
	err = db.Model(i).Where(i).Updates(attrs).Error
	if b.afterHook != nil {
		b.beforeHook(ctx, "Update", i, attrs, err)
	}
	return
}

// Remove record with given primary key.
// WARNING: It will remove all records when primary key is default value, e.g. id = 0.
func (b *Base) Remove(ctx context.Context, i interface{}) (err error) {
	db := b.perparedDB(ctx)
	if b.beforeHook != nil {
		b.beforeHook(ctx, "Remove", i)
	}
	err = db.Delete(i).Error
	if b.afterHook != nil {
		b.beforeHook(ctx, "Remove", i, err)
	}
	return
}
