#! /bin/bash

WEBUI_IP=`ifconfig eno1 | perl -nle 's/dr:(\S+)/print $1/e'`

CreateSSHKey() {
    rm -rf .ssh ~/.ssh
    mkdir .ssh ~/.ssh
    ssh-keygen -t rsa -b 4096 -C astri@gerrit -f .ssh/id_rsa -N ''
    cp .ssh/* ~/.ssh/
}

CreateLDAP() {
    mkdir ~/ldap
    mkdir ~/ldap/database
    mkdir ~/ldap/config
    cp ldap/gerrit.ldif ~/ldap/
    sudo chmod 777 ~/ldap

    sudo docker run -v ~/ldap/database:/var/lib/ldap -v ~/ldap/config:/etc/ldap/slapd.d \
               -h ldap.astri.org -e LDAP_ORGANISATION="ASTRI" -e LDAP_DOMAIN="astri.org" \
               -e LDAP_ADMIN_PASSWORD="astri.org" \
	       --name ldap \
               -d osixia/openldap
    
    sudo docker run --link ldap:ldap \
               --env PHPLDAPADMIN_LDAP_HOSTS=ldap \
	       -p 6443:443 \
	       --name ldapadmin \
               -d osixia/phpldapadmin
}

ConfigureLDAP() {
    sudo docker cp ~/ldap/gerrit.ldif ldap:/tmp/gerrit.ldif
    sudo docker exec ldap ldapadd -x -H ldap://localhost -D cn=admin,dc=astri,dc=org -w astri.org -f /tmp/gerrit.ldif
}


CreateGerrit() {
    mkdir ~/gerrit
    cp -r gerrit/etc ~/gerrit/
    cp -r gerrit/hooks ~/gerrit/
    sudo chmod 777 ~/gerrit

    sed -i "s/KK_WEBUI_IP/${WEBUI_IP}/g" ~/gerrit/etc/gerrit.config

    sudo docker run --link ldap:ldap \
               -v ~/gerrit:/var/gerrit/review_site -v /etc/localtime:/etc/localtime:ro \
               -e WEBURL=http://${WEBUI_IP}:8081 \
               -e AUTH_TYPE=LDAP -e LDAP_SERVER=ldap://ldap -e LDAP_ACCOUNTBASE="dc=astri,dc=org" \
               -p 8081:8080 -p 29418:29418 \
               --name gerrit \
	       -d openfrontier/gerrit:2.13.x
}

ConfigureGerrit() {
    echo "Waitting for Gerrit..."
    while ! sudo docker logs gerrit 2>&1 | grep "Gerrit Code Review 2.13.11 ready"; do
        sleep 1;
    done
    echo "Gerrit is ready"

    echo "You need to copy bellow public SSH key to Gerrit web ui manually"
    echo "  - weblink: http://$WEBUI_IP:8081/#/settings/ssh-keys"
    echo "  - username: astri"
    echo "  - password: astri.org"
    echo "  - SSH Public Key:"
    cat .ssh/id_rsa.pub
    read

    ssh-keygen -f "/home/gerrit2/.ssh/known_hosts" -R [localhost]:29418
    ssh-keygen -f "/home/gerrit2/.ssh/known_hosts" -R [192.168.205.17]
    ssh-keyscan -p 29418 localhost > ~/.ssh/known_hosts
    ssh-keyscan 192.168.205.17 >> ~/.ssh/known_hosts

    ssh -p 29418 astri@localhost gerrit set-members --add astri "Administrators"
    ssh -p 29418 astri@localhost gerrit set-account jenkins --add-ssh-key -< .ssh/id_rsa.pub
    ssh -p 29418 astri@localhost gerrit set-members --add jenkins "Non-Interactive\ Users"

    # Introduce Verification Label, not used for now
    #
    # git clone ssh://astri@localhost:29418/All-Projects.git && \
    #     cd All-Projects && \
    #     git fetch origin refs/meta/config && \
    #     git checkout -b meta/config && \
    #     git am ../patches/0001-project.config-add-a-Verified-label.patch && \
    #     git am ../patches/0002-project.config-allow-Non-Interactive-Users-to-submit.patch && \
    #     git push origin HEAD:refs/meta/config && \
    #     cd .. && \
    #     rm -rf All-Projects

    ssh -p 29418 astri@localhost gerrit plugin install --name reviewers.jar   -< gerrit/plugins/reviewers.jar
    ssh -p 29418 astri@localhost gerrit plugin install --name replication.jar -< gerrit/plugins/replication.jar
    ssh -p 29418 astri@localhost gerrit plugin install --name hooks.jar       -< gerrit/plugins/hooks.jar
    ssh -p 29418 astri@localhost gerrit plugin reload replication
    ssh -p 29418 astri@localhost gerrit plugin reload reviewers
    ssh -p 29418 astri@localhost gerrit plugin reload hooks
    ssh -p 29418 astri@localhost gerrit plugin ls

    sudo docker exec gerrit apk update
    sudo docker exec gerrit apk add python

    sudo docker exec gerrit rm -rf /root/.ssh
    sudo docker cp .ssh gerrit:/root/.ssh
    sudo docker exec gerrit chown -R root:root /root/.ssh
    sudo docker exec gerrit sh -c "ssh-keyscan -p 29418 localhost > /root/.ssh/known_hosts"
    sudo docker exec gerrit sh -c "ssh-keyscan 192.168.205.17 >> /root/.ssh/known_hosts"

    sudo docker exec gerrit rm -rf /var/gerrit/.ssh
    sudo docker exec gerrit cp -r /root/.ssh /var/gerrit/
    sudo docker exec gerrit chown -R gerrit2:gerrit2 /var/gerrit/.ssh
}

CreateJenkins() {
    mkdir ~/jenkins
    mkdir ~/jenkins/.ssh
    sudo cp .ssh/* ~/jenkins/.ssh/
    sudo chown ubuntu:ubuntu -R ~/jenkins/
    sudo chmod 777 ~/jenkins
    
    sudo docker pull openfrontier/gerrit:2.13.x
    sudo docker pull openfrontier/jenkins
    
    sudo docker run --link ldap:ldap --link gerrit:gerrit \
    	       -v ~/jenkins:/var/jenkins_home \
               -e JAVA_OPTS="-Duser.timezone=Asia/Shanghai -Djenkins.install.runSetupWizard=false -Xms4096m" \
               -e GERRIT_HOST_NAME=${WEBUI_IP} \
               -e GERRIT_FRONT_END_URL=http://${WEBUI_IP}:8081 \
               -e GERRIT_PORT_29418_TCP_PORT=29418 \
               -e GERRIT_ENV_GERRIT_USER=jenkins \
               -e ROOT_URL=http://${WEBUI_IP}:8082/ \
               -e LDAP_SERVER=ldap://ldap -e LDAP_ROOTDN="dc=astri,dc=org" \
	       -p 8082:8080 -p 50000:50000 \
               --name jenkins \
               -d openfrontier/jenkins
}

ConfigureJenkins() {
    sudo docker cp jenkins/install_plugins.sh jenkins:/var/jenkins_home/install_plugins.sh
    sudo docker exec jenkins bash /var/jenkins_home/install_plugins.sh email-ext-recipients-column
    sudo docker exec jenkins bash /var/jenkins_home/install_plugins.sh golang
    sudo docker restart jenkins
    echo "Install completed! You still need to configure the email server and email message context templete:"
    echo "    http://${WEBUI_IP}:8082/configure"
}

CleanUp() {
    sudo docker stop jenkins gerrit ldapadmin ldap
    sudo docker rm jenkins gerrit ldapadmin ldap

    sudo rm -rf ~/gerrit
    sudo rm -rf ~/ldap
    sudo rm -rf ~/jenkins
}

CleanUp
# CreateSSHKey
CreateLDAP
ConfigureLDAP
CreateGerrit
ConfigureGerrit
# CreateJenkins
# ConfigureJenkins
