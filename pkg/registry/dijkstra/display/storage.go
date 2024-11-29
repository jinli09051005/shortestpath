package display

import (
	"context"
	"fmt"

	dijkstraclient "jinli.io/shortestpath/generated/client/clientset/internalversion/typed/dijkstra/internalversion"
	"jinli.io/shortestpath/pkg/apis/dijkstra"
	printerstorage "jinli.io/shortestpath/pkg/utils/printers/storage"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/kubernetes/pkg/printers"
)

type REST struct {
	*genericregistry.Store
	dijkstraClient dijkstraclient.DijkstraInterface
}

var _ rest.ShortNamesProvider = &REST{}
var _ rest.CategoriesProvider = &REST{}

// ShortNames implements the ShortNamesProvider interface. Returns a list of short names for a resource.
func (r *REST) ShortNames() []string {
	return []string{"dp"}
}

// Categories implements the CategoriesProvider interface. Returns a list of categories a resource is part of.
func (r *REST) Categories() []string {
	return []string{"all"}
}

type Storage struct {
	Display *REST
	Status  *StatusREST
}

// display资源REST策略等实现
func NewStorage(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter, dijkstraclient dijkstraclient.DijkstraInterface) (*Storage, error) {
	strategy := NewStrategy(scheme, dijkstraclient)

	store := &genericregistry.Store{
		NewFunc:                   func() runtime.Object { return &dijkstra.Display{} },
		NewListFunc:               func() runtime.Object { return &dijkstra.DisplayList{} },
		PredicateFunc:             MatchDisplay,
		DefaultQualifiedResource:  dijkstra.Resource("displays"),
		SingularQualifiedResource: dijkstra.Resource("dispaly"),

		CreateStrategy:      strategy,
		UpdateStrategy:      strategy,
		DeleteStrategy:      strategy,
		ResetFieldsStrategy: strategy,
		// 将API响应对象转换成表格格式,使得输出更加清晰和易于理解
		TableConvertor: printerstorage.TableConvertor{TableGenerator: printers.NewTableGenerator().With(AddHandlers)},
	}

	options := &generic.StoreOptions{
		RESTOptions: optsGetter,
		AttrFunc:    GetAttrs,
		Indexers:    Indexers(),
	}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}

	statusStrategy := NewStatusStrategy(strategy)
	statusStore := *store
	statusStore.UpdateStrategy = statusStrategy
	statusStore.ResetFieldsStrategy = statusStrategy

	return &Storage{
		Display: &REST{
			store,
			dijkstraclient,
		},
		Status: &StatusREST{
			store: &statusStore,
		},
	}, nil
}

type StatusREST struct {
	store *genericregistry.Store
}

var _ rest.Storage = &StatusREST{}
var _ rest.Getter = &StatusREST{}
var _ rest.Patcher = &StatusREST{}
var _ rest.Updater = &StatusREST{}
var _ rest.TableConvertor = &StatusREST{}

// genericregistry.Store实现了CRUD，如果想定义自己的，可以在这里覆盖
// New returns empty Display object.
func (r *StatusREST) New() runtime.Object {
	return r.store.NewFunc()
	// return &dijkstra.Display{}
}

// Destroy cleans up resources on shutdown.
func (r *StatusREST) Destroy() {
	// Given that underlying store is shared with REST,
	// we don't destroy it here explicitly.
}

// Get retrieves the object from the storage. It is required to support Patch.
func (r *StatusREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return r.store.Get(ctx, name, options)
}

// Update alters the status subset of an object.
func (r *StatusREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	// We are explicitly setting forceAllowCreate to false in the call to the underlying storage because
	// subresources should never allow create on update.
	return r.store.Update(ctx, name, objInfo, createValidation, updateValidation, false, options)
}

func (r *StatusREST) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	return r.store.ConvertToTable(ctx, object, tableOptions)
}

func RESTInPeace(storage *Storage, err error) *Storage {
	if err != nil {
		err = fmt.Errorf("unable to create REST storage for a resource,due to %v,will die", err)
		panic(err)
	}

	return storage
}
