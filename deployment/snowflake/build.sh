VERSION=v1
docker build -t guptang/tiktok-snowflake:$VERSION -f ./Dockerfile ../../
docker push guptang/tiktok-snowflake:$VERSION