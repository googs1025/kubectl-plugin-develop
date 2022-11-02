package main

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"log"
)

func main() {

	configFlags := genericclioptions.NewConfigFlags(true)
	config, err := configFlags.ToRawKubeConfigLoader().ClientConfig() // 可以注意一下里面的ClientConfig()与RawConfig()方法
	if err != nil {
		log.Fatalln(err)
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	podList, err := client.CoreV1().Pods("default").List(ctx, v1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	for _, pod := range podList.Items {
		fmt.Printf("pod Name:%s, pod Namespace:%s \n", pod.Name, pod.Namespace)
	}



}