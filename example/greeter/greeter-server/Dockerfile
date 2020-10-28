FROM alpine:3.12
COPY ./greeter-server /

ENV PORT 9000
EXPOSE $PORT

RUN chmod +x /greeter-server
CMD ["/greeter-server"]