FROM robertasolimandonofreo/debian_cherokee:latest 

ENV                 GRAFANA_VERSION=8.2.5
WORKDIR /root

RUN apt-get -y update \
    && apt-get -y dist-upgrade \
    && apt-get -y --force-yes install \
    && wget https://dl.grafana.com/oss/release/grafana_${GRAFANA_VERSION}_amd64.deb \
    && dpkg -i grafana_${GRAFANA_VERSION}_amd64.deb \
    && rm grafana_${GRAFANA_VERSION}_amd64.deb \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

EXPOSE 3000