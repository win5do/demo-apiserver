
/*
Copyright YEAR The Kubernetes Authors.

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


package catadmission

import (
	"context"
	aggregatedadmission "demo.com/ext-apiserver/plugin/admission"
	aggregatedinformerfactory "demo.com/ext-apiserver/pkg/client/informers_generated/externalversions"
	aggregatedclientset "demo.com/ext-apiserver/pkg/client/clientset_generated/clientset"
	genericadmissioninitializer "k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/apiserver/pkg/admission"
)

var _ admission.Interface 											= &catPlugin{}
var _ admission.MutationInterface 									= &catPlugin{}
var _ admission.ValidationInterface 								= &catPlugin{}
var _ genericadmissioninitializer.WantsExternalKubeInformerFactory 	= &catPlugin{}
var _ genericadmissioninitializer.WantsExternalKubeClientSet 		= &catPlugin{}
var _ aggregatedadmission.WantsAggregatedResourceInformerFactory 	= &catPlugin{}
var _ aggregatedadmission.WantsAggregatedResourceClientSet 			= &catPlugin{}

func NewCatPlugin() *catPlugin {
	return &catPlugin{
		Handler: admission.NewHandler(admission.Create, admission.Update),
	}
}

type catPlugin struct {
	*admission.Handler
}

func (p *catPlugin) ValidateInitialization() error {
	return nil
}

func (p *catPlugin) Admit(ctx context.Context, a admission.Attributes, o admission.ObjectInterfaces) error {
	return nil
}

func (p *catPlugin) Validate(ctx context.Context, a admission.Attributes, o admission.ObjectInterfaces) error {
	return nil
}

func (p *catPlugin) SetAggregatedResourceInformerFactory(aggregatedinformerfactory.SharedInformerFactory) {}

func (p *catPlugin) SetAggregatedResourceClientSet(aggregatedclientset.Interface) {}

func (p *catPlugin) SetExternalKubeInformerFactory(informers.SharedInformerFactory) {}

func (p *catPlugin) SetExternalKubeClientSet(kubernetes.Interface) {}
