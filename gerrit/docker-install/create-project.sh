#! /bin/bash

WEBUI_IP=`ifconfig eno1 | perl -nle 's/dr:(\S+)/print $1/e'`

PROJECT_NAME=$1
PROJECT_GITLAB=$2
GERRIT_WEBUI=http://${WEBUP_IP}:8081

git clone --bare $PROJECT_GITLAB

sudo chown -R ubuntu:ubuntu ${PROJECT_NAME}.git/
sudo rm -rf ~/gerrit/git/${PROJECT_NAME}.git/
sudo mv ${PROJECT_NAME}.git/ ~/gerrit/git/

ssh -p 29418 astri@localhost gerrit flush-caches
ssh -p 29418 astri@localhost gerrit ls-projects

# Just for Jenkins
#
# ## Create Jenkins Project
# cp jenkins/job.xml /tmp/${PROJECT_NAME}-job.xml
# sed -i "s~KK_GERRIT_PROJECT_WEBUI_ADDR~${GERRIT_WEBUI}/p/${PROJECT_NAME}~g" /tmp/${PROJECT_NAME}-job.xml
# sed -i "s~KK_GERRIT_PROJECT_NAME~${PROJECT_NAME}~g" /tmp/${PROJECT_NAME}-job.xml
# curl -X POST -H "Content-Type: application/xml" -d @/tmp/${PROJECT_NAME}-job.xml http://192.168.205.16:8082/createItem?name=${PROJECT_NAME}
# 
# ## Delete Jenkins Project
# # ssh -p 29418 astri@localhost deleteproject delete --yes-really-delete ${PROJECT_NAME}
# # curl -X POST http://192.168.205.16:8082/job/${PROJECT_NAME}/doDelete
