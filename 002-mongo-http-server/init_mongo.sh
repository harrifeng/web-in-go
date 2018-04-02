docker-compose stop
docker-compose rm -vf
docker-compose up -d
sleep 3
mongo localhost:27027/web-in-go --quiet insert_member.js
