#!/bin/sh

export gtm_dist=/opt/fis-gtm/engine
sudo ln -s /opt/fis-gtm/6.3-001A $gtm_dist
mkdir /tmp/gtm
export PATH=$gtm_dist:$PATH
export gtm_chset=UTF-8
export gtmgbldir=/tmp/gtm.gld
export gtm_tmp=/tmp/gtm
export gtm_icu_version=5.2
alias gde="$gtm_dist/mumps -run GDE"
export gtmroutines=$gtm_dist/utf8/libgtmutil.so
