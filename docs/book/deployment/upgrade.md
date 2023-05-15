## Upgrading cns-manager

The easiest way to upgrade cns-manager would be to completely undeploy the current version, checkout the targetted release & use deployment artifacts from the that release. The steps to deploy are already available [here](../../../README.md#deploying-cns-manager).

But if you want to preserve some of the earlier configurations such as oAuth2 configuration,  clusters that were already registered etc., you can perform below steps:

**1.** Update Swagger config to reflect the newly added API endpoints(if any) in Swagger UI.  
This can be done using following commands.

```
> git checkout <release-tag>
> kubectl -n <cns-manager-namespace> delete configmap swagger-api
> export CNS_MANAGER_ENDPOINT=<endpoint_url>
> sed "s/%CNS_MANAGER_ENDPOINT%/$CNS_MANAGER_ENDPOINT/g" deploy/swagger-template.yaml > swagger.yaml
> kubectl -n <cns-manager-namespace> create configmap swagger-api --from-file=swagger.yaml
> rm swagger.yaml
```

Here `CNS_MANAGER_ENDPOINT` will be the endpoint over which CNS manager service is accessible (<<ip:port>> or FQDN set during deployment).

**2.** Update nginx config to reflect any changes done for nginx proxy.
```
> git checkout <release-tag>
> kubectl -n <cns-manager-namespace> delete configmap nginx-conf
> kubectl -n <cns-manager-namespace> create configmap nginx-conf --from-file=<auth-folder>/nginx.conf
```

Here `auth-folder` is the folder corresponding to your deployment type - `deploy/basic-auth` or `deploy/oauth2`.

**3.** Check if orphan volume auto-deletion is disabled. It's recommended to keep it disabled until you fully understand its usage (Read the [orphan volume feature documentation](../features/orphan_volumes.md#setting-up-auto-monitoring-for-orphan-volumes-deletion) for details).  

The desired value can be set in `auto-delete-ov` field in `cnsmanager-config` configmap.

```
kubectl edit configmap cnsmanager-config -n <cns-manager-namespace>
```

**4.** Update the new release image in cns-manager deployment. For instance, for upgrading to release 0.2.0:
```
kubectl set image deployment/cns-manager cns-manager=projects.registry.vmware.com/cns_manager/cns-manager:r0.2.0 -n <cns-manager-namespace>
```

This will restart the deployment including updating the nginx config as well as new API endpoints in Swagger UI.