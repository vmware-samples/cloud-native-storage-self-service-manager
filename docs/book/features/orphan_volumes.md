## Orphan volumes detection and clean-up
Orphan volumes are vSphere volumes that are present on a vSphere datastore but there is no corresponding PersistentVolume in Kubernetes clusters on the vCenter.

Orphan volumes are often created when CNS solution creates more than one vSphere volume for a Persistent Volume in the Kubernetes cluster. This could occur when the vCenter components are slow, storage is slow, vCenter service restarts, connectivity issues between vCenter and ESXi hosts etc. Since these orphan volumes occupy space in the datastore and are not really used in Kubernetes, it's useful to identify and cleanup orphan volumes periodically.

This functionality provides a set of APIs to detect and delete orphan volumes on-demand, and also provide an option to turn on automatic deletion periodically.

### What qualifies as an orphan volume ?
A volume qualifies as an orphan volume if it meets all the conditions listed below:

1) A PersistentVolume in any of the registered kubernetes clusters is not using the volume.
2) The volume was dynamically provisioned using vSphere CSI driver.
3) The volume is no longer classified as a container volume in vCenter.
4) Volume exists for more than 50 minutes(or `orphan-volume-detection-interval-mins` value configured in `cnsmanager-config` configmap) at the time of orphan volumes detection.  
[**Note:** The worst case time for a volume to be detected as an orphan after its creation is twice `orphan-volume-detection-interval-mins`  (i.e. 100 minutes by default).]

### Which orphan volumes are skipped from detection/deletion ?
Some volumes, even if a kubernetes PersistentVolume doesn't map to them, will not be considered during orphan volume detection/deletion. These include:
* Volumes created out of band (not using vSphere CSI driver).
* Statically provisioned CNS volume whose name doesn't start with `pvc-`.
* File volumes.
* Orphan volumes that have snapshots will be detected as orphans, but they cannot be deleted using Orphan volume delete API.

### A reminder to register all Kubernetes clusters!
Before you start using orphan volume functionality, it's imperative that you [register all the Kubernetes clusters](../../../README.md#register-kubernetes-clusters-before-you-start) in vCenter with CNS Manager, so that orphan volumes are detected correctly. Any newly added kubernetes cluster should also be immediately registered.

### APIs provided
1. *GET /orphanvolumes* 

This API takes optional parameters, datacenter & list of datastores, and returns orphan volumes for them. 
- If datacenter is not specified, then it returns all orphan volumes in the vCenter (all datastores on all datacenters).
- If only datacenter is specified, then it returns orphan volumes in all datastores in the datacenter.
- If both datacenter & list of datastores is specified, it returns orphan volumes in specified datastores on the datacenter.

This API also takes optional parameter includeDetails, it should be set to `true` to get a detailed dump of the orphan volumes.

There could be hundreds of orphan volumes, so this API supports pagination as well. It takes optional parameters limit and offset. Limit specifies the maximum entries that should be displayed in single request, whereas offset specifies the starting point of the result set.

Detection of orphan volumes can be a time-consuming operation if there are large number of orphans. Hence, it is performed asynchronously at regular intervals and the response is cached. This API returns list of orphan volumes computed in the last run, along with the next operation interval(`RetryAfterMinutes`).  
**Note:** For newly deployed CNS manager application when orphan volumes are being computed in the background for the first time, the API may return no orphan volumes. It should then be re-tried after `RetryAfterMinutes` to get orphan volumes computed in the latest run.


2. *DELETE /orphanvolumes*

This API is used to delete orphan volumes. It also takes optional parameters, datacenter & list of datastores, and deletes orphan volumes from them.  
You can also specify whether you want to delete orphan volumes attached to a virtual machine or not. If set to `true`, the API will detach the orphan volume from the VM before deleting it.  

Please note if there are large number of orphan volumes in the system or if there's a slowness in vCenter networking/storage, the orphan deletion may take longer. If it takes longer than 30 minutes, the API client will time out. But be assured that orphan volumes are being deleted in the background which can also be verified by listing the orphans again using `GET /orphanvolumes` API.


### Setting up auto-monitoring for orphan volumes deletion  
There's also an option to automatically monitor and delete orphan volumes periodically. It's controlled using `auto-delete-ov` configuration in `cnsmanager-config` configmap. It can take one of the 3 values:  
    a. `disable`:  Orphan volumes will not be deleted automatically.  
    b. `enable-only-detached-ov`: Delete only detached orphan volumes.  
    c. `enable`:  Delete all the detected orphan volumes(both attached & detached).  
    
By default, it is disabled.

## Known Issues

### False Positives Due to Non-CSI Created Volumes
CNS Manager relies on naming conventions to identify orphan volumes, which can lead to false positives in certain cases.
- If a volume is created by something other than the CSI driver but has a name prefixed with pvc-, it may be mistakenly classified as an orphan.
- For example, First Class Disks (FCDs) backing Image Disks are created by directly invoking the VPXD endpoint and do not contain metadata in CNS or VSLM.
- If any FCD backing an Image Disk has a name starting with pvc-, CNS Manager might incorrectly detect it as an orphan volume.
##### Recommendation:
Before deleting orphan volumes, VI administrators should verify that no such FCDs exist to prevent accidental data loss. 

### Undetectable Orphan Volumes from Unresponsive Clusters
If a Kubernetes cluster becomes unresponsive, the CNS metadata for its volumes remains intact. As a result, CNS Manager is unable to detect these volumes as orphans, even if they are no longer in use.

⚠ Limitation: Currently, there is no automated way to identify and clean up such volumes within CNS Manager.

### Deletion Failure for Orphan Volumes with Snapshots
If an orphan volume has snapshots attached, the delete operation will fail because CNS Manager does not currently account for volumes with existing snapshots.

#### Workaround:
Users must manually delete the snapshots before removing the orphan volume by invoking:

- DELETE /volumes/{volumeId}/snapshots (to delete all snapshots for a volume)

- DELETE /volumes/{volumeId}/snapshots/{snapshotId} (to delete a specific snapshot)

#### 🚀 Planned Improvement: The detection algorithm will be enhanced to exclude such volumes in future updates.

### Slow Deletion on vCenter with Many Orphan Volumes
On a vCenter with a large number of orphan volumes, the delete operation may take a long time to complete, potentially causing a timeout when invoking the endpoint.

- Despite the timeout, CNS Manager continues deleting the volumes in the background.
- Users may not receive immediate confirmation of successful deletion.
#### Workaround:
To track progress, periodically invoke the LIST orphan volumes endpoint to check the current state.

#### 🚀 Planned Improvement: There are plans to convert this process into an asynchronous job to improve efficiency and avoid timeouts.
