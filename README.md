# Sop Search 

## Description
This program takes the SOP documents from the openshift repository and indexes them into an
elasticsearch container as a way of making it easier to find an SOP document. This information is then displayed in a web interface which allows you to easily search for the SOP document you want.

## How to Use 
After launching the UI, simply type into the searchbox to find results. You can also choose to filter by the tags (aka the possible location of your document), author, and the name of the SOP to help narrow your search. You can further choose to sort your results by the last updated or by the most relevant. 

## Building the Docker Images 
To build the images used in deployment.yml:

### The Elasticsearch Image

1. build the image using the dockerfile in the elasticsearch folder
2. push image to your repo

### The Sop Search Image 

1. build the image using the dockerfile
2. push image to your repo

### The UI Image 

The UI image is slightly different as you have to make sure to update the build folder each time you change anything in the UI. 

1. make sure your build folder has been updated
   - `npm run build`
   - if above fails, try `npm install` or `npm update` and then trying the above command again
2. build the image using dockerfile in the ui folder
3. push image to your repo

## Deploying
1. Create the ssh secret yml file
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
2. Create the configmap file
```
apiVersion: v1
kind: ConfigMap
metadata:
  name: configmap
  namespace: sop-search
data:
  time: "5" #number of minutes before restarting routine
  elastic: http://localhost:9200 #location of elasticsearch data
  repourl: https://github.com/openshift/ops-sop-search/blob/master/ #repo url used for creating links
  reponame: ops-sop-search #name of the repository you're indexing
  gitscript: script.sh #location of the shell script
  giturl: git@github.com:openshift/ops-sop-search.git #clone with ssh
```
3. Create the service_account, role, and role_binding (using files from deploy folder)
4. Create the services (one for the UI and one for the elasticsearch using files from deploy folder)
5. Create the routes (one for UI and one for elasticsearch using files from deploy folder)
6. Deploy the deployment file
7. Access the application via the web address given in the UI route