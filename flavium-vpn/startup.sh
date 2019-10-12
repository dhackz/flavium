#!/bin/sh
cd /etc/openvpn
openvpn mullvad_ch.conf

#service openvpn start           # Start openvpn.
#sleep 10                         # Give time to take effect.
#curl https://am.i.mullvad.net/connected     # Are we connected?
