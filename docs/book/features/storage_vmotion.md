## Storage vMotion for CNS volumes
This feature involves moving volumes from one datastore to another. It supports both offline and online migration of CNS volumes. Also to complement this functionality, there is a set of APIs provided that includes: 
1. Querying resources on a datastore.
2. Suspend CNS volume provisioning on a datastore.
3. Migrate volumes across datastores.
4. Resume CNS volume provisioning on a datastore.

### Prerequisites
* vSphere 7.0.u3
* CSI driver version 2.5+ for Suspend/Resume volume provisioning operations to work.  
Volume migration alone can work with CSI driver 2.4+.
* Target datastore to be mounted on all the hosts from which source datastore is accessible.

### Functionality Not Supported
* The following are not supported with current release of CNS manager, hence please ensure that storage vMotion be not performed for:  
  1. Migration of non-CNS volumes
  2. Migration of file volumes

* While migration of VMs is not possible with this tool yet, the worker node VMs can be migrated using vMotion from VCenter UI. This should be done after all the volumes attached to a worker VM have been migrated to destination datastore.

#### Caveats
* If vcp-to-csi migrated volumes are relocated to a different datastore, then it won't be possible to switch back to in-tree VCP plug-in as the volumes will lose mapping with VCP.
* While the volume migration is going on, the attach/detach operations on the volume will be waiting. But Kubernetes will continue to retry the operation which should succeed after the volume has been relocated.

### Concurrency Limits & Scale
While some concurrency is possible to achieve while migrating volumes, with limited scale testing, we recommend to migrate only one volume at a time using this tool.

Nevertheless on CNS manager application level, there can be 8 volume migrations that can be invoked in parallel across all clusters.  
And on vCenter level, the limits for simultaneous migrations can be be derived from this document - https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vcenterhost.doc/GUID-25EA5833-03B5-4EDD-A167-87578B8009B3.html

This translates to concurrent migration limits to be 1 per VM(for attached volumes), 2 per host & 8 per datastore. If there are parallel FCD migrations invoked beyond these limits, they will be queued based on the limits for each type of resource.

### Using the APIs
1. **Identifying resources on a datastore**  
  We're providing a helper API for the user to know the volume details on a datastore level. This API can be used to know what are the resources on a datastore that need to be moved if a datastore needs to be decommissioned.

   API Details:  
   */datastoreresources*  
   This gives details of the resources on the datastore. The resources can include
 
     a. List of CNS volumes  
     b. List of non-container volumes(created out of band)  
     c. List of VMs on the datastore.

2. **Suspend volume provisioning on a datastore**  
Before migrating volumes from a datastore, we don't want to create any new volumes on that datastore. So it will be better to put the datastore in a mode that suspends volume provisioning from CSI.  
This can be achieved by invoking `SuspendVolumeProvisioning` API. This needs all k8s clusters registered with CNS manager to be upgraded to CSI 2.5, otherwise the operation will fail.

    API Details:  
   */suspendvolumeprovisioning*

   For **file volumes**, the provisioning suspension takes a few minutes to come into effect. This time interval depends on the value for the parameter `csi-auth-check-intervalinmin` set in vsphere csi configuration file while setting up vsphere-csi-driver. The default time interval is 5 minutes.  
   For **block volumes**, provisioning is suspended immediately.

3. **Migrate Volumes**  
Once the list of volumes to be migrated is identified either from the VC UI or the helper API, the fcdIds can be inputted to the `MigrateVolumes` API along with target datastore.

    API Details:  
   */migratevolumes*

   This API is an asynchronous API and will return a job Id that's migrating volumes in the background. The status of this job can be queried anytime using `getjobstatus` API.

4. **Checking status of the volume migration job**  
Two APIs are being offered to check the status of the volume migration job (A job corresponds to a single invocation of the API, that can have multiple tasks corresponding to each volume Id passed in the API request)

    * */getjobstatus*  
    This returns the current status of the job. A job can be in one of the following status:

        * Queued - Job has been created but hasn't started processing.  
        * Running - Job is currently executing.
        * Success - Job has completed successfully with all tasks succeeding.
        * Error - Job ran but some or all of its tasks failed.
    

    * */waitforjob*  
    This is a blocking API that waits for job to be successful or fail.  
    Unlike `getjobstatus` API, this will wait for the job to finish before returning the job result response.