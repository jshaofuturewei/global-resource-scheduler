/*
Copyright 2020 Authors of Arktos.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// These const variables are used in our custom controller.
const (
	GroupName string = "globalscheduler.com"
	Kind      string = "Allocation"
	Version   string = "v1"
	Plural    string = "allocations"
	Singluar  string = "allocation"
	ShortName string = "alloc"
	Name      string = Plural + "." + GroupName
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// Allocation describes a Cluster custom resource.
type Allocation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AllocationSpec `json:"spec"`
	Status            string         `json:"status"`
}

// AllocationSpec is the spec for a Allocation resource.
//This is where you would put your custom resource data
type AllocationSpec struct {
	ResourceGroup ResourceGroup `json:"resource_group"`
	Selector      Selector      `json:"selector"`
	Replicas      int           `json:"replicas"`
}

type ResourceGroup struct {
	Name      string      `json:"name"`
	Resources []Resources `json:"resources"`
}

type Resources struct {
	Name         string   `json:"name"`
	ResourceType string   `json:"resource_type"`
	Flavors      []Flavor `json:"flavors"`
	Storage      Volume   `json:"storage,omitempty"`
	NeedEip      bool     `json:"need_eip,omitempty"`
	//Count        int      `json:"count,omitempty"`
}

type Flavor struct {
	FlavorID string `json:"flavor_id,omitempty"`
	Spot     int64  `json:"spot,omitempty"`
}

type Spot struct {
	MaxPrice           string `json:"max_price,omitempty"`
	SpotDurationHours  int    `json:"spot_duration_hours,omitempty"`
	SpotDurationCount  int    `json:"spot_duration_count,omitempty"`
	InterruptionPolicy string `json:"interruption_policy,omitempty"`
}

type Volume struct {
	SATA int `json:"sata,omitempty"`
	SAS  int `json:"sas,omitempty"`
	SSD  int `json:"ssd,omitempty"`
}

type Selector struct {
	GeoLocation GeoLocation `json:"geo_location,omitempty"`
	Regions     []Region    `json:"regions,omitempty"`
	Operator    string      `json:"operator,omitempty"`
	Strategy    Strategy    `json:"strategy,omitempty"`
}

type GeoLocation struct {
	City     string `json:"city"`
	Province string `json:"province"`
	Area     string `json:"area"`
	Country  string `json:"country"`
}

type Region struct {
	Region           string   `json:"region"`
	AvailabilityZone []string `json:"availability_zone,omitempty"`
}

type Strategy struct {
	// Location Strategy（centralize，discrete）
	LocationStrategy string `json:"location_strategy,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AllocationList is a list of Allocation resources.
type AllocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Allocation `json:"items"`
}
