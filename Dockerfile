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


ADD conf/provisioning/datasources/conf.yaml /etc/grafana/provisioning/datasources
ADD conf/provisioning/dashboards/conf.yaml /etc/grafana/provisioning/dashboards
ADD conf/provisioning/notifiers/conf.yaml /etc/grafana/provisioning/notifiers
ADD conf/provisioning/plugins/conf.yaml /etc/grafana/provisioning/plugins
ADD conf/provisioning/grafana.ini /etc/grafana/
ADD conf/provisioning/ldap.toml /etc/grafana/

COPY ./dashboards/Disponibilidade /var/lib/grafana/dashboards/Disponibilidade
COPY ./dashboards/DEV01 /var/lib/grafana/dashboards/DEV01
COPY ./dashboards/Livelo /var/lib/grafana/dashboards/Livelo
COPY ./dashboards/SOC /var/lib/grafana/dashboards/SOC
COPY ./dashboards/Geral /var/lib/grafana/dashboards/Geral
COPY ./dashboards/DevOps /var/lib/grafana/dashboards/DevOps


EXPOSE 3000