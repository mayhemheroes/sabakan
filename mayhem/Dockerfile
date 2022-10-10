FROM fuzzers/go-fuzz:1.2.0 as builder

COPY . /sabakan
WORKDIR /sabakan/mayhem

RUN go-fuzz-build -libfuzzer -o fuzz.a && \
    clang -fsanitize=fuzzer fuzz.a -o fuzz.libfuzzer

FROM debian:bullseye-slim
COPY --from=builder /sabakan/mayhem/fuzz.libfuzzer /