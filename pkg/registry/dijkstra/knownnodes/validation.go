package knownnodes

import (
	"fmt"
	"reflect"

	"jinli.io/shortestpath/pkg/apis/dijkstra"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func ValidateKnownNodes(kn *dijkstra.KnownNodes) field.ErrorList {
	errors := field.ErrorList{}

	if fmt.Sprintf("%v", reflect.TypeOf(kn.Spec.NodeIdentity)) != "string" {
		errors = append(errors, field.Invalid(field.NewPath("spec", "nodeIdentity"), kn.Spec.NodeIdentity, "must be string"))
	}
	return errors
}

func ValidateKnownNodesUpdate(new, old *dijkstra.KnownNodes) field.ErrorList {
	errors := field.ErrorList{}
	// 不允许修改nodeIdentity
	if new.Spec.NodeIdentity != old.Spec.NodeIdentity {
		errors = append(errors, field.Invalid(field.NewPath("spec", "nodeIdentity"), old.Spec.NodeIdentity, "No modifications allowed"))
	}

	return errors
}
