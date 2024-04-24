FROM alpine

ADD ./bin/getip /root/getip

WORKDIR /root

RUN pwd
RUN ls -la /root

ENTRYPOINT ["/bin/sh","-c","/root/getip"]