# syntax=docker/dockerfile:1
FROM ubuntu:22.04

# IMPORTANT !!!!!!!!!!!!!
# This Dockerfile should be in the same directory as 'main.py'

# install app dependencies
RUN apt-get update
RUN apt-get install -y python3 python3-pip

RUN python3 -m pip install -r ./requirements.txt

COPY playbook.yml /etc/ansible/playbook.yml
COPY main.py /etc/ansible/index.py

RUN mkdir -m 755 /opt/demo-app

RUN ansible-playbook /etc/ansible/playbook.yml --connection=local

EXPOSE 8000
