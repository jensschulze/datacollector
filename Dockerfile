ARG BASE_IMAGE_BUILD=golang:1.17-alpine
ARG BASE_IMAGE_RELEASE=alpine:3.14

FROM ${BASE_IMAGE_BUILD} AS build-env

WORKDIR /src

COPY go.* .
RUN go mod download

COPY . .
RUN go build -o datacollector

FROM ${BASE_IMAGE_RELEASE}

LABEL org.opencontainers.image.authors="Jens Schulze"

COPY --from=build-env /src/datacollector /usr/local/bin/

CMD ["/usr/local/bin/datacollector"]

