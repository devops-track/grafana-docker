FROM prom/prometheus:latest

VOLUME /etc/prometheus
VOLUME /var/lib/prometheus
ADD prometheus.yml /etc/prometheus/
ADD rules.yml /etc/prometheus/

EXPOSE 9090