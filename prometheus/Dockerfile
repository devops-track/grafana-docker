FROM robertasolimandonofreo/debian_cherokee:latest 
ENV prometheus_version=2.31.1
RUN adduser -s /bin/false -D -H prometheus \
    && adduser -s /bin/false -D -H node_exporter \
    && apt-get update \
    && wget https://github.com/prometheus/prometheus/releases/download/v${prometheus_version}/prometheus-${prometheus_version}.linux-amd64.tar.gz \
    && tar -xvzf prometheus-${prometheus_version}.linux-amd64.tar.gz \
    && mkdir -p /etc/prometheus /var/lib/prometheus \
    && cp prometheus-${prometheus_version}.linux-amd64/promtool /usr/local/bin/ \
    && cp prometheus-${prometheus_version}.linux-amd64/prometheus /usr/local/bin/ \
    && cp -R prometheus-${prometheus_version}.linux-amd64/console_libraries/ /etc/prometheus/ \
    && cp -R prometheus-${prometheus_version}.linux-amd64/consoles/ /etc/prometheus/ \
    && rm -rf prometheus-${prometheus_version}.linux-amd64*

VOLUME /etc/prometheus
VOLUME /var/lib/prometheus
ADD prometheus.yml /etc/prometheus/
ADD alert.rules /etc/prometheus/


ENTRYPOINT /usr/local/bin/prometheus \ 
            --config.file /etc/prometheus/prometheus.yml \ 
            --storage.tsdb.path /var/lib/prometheus/ \
            --web.console.libraries=/usr/share/prometheus/console_libraries \
            --web.console.templates=/usr/share/prometheus/consoles

EXPOSE 9090