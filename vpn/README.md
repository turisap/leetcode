### VPS

* ru per.cloud
* Estonia weasel
* ```
  ssh-copy-id root@<RU_IP>
  ssh-copy-id root@<EU_IP>
  ```
* very cool

```
Host ru
    HostName <RU_IP>
    User root
    IdentityFile ~/.ssh/id_ed25519

Host eu
    HostName <EU_IP>
    User root
    IdentityFile ~/.ssh/id_ed25519
```

* no password auth

```
cat > /etc/ssh/sshd_config.d/00-hardening.conf <<'EOF'
PasswordAuthentication no
KbdInteractiveAuthentication no
PermitRootLogin prohibit-password
EOF
sshd -t && systemctl restart ssh
```

* firewall

```shell
ufw default deny incoming     # block all incoming connections...
ufw default allow outgoing    # ...but the server can connect anywhere
ufw allow 22/tcp              # except SSH
ufw --force enable 
```

* BPR

```
cat > /etc/sysctl.d/99-bbr.conf <<'EOF'
net.core.default_qdisc=fq
net.ipv4.tcp_congestion_control=bbr
EOF
sysctl --system >/dev/null
sysctl net.ipv4.tcp_congestion_control net.core.default_qdisc
```

### Exit node

```shell
EXIT_DOMAIN=<exit-node>.duckdns.org

mkdir -p /var/www/site
cat > /var/www/site/index.html <<'EOF'
<!doctype html><html><head><meta charset="utf-8"><title>Notes</title></head>
<body><h1>Field notes</h1><p>Photos and notes from trips. Updated occasionally.</p></body></html>
EOF

# port 80: ACME challenge + redirect
cat > /etc/nginx/sites-available/default <<EOF
server {
    listen 80 default_server;
    server_name ${EXIT_DOMAIN};
    location /.well-known/acme-challenge/ { root /var/www/html; }
    location / { return 301 https://\$host\$request_uri; }
}
EOF
nginx -t && systemctl reload nginx

certbot certonly --webroot -w /var/www/html -d "$EXIT_DOMAIN" \
  --agree-tos --register-unsafely-without-email \
  --deploy-hook "systemctl reload nginx"
```

```
cat >> /etc/nginx/sites-available/default <<EOF

server {
    listen 127.0.0.1:8443 ssl http2;
    server_name ${EXIT_DOMAIN};
    ssl_certificate     /etc/letsencrypt/live/${EXIT_DOMAIN}/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/${EXIT_DOMAIN}/privkey.pem;
    ssl_protocols TLSv1.2 TLSv1.3;
    root /var/www/site;
    index index.html;
}
EOF
nginx -t && systemctl reload nginx
curl -sk --resolve ${EXIT_DOMAIN}:8443:127.0.0.1 https://${EXIT_DOMAIN}:8443 | head -3
```

```shell
Xray v26.11.9
```


### Bridge node
```shell
BRIDGE_DOMAIN=<bridge_node>.org
echo "$BRIDGE_DOMAIN"        # check the spelling this time 🙂

mkdir -p /var/www/site
cat > /var/www/site/index.html <<'EOF'
<!doctype html><html lang="ru"><head><meta charset="utf-8"><title>Заметки</title></head>
<body><h1>Заметки</h1><p>Личный блог: книги, походы, рецепты. Обновляется нерегулярно.</p></body></html>
EOF

cat > /etc/nginx/sites-available/default <<EOF
server {
    listen 80 default_server;
    server_name ${BRIDGE_DOMAIN};
    location /.well-known/acme-challenge/ { root /var/www/html; }
    location / { return 301 https://\$host\$request_uri; }
}
EOF
nginx -t && systemctl reload nginx

certbot certonly --webroot -w /var/www/html -d "$BRIDGE_DOMAIN" \
  --agree-tos --register-unsafely-without-email \
  --deploy-hook "systemctl reload nginx"

cat >> /etc/nginx/sites-available/default <<EOF

server {
    listen 127.0.0.1:8443 ssl http2;
    server_name ${BRIDGE_DOMAIN};
    ssl_certificate     /etc/letsencrypt/live/${BRIDGE_DOMAIN}/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/${BRIDGE_DOMAIN}/privkey.pem;
    ssl_protocols TLSv1.2 TLSv1.3;
    root /var/www/site;
    index index.html;
}
EOF
nginx -t && systemctl reload nginx
curl -sk --resolve ${BRIDGE_DOMAIN}:8443:127.0.0.1 https://${BRIDGE_DOMAIN}:8443 | head -3
```

