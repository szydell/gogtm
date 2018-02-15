#!/bin/sh
# Script for installing the gt.m and its minimal configuration needed to compile gogtm
# Version for ubuntu and travis

env

echo "exit" | $gtm_dist/mumps -run GDE
$gtm_dist/mupip create
