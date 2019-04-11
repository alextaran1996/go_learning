#!/bin/bash
# Install fedora repo
yum -y install epel-release
# Update the system
yum -y update
#Install additional packages
yum install -y yum-utils device-mapper-persistent-data lvm2 perl gcc dkms make bzip2 patch 
yum -y install kernel-headers-`(uname -r)`
# Docker installation
yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
yum -y install docker-ce
usermod -aG docker vagrant
usermod -aG docker root
systemctl enable docker.service
# Install ansible 
yum -y install ansible 
# Golang installation
curl -O https://dl.google.com/go/go1.12.3.linux-amd64.tar.gz
tar -C /usr/local -xvzf go1.12.3.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile
rm -rf go1.12.3.linux-amd64.tar.gz
# Install VBoxGuestAdditions
echo "########################################Check here"
yum -y install kernel-devel-`(uname -r)`
sudo yum -y install kernel-headers kernel-devel
echo "####################################################"

