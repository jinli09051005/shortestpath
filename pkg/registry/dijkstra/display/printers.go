package display

import (
	"jinli.io/shortestpath/pkg/apis/dijkstra"
	v2 "jinli.io/shortestpath/pkg/apis/dijkstra/v2"
	printerstorage "jinli.io/shortestpath/pkg/utils/printers/storage"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/printers"
)

// AddHandlers adds print handlers for default Kubernetes types dealing with internal versions.
func AddHandlers(h printers.PrintHandler) {
	dpColumnDefinitions := []metav1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "NodeIdentity", Type: "string", Description: v2.DisplaySpec{}.SwaggerDoc()["nodeIdentity"]},
		{Name: "Algorithm", Type: "string", Description: v2.DisplaySpec{}.SwaggerDoc()["algorithm"]},
		{Name: "StartNodeID", Type: "string", Description: v2.KnownNodesSpec{}.SwaggerDoc()["startNode"]},
		{Name: "Age", Type: "string", Description: metav1.ObjectMeta{}.SwaggerDoc()["creationTimestamp"]},
		{Name: "ComputeStatus", Type: "string", Description: v2.DisplayStatus{}.SwaggerDoc()["computeStatus"]},
		{Name: "Update", Type: "string", Description: v2.DisplayStatus{}.SwaggerDoc()["lastUpdate"]},
	}

	// Errors are suppressed as TableHandler already logs internally
	_ = h.TableHandler(dpColumnDefinitions, printDp)
	_ = h.TableHandler(dpColumnDefinitions, printDpList)
}

func printDp(obj *dijkstra.Display, options printers.GenerateOptions) ([]metav1.TableRow, error) {
	row := metav1.TableRow{
		Object: runtime.RawExtension{Object: obj},
	}
	row.Cells = append(row.Cells, obj.Name, obj.Spec.NodeIdentity, obj.Spec.Algorithm, obj.Spec.StartNode.ID, printerstorage.TranslateTimestamp(obj.CreationTimestamp), obj.Status.ComputeStatus, printerstorage.TranslateTimestamp(obj.Status.LastUpdate))
	return []metav1.TableRow{row}, nil
}

func printDpList(list *dijkstra.DisplayList, options printers.GenerateOptions) ([]metav1.TableRow, error) {
	rows := make([]metav1.TableRow, 0, len(list.Items))
	for i := range list.Items {
		r, err := printDp(&list.Items[i], options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}
