FROM alpine:latest

RUN mkdir /app
RUN mkdir /app/templates
RUN mkdir /app/static

ADD templates/* /app/templates/
ADD static/* /app/static/
ADD jamespants /app

WORKDIR /app

CMD ["./jamespants"]

EXPOSE 8080