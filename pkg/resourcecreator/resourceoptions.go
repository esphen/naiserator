package resourcecreator

// ResourceOptions defines customizations for resource objects.
type ResourceOptions struct {
	NumReplicas                 int32
	AccessPolicy                bool
	AccessPolicyNotAllowedCIDRs []string
	NativeSecrets               bool
	GoogleCluster               bool
	GoogleProjectId				string
}

// NewResourceOptions creates a struct with the default resource options.
func NewResourceOptions() ResourceOptions {
	return ResourceOptions{
		NumReplicas: 1,
	}
}
