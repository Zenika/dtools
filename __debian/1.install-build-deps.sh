#!/usr/bin/env bash

# NOTE:
# -----
# This docker container has been stripped down as much as possible, and works for GO software
# The following extra packages might be needed for languages other than GO :
# apt install -y g++ fakeroot devscripts build-essential

echo "Installing dependencies";echo
apt-get update && apt update -y
echo;echo;echo "Done. Now installing the Go binaries"
/opt/bin/install_golang.sh 1.22.4 amd64
