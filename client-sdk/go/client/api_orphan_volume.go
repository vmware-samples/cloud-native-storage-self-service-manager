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

type OrphanVolumeApiService service

/*
OrphanVolumeApiService Delete orphan volumes.
Delete the orphan volumes for the given criteria.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param deleteAttachedOrphans Set to &#x60;true&#x60; to delete attached orphans. When set to &#x60;true&#x60;,  the API will detach the orphan volume from the VM before deleting it.
 * @param optional nil or *OrphanVolumeApiOrphanVolumeDeleteOpts - Optional Parameters:
     * @param "Datacenter" (optional.String) -  (Optional) Datacenter name to narrow down the deletion of orphan volumes to.
     * @param "Datastores" (optional.String) -  (Optional) List of comma-separated datastores to narrow down the deletion of orphan volumes to. Specify only if the &#x60;datacenter&#x60; param is specified.
@return OrphanVolumeDeleteResult
*/

type OrphanVolumeApiOrphanVolumeDeleteOpts struct {
	Datacenter optional.String
	Datastores optional.String
}

func (a *OrphanVolumeApiService) OrphanVolumeDelete(ctx context.Context, deleteAttachedOrphans bool, localVarOptionals *OrphanVolumeApiOrphanVolumeDeleteOpts) (OrphanVolumeDeleteResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Delete")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue OrphanVolumeDeleteResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/orphanvolumes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("deleteAttachedOrphans", parameterToString(deleteAttachedOrphans, ""))
	if localVarOptionals != nil && localVarOptionals.Datacenter.IsSet() {
		localVarQueryParams.Add("datacenter", parameterToString(localVarOptionals.Datacenter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Datastores.IsSet() {
		localVarQueryParams.Add("datastores", parameterToString(localVarOptionals.Datastores.Value(), ""))
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
			var v OrphanVolumeDeleteResult
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
OrphanVolumeApiService List all the orphan volumes.
Returns a list of orphan volumes for the given input parameters, which could be attached or detached. Since the detection of orphan volumes is an expensive operation, the operation is performed asynchronously at regular intervals. This API returns the list of orphan volumes found in the last run of the operation.  If the request is successful, the response will contain the following: 1. &#x60;TotalOrphans&#x60; - The total number of orphan volumes found. 2. &#x60;OrphanVolumes&#x60; - The list of orphan volumes found. 3. &#x60;RetryAfterMinutes&#x60; - The time in minutes after which the next retry should be attempted to get the updated orphan volume list. 4. &#x60;TotalOrphansAttached&#x60; - The total number of orphan volumes found that are attached to a VM. 5. &#x60;TotalOrphansDetached&#x60; - The total number of orphan volumes found that are detached. 6. &#x60;Limit&#x60; - The maximum number of orphan volumes to be returned. 7. &#x60;NextOffset&#x60; - The offset of the next page if there are more orphan volumes to query. Orphan volumes are safe to delete since there is no PersistentVolume in the Kubernetes cluster referring to them.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *OrphanVolumeApiOrphanVolumeListOpts - Optional Parameters:
     * @param "IncludeDetails" (optional.Bool) -  (Optional) Set to \&quot;true\&quot; to get a detailed dump of the orphan volume.
     * @param "Datacenter" (optional.String) -  (Optional) Datacenter name to narrow down the orphan volume search.
     * @param "Datastores" (optional.String) -  (Optional) List of comma-separated datastores. Specify only if the &#x60;datacenter&#x60; param is specified.
     * @param "Offset" (optional.Int32) -  (Optional) The offset indicates the starting point of the result set.
     * @param "Limit" (optional.Int32) -  (Optional) The limit indicates the maximum number of orphan volumes to be returned.
@return OrphanVolumeResult
*/

type OrphanVolumeApiOrphanVolumeListOpts struct {
	IncludeDetails optional.Bool
	Datacenter     optional.String
	Datastores     optional.String
	Offset         optional.Int32
	Limit          optional.Int32
}

func (a *OrphanVolumeApiService) OrphanVolumeList(ctx context.Context, localVarOptionals *OrphanVolumeApiOrphanVolumeListOpts) (OrphanVolumeResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue OrphanVolumeResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/orphanvolumes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.IncludeDetails.IsSet() {
		localVarQueryParams.Add("includeDetails", parameterToString(localVarOptionals.IncludeDetails.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Datacenter.IsSet() {
		localVarQueryParams.Add("datacenter", parameterToString(localVarOptionals.Datacenter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Datastores.IsSet() {
		localVarQueryParams.Add("datastores", parameterToString(localVarOptionals.Datastores.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
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
			var v OrphanVolumeResult
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
