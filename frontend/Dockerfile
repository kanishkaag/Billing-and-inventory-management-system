FROM nginx

COPY ./*.html /var/www/html/app/

COPY *.nginx /etc/nginx/conf.d/

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
