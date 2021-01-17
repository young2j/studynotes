cd /dockers
docker-compose down
docker-compose pull
docker-compose up -d dev
noneimgs=$(docker images -aq --filter dangling=true)
$noneimgs && docker rmi $noneimgs