#!/bin/bash
cp .env.example .env
sudo docker stop $(docker ps -q) && docker rm $(docker ps -aq)
sudo docker-compose down
sudo docker-compose up --build -d
