package gormodel

import "context"

// Creator for create actions.
type Creator interface {
	Create(context.Context, interface{}) error
	FirstOrCreate(context.Context, interface{}) error
}

// Finder for find actions.
type Finder interface {
	GetByID(context.Context, interface{}, ...uint) error
	Get(context.Context, interface{}) error
}

// Updater for update actions.
type Updater interface {
	UpdateOrCreate(context.Context, interface{}) error
	UpdateAttrs(context.Context, interface{}, map[string]interface{}) error
}

// Remover for remove actions.
type Remover interface {
	Remove(context.Context, interface{}) error
}

// Actions interface for CRUD
type Actions interface {
	Creator
	Finder
	Updater
	Remover
}

// Modeler for table name.
type Modeler interface {
	TableName() string
}
