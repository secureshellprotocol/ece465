# syntax=docker/dockerfile:1
FROM ubuntu:22.04

# IMPORTANT !!!!!!!!!!!!!
# This Dockerfile should be in the same directory as 'main.py'

# install app dependencies
RUN apt-get update && apt-get install -y python3 python3-pip
COPY . /app
WORKDIR /app
RUN python3 -m pip install -r ./requirements.txt

# final configuration
ENV FLASK_APP=main
EXPOSE 8000
