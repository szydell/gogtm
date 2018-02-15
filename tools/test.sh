#!/bin/sh
# Script for installing the gt.m and its minimal configuration needed to compile gogtm
# Version for ubuntu and travis

set -ex
sudo rm -rf /tmp/gtm_install
mkdir /tmp/gtm_install
cd /tmp/gtm_install
wget https://downloads.sourceforge.net/project/fis-gtm/GT.M-amd64-Linux/V6.3-001A/gtm_V63001A_linux_x8664_pro.tar.gz
tar -xzf gtm_V63001A_linux_x8664_pro.tar.gz
#sudo apt-get install -y libicu-dev icu-devtools
sudo rm -rf /opt/fis-gtm/6.3-001A
sudo ./configure << EOF
bin
bin
n
/opt/fis-gtm/6.3-001A
y
y
y
5.0
y
y
y
EOF
rm -rf /tmp/gtm_install
# Create symbolic link used for gogtm compilation
export gtm_dist=/opt/fis-gtm/engine
sudo ln -s /opt/fis-gtm/6.3-001A $gtm_dist
mkdir -p /tmp/gtm
