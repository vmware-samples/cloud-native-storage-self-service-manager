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

type SnapshotDeletionJobStatus struct {
	// Total snapshots which will be deleted as part of this job.
	TotalSnapshotsPlannedForDeletion int32 `json:"totalSnapshotsPlannedForDeletion,omitempty"`
	// Total snapshots which got successfully deleted as part of this job.
	TotalSnapshotsSuccessfullyDeleted int32 `json:"totalSnapshotsSuccessfullyDeleted,omitempty"`
	// Total snapshots whose deletion failed as part of this job.
	TotalSnapshotsWithFailedDeletion int32 `json:"totalSnapshotsWithFailedDeletion,omitempty"`
	// Time at which the job started processing.
	StartTime time.Time `json:"startTime,omitempty"`
	// Time at which the job completed processing.
	EndTime time.Time `json:"endTime,omitempty"`
	// Array of status of individual snapshot deletion tasks in the job.
	SnapshotDeletionTasks []SnapshotDeletionTaskStatus `json:"snapshotDeletionTasks,omitempty"`
}
