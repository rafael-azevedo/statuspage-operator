package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StatusPageCustomerSpec defines the desired state of StatusPageCustomer
// +k8s:openapi-gen=true
type StatusPageCustomerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// StatusPageCustomerStatus defines the observed state of StatusPageCustomer
// +k8s:openapi-gen=true
type StatusPageCustomerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StatusPageCustomer is the Schema for the statuspagecustomers API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type StatusPageCustomer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StatusPageCustomerSpec   `json:"spec,omitempty"`
	Status StatusPageCustomerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StatusPageCustomerList contains a list of StatusPageCustomer
type StatusPageCustomerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StatusPageCustomer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StatusPageCustomer{}, &StatusPageCustomerList{})
}
