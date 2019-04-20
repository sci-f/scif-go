FROM iron/go:dev as builder

# multistage build alpine image with the go tools
# docker build -t vanessa/scif-go .

WORKDIR /code
ENV SRC_DIR=/go/src/github.com/sci-f/scif-go
ADD . $SRC_DIR
WORKDIR $SRC_DIR

# Dependencies
RUN make deps && \
    make build && \
    cp bin/scif /usr/local/bin

FROM alpine:3.7
ENV SCIF_SHELL /bin/sh
LABEL Maintainer vsochat@stanford.edu
COPY --from=builder /usr/local/bin/scif /usr/local/bin/scif

ENTRYPOINT ["/usr/local/bin/scif"]
