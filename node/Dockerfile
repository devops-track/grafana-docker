FROM robertasolimandonofreo/debian_cherokee:latest 
ENV node_exporter_version=1.3.0 
ADD node_exporter.service etc/systemd/system/
RUN apt-get update \
    && wget https://github.com/prometheus/node_exporter/releases/download/v${node_exporter_version}/node_exporter-${node_exporter_version}.linux-amd64.tar.gz \
    && tar -xzvf node_exporter-${node_exporter_version}.linux-amd64.tar.gz \
    && cp /node_exporter-1.3.0.linux-amd64/node_exporter /usr/local/bin/ \
    && useradd --no-create-home --shell /bin/false node_exporter \
    && chown node_exporter:node_exporter /usr/local/bin/node_exporter 
ENTRYPOINT /usr/local/bin/node_exporter
EXPOSE 9100