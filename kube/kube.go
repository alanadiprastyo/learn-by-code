package kube

type K8 struct {
	Deployment string
	Service    string
	Ingress    string
	Pod        int
}

type Cluster struct {
	Name     string
	Location string
}
