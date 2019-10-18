FROM anapsix/alpine-java
RUN apk add --update \
    transmission-cli \
    curl \
    rsync

COPY tar.sh /tmp/tar.sh
RUN /tmp/tar.sh

COPY .filebot.sm /tmp/filebot.sm

COPY main /var/www/main
COPY .filebot.sm /tmp/.filebot.sm
RUN filebot --license /tmp/.filebot.sm
