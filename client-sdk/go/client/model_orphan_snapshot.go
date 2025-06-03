/*
Copyright 2025 VMware, Inc.

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

type OrphanSnapshot struct {
	// FCD Id of the orphan snapshot.
	VolumeId string `json:"volumeId,omitempty"`
	// Snapshot Id of the orphan snapshot.
	VolumeSnapshotId string `json:"volumeSnapshotId,omitempty"`
	// Datacenter where the orphan snapshot is located.
	Datacenter string `json:"datacenter,omitempty"`
	// Datastore where the orphan snapshot is located.
	Datastore string `json:"datastore,omitempty"`
	// Create time of the orphan snapshot.
	CreateTime string `json:"createTime,omitempty"`
	// Description of orphan snapshot
	SnapshotDescription string `json:"snapshotDescription,omitempty"`
}
