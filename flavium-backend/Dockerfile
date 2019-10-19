FROM alpine:3.10
RUN apk add --update \
    transmission-cli \
    curl \
    rsync \
    openjdk8

COPY install_filebot.sh /tmp/install_filebot.sh
RUN /tmp/install_filebot.sh && rm /tmp/install_filebot.sh

COPY main /var/www/main
COPY .filebot.sm /tmp/.filebot.sm
RUN filebot --license /tmp/.filebot.sm
