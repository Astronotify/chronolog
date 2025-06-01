package entries

import "context"

// K8SLogEntry represents a structured log entry containing Kubernetes-specific context.
//
// This log type extends LogEntry with metadata related to the Kubernetes execution environment,
// allowing better correlation, debugging, and traceability across clusters, pods, and containers.
//
// Fields:
//
//   - ClusterName: the name of the Kubernetes cluster where the log was generated.
//   - Namespace: the Kubernetes namespace of the pod emitting the log.
//   - PodName: the name of the pod responsible for the log entry.
//   - Container: the name of the container within the pod emitting the log.
//   - NodeName: the name of the Kubernetes node hosting the pod.
type K8SLogEntry struct {
	LogEntry
	ClusterName string `json:"cluster_name"`
	Namespace   string `json:"namespace"`
	PodName     string `json:"pod_name"`
	Container   string `json:"container"`
	NodeName    string `json:"node_name"`
}

// NewK8SLogEntry creates a new K8SLogEntry enriched with Kubernetes context.
//
// This constructor extends the base log structure with Kubernetes-specific fields, while also
// applying the standard metadata enrichment (timestamp, trace, build info, etc.) from LogEntry.
//
// Parameters:
//
//   - ctx (context.Context): the execution context for trace/build metadata extraction.
//   - clusterName (string): the name of the Kubernetes cluster.
//   - namespace (string): the Kubernetes namespace of the emitting pod.
//   - podName (string): the name of the pod that produced the log.
//   - container (string): the name of the container inside the pod.
//   - nodeName (string): the name of the node hosting the pod.
//   - level (LogLevel): the severity level of the log.
//   - message (string): the log message content.
//   - additionalData (...map[string]any): optional structured metadata for enrichment.
//
// Returns:
//
//   - K8SLogEntry: a structured and context-rich log entry for Kubernetes environments.
func NewK8SLogEntry(
	ctx context.Context,
	clusterName, namespace, podName, container, nodeName string,
	level LogLevel, message string,
	additionalData ...map[string]any,
) K8SLogEntry {
	entry := K8SLogEntry{
		LogEntry:    NewLogEntry(ctx, level, message, additionalData...),
		ClusterName: clusterName,
		Namespace:   namespace,
		PodName:     podName,
		Container:   container,
		NodeName:    nodeName,
	}
	entry.EventType = "K8SLogEntry"
	return entry
}
