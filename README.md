# Docker Machine LXD Driver


## System requirements
* OS - Ubuntu 20.04
* LXD storage driver - btrfs


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
$ docker-machine create -d lxd docker1
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


## Using remote LXD
Docker Machine cannot connect SSH in a remote LXD bridge network (eg `lxdbr0`).

You can specify the public network adapter for the LXD host by using the `--lxd-external-network` command line argument. The network is exposed using `macvlan`.

First, set up to connect to the remote LXD.

```
$ lxc remote add remotelxd 192.168.1.120
$ lxc remote set-default remotelxd
```

Run `docker-machine create` with the remote LXD public network adapter.

```
$ docker-machine create -d lxd --lxd-external-network --lxd-external-network eth0 docker1
```


## LICENSE
[MIT LICENSE](./LICENSE.md)