FROM nginx:1.18

WORKDIR /library-portal

ADD dist/ /library-portal/

ADD portal-nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80