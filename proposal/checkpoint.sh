#!/bin/bash
podman save --output demo-app.tar demo-app
podman container checkpoint demo-app --export /vagrant/demo-app-cp.tar.gz
