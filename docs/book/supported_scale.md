## Configuration Limits for CNS Manager

This topic provides the configuration limits for CNS manager and the features it offers. When you use CNS manager in your environment, stay within the supported and recommended limits.

**Number of Kubernetes clusters in vCenter registered with CNS manager**  
32   


**Number of concurrent PV migrations**   
On CNS manager application level, there can be 8 volume migrations that can be invoked in parallel across all clusters.  
And on vCenter level, the limits for simultaneous migrations can be be derived from this document - https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vcenterhost.doc/GUID-25EA5833-03B5-4EDD-A167-87578B8009B3.html

This translates to concurrent migration limits to be 1 per VM(for attached volumes), 2 per host & 8 per datastore.   
If there are parallel FCD migrations invoked beyond these limits, they will be queued based on the limits for each type of resource. CNS manager supports queueing upto 400 volumes at any given time.