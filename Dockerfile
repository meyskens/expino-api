ARG arch
FROM multiarch/alpine:${arch}-edge

COPY ./api /api

CMD /api