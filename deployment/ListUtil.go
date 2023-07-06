package deployment

import (
	"context"
	"github.com/captain/lib"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Deployment struct {
	Name     string
	Replicas [3]int32
	Images   string
}

// 显示 所有

func ListAll(namespace string) (ret []*Deployment) {
	ctx := context.Background()
	listopt := metav1.ListOptions{}
	depList, err := lib.K8sClient.AppsV1().Deployments(namespace).List(ctx, listopt)
	lib.CheckError(err)
	for _, dep := range depList.Items {
		image := ""
		for _, container := range dep.Spec.Template.Spec.Containers {
			image += container.Image + " "
		}
		ret = append(ret, &Deployment{
			Name: dep.Name,
			Replicas: [3]int32{
				dep.Status.Replicas,
				dep.Status.AvailableReplicas,
				dep.Status.UnavailableReplicas,
			},
			Images: image,
		})
	}
	return
}
