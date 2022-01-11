package main

import (
	"fmt"

	"github.com/alanadiprastyo/learn-by-code/kube"
)

func main() {
	mykube := kube.K8{
		Deployment: "WebService",
		Service:    "web-service",
		Ingress:    "webservice.apps.awancloud.io",
		Pod:        2,
	}

	myCluster := kube.Cluster{
		Name:     "Jakarta-zone-1",
		Location: "Indonesia",
	}

	fmt.Println(fmt.Sprintf("Halo Team we have Deployment %s , Service : %s, Ingress : %s, Pod : %d, running on the cluster : %s, Location : %s", mykube.Deployment, mykube.Service, mykube.Ingress, mykube.Pod, myCluster.Name, myCluster.Location))

}
