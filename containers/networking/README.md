```sh
docker network create --driver=bridge app-net
docker run -d --network=app-net -p 27017:27017 --name=db --rm mongo:3

docker build -t app-with-mongo .
docker run -p 3000:3000 --network=app-net --env MONGO_CONNECTION_STRING=mongodb://db:27017 app-with-mongo
```