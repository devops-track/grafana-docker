FROM grafana/grafana:latest 

RUN grafana-cli plugins install grafana-piechart-panel

ADD conf/provisioning/datasources/conf.yaml /etc/grafana/provisioning/datasources
ADD conf/provisioning/dashboards/conf.yaml /etc/grafana/provisioning/dashboards
ADD conf/provisioning/notifiers/conf.yaml /etc/grafana/provisioning/notifiers
ADD conf/provisioning/plugins/conf.yaml /etc/grafana/provisioning/plugins


COPY ./dashboards/Disponibilidade /var/lib/grafana/dashboards/Disponibilidade
COPY ./dashboards/DEV01 /var/lib/grafana/dashboards/DEV01
COPY ./dashboards/Livelo /var/lib/grafana/dashboards/Livelo
COPY ./dashboards/SOC /var/lib/grafana/dashboards/SOC
COPY ./dashboards/Geral /var/lib/grafana/dashboards/Geral

EXPOSE 3000