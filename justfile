podman-build:
  podman build -f Dockerfile -t test:local .

podman-run:
  podman run --rm -v=./test.txt:/app/test.txt:Z test:local /go/bin/app /app/test.txt

docker-build:
  docker build -f Dockerfile -t test:local .

docker-run:
    docker run --rm -v=./test.txt:/app/test.txt:Z test:local /go/bin/app /app/test.txt
