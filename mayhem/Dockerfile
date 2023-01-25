FROM golang:1.19 as builder

COPY . /sabakan
WORKDIR /sabakan/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN go get github.com/AdaLogics/go-fuzz-headers
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzzIsValidImageID.a && \
    clang -fsanitize=fuzzer fuzzIsValidImageID.a -o fuzzIsValidImageID.libfuzzer

FROM debian:bullseye-slim
COPY --from=builder /sabakan/mayhem/fuzzIsValidImageID.libfuzzer /