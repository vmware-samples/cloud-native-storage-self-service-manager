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

type SnapshotDeletionJobParameters struct {
	// fcdId is the identifier of FCD.
	FcdId string `json:"fcdId,omitempty"`
	// snapshotId is the identifier of snapshot.
	SnapshotId string `json:"snapshotId,omitempty"`
	// datacenter to which snapshots belong.
	Datacenter string `json:"datacenter,omitempty"`
	// datastores to which snapshots belong.
	Datastores []string `json:"datastores,omitempty"`
	// snapshotPrefix is the prefix used in the snapshot description.
	SnapshotPrefix string `json:"snapshotPrefix,omitempty"`
}
