FROM alpine:latest AS final

RUN apk --update add \
    ca-certificates \
    tzdata \
    && update-ca-certificates

ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser
USER appuser

# Copy server binary
COPY --from=build /bin/server /bin/

# ✅ Copy static frontend
COPY ./out /out

EXPOSE 8080

ENTRYPOINT [ "/bin/server" ]
