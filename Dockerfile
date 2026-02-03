FROM alpine:latest

ENV KAGE_DOWNLOADER_PATH=/downloads
ENV KAGE_USER=abc
ENV KAGE_UID=1000
ENV KAGE_GID=1000

WORKDIR "/config"
RUN mkdir -p "${KAGE_DOWNLOADER_PATH}" && addgroup -g "${KAGE_GID}" "${KAGE_USER}" && adduser \
    --disabled-password \
    --gecos "" \
    --home "$(pwd)" \
    --ingroup "${KAGE_USER}" \
    --no-create-home \
    --uid "${KAGE_UID}" \
    "${KAGE_USER}" && \
    chown abc:abc /config "${KAGE_DOWNLOADER_PATH}"

COPY kage /usr/local/bin/kage
RUN chmod +x /usr/local/bin/kage
USER "${KAGE_USER}"
ENTRYPOINT ["/usr/local/bin/kage"]
