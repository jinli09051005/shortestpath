package knownnodes

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
	"k8s.io/klog"
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
var _ rest.ResetFieldsStrategy = &Strategy{}

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
	kn, ok := obj.(*dijkstra.KnownNodes)
	if !ok {
		return nil, fmt.Errorf("not a kn")
	}
	return []string{kn.Spec.NodeIdentity}, nil
}

// Indexers returns the indexers for pod storage.
func Indexers() *cache.Indexers {
	return &cache.Indexers{
		storage.FieldIndex("spec.nodeIdentity"): NodeIdentityIndexFunc,
	}
}

// rest.RESTCreateStrategy
func (s Strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	knownNodes := obj.(*dijkstra.KnownNodes)
	knownNodes.Status = dijkstra.KnownNodesStatus{}
	// 添加标签
	labels := make(map[string]string)
	labels["nodeIdentity"] = knownNodes.Spec.NodeIdentity
	knownNodes.Labels = labels
}

func (s Strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	kn := obj.(*dijkstra.KnownNodes)
	var fielderr *field.Error
	// 使用validation,字段验证在前
	errList := ValidateKnownNodes(kn)

	// 设置要查询的字段选择器
	fieldSelector := fields.ParseSelectorOrDie("spec.nodeIdentity=" + kn.Spec.NodeIdentity)

	//不允许创建NodeIdentity相同KnownNodes
	kns, err := s.dijkstraClient.KnownNodeses(kn.Namespace).List(ctx, metav1.ListOptions{
		FieldSelector: fieldSelector.String(),
	})
	if err != nil {
		klog.Error(err)
		fielderr = field.Invalid(field.NewPath("spec", "nodeIdentity"), kn.Spec.NodeIdentity, "KnownNodes with the same nodeIdentity could not be found!")
	} else if len(kns.Items) != 0 {
		fielderr = field.Invalid(field.NewPath("spec", "nodeIdentity"), kn.Spec.NodeIdentity, "Exist knownNodes with the same nodeIdentity!")
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
	newKnownNodes := obj.(*dijkstra.KnownNodes)
	oldKnownNodes := old.(*dijkstra.KnownNodes)

	return ValidateDisplayUpdate(newKnownNodes, oldKnownNodes)
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
	kn, ok := obj.(*dijkstra.KnownNodes)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a KnownNodes")
	}
	return labels.Set(kn.ObjectMeta.Labels), SelectableFields(kn), nil
}

func SelectableFields(obj *dijkstra.KnownNodes) fields.Set {
	// knSpecificFieldsSet := make(fields.Set, 1)
	// knSpecificFieldsSet["spec.nodeIdentity"] = string(obj.Spec.NodeIdentity)
	// return generic.AddObjectMetaFieldsSet(knSpecificFieldsSet, &obj.ObjectMeta, true)

	objectMetaFieldsSet := generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
	knSpecificFieldsSet := fields.Set{
		"spec.nodeIdentity": string(obj.Spec.NodeIdentity),
	}
	return generic.MergeFieldsSets(objectMetaFieldsSet, knSpecificFieldsSet)
}

func MatchKnownNodes(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
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

var _ rest.RESTUpdateStrategy = &StatusStrategy{}
var _ rest.ResetFieldsStrategy = &StatusStrategy{}

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
	newKnownNodes := obj.(*dijkstra.KnownNodes)
	oldKnownNodes := old.(*dijkstra.KnownNodes)
	newKnownNodes.Spec = oldKnownNodes.Spec
	newKnownNodes.Labels = oldKnownNodes.Labels
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
