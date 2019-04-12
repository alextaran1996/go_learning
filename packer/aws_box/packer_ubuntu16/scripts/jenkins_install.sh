#!/bin/bash
apt -y update 
apt -y install openjdk-8-jdk
wget -q -O - https://pkg.jenkins.io/debian/jenkins.io.key | sudo apt-key add -
sh -c 'echo deb http://pkg.jenkins.io/debian-stable binary/ > /etc/apt/sources.list.d/jenkins.list'
apt -y update 
apt -y install jenkins 
systemctl enable jenkins
systemctl start jenkins