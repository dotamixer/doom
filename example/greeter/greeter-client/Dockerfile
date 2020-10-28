FROM alpine:3.12
COPY ./greeter-client /

ENV PORT 9000
EXPOSE $PORT

RUN chmod +x /greeter-client
CMD ["/greeter-client"]