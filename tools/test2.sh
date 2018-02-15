#!/bin/sh
# Script for installing the gt.m and its minimal configuration needed to compile gogtm
# Version for ubuntu and travis
export gtm_dist=/opt/fis-gtm/engine
export PATH=$gtm_dist:$PATH
export gtm_chset=UTF-8
export gtmgbldir=/tmp/gtm.gld
export gtm_tmp=/tmp/gtm
export gtm_icu_version=5.0
alias gde="$gtm_dist/mumps -run GDE"
export gtmroutines=$gtm_dist/utf8/libgtmutil.so
