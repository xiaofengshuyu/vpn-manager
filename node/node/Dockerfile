FROM hwdsl2/ipsec-vpn-server

RUN apt-get update \
    && apt-get install  -yq rsync cron \
    && rm -rf /var/lib/apt/lists/*

COPY ./docker-entrypoint.sh /docker-entrypoint.sh

CMD ["bash","/docker-entrypoint.sh"]