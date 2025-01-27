package main

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func main() {

	fmt.Println("Hello, Kubernetes!")
	deployment := appsv1.Deployment{
		Spec: appsv1.DeploymentSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{Name: "nginx", Image: "nginx:1.14.2"},
					},
				},
			},
		},
	}
	fmt.Printf("Deployment: %v\n", deployment)

}
