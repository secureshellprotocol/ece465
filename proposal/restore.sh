#!/bin/bash
podman rm -a
podman load --input /vagrant/demo-app.tar
podman container restore --import /vagrant/demo-app-cp.tar.gz
