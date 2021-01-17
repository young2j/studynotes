#!/bin/bash

docker-compose down
docker-compose pull
docker-compose up -d dev
noneimgs=$(docker images -aq --filter dangling=true)
len=$(echo -n "$noneimgs" | wc -c)
if (( $len>0 ))
then
  docker rmi $noneimgs
else
  echo "there is no images tagged none."
fi