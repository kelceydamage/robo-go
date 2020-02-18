FROM alpine:3.7

COPY {{InputBuildArtifact Product "darwin-amd64"}} /usr/local/bin/

ENTRYPOINT [ "/usr/local/bin/robo-go" ]
