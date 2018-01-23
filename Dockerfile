ARG arch
FROM multiarch/alpine:${arch}-edge

COPY ./expino-api /expino-api

CMD /expino-api