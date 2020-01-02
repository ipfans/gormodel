package gormodel_test

import (
	"context"
	"testing"

	"github.com/ipfans/gormodel"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestExtractAndBind(t *testing.T) {
	db := &gorm.DB{}
	assert.NotNil(t, db)
	ctx := gormodel.BindContext(context.TODO(), db)
	newdb := gormodel.ExtractContext(ctx)
	assert.NotNil(t, newdb)
}
