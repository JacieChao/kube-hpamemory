/*
Copyright 2018 The Openshift Evangelists

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
package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1 "kube-hpamemory/pkg/client/clientset/versioned/typed/hor/v1"
)

type FakeHorV1 struct {
	*testing.Fake
}

func (c *FakeHorV1) HORHPAMemories(namespace string) v1.HORHPAMemoryInterface {
	return &FakeHORHPAMemories{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeHorV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
