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
package main

import (
	"context"

	apiclient "cns.vmware.com/cns-manager/client"
	"github.com/antihax/optional"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func main() {
	ctx := context.TODO()
	devLog, _ := zap.NewDevelopment()
	logger := zapr.NewLogger(devLog)
	logger = logger.WithName("Main")
	ctx = log.IntoContext(ctx, logger)

	//Set BasicAuth credentials in ctx.
	ctx = context.WithValue(ctx, apiclient.ContextBasicAuth, apiclient.BasicAuth{
		UserName: "Administrator",
		Password: "Admin123@",
	})
	//Set API server basepath and create a client.
	cfg := &apiclient.Configuration{
		BasePath: "http://10.185.227.74:30008/1.0.0",
	}
	client := apiclient.NewAPIClient(cfg)

	//================API Invocation Examples===========

	//======Datastore Resources============
	// datacenter := "VSAN-DC"
	// datastore := "vsanDatastore"
	// res, resp, err := client.DatastoreOperationsApi.GetDatastoreResources(ctx, datacenter, datastore)
	// if err != nil {
	// 	logger.Error(err, "failed to get datastore resources")
	// }

	// logger.Info("Result", "result", res) //This gives the result in json format
	// logger.Info("HTTP status", "status", resp.Status)
	//=======================================================

	//====== Migrate Volumes========
	// datacenter := "VSAN-DC"
	// srcDatastore := "vsanDatastore"
	// targetDatastore := "agohil-nfs"
	// fcdIdsToMigrate := &apiclient.DatastoreOperationsApiMigrateVolumesOpts{
	// 	FcdIdsToMigrate: optional.NewInterface([]string{"ae903e93-b589-4cda-a3af-e59ef4660325"}),
	// }

	// res, resp, err := client.DatastoreOperationsApi.MigrateVolumes(ctx, datacenter, srcDatastore, targetDatastore, fcdIdsToMigrate)
	// if err != nil {
	// 	logger.Error(err, "failed to migrate volumes")
	// }

	// logger.Info("Result", "result", res) //This gives the result in json format
	// logger.Info("HTTP status", "status", resp.Status)
	//=======================================================

	//====== Job status ========
	// jobId := "volumemigrationjob-64248e96-7cf4-11ec-94b8-165555b18a9c"

	// res, resp, err := client.JobDetailsApi.GetJobStatus(ctx, jobId)
	// if err != nil {
	// 	logger.Error(err, "failed to get job status")
	// }

	// logger.Info("Result", "result", res) //This gives the result in json format
	// logger.Info("HTTP status", "status", resp.Status)
	//=======================================================

	//======Suspend Volume Provisioning on Datastore======
	// datacenter := "VSAN-DC"
	// datastore := "local-0"
	//
	// res, resp, err := client.DatastoreOperationsApi.SuspendVolumeProvisioning(ctx, datacenter, datastore)
	// if err != nil {
	//	logger.Error(err, "failed to suspend volume provisioning on datastore")
	// }
	// logger.Info("Result", "result", res) //This gives the result in json format
	// logger.Info("HTTP status", "status", resp.Status)
	//=======================================================

	//======Resume Volume Provisioning on Datastore======
	// datacenter := "VSAN-DC"
	// datastore := "local-0"
	//
	// res, resp, err := client.DatastoreOperationsApi.ResumeVolumeProvisioning(ctx, datacenter, datastore)
	// if err != nil {
	// 	logger.Error(err, "failed to resume provisioing volumes on datastore")
	// }
	// logger.Info("Result", "result", res) //This gives the result in json format
	// logger.Info("HTTP status", "status", resp.Status)
	//=======================================================

	//======List orphan volumes======
	includeDetails := true
	opts := &apiclient.OrphanVolumeApiOrphanVolumeListOpts{
		Datacenter: optional.NewString("VSAN-DC"),
		Datastores: optional.NewString("vsanDatastore"),
	}

	res, resp, err := client.OrphanVolumeApi.OrphanVolumeList(ctx, includeDetails, opts)
	if err != nil {
		logger.Error(err, "failed to list orphan volumes")
	}
	logger.Info("Result", "result", res) //This gives the result in json format
	logger.Info("HTTP status", "status", resp.Status)
	//=======================================================

	//======Delete orphan volumes======
	// deleteAttachedOrphans := false
	// opts := &apiclient.OrphanVolumeApiOrphanVolumeDeleteOpts{
	// 	Datacenter: optional.NewString("VSAN-DC"),
	// 	Datastores: optional.NewString("vsanDatastore"),
	// }

	// res, resp, err := client.OrphanVolumeApi.OrphanVolumeDelete(ctx, deleteAttachedOrphans, opts)
	// if err != nil {
	// 	logger.Error(err, "failed to delete orphan volumes")
	// }
	// logger.Info("Result", "result", res) //This gives the result in json format
	// logger.Info("HTTP status", "status", resp.Status)
	//=======================================================
}
