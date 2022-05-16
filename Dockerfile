FROM golang:alpine3.13 as builder

LABEL maintainer="tommylike<tommylikehu@gmail.com>"
WORKDIR /app
COPY . /app
RUN go mod download
RUN CGO_ENABLED=0 go build -o omni-repository
RUN ls /app
FROM alpine/git:v2.30.2
ARG user=root
ARG group=root
ARG home=/app
# RUN addgroup -S ${group} && adduser -S ${user} -G ${group} -h ${home}

USER ${user}
WORKDIR ${home}
RUN ls /app
COPY --chown=${user} --from=builder /app/omni-repository /app
RUN ls /app
COPY --chown=${user} ./config/prod.app.toml ./config/app.toml
RUN ls /app
COPY --chown=${user} ./config/prod.env ./config/.env
RUN ls /app
# to fix the directory permission issue
RUN mkdir -p ${home}/logs $$ -p ${home}/data
VOLUME ["${home}/logs","${home}/data"]
RUN ls /app
ENV PATH="${home}:${PATH}"
ENV APP_ENV="prod"
EXPOSE 8080
ENTRYPOINT ["/app/omni-repository"]
