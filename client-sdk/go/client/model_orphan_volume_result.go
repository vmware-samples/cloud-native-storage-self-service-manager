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

type OrphanVolumeResult struct {
	// The total number of orphan volumes detected in the vCenter for the input.
	TotalOrphans int32 `json:"totalOrphans"`
	// The total number of orphan volumes that are attached to a VM. This field is returned only if includeDetails is set to true.
	TotalOrphansAttached int32 `json:"totalOrphansAttached,omitempty"`
	// The total number of orphan volumes that are not attached to a VM. This field is returned only if includeDetails is set to true.
	TotalOrphansDetached int32 `json:"totalOrphansDetached,omitempty"`
	// Array of orphan volumes.
	OrphanVolumes []OrphanVolume `json:"orphanVolumes"`
	// The time in minutes after which the next retry should be attempted to get the updated orphan volume list.
	RetryAfterMinutes int32 `json:"retryAfterMinutes"`
	// The maximum number of orphan volumes returned.
	Limit int32 `json:"limit,omitempty"`
	// The offset of the next page if there are more orphan volumes to query.
	NextOffset int32 `json:"nextOffset,omitempty"`
}
