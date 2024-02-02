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

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type SnapshotForVolumeApiService service
/*
SnapshotForVolumeApiService Delete all the snapshots for a specific volume
Use this API to delete all the snapshots for a specific volume. Snapshot deletion operation is performed asynchronously. It returns a job id, the status of which can be retrieved using &#x60;jobStatus&#x60; API.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param volumeId Unique Id of the volume.
 * @param datacenter Datacenter name to which volume and snapshots belong to.
 * @param datastore Datastore name to which volume and snapshots belong to.
 * @param optional nil or *SnapshotForVolumeApiDeleteAllSnapshotsForVolumeOpts - Optional Parameters:
     * @param "SnapshotPrefix" (optional.String) -  (Optional) The snapshot prefix indicates the prefix used in snapshot description. If snapshot prefix is not specified, then it will delete all snapshots of a volume.
@return SnapshotDeleteResult
*/

type SnapshotForVolumeApiDeleteAllSnapshotsForVolumeOpts struct {
    SnapshotPrefix optional.String
}

func (a *SnapshotForVolumeApiService) DeleteAllSnapshotsForVolume(ctx context.Context, volumeId string, datacenter string, datastore string, localVarOptionals *SnapshotForVolumeApiDeleteAllSnapshotsForVolumeOpts) (SnapshotDeleteResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue SnapshotDeleteResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/volumes/{volumeId}/snapshots"
	localVarPath = strings.Replace(localVarPath, "{"+"volumeId"+"}", fmt.Sprintf("%v", volumeId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("datacenter", parameterToString(datacenter, ""))
	localVarQueryParams.Add("datastore", parameterToString(datastore, ""))
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
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 202 {
			var v SnapshotDeleteResult
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
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
SnapshotForVolumeApiService Delete one specific snapshot for volume
Use this API to delete the specific snapshot of a volume. Snapshot deletion operation is performed asynchronously. It returns a job id, the status of which can be retrieved using &#x60;jobStatus&#x60; API.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param volumeId Unique Id of the volume.
 * @param snapshotId Unique Id of the snapshot.
 * @param datacenter Datacenter name to which volume and snapshot belongs to.
 * @param datastore Datastore name to which volume and snapshot belongs to.
@return SnapshotDeleteResult
*/
func (a *SnapshotForVolumeApiService) DeleteSnapshotForVolume(ctx context.Context, volumeId string, snapshotId string, datacenter string, datastore string) (SnapshotDeleteResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue SnapshotDeleteResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/volumes/{volumeId}/snapshots/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"volumeId"+"}", fmt.Sprintf("%v", volumeId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", fmt.Sprintf("%v", snapshotId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("datacenter", parameterToString(datacenter, ""))
	localVarQueryParams.Add("datastore", parameterToString(datastore, ""))
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
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 202 {
			var v SnapshotDeleteResult
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
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
SnapshotForVolumeApiService List all the snapshots for a specific volume
Use this API to retrieve all the snapshots for a specific volume.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param volumeId Unique Id of the volume.
 * @param datacenter Datacenter name to which volume and snapshots belong to.
 * @param datastore Datastore name to which volume and snapshots belong to.
 * @param optional nil or *SnapshotForVolumeApiListAllSnapshotsForVolumeOpts - Optional Parameters:
     * @param "SnapshotPrefix" (optional.String) -  (Optional) The snapshot prefix indicates the prefix used in snapshot description. If snapshot prefix is not specified, then it will list all snapshots of a volume.
@return SnapshotsForVolumeResult
*/

type SnapshotForVolumeApiListAllSnapshotsForVolumeOpts struct {
    SnapshotPrefix optional.String
}

func (a *SnapshotForVolumeApiService) ListAllSnapshotsForVolume(ctx context.Context, volumeId string, datacenter string, datastore string, localVarOptionals *SnapshotForVolumeApiListAllSnapshotsForVolumeOpts) (SnapshotsForVolumeResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue SnapshotsForVolumeResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/volumes/{volumeId}/snapshots"
	localVarPath = strings.Replace(localVarPath, "{"+"volumeId"+"}", fmt.Sprintf("%v", volumeId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("datacenter", parameterToString(datacenter, ""))
	localVarQueryParams.Add("datastore", parameterToString(datastore, ""))
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
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v SnapshotsForVolumeResult
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
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
