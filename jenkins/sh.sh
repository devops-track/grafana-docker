#!/bin/bash
docker stop grafana || true && docker rm grafana || true
docker rmi grafana_grafana
repos=( 
  "/dados1/docker-config/grafana/grafana-docker"
)

echo ""
echo "Getting latest for" ${#repos[@]} "repositories using pull --rebase"

for repo in "${repos[@]}"
do
  echo ""
  echo "****** Getting latest for" ${repo} "******"
  cd "${repo}"
  git pull --rebase
  echo "******************************************"
done
cd "${repo}"
docker-compose up