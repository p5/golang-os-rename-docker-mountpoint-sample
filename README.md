# golang-os-rename-docker-mountpoint-sample

An example repository showing how we cannot use the `os.Rename` function to rename a file on a Docker mountpoint.

## How to run

```console
$ docker build -t test:local .
$ docker run --rm -v=./test.txt:/app/test.txt test:local /go/bin/app /app/test.txt
```

## Expected result

```console
❯ docker build -t test:local .
[+] Building 3.2s (9/9) FINISHED                                                                                                                                                                                     docker:default
 => [internal] load build definition from Dockerfile                                                                                                                                                                           0.0s
 => => transferring dockerfile: 201B                                                                                                                                                                                           0.0s
 => [internal] load .dockerignore                                                                                                                                                                                              0.0s
 => => transferring context: 108B                                                                                                                                                                                              0.0s
 => [internal] load metadata for docker.io/library/golang:1.20                                                                                                                                                                 1.2s
 => [1/4] FROM docker.io/library/golang:1.20@sha256:685846c2eb68f6ae4d784d56125ea9894981b3408340aa7d89362fef83e1aebb                                                                                                           0.0s
 => [internal] load build context                                                                                                                                                                                              0.0s
 => => transferring context: 5.66kB                                                                                                                                                                                            0.0s
 => CACHED [2/4] WORKDIR /go/src/app                                                                                                                                                                                           0.0s
 => [3/4] COPY . .                                                                                                                                                                                                             0.1s
 => [4/4] RUN go build -o /go/bin/app main.go                                                                                                                                                                                  1.8s
 => exporting to image                                                                                                                                                                                                         0.1s
 => => exporting layers                                                                                                                                                                                                        0.1s
 => => writing image sha256:07fd39954bafc88cb49554d0e49d55fa8867d93c752e308d0af9c1a2c43df3d0                                                                                                                                   0.0s
 => => naming to docker.io/library/test:local                                                                                                                                                                                  0.0s

❯ docker run --rm -v=./test.txt:/app/test.txt test:local /go/bin/app /app/test.txt
File name: /app/test.txt
New file created: newfile.txt
rename newfile.txt /app/test.txt: device or resource busy
```
