FROM anapsix/alpine-java
RUN apk add --update \
    transmission-cli \
    curl

RUN mkdir filebot-portable
WORKDIR filebot-portable

COPY tar.sh /tmp/tar.sh
RUN /tmp/tar.sh

COPY .filebot.sm /tmp/filebot.sm

COPY main /var/www/main
