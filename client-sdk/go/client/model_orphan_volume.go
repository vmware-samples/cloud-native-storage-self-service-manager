/*
Copyright 2023 VMware, Inc.

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

package swagger

type OrphanVolume struct {
	// ID of the orphan volume.
	VolumeId string `json:"volumeId,omitempty"`
	// Name of the orphan volume.
	VolumeName string `json:"volumeName,omitempty"`
	// Datastore where the orphan volume is located.
	Datastore string `json:"datastore,omitempty"`
	// Create time of the orphan volume.
	CreateTime string `json:"createTime,omitempty"`
	// Capacity of the orphan volume.
	CapacityInMb int64                `json:"capacityInMb,omitempty"`
	Details      *OrphanVolumeDetails `json:"details,omitempty"`
}
