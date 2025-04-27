#!/bin/bash
docker build -t test:latest .
docker run -p 127.0.0.1:8000:8000 test:latest
