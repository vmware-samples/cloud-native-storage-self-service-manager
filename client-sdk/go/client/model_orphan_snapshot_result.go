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

type OrphanSnapshotResult struct {
	// The total orphan snapshots returned.
	TotalOrphanSnapshots int64 `json:"totalOrphanSnapshots,omitempty"`
	// Limit specifies the maximum number of entries that are displayed in single request.
	Limit int64 `json:"limit,omitempty"`
	// Offset specifies the starting point of the result set.
	Offset int64 `json:"offset,omitempty"`
	// Array of orphan snapshots
	OrphanSnapshots []OrphanSnapshot `json:"orphanSnapshots,omitempty"`
	// Since the detection of orphan snapshots is an expensive operation, the operation is performed asynchronously at regular intervals. This API returns the list of orphan snapshots found in the last run of the operation. `retryAfterMinutes` indicates the time in minutes after which the next retry should be attempted to get the updated orphan snapshot list.
	RetryAfterMinutes int64 `json:"retryAfterMinutes,omitempty"`
}
