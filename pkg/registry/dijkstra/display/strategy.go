package display

import (
	"context"
	"fmt"

	dijkstraclient "jinli.io/shortestpath/generated/client/clientset/internalversion/typed/dijkstra/internalversion"
	"jinli.io/shortestpath/pkg/apis/dijkstra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
	"sigs.k8s.io/structured-merge-diff/v4/fieldpath"
)

type Strategy struct {
	runtime.ObjectTyper
	names.NameGenerator
	dijkstraClient dijkstraclient.DijkstraInterface
}

var _ rest.RESTCreateStrategy = &Strategy{}
var _ rest.RESTUpdateStrategy = &Strategy{}
var _ rest.RESTDeleteStrategy = &Strategy{}
var _ rest.GarbageCollectionDeleteStrategy = &Strategy{}

// NewStrategy creates and returns a updateconfigStrtegy instance
func NewStrategy(typer runtime.ObjectTyper, dijkstraClient dijkstraclient.DijkstraInterface) Strategy {
	return Strategy{typer, names.SimpleNameGenerator, dijkstraClient}
}

// DefaultGarbageCollectionPolicy returns DeleteDependents for all currently served versions.
func (s Strategy) DefaultGarbageCollectionPolicy(ctx context.Context) rest.GarbageCollectionPolicy {
	return rest.DeleteDependents
}

// Generic
func (s Strategy) NamespaceScoped() bool {
	return true
}

func (s Strategy) GetResetFields() map[fieldpath.APIVersion]*fieldpath.Set {
	fields := map[fieldpath.APIVersion]*fieldpath.Set{
		"dijkstra.jinli.io/v1": fieldpath.NewSet(
			fieldpath.MakePathOrDie("status"),
		),
		// 新增v2版本
		"dijkstra.jinli.io/v2": fieldpath.NewSet(
			fieldpath.MakePathOrDie("status"),
		),
	}

	return fields
}

func (s Strategy) Canonicalize(obj runtime.Object) {
	//TODO
}

// NodeIdentityIndexFunc return value spec.nodeIdentity of given object.
func NodeIdentityIndexFunc(obj interface{}) ([]string, error) {
	dp, ok := obj.(*dijkstra.Display)
	if !ok {
		return nil, fmt.Errorf("not a dp")
	}
	return []string{dp.Spec.NodeIdentity}, nil
}

// Indexers returns the indexers for dp storage.
func Indexers() *cache.Indexers {
	return &cache.Indexers{
		storage.FieldIndex("spec.nodeIdentity"): NodeIdentityIndexFunc,
	}
}

// rest.RESTCreateStrategy
func (s Strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	display := obj.(*dijkstra.Display)
	display.Status = dijkstra.DisplayStatus{}
	// 添加标签
	labels := make(map[string]string)
	labels["nodeIdentity"] = display.Spec.NodeIdentity
	display.Labels = labels
}

func (s Strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	dp := obj.(*dijkstra.Display)
	var fielderr *field.Error
	// 使用validation,字段验证在前
	errList := ValidateDisplay(dp)

	// 设置要查询的字段选择器
	fieldSelector := fields.ParseSelectorOrDie("spec.nodeIdentity=" + dp.Spec.NodeIdentity)

	//创建Display前相同NodeIdentity的KnownNodes需要创建
	kns, err := s.dijkstraClient.KnownNodeses(dp.Namespace).List(ctx, metav1.ListOptions{
		FieldSelector: fieldSelector.String(),
	})
	if err != nil {
		klog.Error(err)
		fielderr = field.Invalid(field.NewPath("spec", "nodeIdentity"), dp.Spec.NodeIdentity, "KnownNodes with the same nodeIdentity could not be found!")
	} else if len(kns.Items) == 0 {
		fielderr = field.Invalid(field.NewPath("spec", "nodeIdentity"), dp.Spec.NodeIdentity, "KnownNodes with the same nodeIdentity need to be created before creating Display!")
	}

	// 相同NodeIdentity的StartNode不允许相同
	dps, err := s.dijkstraClient.Displays(dp.Namespace).List(ctx, metav1.ListOptions{
		FieldSelector: fieldSelector.String(),
	})
	if err != nil {
		klog.Error(err)
		fielderr = field.Invalid(field.NewPath("spec", "nodeIdentity"), dp.Spec.NodeIdentity, "Display with the same nodeIdentity could not be found!")
	} else if len(dps.Items) != 0 {
		for i := range dps.Items {
			if dps.Items[i].Spec.StartNode.ID != dp.Spec.StartNode.ID {
				continue
			}
			fielderr = field.Invalid(field.NewPath("spec", "startNode"), dp.Spec.StartNode, "Display with the same nodeIdentity and startNode.ID not allow!")
			break
		}
	}

	if fielderr == nil {
		return errList
	}
	return append(errList, fielderr)
}

