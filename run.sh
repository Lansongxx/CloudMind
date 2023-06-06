docker stop $(docker ps -aq)
docker-compose -f docker-compose-env.yml up -d
echo "wait 10s ..."
sleep 10
docker-compose up -d


