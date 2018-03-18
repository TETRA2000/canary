FROM node:6

COPY repo /opt/site
WORKDIR /opt/site
RUN npm install
RUN npm run build

FROM httpd:2.4-alpine
COPY --from=0 /opt/site/public /usr/local/apache2/htdocs/
