package gormodel

import (
	"context"

	"github.com/jinzhu/gorm"
)

type contextKey struct{}

var ctxKey = contextKey{}

// ExtractContext extracts db from context.Context.
func ExtractContext(ctx context.Context) (db *gorm.DB) {
	if v := ctx.Value(ctxKey); v != nil {
		db = v.(*gorm.DB)
	}
	return
}

// BindContext bind db to context.Context.
func BindContext(ctx context.Context, db *gorm.DB) (c context.Context) {
	c = context.WithValue(ctx, ctxKey, db)
	return
}
