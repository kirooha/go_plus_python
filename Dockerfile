FROM ubuntu:18.10

WORKDIR /app
ADD hello.py /app
ADD main /app
ENV PYTHONPATH=/app

RUN set -x \
    && apt-get update \
    && apt-get install -y \
                python2.7 \
                python2.7-dev
