# FROM golang:latest

# WORKDIR /letargico

# COPY . .

# RUN go build -o /go/bin/letargico server/main.go

# CMD ["webgo"]

# FROM heroku/heroku:18-build as build

# COPY . /app
# WORKDIR /app

# # Setup buildpack
# RUN mkdir -p /tmp/buildpack/heroku/go /tmp/build_cache /tmp/env
# RUN curl https://codon-buildpacks.s3.amazonaws.com/buildpacks/heroku/go.tgz | tar xz -C /tmp/buildpack/heroku/go

# #Execute Buildpack
# RUN STACK=heroku-18 /tmp/buildpack/heroku/go/bin/compile /letargico /tmp/build_cache /tmp/env

# # Prepare final, minimal image
# FROM heroku/heroku:18

# COPY --from=build /letargico /letargico
# ENV HOME /app
# WORKDIR /app
# RUN useradd -m heroku
# USER heroku
# CMD /app/bin/letargico