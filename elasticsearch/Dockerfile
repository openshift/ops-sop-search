FROM elasticsearch:7.7.0

RUN echo -e "http.cors.enabled: true"  >> /usr/share/elasticsearch/config/elasticsearch.yml
RUN echo -e "http.cors.allow-origin: \"*\"" >> /usr/share/elasticsearch/config/elasticsearch.yml
RUN echo -e "http.cors.allow-methods: OPTIONS, HEAD, GET, POST, PUT, DELETE" >> /usr/share/elasticsearch/config/elasticsearch.yml
RUN echo -e "http.cors.allow-headers: X-Requested-With,X-Auth-Token,Content-Type, Content-Length" >> /usr/share/elasticsearch/config/elasticsearch.yml

RUN cat /usr/share/elasticsearch/config/elasticsearch.yml