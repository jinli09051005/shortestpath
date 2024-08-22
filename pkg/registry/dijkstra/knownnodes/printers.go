package knownnodes

import (
	"jinli.io/shortestpath/pkg/apis/dijkstra"
	knv2 "jinli.io/shortestpath/pkg/apis/dijkstra/v2"
	printerstorage "jinli.io/shortestpath/pkg/utils/printers/storage"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/printers"
)

// AddHandlers adds print handlers for default Kubernetes types dealing with internal versions.
func AddHandlers(h printers.PrintHandler) {
	knColumnDefinitions := []metav1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "NodeIdentity", Type: "string", Description: knv2.KnownNodesSpec{}.SwaggerDoc()["nodeIdentity"]},
		{Name: "CostUnit", Type: "string", Description: knv2.KnownNodesSpec{}.SwaggerDoc()["costUnit"]},
		{Name: "Nodes", Type: "string", Description: knv2.KnownNodesSpec{}.SwaggerDoc()["nodes"]},
		{Name: "Age", Type: "string", Description: metav1.ObjectMeta{}.SwaggerDoc()["creationTimestamp"]},
		{Name: "Update", Type: "string", Description: knv2.KnownNodesStatus{}.SwaggerDoc()["lastUpdate"]},
	}

	// Errors are suppressed as TableHandler already logs internally
	_ = h.TableHandler(knColumnDefinitions, printKn)
	_ = h.TableHandler(knColumnDefinitions, printKnList)
}

func printKn(obj *dijkstra.KnownNodes, options printers.GenerateOptions) ([]metav1.TableRow, error) {
	row := metav1.TableRow{
		Object: runtime.RawExtension{Object: obj},
	}
	row.Cells = append(row.Cells, obj.Name, obj.Spec.NodeIdentity, obj.Spec.CostUnit, len(obj.Spec.Nodes), printerstorage.TranslateTimestamp(obj.CreationTimestamp), printerstorage.TranslateTimestamp(obj.Status.LastUpdate))
	return []metav1.TableRow{row}, nil
}

func printKnList(list *dijkstra.KnownNodesList, options printers.GenerateOptions) ([]metav1.TableRow, error) {
	rows := make([]metav1.TableRow, 0, len(list.Items))
	for i := range list.Items {
		r, err := printKn(&list.Items[i], options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}
