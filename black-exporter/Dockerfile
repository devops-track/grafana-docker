FROM robertasolimandonofreo/debian_cherokee:latest 

ENV blackbox_exporter_version=0.19.0
RUN apt-get update \
    && wget https://github.com/prometheus/blackbox_exporter/releases/download/v${blackbox_exporter_version}/blackbox_exporter-${blackbox_exporter_version}.linux-amd64.tar.gz \
    && tar -xvzf /blackbox_exporter-0.19.0.linux-amd64.tar.gz \
    && mv blackbox_exporter-0.19.0.linux-amd64/blackbox_exporter /usr/local/bin \
    && mkdir -p /etc/blackbox \
    && mv blackbox_exporter-0.19.0.linux-amd64/blackbox.yml /etc/blackbox \
    && useradd -rs /bin/false blackbox \
    && chown blackbox:blackbox /usr/local/bin/blackbox_exporter \
    && chown -R blackbox:blackbox /etc/blackbox/* \
    &&  rm -rf blackbox_exporter-0.19.0.linux-amd64 blackbox_exporter-0.19.0.linux-amd64.tar.gz
ADD blackbox.service /lib/systemd/system

ENTRYPOINT /usr/local/bin/blackbox_exporter --config.file /etc/blackbox/blackbox.yml

EXPOSE 9115