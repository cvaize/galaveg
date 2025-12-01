FROM rustlang/rust:nightly as build_rust

RUN cargo install watchexec-cli

FROM golang:1.25.3 as build

ARG WWWGROUP=1000
ARG WWWUSER=1000

RUN mkdir -p /app
WORKDIR /app
COPY --from=build_rust "/usr/local/cargo/bin/watchexec" /bin/watchexec

RUN #userdel -r ubuntu
RUN groupadd --force -g $WWWGROUP web
RUN useradd -ms /bin/bash --no-user-group -g $WWWGROUP -u $WWWUSER web
USER web

CMD ["sleep", "infinity"]
#CMD ["watchexec", "-r", "-e", "go,gohtml,html", "go", "run", ".", "serve"]
