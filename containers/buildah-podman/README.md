```
docker run -it --rm -p 3000:3000 --privileged --mount type=bind,source="$(pwd)",target=/src  --mount type="volume",src=podman-data,target=/var/lib/containers tomkukral/buildah bash

# buildah bud -f ./Dockerfile -t my-app-buildah .
buildah build-using-dockerfile -f ./Dockerfile -t my-app-buildah .


buildah from my-app-buildah

buildah run --net host my-app-buildah-working-container -- sh

buildah rm my-app-buildah-working-container
```

```
# use podman to run builda image
podman run --cgroup-manager cgroupfs -p 3000:3000 localhost/my-app-buildah
```


```
# run a buildah container and bind mount docker into container
docker run -it --rm -p 3000:3000 --privileged --mount type=bind,source="$(pwd)",target=/src -v /var/run/docker.sock:/var/run/ docker.sock  --mount type="volume",src=podman-data,target=/var/lib/containers tomkukral/buildah bash

# use buildah to build OCI(Open Container Initiative) image
buildah bud -f ./Dockerfile -t my-app-buildah .

# push to a local Docker daemon (background process.)
buildah push localhost/my-app-buildah docker-daemon:my-app-buildah:latest

# run OCI image with docker
docker run -it my-app-buildah bash
```

