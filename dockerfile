FROM golang:1.15.2-alpine3.12 as dev

# installing git
RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN go get github.com/sirupsen/logrus
RUN go get github.com/gin-gonic/gin
RUN go get github.com/google/uuid
RUN go get github.com/drhodes/golorem
RUN go get k8s.io/client-go/rest
RUN go get k8s.io/client-go/tools/clientcmd
RUN go get k8s.io/apimachinery/pkg/apis/meta/v1
RUN go get k8s.io/client-go/kubernetes

WORKDIR /work
COPY . /work/
RUN go build -o app
###########START NEW IMAGE###################

FROM alpine:3.12 as prod
EXPOSE 8080/tcp
COPY --from=dev /work/app /
CMD ./app