* geofiles
```shell
cat > /usr/local/bin/update-geo.sh <<'EOF'
#!/bin/sh
set -e
cd /usr/local/share/xray
BASE=https://github.com/runetfreedom/russia-v2ray-rules-dat/releases/latest/download
curl -fsSL -o geoip_RU.dat.new   "$BASE/geoip.dat"
curl -fsSL -o geosite_RU.dat.new "$BASE/geosite.dat"
mv geoip_RU.dat.new   geoip_RU.dat
mv geosite_RU.dat.new geosite_RU.dat
CFG=/usr/local/etc/xray/config.json
if [ -f "$CFG" ]; then
  xray run -test -c "$CFG" >/dev/null && systemctl restart xray
fi
EOF
chmod +x /usr/local/bin/update-geo.sh
/usr/local/bin/update-geo.sh
ls -lh /usr/local/share/xray/
echo "17 4 * * * root /usr/local/bin/update-geo.sh" > /etc/cron.d/xray-geo
```

* bridge secrets
```shell
# values from the exit (/root/exit-secrets.txt on the EU server)
EXIT_DOMAIN=<exit-node>.duckdns.org
UUID_B=0fa216a3-9960-4b54-abe8-ddaa0ec0a053
PRIV_EXIT=iEbprczlQVXdXRaVUb1FORrHjUjNgrCPVMgSnI3OY0U
PUB_EXIT=4n-Fup3l0Eun81Auc8KKv19ujtwjf2AJeyVxZCy-gRc
SID_EXIT=2aee771a52007866
PATH_EXIT=/54a5cc4bc251

# the bridge's own values
UUID_C=$(xray uuid)
KEYS=$(xray x25519)
PRIV_BRIDGE=$(echo "$KEYS" | awk -F': *' 'tolower($1) ~ /private/ {print $2}')
PUB_BRIDGE=$(echo "$KEYS"  | awk -F': *' 'tolower($1) ~ /public|password/ {print $2}')
SID_BRIDGE=$(openssl rand -hex 8)
PATH_BRIDGE="/$(openssl rand -hex 6)"

cat > /root/bridge-secrets.txt <<EOF
BRIDGE_DOMAIN=$BRIDGE_DOMAIN
UUID_C=$UUID_C
PRIV_BRIDGE=$PRIV_BRIDGE
PUB_BRIDGE=$PUB_BRIDGE
SID_BRIDGE=$SID_BRIDGE
PATH_BRIDGE=$PATH_BRIDGE
EXIT_DOMAIN=$EXIT_DOMAIN
UUID_B=$UUID_B
PUB_EXIT=$PUB_EXIT
SID_EXIT=$SID_EXIT
PATH_EXIT=$PATH_EXIT
EOF
chmod 600 /root/bridge-secrets.txt
cat /root/bridge-secrets.txt
```

