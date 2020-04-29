package kube

import (
	v1 "k8s.io/api/apps/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Deployment struct {
	clientset *kubernetes.Clientset
}

func NewDeployment(clientset *kubernetes.Clientset) *Deployment {
	return &Deployment{clientset: clientset}
}

func (c *Deployment) Create(namespace string, deployment *v1.Deployment) (err error) {
	deployment, err = c.clientset.AppsV1().Deployments(namespace).Create(deployment)
	return
}

func (c *Deployment) Delete(name, namespace string) (err error) {
	ops := &meta_v1.DeleteOptions{}
	err = c.clientset.AppsV1().Deployments(namespace).Delete(name, ops)
	return
}

func (c *Deployment) Update(name, namespace string) (err error) {
	ops := &meta_v1.DeleteOptions{}
	err = c.clientset.AppsV1().Deployments(namespace).Delete(name, ops)
	return
}

func (c *Deployment) Get(name, namespace string) (deployment *v1.Deployment, err error) {
	ops := meta_v1.GetOptions{}
	deployment, err = c.clientset.AppsV1().Deployments(namespace).Get(name, ops)

	return
}

func (c *Deployment) List(namespace string, ops meta_v1.ListOptions) (deploymentList *v1.DeploymentList, err error) {
	deploymentList, err = c.clientset.AppsV1().Deployments(namespace).List(ops)
	return
}
