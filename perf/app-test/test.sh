#!/bin/bash

build() {
    GOOS=linux GOARCH=amd64 go build -o pjob main.go
    echo "FROM alpine:latest" > Dockerfile
    echo "RUN apk --no-cache add curl bash" >> Dockerfile
    docker build -t="leaky" .
}

run() {
    run-consul
    run-leaky
    run-profiler
}

# ----------------------------------------

run-leaky() {
    docker run -d \
           --name leaky \
           -p 6060:6060 \
           --link leaky_consul:consul \
           -v "$(pwd)/containerpilot:/bin/containerpilot" \
           -v "$(pwd)/containerpilot.json5:/etc/containerpilot.json5" \
           -v "$(pwd)/pjob:/bin/pjob" \
           leaky \
           /bin/containerpilot -config /etc/containerpilot.json5
}

run-consul() {
    docker run -d \
           --name leaky_consul \
           -m 256m \
		   consul:latest \
           agent -dev -client 0.0.0.0 -bind=0.0.0.0
}

run-profiler() {
    mkdir -p profile
    docker run -d \
           --name leaky_profiler \
           --link leaky:leaky \
           -v "$(pwd)/profile:/profile" \
           -v "$(pwd)/test.sh:/test.sh" \
           leaky \
           /test.sh do-profile
}

stop() {
    docker rm -f leaky_consul
    docker rm -f leaky
    docker rm -f leaky_profiler
}

# ----------------------------------------

do-profile() {
    while true; do
        now=$(date +%s)
        curl -so "/profile/heap-$now" "http://leaky:6060/debug/pprof/heap?debug=1"
        sleep 60
    done
}

cmd=$1
$cmd
