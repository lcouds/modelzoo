package consts

const (
	ResourceNvidiaGPU = "nvidia.com/gpu"

	LabelInferenceName      = "inference"
	LabelInferenceNamespace = "inference-namespace"

	AnnotationBuilding = "ai.tensorchord.building"

	TolerationGPU              = "ai.tensorchord.gpu"
	TolerationNvidiaGPUPresent = "nvidia.com/gpu"

	//OrchestrationIdentifier identifier string for provider orchestration
	OrchestrationIdentifier = "kubernetes"
	//ProviderName name of the provider
	ProviderName = "modelzooetes"

	DefaultServicePrefix = "mdz-"

	DefaultHTTPProbePath = "/"

	// MaxReplicas is the maximum number of replicas that can be set for a inference.
	MaxReplicas = 5
)
