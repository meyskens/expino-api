ARG arch
FROM multiarch/alpine:${arch}-edge

COPY ./expino-api /expino-api


ENV INFLUXURL=""

CMD /expino-api