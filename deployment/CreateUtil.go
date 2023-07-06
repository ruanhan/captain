package deployment

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/captain/lib"
	"io/ioutil"
	"k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func create() {
	ctx := context.Background()
	ngxDep := &v1.Deployment{} //我们要创建的deployment
	b, _ := ioutil.ReadFile("yamls/deployment.yaml")
	ngxJson, _ := yaml.ToJSON(b)
	lib.CheckError(json.Unmarshal(ngxJson, ngxDep))

	createopt := metav1.CreateOptions{}

	_, err := lib.K8sClient.AppsV1().Deployments("default").
		Create(ctx, ngxDep, createopt)

	if err != nil {
		fmt.Println(err)
	}
}
