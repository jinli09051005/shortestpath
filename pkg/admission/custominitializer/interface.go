package custominitializer

import (
	informers "jinli.io/shortestpath/generated/client/informers/externalversions"
	"k8s.io/apiserver/pkg/admission"
)

const (
	KN = 10000
	DP = 1000000
)

type WantsInformerFactory interface {
	SetInformerFactory(informers.SharedInformerFactory)
	admission.InitializationValidator
	admission.ValidationInterface
}
