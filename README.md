# AccountEd

### Overview

AccountEd (AE) is a platform-agnostic data processing framework consisting of a learning management system and educator tools. This framework is integrated with blockchain technologies that are used to validate data integrity at any point in time.

### System Requirements

This system requires an environment with Docker (or Docker compatible environment), [NodeJS][nodejs-download] & [Go][golang-download] installed.

### Installation

`Local Development`: 

- [Docker Desktop][docker-desktop] or any other means of ensuring `docker compose v2` is running on your system
- `Mac`: Add the following to your `/etc/hosts` file to connect to the local MongDB replica set from a GUI such as MongoDB Compass
```bash
# /etc/hosts
# local mongo replica
127.0.0.1 db
127.0.0.1 db1
127.0.0.1 db2
```
- Open a terminal and run `make init-local-environment` from the root of the project

### Getting Started

### Architecture

### Contributing 

Anyone is encouraged to contribute to the repository by forking and submitting a pull request. (If you are new to GitHub, you might start with a basic tutorial.) By contributing to this project, you grant a world-wide, royalty-free, perpetual, irrevocable, non-exclusive, transferable license to all users under the terms of the Apache Software License v2 or later.


[docker-desktop]: https://www.docker.com/products/docker-desktop/
[golang-download]: https://go.dev/dl/
[nodejs-download]: https://nodejs.org/en/download/