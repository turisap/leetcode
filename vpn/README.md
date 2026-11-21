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


