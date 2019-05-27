# Wake on LAN Container

Containerized server and UI for Wake on LAN.

## Description
Wake on LAN is not possible over most VPN connections. One way to work around this limitation is to run a helper service on the local area network that sends the magic packet on your behalf.

This project seeks to provide the minimal container necessary to enable wake on LAN via a VPN connection.

## Development
### Dependencies
- [`docker`](https://docs.docker.com/install/)
- optional: [GNU `coreutils`](https://www.gnu.org/software/coreutils/coreutils.html) ([`numfmt`](https://www.gnu.org/software/coreutils/manual/html_node/numfmt-invocation.html) is used in `scripts/size`)

### Scripts
The below scripts can be executed by running `./scripts/<name>`, for example `./scripts/start`.
- `build`: build `Dockerfile` into an image tagged `<your username>/wake-on-lan`
- `log`: follow logs from running container
- `size`: log size of image in human-readable format
- `start`: start container in detached mode, must run `build` first
- `stop`: stop container
