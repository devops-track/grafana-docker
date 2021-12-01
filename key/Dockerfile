FROM jboss/keycloak:10.0.2

ADD https://github.com/aerogear/keycloak-metrics-spi/releases/download/2.0.1/keycloak-metrics-spi-2.0.1.jar \
    /opt/jboss/keycloak/providers/keycloak-metrics-spi-2.0.1.jar

USER root
RUN chown -R jboss:root /opt/jboss/keycloak/providers/ &&\
    chmod +x /opt/jboss/keycloak/providers/keycloak-metrics-spi-2.0.1.jar
USER jboss