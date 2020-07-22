# Sop Search 

## Description
This program takes the SOP documents from the openshift repository and indexes them into an
elasticsearch container as a way of making it easier to find an SOP document.

## How to Use 
After the program indexes the SOP documents into the elasticsearch container, the user can then search for specific documents based on the name, the authors, the content, etc. using curl commands once inside the elasticsearch container.

** Example ** 
`curl -X GET 'http://localhost:9200/sop/_search?q=_id%3Alogging'`

## How to Deploy

** Additional Files Not Included **
1. secret.yml file 

```
apiVersion: v1
data:
  ssh-privatekey:
  ssh-publickey:
kind: Secret
metadata:
  name: ssh
  namespace: sop-search
type: Opaque
```

2. configmap.yml file 

```
apiVersion: v1
kind: ConfigMap
metadata:
  name: configmap
  namespace: sop-search
data:
  time: 
  elastic: 
  repourl: 
  reponame: 
  gitscript: 
```

** Changes in Files **
1. in deployment change quay.io link to your quay.io link

### Building and Running
1. build image
2. push the image to a quay repo location
3. inside the oc environment, apply the service_account, role, and role_binding yml files along with your created configmap and secret files
4. apply the deployment yml file in the deploy folder