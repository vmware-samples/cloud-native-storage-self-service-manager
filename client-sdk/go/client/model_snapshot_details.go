/*
Copyright 2024 VMware, Inc.

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

type SnapshotDetails struct {
	// Id of the snapshot.
	SnapshotId string `json:"snapshotId,omitempty"`
	// Id of the volume.
	VolumeId string `json:"volumeId,omitempty"`
	// Time when snapshot is created.
	CreateTime string `json:"createTime,omitempty"`
	// Description of the snapshot.
	SnapshotDescription string `json:"snapshotDescription,omitempty"`
	// Phase of the sanpshot if it is created by Velero vSphere plugin.
	VelerovSpherePluginSnapshotPhase string `json:"velerovSpherePluginSnapshotPhase,omitempty"`
	// Associated VolumeSnapshotContent name if snapshot is created by CSI driver.
	AssociatedVolumeSnapshotContent string `json:"associatedVolumeSnapshotContent,omitempty"`
	// Owner who created this snapshot. e.g. vSphere CSI driver, Velero vSphere plugin etc.
	Owner string `json:"owner,omitempty"`
}
