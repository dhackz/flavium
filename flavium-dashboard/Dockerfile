FROM nginx:1.17.4

RUN mkdir -p /var/www/
COPY build /var/www
COPY nginx /etc/nginx

ENTRYPOINT ["nginx", "-g", "daemon off;"]
