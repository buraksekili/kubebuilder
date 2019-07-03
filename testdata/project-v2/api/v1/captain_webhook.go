/*
Copyright 2019 The Kubernetes authors.

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

package v1

import (
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var captainlog = logf.Log.WithName("captain-resource")

func (r *Captain) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-crew-testproject-org-v1-captain,mutating=true,failurePolicy=fail,groups=crew.testproject.org,resources=captains,verbs=create;update,versions=v1,name=mcaptain.kb.io

var _ webhook.Defaulter = &Captain{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Captain) Default() {
	captainlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}
