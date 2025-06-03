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

// OrphanVolumeDeleteResult is the result of deleting orphan volumes.
type OrphanVolumeDeleteResult struct {
	// Number of orphan volumes detected.
	TotalOrphansDetected int64 `json:"totalOrphansDetected"`
	// Number of orphan volumes deleted.
	TotalOrphansDeleted int64 `json:"totalOrphansDeleted"`
	// Number of deleted orphan volumes that were detached.
	TotalDetachedOrphansDeleted int64 `json:"totalDetachedOrphansDeleted"`
	// Number of deleted orphan volumes that were attached to a VM.
	TotalAttachedOrphansDeleted int64 `json:"totalAttachedOrphansDeleted"`
	// Array of successfully deleted orphan volume IDs.
	SuccessfulOrphanDeletions []string `json:"successfulOrphanDeletions"`
	// Array of failed orphan volume deletions with the reason for failure for each orphan volume.
	FailedOrphanDeletions []OrphanVolumeDeleteFailure `json:"failedOrphanDeletions,omitempty"`
}
