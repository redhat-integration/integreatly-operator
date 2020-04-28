package controller

import (
	"github.com/redhat-integration/rhi-operator/pkg/controller/installation"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, installation.Add)
}
