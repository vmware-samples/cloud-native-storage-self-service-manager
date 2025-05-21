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

import (
	"time"
)

type SnapshotDeletionTaskStatus struct {
	// Id of the FCD to which snapshot belongs.
	FcdId string `json:"fcdId,omitempty"`
	// Id of the snapshot to be deleted.
	SnapshotId string `json:"snapshotId,omitempty"`
	// datacenter to which snapshot belongs.
	Datacenter string `json:"datacenter,omitempty"`
	// datastore to which snapshot belongs.
	Datastore string `json:"datastore,omitempty"`
	// description of the snapshot.
	SnapshotDescription string `json:"snapshotDescription,omitempty"`
	// Current phase of the snapshot deletion task.
	Phase string `json:"phase,omitempty"`
	// The timestamp at which the task was invoked.
	TaskStartTime time.Time `json:"taskStartTime,omitempty"`
	// The timestamp at which the task finished.
	TaskEndTime time.Time `json:"taskEndTime,omitempty"`
	Error_      *Fault    `json:"error,omitempty"`
}
