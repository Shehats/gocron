FROM mcr.microsoft.com/devcontainers/base:bullseye

ENV PATH=/opt/kafka/bin:$PATH

# add kafka cli tools
RUN mkdir -p /opt/kafka
RUN wget --quiet -O - https://dlcdn.apache.org/kafka/3.2.3/kafka_2.13-3.2.3.tgz \
    | tar -xzv --directory /opt/kafka --strip-components=1

# add postgresql apt repo
RUN wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add - && \
    echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list

# add timescaledb apt repo
RUN wget --quiet -O - https://packagecloud.io/timescale/timescaledb/gpgkey | apt-key add - && \
    echo "deb https://packagecloud.io/timescale/timescaledb/debian/ $(lsb_release -c -s) main" > /etc/apt/sources.list.d/timescaledb.list

# add redis apt repo
RUN curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg && \
    echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/redis.list

RUN apt update && export DEBIAN_FRONTEND=noninteractive \
    && apt -y install --no-install-recommends \
        # psql client
        postgresql-client-14 \
        # timescaledb utils
        timescaledb-tools \
        # redis-cli
        redis-tools \
        # java
        default-jdk \
        # pre-commit
        pre-commit python3-dev \
        # fswatch (for autoreloading)
        fswatch \
        # tab completion for git
        bash-completion

# add grpcurl util
RUN ARCH=$(arch | sed s/aarch64/arm64/) \
    && wget --quiet -O - https://github.com/fullstorydev/grpcurl/releases/download/v1.8.6/grpcurl_1.8.6_linux_${ARCH}.tar.gz \
    | tar -xz --directory /usr/local/bin grpcurl

# add bumpversion util
RUN su vscode -c "pip install bump2version"
