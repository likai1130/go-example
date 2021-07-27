FROM alpine:3.9
COPY ./go-example .
COPY ./application.yaml .
COPY ./tomls ./tomls
EXPOSE 8188
ENTRYPOINT ["/go-example"]
