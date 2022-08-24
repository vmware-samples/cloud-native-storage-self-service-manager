## Deploying CNS Manager with OAuth2

CNS Manager uses oauth2-proxy for providing authentication using OIDC providers such Gitlab, GitHub,Google and others). Please see the list of supported OIDC providers [here](https://oauth2-proxy.github.io/oauth2-proxy/docs/configuration/oauth_provider)

Below are the steps to configure, deploy & run cns-manager on a vanilla Kubernetes cluster with OAuth2.

### Prepare the config
1. Capture kubeconfig of the cluster in which CNS manager is being deployed in a file named sv_kubeconfig.  
Refer to sample config file provided under config folder. The kube config on master VM can be checked using below command:
```
// On master VM
cat ~/.kube/config
```

2. Create a file named vc_creds.json and copy into it the credentials to your VC.  
Refer to sample config file provided under config folder.
```
{
    "vc": "10.187.99.154",
    "user": "vc-user@domain",
    "password": "vc-password"
}
```

3. Choose any of the worker nodes to act as an endpoint for CNS manager APIs. Let's call this worker node's IP as WORKER_NODE_IP.
The nodePort is set to 30008 by default in *deploy-template.yaml*.
So the CNS manager endpoint would be <WORKER_NODE_IP>:30008.

Note : If your cloud provider supports a load balancer, you can choose to deploy a load balancer service instead. In that case, the CNS manager endpoint would be <LB_SVC_EXTERNAL_IP>:30008

Also if you need to change kubeconfig or VC creds after the deployment script has run, then you can either:  
a. Recreate the secrets sv-kubeconfig & vc-creds created from these files and restart the cns- manager deployment, OR  
b. Delete the namespace and run the deployment script again.

### Register an OAuth application and Update configuration
1. Select a provider and register an OAuth application with it.  
Use `http://<CNS_manager_endpoint>:30008/oauth2/callback` for Redirect URL.

2. In provided oauth2 [deployment yaml](../../../deploy/oauth2/deploy-template.yaml) that uses Gitlab as an example, update the provider, client-id & client-secret arguments with corresponding values from your registered OAuth application.  
Check [oauth2-proxy documentation](https://oauth2-proxy.github.io/oauth2-proxy/docs/configuration/oauth_provider/#gitlab-auth-provider) for configuration for each provider.

### Deploy the application
* After preparing the config, use the following command to deploy CNS manager on the cluster.

```
> cd deploy
> ./deploy.sh <namespace> <path-to-sv_kubeconfig> <path-to-vc_creds.json> <CNS manager endpoint> <authType> <(tls flag)true|false> <BasicAuth Username(required with basicauth)> <BasicAuth Password(required with basicauth)> <path-to-tls.key(required if tls enabled)> <path-to-tls.pem(required if tls enabled)>
```

For example
```
> ./deploy.sh cns-manager ../config/sv_kubeconfig ../config/vc_creds.json <WORKER- NODE-IP/LB_IP>:30008 oauth2 false
```

* The deployment script will create a bunch of Kubernetes objects. The sample output should look like as described below.  
Once the deployment is successful, verify that CNS manager pod is running in the namespace.
```
> ./deploy.sh cns-manager ../config/sv_kubeconfig ../config/vc_creds.json 10.184.71.61:30008 oauth2 false

namespace/cns-manager created
secret/sv-kubeconfig created
secret/vc-creds created
secret/basicauth-creds created
configmap/swagger-api created
serviceaccount/cns-manager created 
rolebinding.rbac.authorization.k8s.io/cns-manager created 
customresourcedefinition.apiextensions.k8s.io/orphanvolumestats.cnsmanager.cns.vmware.com configured 
customresourcedefinition.apiextensions.k8s.io/volumemigrationjobs.cnsmanager.cn s.vmware.com configured 
customresourcedefinition.apiextensions.k8s.io/volumemigrationtasks.cnsmanager.c ns.vmware.com configured
clusterrole.rbac.authorization.k8s.io/cns-manager configured 
clusterrolebinding.rbac.authorization.k8s.io/cns-manager-rolebinding unchanged 
configmap/cnsmanager-config created
configmap/nginx-conf created
deployment.apps/cns-manager created 
service/cns-manager created

> kubectl get pods -n cns-manager  
NAME                           READY  STATUS    RESTARTS        AGE 
cns-manager-6ff456dc97-nrj65   4/4    Running       0           37s
```

* After the deployment, the Swagger UI for invoking APIs can be accessed at <CNS_Manager_Endpoint>/ui/  
This will redirect to the OIDC provider for authentication.

After authentication, the APIs can be directly invoked from Swagger UI.