FROM nginx:stable

COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf

WORKDIR /etc/nginx

EXPOSE 80 443
