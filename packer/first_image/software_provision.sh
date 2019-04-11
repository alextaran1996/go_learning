#!/bin/bash
# Install fedora repo
yum -y install epel-release
# Update the system
sudo yum -y update
# Docker installation
yum install -y yum-utils device-mapper-persistent-data lvm2 perl gcc dkms make bzip2 patch
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
yum -y install kernel-headers-$(uname -r) 
yum -y install kernel-devel-$(uname -r)
yum -y install kernel-devel
/sbin/rcvboxadd setup
curl -O http://download.virtualbox.org/virtualbox/6.0.4/VBoxGuestAdditions_6.0.4.iso
mount -o loop ./VBoxGuestAdditions_6.0.4.iso /mnt
sh /mnt/VBoxLinuxAdditions.run
umount /mnt
rm VBoxGuestAdditions_6.0.4.iso
reboot