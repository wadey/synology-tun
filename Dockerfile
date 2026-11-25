# syntax=docker/dockerfile:1

# Usage:
#   docker run --cap-add CAP_SYS_MODULE -v /lib/modules:/lib/modules:ro wadey/synology-tun

ARG BASE_IMAGE=busybox:latest

FROM ${BASE_IMAGE}

CMD ["/bin/sh", "-c", "grep '^tun ' \"${PROC_MODULES:-/proc/modules}\" || insmod /lib/modules/tun.ko"]
