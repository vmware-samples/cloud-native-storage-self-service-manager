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

type DatastoreOperationsApiService service

/*
DatastoreOperationsApiService Get volumes(CNS &amp; non-CNS) and virtual machines on a datastore.
This API returns all the volumes(container volumes and non-CNS) as well as the virtual machines on a particular datastore. It is particularly useful to get this information while decommissioning a datastore.  The fcd ids outputted from this API can then be used as an input parameter in MigrateVolumes API.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param datacenter Datacenter name where the datastore is located. This input is case-sensitive.
  - @param datastore Name of the datastore on which container volumes need to be queried. This input is case-sensitive.

@return DatastoreResourcesResult
*/
func (a *DatastoreOperationsApiService) GetDatastoreResources(ctx context.Context, datacenter string, datastore string) (DatastoreResourcesResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue DatastoreResourcesResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/datastoreresources"

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
			var v DatastoreResourcesResult
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
DatastoreOperationsApiService Migrate volumes from source datastore to target datastore.
Volumes may need to be moved between the different datastores due to various  reasons like retiring older datastores, replacing or disruptive upgrades to existing datastores, saving volumes from failing datastores and so on. This API supports storage vMotion for PVs between different datastores (of same or different types). It returns a job id, the status of which can be retrieved using &#x60;jobStatus&#x60; API.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param datacenter Datacenter name where source and target datastores are located.
 * @param targetDatastore Name of the target datastore.
 * @param optional nil or *DatastoreOperationsApiMigrateVolumesOpts - Optional Parameters:
     * @param "SourceDatastore" (optional.String) -  (Optional) Name of the source datastore. Specify only if all volumes from source datastore need to be migrated to destination datastore, and don&#x27;t specify fcd Ids in that case. If specific list of fcd Ids is provided, then source datastore field will be ignored.
     * @param "FcdIdsToMigrate" (optional.Interface of []string) -  (Optional) Array of FCD ids to migrate. If not specified, all volumes from source datastore will be migrated to destination datastore.
     * @param "SkipPolicyCheck" (optional.Bool) -  (Optional) A flag to skip validation of volume policy with target datastore. Set to \&quot;true\&quot; to skip the policy check and force migrate a volume.
     * @param "SkipVolumeAccessibilityCheck" (optional.Bool) -  (Optional) If this flag is set to &#x27;true&#x27;, it will force migrate the volumes without checking if they will be accessible on target datastore from all cluster nodes(or topology-matching nodes in a topology-aware environment). This may affect the application availability if it gets scheduled on a cluster node that can&#x27;t access the target datastore. So it&#x27;s NOT recommended to set this flag to true.
@return MigrateVolumesResult
*/

type DatastoreOperationsApiMigrateVolumesOpts struct {
	SourceDatastore              optional.String
	FcdIdsToMigrate              optional.Interface
	SkipPolicyCheck              optional.Bool
	SkipVolumeAccessibilityCheck optional.Bool
}

func (a *DatastoreOperationsApiService) MigrateVolumes(ctx context.Context, datacenter string, targetDatastore string, localVarOptionals *DatastoreOperationsApiMigrateVolumesOpts) (MigrateVolumesResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue MigrateVolumesResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/migratevolumes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("datacenter", parameterToString(datacenter, ""))
	if localVarOptionals != nil && localVarOptionals.SourceDatastore.IsSet() {
		localVarQueryParams.Add("sourceDatastore", parameterToString(localVarOptionals.SourceDatastore.Value(), ""))
	}
	localVarQueryParams.Add("targetDatastore", parameterToString(targetDatastore, ""))
	if localVarOptionals != nil && localVarOptionals.FcdIdsToMigrate.IsSet() {
		localVarQueryParams.Add("fcdIdsToMigrate", parameterToString(localVarOptionals.FcdIdsToMigrate.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.SkipPolicyCheck.IsSet() {
		localVarQueryParams.Add("skipPolicyCheck", parameterToString(localVarOptionals.SkipPolicyCheck.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SkipVolumeAccessibilityCheck.IsSet() {
		localVarQueryParams.Add("skipVolumeAccessibilityCheck", parameterToString(localVarOptionals.SkipVolumeAccessibilityCheck.Value(), ""))
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
			var v MigrateVolumesResult
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
DatastoreOperationsApiService Resume Create Volume operation on datastore.
This API will unblock creation of new File and Block volumes on the specified datastore. To block volume provisioning, invoke SuspendVolumeProvisioning API.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param datacenter Datacenter name where the datastore is located. This input is case-sensitive.
  - @param datastore Name of the datastore where creation of new volumes has to be resumed. This input is case-sensitive.

@return ResumeVolumeProvisioningResult
*/
func (a *DatastoreOperationsApiService) ResumeVolumeProvisioning(ctx context.Context, datacenter string, datastore string) (ResumeVolumeProvisioningResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ResumeVolumeProvisioningResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/resumevolumeprovisioning"

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
			var v ResumeVolumeProvisioningResult
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
DatastoreOperationsApiService Suspend Create Volume operation on datastore.
This API will block creation of new File and Block volumes on the specified datastore. To unblock volume provisioning, invoke ResumeVolumeProvisioning API.  Other volume operations like attach, detach, delete etc. will not get affected for existing volumes.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param datacenter Datacenter name where the datastore is located. This input is case-sensitive.
  - @param datastore Name of the datastore where creation of new volumes has to be blocked. This input is case-sensitive.

@return SuspendVolumeProvisioningResult
*/
func (a *DatastoreOperationsApiService) SuspendVolumeProvisioning(ctx context.Context, datacenter string, datastore string) (SuspendVolumeProvisioningResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue SuspendVolumeProvisioningResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/suspendvolumeprovisioning"

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
			var v SuspendVolumeProvisioningResult
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
