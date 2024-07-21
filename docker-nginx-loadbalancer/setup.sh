#!/bin/bash
sudo docker stop $(docker ps -q) && docker rm $(docker ps -aq)
sudo docker-compose down
sudo docker-compose up --build -d