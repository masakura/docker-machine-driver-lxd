# Docker Machine LXD Driver


## Getting started
* [Install LXD](https://linuxcontainers.org/lxd/getting-started-cli/#installation).
* [Install Docker Machine](https://docs.docker.com/machine/install-machine/).

Install Docker Machine LXD Driver.

```
$ sudo wget -O /usr/local/bin/docker-machine-driver-lxd https://gitlab.com/masakura/docker-machine-driver-lxd/-/releases/v0.5.0/downloads/bin/docker-machine-driver-lxd-Linux-x86_64
$ sudo chmod +x /usr/local/bin/docker-machine-driver-lxd
```

Create Docker Machine.

```
$ docker-machine create -d lxd --engine-storage-driver btrfs docker1
```

Use Docker Machine.

```
$ eval $(docker-machine env docker1)
$ docker run --rm hello-world
```

Remove Docker Machine.

```
$ docker-machine rm docker1
```

See [Docker Machine documents](https://docs.docker.com/machine/get-started/) for details on how to use it.


## LICENSE
[MIT LICENSE](./LICENSE.md)