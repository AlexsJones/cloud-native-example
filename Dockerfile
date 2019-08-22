# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git
ADD . /src
RUN cd /src && go build 

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/cloud-native-example /app/goapp
ENTRYPOINT ./goapp
