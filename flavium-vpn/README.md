## Flavium VPN Proxy
Acts as a VPN proxy for other containers in network.

### Setup
0. Enable ipv6 for docker according to instructions at: https://docs.docker.com/config/daemon/ipv6/
1. You should already have a tun interface. Otherwise run `mkdir -p /dev/net; mknod /dev/net/tun c 10 200; chmod 600 /dev/net/tun; modprobe tun`
2. Download config files from your favourite VPN provider. We are using https://mullvad.net/en/download/config/?platform=linux for VPN.
3. Put the config files (mullvad_xx.conf, mullvad_userpass.txt, mullvad_ca.txt) in a folder called "vpn" in this directory.
4. Run `docker-compose build --no-cache` and `docker-compose up flavium-vpn`

If everything went as expected you should see the text "You are connected to 
Mullvad VPN" or similar.
