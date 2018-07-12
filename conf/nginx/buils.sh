sudo apt update
sudo apt install certbot


sudo openssl dhparam -out /etc/ssl/certs/dhparam.pem 2048


mkdir -p /var/lib/letsencrypt/.well-known
chgrp www-data /var/lib/letsencrypt
chmod g+s /var/lib/letsencrypt

sudo systemctl reload nginx


sudo certbot certonly --agree-tos --email admin@example.com --webroot -w /var/lib/letsencrypt/ -d example.com -d www.example.com


sudo systemctl reload nginx

0 */12 * * * root test -x /usr/bin/certbot -a \! -d /run/systemd/system && perl -e 'sleep int(rand(3600))' && certbot -q renew --renew-hook "systemctl reload nginx"

43 6 * * * certbot renew --post-hook "systemctl reload nginx"