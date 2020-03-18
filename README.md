# Open Iot

open iot project for keng !

# Design

api split for 3 moudules, `device` `keng` `management`.

## device

api for device is look like `/ip/update` or `/gpio`.

Just for raspi update the `device info` to open_iot project.

## keng

I would make the role, a room can bind a lot of devices.

the keng would bind to a gpio port, which would show the status of keng switch.

## management

Management only for the admin user to add some thing.


# Dependences

## gin_common

https://github.com/tsbxmw/gin_common

## framework sample

https://github.com/tsbxmw/go_gin_sample

## Device Management

api for management!

## Keng Management

api for front-page!

## Device Api

api for device update!
