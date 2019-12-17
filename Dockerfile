FROM golang:1.12-alpine3.9 as build_env
ENV GO111MODULE on
WORKDIR /tmp/applegu_build
COPY . /tmp/applegu_build
RUN apk add upx \
    && export GOPROXY="https://goproxy.cn" \
    && sh /tmp/applegu_build/package.sh

FROM alpine:latest

LABEL repository="https://github.com/afghanistanyn/appLegu"
LABEL homepage="https://github.com/afghanistanyn/appLegu"
LABEL maintainer="afghanistanyn@gmail.com"

WORKDIR /usr/local/applegu

COPY --from=build_env /tmp/applegu_build/applegu.tar.gz /app/

RUN apk --update add openjdk8-jre-base \
    && ls -al /app/ \
    && tar vxzf /app/applegu.tar.gz -C /usr/local/ \
    && chmod u+x /usr/local/applegu/lib/zipalign \
    && rm /app -rf \
    && rm /var/cache/apk/* -rf

CMD [ "/usr/local/applegu/bin/appLegu" ]
