/*
Copyright The KubeDB Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package framework

import (
	"fmt"
	"time"

	api "kubedb.dev/apimachinery/apis/kubedb/v1alpha1"

	shell "github.com/codeskyblue/go-sh"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

const (
	updateRetryInterval = 10 * 1000 * 1000 * time.Nanosecond
	maxAttempts         = 5
)

//func deleteInBackground() *metav1.DeleteOptions {
//	policy := metav1.DeletePropagationBackground
//	return &metav1.DeleteOptions{PropagationPolicy: &policy}
//}

func deleteInForeground() *metav1.DeleteOptions {
	policy := metav1.DeletePropagationForeground
	return &metav1.DeleteOptions{PropagationPolicy: &policy}
}

func (f *Framework) CleanWorkloadLeftOvers() {
	// delete statefulset
	if err := f.kubeClient.AppsV1().StatefulSets(f.namespace).DeleteCollection(deleteInForeground(), metav1.ListOptions{
		LabelSelector: labels.SelectorFromSet(map[string]string{
			api.LabelDatabaseKind: api.ResourceKindMongoDB,
		}).String(),
	}); err != nil && !kerr.IsNotFound(err) {
		fmt.Printf("error in deletion of Statefulset. Error: %v", err)
	}

	// delete pvc
	if err := f.kubeClient.CoreV1().PersistentVolumeClaims(f.namespace).DeleteCollection(deleteInForeground(), metav1.ListOptions{
		LabelSelector: labels.SelectorFromSet(map[string]string{
			api.LabelDatabaseKind: api.ResourceKindMongoDB,
		}).String(),
	}); err != nil && !kerr.IsNotFound(err) {
		fmt.Printf("error in deletion of PVC. Error: %v", err)
	}
}

func (f *Framework) PrintDebugHelpers() {
	sh := shell.NewSession()
	fmt.Println("======================================[ Describe Job ]===================================================")
	if err := sh.Command("kubectl", "describe", "job", "-n", fmt.Sprintf("%v", f.Namespace())).Run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("======================================[ Describe Pod ]===================================================")
	if err := sh.Command("kubectl", "describe", "po", "-n", fmt.Sprintf("%v", f.Namespace())).Run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("======================================[ Describe Mongo ]===================================================")
	if err := sh.Command("kubectl", "describe", "mg", "-n", fmt.Sprintf("%v", f.Namespace())).Run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("======================================[ Describe RestoreSession ]==========================================")
	if err := sh.Command("kubectl", "describe", "restoresession", "-n", fmt.Sprintf("%v", f.Namespace())).Run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("======================================[ Describe Nodes ]===================================================")
	if err := sh.Command("kubectl", "describe", "nodes").Run(); err != nil {
		fmt.Println(err)
	}
}
