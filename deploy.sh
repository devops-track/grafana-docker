#!/bin/bash
#Variables
CONTAINER=grafana
CONTAINER_NAME=grafana_grafana
SERVICE=grafana
REPO=/dados1/docker-config/grafana/grafana-docker

docker stop "${CONTAINER}" || true && docker rm "${CONTAINER}" || true
docker rmi "${CONTAINER_NAME}"
docker system prune -a -f 

  cd "${REPO}"
  git pull --rebase
  docker-compose up -d ${SERVICE}
