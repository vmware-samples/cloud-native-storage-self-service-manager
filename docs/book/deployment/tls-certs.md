
## Generating TLS key and certificate

* Pre-requisites - An internal/private CA(Certificate Authority) that is trusted by all machines in your private network.

The Internal CA section below is just used for demonstration purposes. Skip to [Generate CSR](#generate-csr) section directly if you already have a functioning CA that can issue certificates.

### Set up an internal CA  
First step is to set up an internal CA. Ignore this section if you already have a CA.  

Steps:  
1. Generate private key for for the internal CA.
```
openssl genrsa -des3 -out myCA.key 2048
```

2. Generate root certificate for the CA using the private key above.
```
openssl req -x509 -new -nodes -key myCA.key -sha256 -days 1825 -out myCA.pem
```

This root CA needs to be added to client machine's trust store so that browsers recognise this as a valid root CA.


### Generate CSR
Now that the CA is ready, generate a CSR which can be sent to the Internal CA to issue a valid certificate.  

Steps:  
1. Generate private key.
```
openssl genrsa -out helloabc.com.key 2048
```
Private key is stored in helloabc.com.key file

2. Create CSR with the above private key
```
openssl req -new -key helloabc.com.key -out helloabc.com.csr
```
The CSR is stored in helloabc.com.csr file

Make sure to give a proper Common Name(CN), that is, the domain you expect CNS manager to be accessible on. For example, if you give Common Name as helloabc.com then CNS manager will be available on https://helloabc.com

3. Create a .ext file, say `helloabc.com.ext` having config details and a Subject Alternative Name. The SAN should be same as CN provided in the above step.
```
authorityKeyIdentifier=keyid,issuer
basicConstraints=CA:FALSE
keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
subjectAltName = @alt_names
 
[alt_names]
DNS.1 = helloabc.com
```

### Generate certificate
Generate the certificate with the following command:
```
openssl x509 -req -in helloabc.com.csr -CA myCA.pem -CAkey myCA.key -CAcreateserial -out helloabc.com.pem -days 825 -sha256 -extfile helloabc.com.ext
```

It takes as input:  
- CSR - helloabc.com.csr
- Internal CA certificate -  myCA.pem
- Private key - myCA.key
- extfile - helloabc.com.ext

Output(issued certificate) is stored in file `helloabc.com.pem`


### Important Notes
1. The key and certificate files generated `helloabc.com.key` and `helloabc.com.pem` can be used during CNS manager deployment to enable TLS.  
2. CNS manager currently does not manage certificate lifecycle. Thus, the responsibility to renew certificates must be taken care of manually or by some other preferred way.  
3. After deployment, ensure there's a DNS entry for the the domain name used here to the CNS manager service ExternalIP/NodePort running in Kubernetes cluster.