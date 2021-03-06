package controller

import (
	"github.com/redhat-developer/openshift-jenkins-operator/pkg/controller/controllerutil"
	"github.com/redhat-developer/openshift-jenkins-operator/pkg/controller/jenkins"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	controllerutil.AddToManagerFuncs = append(controllerutil.AddToManagerFuncs, jenkins.Add)
}
