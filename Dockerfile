FROM registry.docker-cn.com/library/golang:alpine
RUN apk add -U tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" >  /etc/timezone
ADD main /

EXPOSE 8080
CMD /main