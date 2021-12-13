FROM alpine:3.7

EXPOSE 8080

ENV APP_BIN=/bin \
    PATH=${APP_BIN}:$PATH \
    TZ='Asia/Shanghai'

RUN  mkdir -p ${APP_BIN} ${APP_ROOT} \
     && sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
     && apk update \
     && apk upgrade \
     && apk --no-cache add ca-certificates iputils\
     && apk add -U tzdata ttf-dejavu busybox-extras curl bash git\
     && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY ./grpc-logging ${APP_BIN}

CMD ["grpc-logging"]