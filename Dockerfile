FROM openresty/openresty:alpine

WORKDIR /app

COPY nginx.conf /usr/local/openresty/nginx/conf/nginx.conf
COPY wake-on-lan.lua .
