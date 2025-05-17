#!/bin/bash

# ensure ansible, vagrant are installed

## gen a private ssh key
#if [ ! -d ./.vagrant.d ]; then
#    mkdir ./.vagrant.d
#fi
#
#if [ ! -e ./.vagrant.d/worker_key ]; then
#    ssh-keygen -t ed25519 -f ./.vagrant.d/worker_key -N ""
#else
#    echo "SSH Key exists for hosts already."
#fi
#echo "Done generating SSH Key."

# install ansible requirements
ansible-galaxy install -r requirements.yml
echo "Done setting up ansible-galaxy dependencies"
