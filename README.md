synology-tun
============

This small Docker image allows you to load the tun kernel module on Synology
hosts without having to set up manual scripts or scheduled tasks.

You can use this in docker compose to have a service that depends on the tun
module being loaded.

Example
-------

```yaml
version: '3.8'
services:
    synology-tun:
        image: wadey/synology-tun:latest
        volumes:
            - /lib/modules:/lib/modules:ro
        cap_add:
            - CAP_SYS_MODULE
        network_mode: none
    dnclient:
        depends_on:
            synology-tun:
                condition: service_completed_successfully
        image: definednet/dnclient:latest
        volumes:
            - /volume1/docker/dnclient/etc:/etc/defined
        cap_add:
            - NET_ADMIN
        network_mode: host
        restart: always
```
