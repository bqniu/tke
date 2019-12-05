/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	internalversion "tkestack.io/tke/api/client/clientset/internalversion/typed/business/internalversion"
)

type FakeBusiness struct {
	*testing.Fake
}

func (c *FakeBusiness) ConfigMaps() internalversion.ConfigMapInterface {
	return &FakeConfigMaps{c}
}

func (c *FakeBusiness) ImageNamespaces(namespace string) internalversion.ImageNamespaceInterface {
	return &FakeImageNamespaces{c, namespace}
}

func (c *FakeBusiness) Namespaces(namespace string) internalversion.NamespaceInterface {
	return &FakeNamespaces{c, namespace}
}

func (c *FakeBusiness) Platforms() internalversion.PlatformInterface {
	return &FakePlatforms{c}
}

func (c *FakeBusiness) Portals() internalversion.PortalInterface {
	return &FakePortals{c}
}

func (c *FakeBusiness) Projects() internalversion.ProjectInterface {
	return &FakeProjects{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeBusiness) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}