/*
Copyright 2023 VMware, Inc.

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
 * @param deleteAttachedOrphans Set to &#x60;true&#x60; to delete attached orphans. When set to &#x60;true&#x60;, the API will detach the orphan volume from the VM before deleting it.
 * @param optional nil or *OrphanVolumeApiOrphanVolumeDeleteOpts - Optional Parameters:
     * @param "Datacenter" (optional.String) -  (Optional) Datacenter name to narrow down the deletion of orphan.
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
Orphan volumes are created when the CNS solution creates more than one volume for a Persistent Volume in the Kubernetes cluster. This occurs when the vCenter components are slow, storage is slow, vCenter service restarts, connectivity issues between vCenter and ESXi hosts etc.  Orphan volumes are the volumes that are present in the vSphere datastore but there is no corresponding PersistentVolume in the Kubernetes cluster. This API detects the orphan volumes for the given input parameters and returns a list of orphan volumes.   The orphan volumes could be attached or detached. 1. &#x60;Attached orphan volumes&#x60; - These would have the details set and will have info on the VM it is attached to. 2. &#x60;Detached orphan volumes&#x60; - These are the orphan volumes that do not have the details set.  Orphan volumes are safe to detach since there is no &#x60;PersistentVolume&#x60; in the Kubernetes cluster referring it.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param includeDetails Set to \&quot;true\&quot; to get a detailed dump of the orphan volume.
 * @param optional nil or *OrphanVolumeApiOrphanVolumeListOpts - Optional Parameters:
     * @param "Datacenter" (optional.String) -  (Optional) Datacenter name to narrow down the orphan volume search.
     * @param "Datastores" (optional.String) -  (Optional) List of comma-separated datastores. Specify only if the &#x60;datacenter&#x60; param is specified.
@return OrphanVolumeResult
*/

type OrphanVolumeApiOrphanVolumeListOpts struct {
	Datacenter optional.String
	Datastores optional.String
}

func (a *OrphanVolumeApiService) OrphanVolumeList(ctx context.Context, includeDetails bool, localVarOptionals *OrphanVolumeApiOrphanVolumeListOpts) (OrphanVolumeResult, *http.Response, error) {
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

	localVarQueryParams.Add("includeDetails", parameterToString(includeDetails, ""))
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
