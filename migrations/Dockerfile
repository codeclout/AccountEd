FROM node:18-bullseye

ENV NODE_ENV=production

RUN set -ex \
    && apt update \
    && apt upgrade -y \
    && apt full-upgrade -y \
    && apt --purge autoremove \
    && adduser --disabled-password --gecos "" ci-svc-accountEd --force-badname \
    && usermod -aG users ci-svc-accountEd \
    && echo "ci-svc-accountEd ALL=(ALL:ALL) NOPASSWD:ALL" >> /etc/sudoers \
    && chmod 0440 /etc/sudoers

RUN su - ci-svc-accountEd

WORKDIR /usr/local/migrate
CMD [ "npx", "migrate-mongo", "up" ]