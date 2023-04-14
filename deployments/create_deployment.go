// This script create a namespace named 'test' and a deployment of 3 'busybox' replicas inside it

package main

import (
	"context"
	"fmt"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	} else {
		kubeconfig = ""
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// create namespace
	nsClient := clientset.CoreV1().Namespaces()
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}
	if _, err := nsClient.Create(context.TODO(), ns, metav1.CreateOptions{}); err != nil {
		panic(err.Error())
	}

	// create deployment
	deployClient := clientset.AppsV1().Deployments("test")

	image := "busybox"
	replicas := int32(3)
	command := []string{"/bin/sh", "-c", "while true; do echo 'hello'; done"}

	deploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "busybox",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "busybox",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "busybox",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:    "busybox",
							Image:   image,
							Command: command,
						},
					},
				},
			},
		},
	}

	if _, err := deployClient.Create(context.TODO(), deploy, metav1.CreateOptions{}); err != nil {
		panic(err.Error())
	}

	fmt.Println("Deployment created successfully!")
}
