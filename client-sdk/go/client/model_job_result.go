/*
Copyright 2021 VMware, Inc.

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

type JobResult struct {
	// ID of the job.
	JobId string `json:"jobId,omitempty"`
	// Current phase of the job.
	Phase string `json:"phase,omitempty"`
	// Input parameters of the job.
	JobParameters *OneOfJobResultJobParameters `json:"jobParameters,omitempty"`
	// Status of individual tasks and the overall job status.
	JobStatus *OneOfJobResultJobStatus `json:"jobStatus,omitempty"`
}
