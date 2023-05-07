VERSION=v1
docker build -t guptang/tiktok-user:$VERSION -f ./Dockerfile ../../
docker push guptang/tiktok-user:$VERSION