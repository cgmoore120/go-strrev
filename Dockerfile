FROM scratch

MAINTAINER Chris Moore <chris@cloudspace.com>

WORKDIR /
ADD go-strrev go-strrev

ADD ./microservice.yml microservice.yml 
