FROM ubuntu:18.04

RUN apt-get update \
 && apt-get install -y openvpn \
 && apt-get install -y curl

# vpn folder should contain vpn.conf, userpass.txt, and vpn.crt files.
COPY ./vpn/* /etc/openvpn/

WORKDIR /var/flavium/
COPY ./startup.sh .

RUN chmod 755 startup.sh
CMD /bin/sh startup.sh
# At this point we should start the proxy but that is a problem for another day
