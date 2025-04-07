## Orphan snapshots detection and clean-up
From vSphere CSI driver's perspective, orphan snapshots are vSphere snapshots that were initiated through the vSphere CSI driver but do not have a corresponding VolumeSnapshotContent object in Kubernetes clusters on the vCenter.

From Velero vSphere plugin's perspective, orphan snapshots are vSphere snapshots whose local deletion is failing after successful upload to the remote repository.  
Velero vSphere plugin performs following steps as part of backup:
1. Take a snapshot of the volume.
2. Upload the snapshot to a remote repository.
3. Delete the snapshot.

In these steps, deletion of local snapshot could fail, which leaves behind a snapshot. It can be considered as an orphan snapshot.

Orphan snapshot detection algorithm considers snapshot prefix for listing a snapshot as an orphan snapshot. Snapshot prefix is the prefix used in the snapshot description (snapshot description can be fetched using the GET /volumes/{volumeId}/snapshots API or will be available in CNS spec).  
For example, vSphere CSI driver creates snapshots with default prefix of “snapshot”. So, all these snapshot’s description has prefix “snapshot”. Similarly, Velero vSphere plugin uses snapshot prefix as “AstrolabeSnapshot”. 
CNS manager by default considers [“snapshot”, “AstrolableSnapshot”] as snapshot prefixes. These values are configurable using `snapshot-prefixes` parameter in the `cnsmanager-config` ConfigMap.

Since these orphan snapshots occupy space in the datastore and are not really used in Kubernetes, it's useful to identify and cleanup orphan snapshots periodically.
This functionality provides a set of APIs to detect and delete orphan snapshots on-demand.

### Supported versions
* vSphere version: 7.0.u3+
* Kubernetes version: 1.26+
* vSphere CSI driver version: 2.6+
* Velero vSphere plugin version: v1.4.3+

### Which orphan snapshots are skipped from detection/deletion ?
Some snapshots will not be considered during orphan snapshot detection/deletion. These include:
* Snapshots created out of band (not using vSphere CSI driver or Velero vSphere plugin).
* Statically provisioned CNS snapshots whose description doesn't start with snapshotPrefix(es) registered with CNS manager.
* Snapshots belonging to Kubernetes clusters which are not registered with CNS manager.

### A reminder to register all Kubernetes clusters!
Before you start using orphan snapshot functionality, it's imperative that you [register all the Kubernetes clusters](../../../README.md#register-kubernetes-clusters-before-you-start) in vCenter with CNS Manager, so that orphan snapshots are detected correctly. Any newly added Kubernetes cluster should also be immediately registered.

### APIs provided
1. *GET /orphansnapshots*

This API is used to get list of orphan snapshots. It takes optional parameters, datacenter & list of datastores, and returns orphan snapshots for them.
- If datacenter is not specified, then it returns all orphan snapshots in the vCenter (all datastores on all datacenters ).
- If only datacenter is specified, then it returns orphan snapshots in all datastores in the datacenter.
- If both datacenter & list of datastores is specified, it returns orphan snapshots in specified datastores on the datacenter.

This API also takes optional parameter snapshotPrefix, which indicates the prefix used in the snapshot description. So, this API returns orphan snapshots whose description’s prefix matches with the specified snapshot prefix. Default value of this parameter is “snapshot”.

There could be hundreds of orphan snapshots, so this API supports pagination as well. It takes optional parameters limit (default value 50) and offset (default value 0). Limit specifies the maximum entries that should be displayed in single request, whereas offset specifies the starting point of the result set.

Detection of orphan snapshots can be a time-consuming operation if there are large number of orphans. Hence it is performed asynchronously at regular intervals and the response is cached. This API returns list of orphan snapshots computed in the last run, along with the next operation interval (`RetryAfterMinutes`).

**Note:** For newly deployed CNS manager application, when orphan snapshots are being computed in the background for the first time, the API may return no orphan snapshots. It should then be re-tried after `RetryAfterMinutes` to get orphan snapshots computed in the latest run.

2. *DELETE /orphansnapshots*

This API is used to delete orphan snapshots. It takes optional parameters, datacenter & list of datastores, and deletes orphan snapshots from them with the same logic explained above.  
This API also takes optional parameter snapshotPrefix, which indicates the prefix used in the snapshot description.

Please note if there are large number of orphan snapshots in the system or if there's a slowness in vCenter networking/storage, the orphan snapshot deletion may take longer. That is why, orphan snapshot deletion operation is performed asynchronously.  
It creates a job to delete these orphan snapshots in the background and returns the job Id to user, the status of this job can be retrieved using `getjobstatus` API.


Apart from APIs provided for detection/deletion of orphan snapshots, there are separate APIs to list/delete snapshots of a specific volume or to delete an individual snapshot.

3. *GET /volumes/{volumeId}/snapshots*

This API lists all the snapshots for a specific volume. It takes required parameters datacenter and datastore name to which volume and snapshots belong to.  
This API also takes optional parameter snapshotPrefix which indicates the prefix used in snapshot description. If snapshot prefix is not specified, then it lists all snapshots of a volume. Otherwise it lists only those snapshots of volume whose description’s prefix matches with the specified snapshot prefix.

4. *DELETE /volumes/{volumeId}/snapshots*

This API deletes all snapshots of a specific volume. Similar to GET API, it takes required parameters datacenter and datastore and optional parameter snapshotPrefix.  
Deletion of all snapshots of a volume is performed asynchronously. It creates a job to delete these snapshots in the background and returns the job Id to user, the status of this job can be retrieved using `getjobstatus` API.

5. *DELETE /volumes/{volumeId}/snapshots/{snapshotId}*

This API deletes a specific snapshot belonging to a volume. It takes required parameters datacenter and datastore name to which volume and snapshot belongs to.  
Deletion of a snapshot is performed asynchronously. It creates a job to delete this snapshot in the background and returns the job Id to user, the status of this job can be retrieved using `getjobstatus` API.

### Checking status of the snapshot deletion job
Two APIs are being offered to check the status of the snapshot deletion job (A job corresponds to a single invocation of the API, that can have multiple tasks corresponding to separate snapshot deletion).

1. *GET /getjobstatus*

This API returns the current status of the job. A job can be in one of the following status:

* Queued - Job has been created but hasn't started processing.
* Running - Job is currently executing.
* Success - Job has completed successfully with all tasks succeeding.
* Error - Job ran but some or all of its tasks failed.

2. *GET /waitforjob*

This is a blocking API that waits for job to be successful or fail.  
Unlike `getjobstatus` API, this will wait for the job to finish before returning the job result response.

## Known Issues

### Deletion Failure for Orphan Snapshots of Orphan Volumes
When attempting to delete an orphan snapshot that belongs to an orphan volume using the DELETE /orphansnapshots API, the operation fails with the error:

> "Volume with ID xxxx is not registered as a CNS Volume."

#### Workaround:
To delete orphan snapshots, use the following APIs instead:

- DELETE /volumes/{volumeId}/snapshots (to delete all snapshots for a volume)

- DELETE /volumes/{volumeId}/snapshots/{snapshotId} (to delete a specific snapshot)

#### 🚀 Planned Improvement: The detection algorithm will be enhanced to handle these scenarios without having to manually invoke the above APIs in future updates.
