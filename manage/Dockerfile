FROM ubuntu:14.04

COPY manage /var/run/manage
# COPY conf.yml /var/run/conf.yml

# need to set 
#ENV key=value

COPY docker-entrypoint.sh /

EXPOSE 8086
EXPOSE 8087

CMD ["bash","/docker-entrypoint.sh"]