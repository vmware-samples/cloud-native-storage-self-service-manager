# cloud-native-storage-self-service-manager

CNS Manager is a diagnostic and self-service tool that helps detect and auto-remediate some of the known issues in storage control plane. It also provides certain table stake features, such as svMotion and datastore decomission to complement the CNS solution.
CNS Manager exposes APIs that can be invoked by authorized users to detect issues 


This repository provides artifacts for deploying CNS manager in vanilla Kubernetes cluster, as well as the client sdk to invoke its endpoints.

## Deploy cns-manager
CNS manager needs to be deployed in one of the kubernetes clusters in the vCenter.  
If there are multiple kubernetes clusters in a vCenter, it's recommended that it be deployed in a dedicated admin-managed cluster, but it's not a must.
However, the admin should be responsible to secure the Kubernetes cluster where CNS manager is deployed since it will have credentials to vCenter and the Kubernetes cluster.

To limit who can access CNS manager APIs, the deployment is supported with two authentication mechanisms:
1. Basic Auth - The CNS manager admin can choose fixed credentials at the time of deployment. This auth mechanism is less secure than OAuth2 and is not recommended to be used in Production. Nevertheless, it can be used for a quick deployment to test the application and in air-gapped environments where the vCenter is not connected to the internet.
See these [instructions](docs/book/deployment/basicauth.md) for basic auth deployment.

2. OAuth2 - With OAuth2, the authentication is delegated to an OIDC provider such as Gitlab, Github,Google etc. It does require creating an OAuth application on the OIDC provider before deploying CNS manager.  
See these [instructions](docs/book/deployment/oauth2.md) for OAuth2 deployment.

## Enable TLS for your deployment
You can enable TLS for your CNS Manager deployment with a few tweaks, so that the communication is encrypted between client(a browser, for instance) and the application.  
See these [deployment changes](docs/book/deployment/tls-enable.md) to enable TLS on CNS Manager. It can be done for both basicauth & OAuth2 deployments, and assumes you have the TLS key and certificate generated.
## Functionalities currently offered through cns-manager

* **Storage vMotion for CNS volumes**   
This feature allows migrating volumes from one datastore to another. Read [here](docs/book/features/storage_vmotion.md) for more details about the feature.


## Contributing

The cloud-native-storage-self-service-manager project team welcomes contributions from the community. Before you start working with cloud-native-storage-self-service-manager, please
read our [Developer Certificate of Origin](https://cla.vmware.com/dco). All contributions to this repository must be
signed as described on that page. Your signature certifies that you wrote the patch or have the right to pass it on
as an open-source patch. For more detailed information, refer to [CONTRIBUTING.md](CONTRIBUTING.md).

