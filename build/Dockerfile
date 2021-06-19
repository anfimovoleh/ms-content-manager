FROM golang:1.16.5 AS build-env
ADD . /src
RUN cd /src && ls && make build

FROM alpine
WORKDIR /app
COPY --from=build-env /src/ms-content-manager /app/