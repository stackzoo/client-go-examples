// This script retrieves the metrics of pods in the 'test' namespace

package main

import (
	"context"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
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
	metricsclientset := metricsv.NewForConfigOrDie(config)
	if err != nil {
		panic(err.Error())
	}
	podMetrics, err := metricsclientset.MetricsV1beta1().PodMetricses("test").List(context.TODO(), metav1.ListOptions{})
	fmt.Printf("SHOWING PODS METRICS\n")
	for _, v := range podMetrics.Items {
		fmt.Printf("pod name: %s\n", v.GetName())
		fmt.Printf("namespace: %s\n", v.GetNamespace())
		fmt.Printf("cpu usage (millicore): %vm\n", v.Containers[0].Usage.Cpu().MilliValue())
		fmt.Printf("memory usage: %vMi\n\n", v.Containers[0].Usage.Memory().Value()/(1024*1024))
	}
}
