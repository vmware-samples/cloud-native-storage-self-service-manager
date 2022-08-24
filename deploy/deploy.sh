#!/bin/bash

if [ $# -lt 6 ]
then
	echo "./deploy.sh <namespace> <SV kubeconfig file> <VC creds file> <CNS manager endpoint> <basicauth|oauth2> <(tls flag)true|false>
	<BasicAuth Username(required with basicauth)> <BasicAuth Password(required with basicauth)>
	<path-to-tls.key(required if tls enabled)> <path-to-tls.pem(required if tls enabled)>"
	exit 1
fi

NAMESPACE=$1
SV_KUBECONFIG_FILE=$2
VC_CREDS_FILE=$3
CNS_MANAGER_ENDPOINT=$4
AUTH_MECHANISM=$5
TLS_FLAG=$6

if ! [ $AUTH_MECHANISM == "basicauth" -o $AUTH_MECHANISM == "oauth2" ]
then
	echo "Auth mechanism needs to be either basicauth or auth2."
	exit 1
fi

# Set variables depending on the values of auth mechanism and tls flag.
if [ $AUTH_MECHANISM == "basicauth" ]
then
	BASICAUTH_USERNAME=$7
	BASICAUTH_PASSWORD=$8

	if [ $TLS_FLAG == "true" ]
	then
		TLS_KEY=$9
		TLS_CERT=${10}
	fi
else
	if [ $TLS_FLAG == "true" ]
	then
		TLS_KEY=$7
		TLS_CERT=$8
	fi
fi

# Create a namespace where CNS manager will be deployed.
kubectl create ns $NAMESPACE
if [ $? -ne 0 ]
then
	echo "Failed to create namespace $NAMESPACE."
	exit 1
fi

# Create a secret that has SV cluster's admin kubeconfig.
# Edit the sv_kubeconfig file as per your deployment.
kubectl -n $NAMESPACE create secret generic sv-kubeconfig --from-file=$SV_KUBECONFIG_FILE
if [ $? -ne 0 ]
then
	echo "Failed to create sv-kubeconfig secret."
	exit 1
fi

# Create a secret that has the vCenter's admin creds.
# Edit vc_creds.json as per your requirement.
kubectl -n $NAMESPACE create secret generic vc-creds --from-file=$VC_CREDS_FILE
if [ $? -ne 0 ]
then
	echo "Failed to create vc-creds secret."
	exit 1
fi

# If auth mechanism is basicauth, create a secret that has basic auth credentials hashed using MD5-based
# password algorithm. Create a temp file with credentials and remove it after creating the secret.
if [ $AUTH_MECHANISM == "basicauth" ]
then
    # Assign manifest folder
    MANIFEST_FOLDER="basic-auth"
    echo -n $BASICAUTH_USERNAME: >> basicauth_creds
    openssl passwd -apr1 $BASICAUTH_PASSWORD >> basicauth_creds
    kubectl -n $NAMESPACE create secret generic basicauth-creds --from-file=basicauth_creds
    if [ $? -ne 0 ]
    then
        echo "Failed to create basicauth-creds secret."
        exit 1
    fi
    rm basicauth_creds
else
    # Assign manifest folder
    MANIFEST_FOLDER="oauth2"
fi

# Create a config map for nginx-conf
kubectl -n $NAMESPACE create configmap nginx-conf --from-file=$MANIFEST_FOLDER/nginx.conf
if [ $? -ne 0 ]
then
	echo "Failed to create nginx-conf config map."
	exit 1
fi

# If ssl is set to true, create a secret to store ssl key and cert.
if [ $TLS_FLAG == "true" ]
then
	kubectl -n $NAMESPACE create secret tls cnsmanager-tls --key $TLS_KEY --cert $TLS_CERT	
	if [ $? -ne 0 ]	
	then	
		echo "Failed to create cnsmanager-tls secret."	
		exit 1	
	fi
fi

# Create a config map for the CNS manager swagger API spec.
sed "s/%CNS_MANAGER_ENDPOINT%/$CNS_MANAGER_ENDPOINT/g" swagger-template.yaml > swagger.yaml
kubectl -n $NAMESPACE create configmap swagger-api --from-file=swagger.yaml
if [ $? -ne 0 ]
then
	echo "Failed to create swagger-api config map."
	exit 1
fi
rm swagger.yaml

# Deploy CNS manager
sed -e "s#%CNS_MANAGER_ENDPOINT%#$CNS_MANAGER_ENDPOINT#g" \
    -e "s#%CNS_MANAGER_NAMESPACE%#$NAMESPACE#g" \
		$MANIFEST_FOLDER/deploy-template.yaml > deploy.yaml
kubectl -n $NAMESPACE apply -f deploy.yaml
if [ $? -ne 0 ]
then
	echo "Failed to deploy CNS manager."
	exit 1
fi
rm deploy.yaml