* bridge  Xray
```shell
cat > /usr/local/etc/xray/config.json <<EOF
{
  "log": { "loglevel": "warning" },
  "dns": {
    "tag": "dns-internal",
    "queryStrategy": "UseIPv4",
    "servers": [
      {
        "address": "77.88.8.8",
        "domains": ["domain:ru", "domain:su", "domain:xn--p1ai",
                    "ext:geosite_RU.dat:ru-available-only-inside"],
        "skipFallback": true
      },
      "1.1.1.1"
    ]
  },
  "inbounds": [{
    "tag": "clients",
    "listen": "0.0.0.0",
    "port": 443,
    "protocol": "vless",
    "settings": {
      "clients": [{ "id": "${UUID_C}", "email": "kirill" }],
      "decryption": "none"
    },
    "streamSettings": {
      "network": "xhttp",
      "xhttpSettings": { "path": "${PATH_BRIDGE}", "mode": "auto" },
      "security": "reality",
      "realitySettings": {
        "target": "127.0.0.1:8443",
        "serverNames": ["${BRIDGE_DOMAIN}"],
        "privateKey": "${PRIV_BRIDGE}",
        "shortIds": ["${SID_BRIDGE}"]
      }
    },
    "sniffing": { "enabled": true, "destOverride": ["http", "tls", "quic"], "routeOnly": true }
  }],
  "outbounds": [
    {
      "tag": "eu",
      "protocol": "vless",
      "settings": {
        "vnext": [{
          "address": "${EXIT_DOMAIN}",
          "port": 443,
          "users": [{ "id": "${UUID_B}", "encryption": "none" }]
        }]
      },
      "streamSettings": {
        "network": "xhttp",
        "xhttpSettings": { "path": "${PATH_EXIT}", "mode": "auto" },
        "security": "reality",
        "realitySettings": {
          "serverName": "${EXIT_DOMAIN}",
          "fingerprint": "chrome",
          "publicKey": "${PUB_EXIT}",
          "shortId": "${SID_EXIT}"
        }
      }
    },
    { "tag": "direct", "protocol": "freedom", "settings": { "domainStrategy": "UseIPv4" } },
    { "tag": "block",  "protocol": "blackhole" }
  ],
  "routing": {
    "domainStrategy": "IPIfNonMatch",
    "rules": [
      { "ip": ["77.88.8.8"], "outboundTag": "direct" },
      { "inboundTag": ["dns-internal"], "outboundTag": "eu" },

      { "ip": ["geoip:private"], "outboundTag": "block" },
      { "protocol": ["bittorrent"], "outboundTag": "block" },

      { "domain": ["ext:geosite_RU.dat:ru-blocked"], "outboundTag": "eu" },
      { "ip": ["ext:geoip_RU.dat:ru-blocked"], "outboundTag": "eu" },

      { "domain": ["domain:ru", "domain:su", "domain:xn--p1ai",
                   "ext:geosite_RU.dat:ru-available-only-inside"], "outboundTag": "direct" },
      { "ip": ["geoip:ru"], "outboundTag": "direct" },

      { "network": "tcp,udp", "outboundTag": "eu" }
    ]
  }
}
EOF
xray run -test -c /usr/local/etc/xray/config.json && systemctl restart xray && systemctl enable xray
systemctl status xray --no-pager | head -5
ss -tlnp | grep -E ':443|:8443'
journalctl -u xray -n 20 --no-pager
```


```shell
EXIT_DOMAIN=<exit-node>.duckdns.org

rm -rf /var/www/site && mkdir -p /var/www/site
tar -xzf /tmp/decoy-exit.tar.gz -C /var/www/site
chown -R www-data:www-data /var/www/site

echo 'server_tokens off;' > /etc/nginx/conf.d/hardening.conf

cat > /etc/nginx/sites-available/default <<EOF
server {
listen 80 default_server;
server_name ${EXIT_DOMAIN};
location /.well-known/acme-challenge/ { root /var/www/html; }
location / { return 301 https://\$host\$request_uri; }
}

server {
listen 127.0.0.1:8443 ssl http2;
server_name ${EXIT_DOMAIN};
ssl_certificate     /etc/letsencrypt/live/${EXIT_DOMAIN}/fullchain.pem;
ssl_certificate_key /etc/letsencrypt/live/${EXIT_DOMAIN}/privkey.pem;
ssl_protocols TLSv1.2 TLSv1.3;

    root /var/www/site;
    index index.html;
    charset utf-8;
    error_page 404 /404.html;

    location / { try_files \$uri \$uri/ =404; }
    location ~* \.(css|svg)$ { expires 7d; add_header Cache-Control "public"; }
}
EOF
nginx -t && systemctl reload nginx
```