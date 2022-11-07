```sh
docker build --tag=incrementor .
docker run incrementor
docker run --env DATA_PATH=/data/num.txt --mount type=volume,src=incrementor-data,target=/data incrementor
```