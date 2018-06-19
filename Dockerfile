FROM alpine:3.4

RUN apk -U add ca-certificates

EXPOSE 8080

ADD jamespants /bin/jamespants
ADD config.yml.dist /etc/jamespants/config.yml

CMD ["jamespants", "-config", "/etc/jamespants/config.yml"]