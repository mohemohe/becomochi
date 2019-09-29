FROM golang:alpine as builder

ARG GOLANG_NAMESPACE="github.com/mohemohe/becomochi"
ENV GOLANG_NAMESPACE="$GOLANG_NAMESPACE"
ARG DRONE_BRANCH=""
ENV DRONE_BRANCH="$DRONE_BRANCH"
ARG DRONE_COMMIT_SHA=""
ENV DRONE_COMMIT_SHA="$DRONE_COMMIT_SHA"

RUN apk --no-cache add alpine-sdk coreutils git tzdata upx util-linux zsh
SHELL ["/bin/zsh", "-c"]
RUN cp -f /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
RUN go get -u -v github.com/pwaller/goupx
ADD ./.drone/version.sh /tmp/
RUN chmod +x /tmp/version.sh
WORKDIR /go/src/$GOLANG_NAMESPACE/server
ADD ./server/go.mod /go/src/$GOLANG_NAMESPACE/server/
ADD ./server/go.sum /go/src/$GOLANG_NAMESPACE/server/
ENV GO111MODULE=on
RUN go mod download
ADD . /go/src/$GOLANG_NAMESPACE/
RUN go build -ldflags "\
      -w \
      -s \
      -X '${GOLANG_NAMESPACE}/util.version=$(/tmp/version.sh)' \
      -X '${GOLANG_NAMESPACE}/util.branch=${DRONE_BRANCH}' \
      -X '${GOLANG_NAMESPACE}/util.hash=${DRONE_COMMIT_SHA:0:8}' \
    " -o /becomochi/app
RUN goupx /becomochi/app

# ====================================================================================

FROM alpine

ARG GOLANG_NAMESPACE="github.com/mohemohe/becomochi"
ENV GOLANG_NAMESPACE="$GOLANG_NAMESPACE"

RUN apk --no-cache add ca-certificates
COPY --from=builder /etc/localtime /etc/localtime
COPY --from=builder /becomochi/app /becomochi/app

EXPOSE 1323
WORKDIR /becomochi
CMD ["/becomochi/app"]