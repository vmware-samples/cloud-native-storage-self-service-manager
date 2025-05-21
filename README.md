# cloud-native-storage-self-service-manager

CNS Manager is a diagnostic and self-service tool that helps detect and auto-remediate some of the known issues in storage control plane in vCenter.
CNS Manager exposes APIs that can be invoked by authorized users to detect issues.

This repository provides artifacts for deploying CNS manager in different versions of vSphere Supervisors starting from 8.0, as well as the client sdk to invoke its endpoints.

## Deploying cns-manager on [vSphere Supervisor 8.0](https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere-supervisor/8-0/vsphere-supervisor-concepts-and-planning/vsphere-iaas-control-plane-concepts/what-is-vsphere-with-tanzu.html)
CNS manager needs to be deployed in one of the namespaces running on the Supervisor in the vCenter.
If there are multiple Kubernetes clusters in a vCenter, it's recommended that it be deployed in a dedicated admin-managed cluster, but it's not a must. However, the admin should be responsible to secure the Kubernetes cluster where CNS manager is deployed since it will have credentials to vCenter and the Kubernetes cluster.  
Also if you want CNS manager to be highly available, deploy it on a Kubernetes cluster that's highly available itself.

Note : To deploy CNS manager from this repo, you can clone it on your machine and then set kubeconfig to point to the remote Kubernetes cluster where CNS manager needs to be deployed. Then follow the instructions for deployment.

The deployment is supported with two authentication mechanisms to limit who can access CNS manager APIs:
1. Basic Auth - The CNS manager admin can choose fixed credentials at the time of deployment. This auth mechanism is less secure than OAuth2 to be used in Production. Nevertheless, it can be used for a quick deployment to test the application and in air-gapped environments where the vCenter is not connected to the internet.
See these [instructions](docs/book/deployment/basicauth.md) for basic auth deployment.

2. OAuth2 - With OAuth2, the authentication is delegated to an OIDC provider such as Gitlab, Github, Google etc. It does require creating an OAuth application on the OIDC provider before deploying CNS manager.  
See these [instructions](docs/book/deployment/oauth2.md) for OAuth2 deployment.

## Enabling TLS for your deployment
You can enable TLS for your CNS Manager deployment with a few tweaks, so that the communication is encrypted between client(a browser, for instance) and the application.  
See these [deployment changes](docs/book/deployment/tls-enable.md) to enable TLS on CNS Manager. It can be done for both basicauth & OAuth2 deployments, and assumes you have the TLS key and certificate generated.

## Register Kubernetes clusters before you start!
CNS manager relies on communicating with Kubernetes clusters for several functionalities it offers. It is therefore a pre-requisite to register all Kubernetes clusters in vCenter with CNS manager.  

Note: CNS manager can support upto 32 Kubernetes clusters per vCenter. Please see [supported scale](docs/book/supported_scale.md) for any recommended configurations for CNS manager.

The following section explains how to register a Kubernetes cluster with CNS manager. These steps are applicable to all Kubernetes clusters in the vCenter.

**1. Generate a kubeconfig with minimal privileges for CNS manager:**  
* The provided script `scripts/get-kubeconfig.sh` generates a kubeconfig for CNS manager with minimal privileges required for its functioning. But if you're fine with providing admin kubeconfig for the cluster to be registered, you can skip kubeconfig generation part mentioned below and directly jump to cluster registration part.  

Note : The script may not work on all Kubernetes distributions if they don't adhere to the [recommended steps](https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/container-storage-plugin/3-0/getting-started-with-vmware-vsphere-container-storage-plug-in-3-0.html) for deploying vSphere CSI driver.

* The script takes 2 mandatory input parameters. First is the path to the cluster's kubeconfig file and the second is the name of the file where the generated kubeconfig file with minimal privileges should be stored. Here is how you can run the script:
```
./get-kubeconfig.sh <kubeconfig_file_path> <output_file_name>
```
* If you have a combined kubeconfig with multiple contexts defined to access multiple servers, you will also need to provide context-name and server URL to access the desired Kubernetes cluster.
Example:
```
./get-kubeconfig.sh <kubeconfig_file_path> <output_file_name> <context name> <server URL>
```
* The script will output the newly generated kubeconfig in the provided output_file_name.

**2. Register cluster with CNS manager using kubeconfig**:
* Invoke `/registercluster` API on CNS manager by uploading the kubeconfig file. You may also modify other input parameters for the API based on your cluster configuration.
The API can also be invoked from command line. Here is an example:
```
curl -X 'POST' "http://CNS-MANAGER-ENDPOINT/1.0.0/registercluster?csiDriverSecretName=vsphere-config-secret&csiDriverSecretNamespace=vmware-system-csi" -H 'accept: application/json' -H 'Content-Type: multipart/form-data' -F 'clusterKubeConfigFile=@output_file_name' -u "Admistrator:Admin123@"
```
* Once the cluster is registered, you may delete this file from the machine.

**Note**: If a registered cluster later gets decommissioned or deleted from the vCenter, don't forget to deregister it from CNS manager as well. This will ensure a smooth execution of functionalities offered through CNS manager.

## Functionalities currently offered through cns-manager

* **Orphan volumes detection & deletion**  
This feature allows detecting/deleting orphan volumes that are not being used in any of the registered Kubernetes clusters on the vCenter. Read [here](docs/book/features/orphan_volumes.md) for more details about this feature.

* **Orphan snapshots detection & deletion**  
This feature allows detecting/deleting orphan snapshots that are not being used in any of the registered Kubernetes clusters on the vCenter. Read [here](docs/book/features/orphan_snapshots.md) for more details about this feature.