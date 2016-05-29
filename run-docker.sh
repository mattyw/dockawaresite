#!/bin/bash
set -e
sudo docker build -t dockerawaresite .
sudo docker run --name dasite6 -d -p 8080:8080 dockerawaresite
