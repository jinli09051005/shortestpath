package v2

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
)

func addConversionFuncs(scheme *runtime.Scheme) error {
	funcs := []func(scheme *runtime.Scheme) error{
		AddFieldLabelConversionForKnownNodes,
		AddFieldLabelConversionForDisplay,
	}
	for _, f := range funcs {
		if err := f(scheme); err != nil {
			return err
		}
	}

	return nil
}

// AddFieldLabelConversionForKnownNodes adds a conversion function to convert
// field selectors of KnownNodes from the given version to internal version
// representation.
func AddFieldLabelConversionForKnownNodes(scheme *runtime.Scheme) error {
	return scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("KnownNodes"),
		func(label, value string) (string, string, error) {
			switch label {
			case "spec.nodeIdentity",
				"spec.nodes",
				"metadata.namespace",
				"metadata.name":
				return label, value, nil
			default:
				return "", "", fmt.Errorf("field label not supported: %s for v2", label)
			}
		})
}

// AddFieldLabelConversionForDisplay adds a conversion function to convert
// field selectors of Display from the given version to internal version
// representation.
func AddFieldLabelConversionForDisplay(scheme *runtime.Scheme) error {
	mapping := map[string]string{
		"spec.nodeIdentity":  "spec.nodeIdentity",
		"spec.startNode":     "spec.startNode",
		"spec.algorithm":     "spec.algorithm",
		"metadata.namespace": "metadata.namespace",
		"metadata.name":      "metadata.name",
	}
	return scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("Display"),
		func(label, value string) (string, string, error) {
			mappedLabel, ok := mapping[label]
			if !ok {
				return "", "", fmt.Errorf("field label not supported: %s for v2", label)
			}
			return mappedLabel, value, nil
		},
	)
}
