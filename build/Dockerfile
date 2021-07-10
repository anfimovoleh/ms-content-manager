FROM golang:1.16.5 AS build-env
ADD . /src
RUN cd /src && ls && GOSUMDB=off go build -o ms-content-manager -v ./cmd/api/

FROM alpine
WORKDIR /app
COPY --from=build-env /src/ms-content-manager /app/