package gormodel_test

import (
	"github.com/ipfans/gormodel"
)

func Example() {
	type User struct {
		gormodel.Model

		Name string
	}

	type UserRepository struct {
		gormodel.Base
	}

	type UserRepo interface{
		gormodel.IBase
	}
}
