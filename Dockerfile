# Usage:
#   docker run --cap-add CAP_SYS_MODULE -v /lib/modules:/lib/modules:ro wadey/synology-tun

FROM alpine:3

RUN apk add --no-cache kmod

CMD grep '^tun ' /proc/modules || insmod /lib/modules/tun.ko
