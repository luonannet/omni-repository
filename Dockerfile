FROM golang:alpine3.13 as builder 
LABEL maintainer="tommylike<tommylikehu@gmail.com>"
ARG GIT_TAG
ARG GIT_COMMIT
ARG RELEASED_AT
WORKDIR /app
COPY . /app
RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags "-X main.Tag=$GIT_TAG -X main.CommitID=$GIT_COMMIT -X main.ReleaseAt=$RELEASED_AT" -o omni-repository

FROM alpine/git:v2.30.2
ARG user=app
ARG group=app
ARG home=/app
RUN addgroup -S ${group} && adduser -S ${user} -G ${group} -h ${home}
RUN apk update --no-cache && apk add --no-cache coreutils=8.32-r2

USER ${user}
WORKDIR ${home}
COPY --chown=${user} --from=builder /app/omni-repository .
COPY --chown=${user} ./config/prod.app.toml ./config/app.toml
# to fix the directory permission issue
RUN mkdir -p ${home}/logs $$ -p ${home}/data
VOLUME ["${home}/logs","${home}/data"]
ENV APP_ENV="prod" APP_PORT=8080
EXPOSE 8080
ENTRYPOINT ["/app/omni-repository"]
