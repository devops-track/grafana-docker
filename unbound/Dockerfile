FROM rust:1-buster AS build

ARG VERSION=master

RUN \
    mkdir unbound-telemetry && cd unbound-telemetry && \
    curl -sL https://github.com/svartalf/unbound-telemetry/archive/$VERSION.tar.gz | tar xzf - --strip-components=1 && \
    cargo build -j 1 --release && \
    strip -s target/release/unbound-telemetry

FROM debian:buster-slim

LABEL org.opencontainers.image.authors "Richard Kojedzinszky <richard@kojedz.in>"
LABEL org.opencontainers.image.source https://github.com/kubernetize/unbound-telemetry

RUN apt-get update && \
    apt-get install -y libssl1.1 && \
    apt-get clean

COPY --from=build unbound-telemetry/target/release/unbound-telemetry /usr/local/bin

USER 15353

EXPOSE 9167/tcp

CMD ["unbound-telemetry", "tcp"]