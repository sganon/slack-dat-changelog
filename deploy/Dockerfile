FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
USER 1000:1000
EXPOSE 8080
ENTRYPOINT ["/slack-dat-changelog"]
COPY slack-dat-changelog /
CMD ["--debug", "serve"]
