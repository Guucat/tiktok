VERSION=v1
docker build -t guptang/tiktok-comment:$VERSION -f ./Dockerfile ../../
docker push guptang/tiktok-comment:$VERSION