func (s Strategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	//TODO
	return nil
}

// rest.RESTUpdateStrategy
func (s Strategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	//TODO
}

func (s Strategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	newDisplay := obj.(*dijkstra.Display)
	oldDisplay := old.(*dijkstra.Display)

	return ValidateDisplayUpdate(newDisplay, oldDisplay)
}

func (s Strategy) WarningsOnUpdate(ctx context.Context, obj, old runtime.Object) []string {
	//TODO
	return nil
}

func (s Strategy) AllowCreateOnUpdate() bool {
	return false
}

func (s Strategy) AllowUnconditionalUpdate() bool {
	return false
}

func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	dp, ok := obj.(*dijkstra.Display)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a Display")
	}
	return labels.Set(dp.ObjectMeta.Labels), SelectableFields(dp), nil
}

func SelectableFields(obj *dijkstra.Display) fields.Set {
	// dpSpecificFieldsSet := make(fields.Set, 1)
	// dpSpecificFieldsSet["spec.nodeIdentity"] = string(obj.Spec.NodeIdentity)
	// return generic.AddObjectMetaFieldsSet(dpSpecificFieldsSet, &obj.ObjectMeta, true)

	objectMetaFieldsSet := generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
	dpSpecificFieldsSet := fields.Set{
		"spec.nodeIdentity": string(obj.Spec.NodeIdentity),
	}
	return generic.MergeFieldsSets(objectMetaFieldsSet, dpSpecificFieldsSet)
}

func MatchDisplay(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:       label,
		Field:       field,
		GetAttrs:    GetAttrs,
		IndexLabels: nil,
		IndexFields: []string{
			"spec.nodeIdentity",
			"metadata.namespace",
			"metadata.name"},
		Limit:               0,
		Continue:            "",
		AllowWatchBookmarks: false,
	}
}

type StatusStrategy struct {
	Strategy
}

// NewStrategy creates and returns a updateconfigStrtegy instance
func NewStatusStrategy(strategy Strategy) StatusStrategy {
	return StatusStrategy{strategy}
}

// GetResetFields returns the set of fields that get reset by the strategy
// and should not be modified by the user.
func (st StatusStrategy) GetResetFields() map[fieldpath.APIVersion]*fieldpath.Set {
	return map[fieldpath.APIVersion]*fieldpath.Set{
		"dijkstra.jinli.io/v1": fieldpath.NewSet(
			fieldpath.MakePathOrDie("spec"),
			fieldpath.MakePathOrDie("metadata", "labels"),
		),
		// 新增v2版本
		"dijkstra.jinli.io/v2": fieldpath.NewSet(
			fieldpath.MakePathOrDie("spec"),
			fieldpath.MakePathOrDie("metadata", "labels"),
		),
	}
}

// PrepareForUpdate clears fields that are not allowed to be set by end users on update of status
func (st StatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	newDisplay := obj.(*dijkstra.Display)
	oldDisplay := old.(*dijkstra.Display)
	newDisplay.Spec = oldDisplay.Spec
	newDisplay.Labels = oldDisplay.Labels
}

// // ValidateUpdate is the default update validation for an end user updating status
func (st StatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	//TODO
	return field.ErrorList{}
}

// WarningsOnUpdate returns warnings for the given update.
func (st StatusStrategy) WarningsOnUpdate(ctx context.Context, obj, old runtime.Object) []string {
	return nil
}

func (st StatusStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	dp := obj.(*dijkstra.Display)
	return ValidateDisplayStatus(dp)
}
