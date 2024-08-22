package v2

var map_Display = map[string]string {
		"": "Display",
		"metadata": "Standard object's metadata.",
}
func (Display) SwaggerDoc() map[string]string {
	return map_Display
}

var map_DisplayList = map[string]string {
		"": "DisplayList",
}
func (DisplayList) SwaggerDoc() map[string]string {
	return map_DisplayList
}

var map_DisplaySpec = map[string]string {
		"": "DisplaySpec defines the desired state of Display",
		"nodeIdentity": "A type of node identity",
		"startNode": "Start node information",
		"targetNodes": "Target nodes information",
		"algorithm": "Algorithms used to calculate the shortest path, including dijkstra and floyd algorithms",
}
func (DisplaySpec) SwaggerDoc() map[string]string {
	return map_DisplaySpec
}

var map_DisplayStatus = map[string]string {
		"": "DisplayStatus defines the observed state of Display",
		"lastUpdate": "Last Update Time",
		"computeStatus": "Dispaly  ShortestPath Compute Status",
}
func (DisplayStatus) SwaggerDoc() map[string]string {
	return map_DisplayStatus
}

var map_Edge = map[string]string {
		"id": "To node id",
		"cost": "To node cost",
}
func (Edge) SwaggerDoc() map[string]string {
	return map_Edge
}

var map_KnownNodes = map[string]string {
		"": "KnownNodes KnownNodes资源对象结构定义",
		"metadata": "Standard object's metadata.",
}
func (KnownNodes) SwaggerDoc() map[string]string {
	return map_KnownNodes
}

var map_KnownNodesList = map[string]string {
		"": "KnownNodesList",
}
func (KnownNodesList) SwaggerDoc() map[string]string {
	return map_KnownNodesList
}

var map_KnownNodesSpec = map[string]string {
		"": "KnownNodesSpec defines the desired state of KnownNodes",
		"nodeIdentity": "A type of node identity",
		"costUnit": "To node cost unit",
		"nodes": "Known nodes information",
}
func (KnownNodesSpec) SwaggerDoc() map[string]string {
	return map_KnownNodesSpec
}

var map_KnownNodesStatus = map[string]string {
		"": "KnownNodesStatus defines the observed state of KnownNodes",
		"lastUpdate": "Last Update Time",
}
func (KnownNodesStatus) SwaggerDoc() map[string]string {
	return map_KnownNodesStatus
}

var map_Node = map[string]string {
		"id": "Node id",
		"name": "Node name",
		"preNode": "Previous node",
		"edges": "Node edges",
}
func (Node) SwaggerDoc() map[string]string {
	return map_Node
}

var map_StartNode = map[string]string {
		"id": "Node id",
		"name": "Node name",
}
func (StartNode) SwaggerDoc() map[string]string {
	return map_StartNode
}

var map_TargetNode = map[string]string {
		"id": "Target node id",
		"name": "Target node name",
		"distance": "Start node to target node distance",
		"path": "Start node to target node path",
}
func (TargetNode) SwaggerDoc() map[string]string {
	return map_TargetNode
}

