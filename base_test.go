package gormodel_test

import (
	"testing"

	"github.com/ipfans/gormodel"
)

func ExampleModel() {
	type User struct {
		gormodel.Model

		Name string
	}
}

type User struct {
	gormodel.Model

	Name string
}

func iFunc(i gormodel.Actions) {

}

func TestBase(t *testing.T) {
	iFunc(&gormodel.Base{})
}
