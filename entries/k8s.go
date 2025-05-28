package entries

type K8SLogEntry struct {
	LogEntry
	ClusterName string `json:"cluster_name"`
	Namespace   string `json:"namespace"`
	PodName     string `json:"pod_name"`
	Container   string `json:"container"`
	NodeName    string `json:"node_name"`
}
