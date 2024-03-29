FROM rust:1.76.0-slim-bookworm as builder

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y libssl-dev pkg-config

WORKDIR /
RUN USER=root cargo new --bin cactus-keychain-vault-server
WORKDIR /cactus-keychain-vault-server
ADD ./rust/gen/ ./
RUN cargo build --release --example server

FROM ubuntu:24.04
ARG APP=/usr/src/app


RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y ca-certificates tzdata curl tini
RUN rm -rf /var/lib/apt/lists/*


ENV TZ=Etc/UTC \
    APP_USER=appuser

RUN groupadd $APP_USER \
    && useradd -g $APP_USER $APP_USER \
    && mkdir -p ${APP}

COPY --from=builder /cactus-keychain-vault-server/target/release/examples/server ${APP}/cactus-keychain-vault-server

RUN chown -R $APP_USER:$APP_USER ${APP}

USER $APP_USER
WORKDIR ${APP}

COPY healthcheck.sh /

EXPOSE 8080

ENTRYPOINT ["/usr/bin/tini", "--"]
CMD ["./cactus-keychain-vault-server"]

HEALTHCHECK --interval=1s --timeout=5s --start-period=1s --retries=30 CMD /healthcheck.sh

