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
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type OrphanSnapshotApiService service

/*
OrphanSnapshotApiService Delete orphan snapshots.
Use this API to identify and delete orphan snapshots. From vSphere CSI plugin&#x27;s perspective, orphan snapshots are FCD snapshots that were initiated through the vSphere CSI driver but do not have a corresponding VolumeSnapshotContent object in the Kubernetes cluster. snapshotPrefix is the prefix used in the snapshot description. Its default value is “snapshot”, which is also the default value used by snapshot sidecar in CSI and it can be configured based on prefix used in the snapshot sidecar. Use the &#x60;snapshotPrefix&#x60; parameter to specify alternate prefix.  From Velero vSphere plugin&#x27;s perspective, orphan snapshots are snapshots whose upload is failing with multiple attempts or snapshots whose local deletion is failing after successful upload. For Velero vSphere plugin, user has to specify “AstrolabeSnapshot” as the snapshotPrefix.  Orphan snapshot deletion operation is performed asynchronously. It returns a job id, the status of which can be retrieved using &#x60;jobStatus&#x60; API.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *OrphanSnapshotApiOrphanSnapshotsDeleteOpts - Optional Parameters:
     * @param "Datacenter" (optional.String) -  (Optional) Datacenter name to narrow down the orphan snapshots search.
     * @param "Datastores" (optional.String) -  (Optional) List of comma-separated datastores. Specify only if the &#x60;datacenter&#x60; param is specified.
     * @param "SnapshotPrefix" (optional.String) -  (Optional) The snapshot prefix indicates the prefix used in snapshot description.
@return SnapshotDeleteResult
*/

type OrphanSnapshotApiOrphanSnapshotsDeleteOpts struct {
	Datacenter     optional.String
	Datastores     optional.String
	SnapshotPrefix optional.String
}

func (a *OrphanSnapshotApiService) OrphanSnapshotsDelete(ctx context.Context, localVarOptionals *OrphanSnapshotApiOrphanSnapshotsDeleteOpts) (SnapshotDeleteResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Delete")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue SnapshotDeleteResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/orphansnapshots"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Datacenter.IsSet() {
		localVarQueryParams.Add("datacenter", parameterToString(localVarOptionals.Datacenter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Datastores.IsSet() {
		localVarQueryParams.Add("datastores", parameterToString(localVarOptionals.Datastores.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SnapshotPrefix.IsSet() {
		localVarQueryParams.Add("snapshotPrefix", parameterToString(localVarOptionals.SnapshotPrefix.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 202 {
			var v SnapshotDeleteResult
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
OrphanSnapshotApiService List all the orphan snapshots.
Use this API to identify orphan snapshots. From vSphere CSI plugin&#x27;s perspective, orphan snapshots are FCD snapshots that were initiated through the vSphere CSI driver but do not have a corresponding VolumeSnapshotContent object in the Kubernetes cluster. snapshotPrefix is the prefix used in the snapshot description. Its default value is “snapshot”, which is also the default value used by snapshot sidecar in CSI and it can be configured based on prefix used in the snapshot sidecar. Use the &#x60;snapshotPrefix&#x60; parameter to specify alternate prefix.  From Velero vSphere plugin&#x27;s perspective, orphan snapshots are snapshots whose upload is failing with multiple attempts or snapshots whose local deletion is failing after successful upload. For Velero vSphere plugin, user has to specify “AstrolabeSnapshot” as the snapshotPrefix.  GET API for orphan snapshots support pagination. The response body contains totalOrphanSnapshots, limit and offset values. Also, response header contains X-Limit and X-Next-Offset values. Based on these values user can decide if there are more results to be fetched. Since the detection of orphan snapshots is an expensive operation, the operation is performed asynchronously at regular intervals. This API returns the list of orphan snapshots found in the last run of the operation. &#x60;retryAfterMinutes&#x60; in response body indicates the time in minutes after which the next retry should be attempted to get the updated orphan snapshot list.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *OrphanSnapshotApiOrphanSnapshotsListOpts - Optional Parameters:
     * @param "Datacenter" (optional.String) -  (Optional) Datacenter name to narrow down the orphan snapshots search.
     * @param "Datastores" (optional.String) -  (Optional) List of comma-separated datastores. Specify only if the &#x60;datacenter&#x60; param is specified.
     * @param "SnapshotPrefix" (optional.String) -  (Optional) The snapshot prefix indicates the prefix used in snapshot description.
     * @param "Limit" (optional.Int64) -  (Optional) Limit specifies the maximum entries that should be displayed in single request.
     * @param "Offset" (optional.Int64) -  (Optional) Offset specifies the starting point of the result set.
@return OrphanSnapshotResult
*/

type OrphanSnapshotApiOrphanSnapshotsListOpts struct {
	Datacenter     optional.String
	Datastores     optional.String
	SnapshotPrefix optional.String
	Limit          optional.Int64
	Offset         optional.Int64
}

func (a *OrphanSnapshotApiService) OrphanSnapshotsList(ctx context.Context, localVarOptionals *OrphanSnapshotApiOrphanSnapshotsListOpts) (OrphanSnapshotResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue OrphanSnapshotResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/orphansnapshots"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Datacenter.IsSet() {
		localVarQueryParams.Add("datacenter", parameterToString(localVarOptionals.Datacenter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Datastores.IsSet() {
		localVarQueryParams.Add("datastores", parameterToString(localVarOptionals.Datastores.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SnapshotPrefix.IsSet() {
		localVarQueryParams.Add("snapshotPrefix", parameterToString(localVarOptionals.SnapshotPrefix.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v OrphanSnapshotResult
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
