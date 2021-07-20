FROM quay.io/cdis/golang:1.14-alpine as build-deps

RUN apk update && apk add --no-cache git ca-certificates gcc musl-dev
RUN apk update && apk add bash procps

ENV gen3=/go/src/github.com/uc-cdis/gen3-client
WORKDIR ${gen3}

COPY . .

RUN go mod download

# Populate git version info into the code
RUN printf "package g3cmd\n\nconst (" >gen3-client/g3cmd/gitversion.go \
    && COMMIT=`git rev-parse HEAD` && echo "    gitcommit=\"${COMMIT}\"" >>gen3-client/g3cmd/gitversion.go \
    && VERSION=`git describe --always --tags` && echo "    gitversion=\"${VERSION}\"" >>gen3-client/g3cmd/gitversion.go \
    && echo ")" >>gen3-client/g3cmd/gitversion.go

#RUN go test -v github.com/uc-cdis/gen3-client/tests

RUN go build -ldflags "-linkmode external -extldflags -static" -o bin/gen3-client
WORKDIR ${gen3}/bin
ENV PATH=${PATH}:${gen3}/bin
