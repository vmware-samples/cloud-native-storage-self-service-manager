## Deploying CNS Manager with TLS enabled

Below are the changes you need to make during deployment to enable TLS on CNS manager. This is on top of the deployment instructions for basicauth/oauth2, and assumes you have the TLS key and certificate already generated.  
(If you want to generate self-signed certificates using an internal/private CA, the instructions are provided [here](tls-certs.md) just as a reference.)

### Add ssl settings to nginx config
Make changes in `nginx.conf` to add ssl config.  
For basicauth, make changes in `deploy/basic-auth/nginx.conf`.  
For oauth2, make changes in  `deploy/oauth2/nginx.conf`. 

* Uncomment the ssl-related settings in nginx.conf
```
http {
    server {
        listen       80;
        # Use below instead to configure an https server
        #listen       443 ssl;
        .
        .
        #Uncomment this to enable ssl/tls communication over https. 
        #ssl_certificate     /etc/nginx/tls/tls.crt;
        #ssl_certificate_key /etc/nginx/tls/tls.key;
        .
        .
    .
    }
}
```
You can also add any additional settings in nginx config by following their documentation.


### Add TLS related changes to deployment yaml
Make changes in `deploy-template.yaml` to add TLS related changes. Select the file to change depending on the authentication mechanism.  
For basicauth, make changes in `deploy/basic-auth/deploy-template.yaml`.  
For oauth2, make changes in  `deploy/oauth2/deploy-template.yaml`. 

* Uncomment in deployment yaml the sections to define and mount the volume containing tls certificates.
```
.
.
---
apiVersion: apps/v1
kind: Deployment
.
.
    containers:
      - name: nginx-proxy
        .
        .
        volumeMounts:
          .
          .
        #uncomment below volume mount if tls is enabled.
          #- name: cnsmanager-tls
          #  mountPath:	/etc/nginx/tls
          #  readOnly: true
          .
    .
    volumes:
      .
      .
    #uncomment below volume creation if tls is enabled.
      #- name: cnsmanager-tls
      #  secret:	
      #    secretName: cnsmanager-tls
    .
.
.
```

### Pass additional command line arguments to deploy script
* After making the above changes, add TLS-related arguments while running the deployment script. You need to set `tls-flag` to `true` and provide the filepaths to TLS certificate and key.

```
> cd deploy
> ./deploy.sh <namespace> <path-to-sv_kubeconfig> <path-to-vc_creds.json> <CNS manager endpoint> <authType> <(tls flag)true|false> <BasicAuth Username(required with basicauth)> <BasicAuth Password(required with basicauth)> <path-to-tls.key(required if tls enabled)> <path-to-tls.pem(required if tls enabled)> 
```

For example
```
> cd deploy
> ./deploy.sh cns-manager ../config/sv_kubeconfig ../config/vc_creds.json <WORKER- NODE-IP/LB_IP>:30009 basicauth 'Administrator' 'Admin123@' /etc/tls/helloabc.com.key /etc/tls/helloabc.com.pem
```

For detailed deployment instructions, refer to the parent document for [basicauth](basicauth.md) and [oauth2](oauth2.md) deployments.
