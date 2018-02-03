FROM node:6

COPY repo /opt/site
WORKDIR /opt/site
RUN npm install
RUN $(npm bin)/gulp

FROM httpd:2.4
RUN rm -rf /usr/local/apache2/htdocs
COPY --from=0 /opt/site/dist /opt/site/dist
COPY copy_artifacts.sh /opt/util/copy_artifacts.sh
RUN ln -s /opt/site/dist /usr/local/apache2/htdocs
