FROM scratch

USER 1000:1000

EXPOSE 8080

ENTRYPOINT ["/slack-dat-changelog"]

COPY slack-dat-changelog /

CMD ["--debug", "serve"]
