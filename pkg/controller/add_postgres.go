package controller

import (
	"github.com/ravishankarsrrav/postgres-k8s-operator/pkg/controller/postgres"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, postgres.Add)
}
