FROM golang:latest

# install openssl tool and rsync
RUN apt-get update \
    && apt-get install  -yq openssl rsync \
    && rm -rf /var/lib/apt/lists/*

COPY ./ /go/src/github.com/github.com/xiaofengshuyu/vpn-manager/manage/
# COPY conf.yml /var/run/conf.yml

# need to set
#ENV key=value

COPY docker-entrypoint.sh /

EXPOSE 8086
EXPOSE 8087

EXPOSE 873

CMD ["bash","/docker-entrypoint.sh"]