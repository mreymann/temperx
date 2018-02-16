# temperx :bar_chart:

Munin plugin that monitors temperature and humidity as measured by the TEMPerHUM/TEMPerX USB dongle (413d:2107)

## Info ##

This will only work for devices with ID 413d:2107. Mine identifies itself by "TEMPerX_V3.1".

## Prerequesites ##

This plugin relies entirely on the hid-query binary provided by https://github.com/edorfaus/TEMPered .
Make sure the binary uses libusb not hidraw. On Ubuntu 16.04 I did the following:

* apt install libhidapi-dev
* edit CMakeLists.txt:
```
-       find_library(HIDAPI_LIB NAMES hidapi-hidraw hidapi-libusb
+       find_library(HIDAPI_LIB NAMES hidapi-libusb
```

## Install

* install php-cli
* copy the hid-query binary to /usr/local/bin/
* copy temperx to /etc/munin/plugins
* create file /etc/munin/plugin-conf.d/temperx with this content:
```
[temperx]
user root
```
* restart munin-node

## Example

no screenshot yet

## Troubleshooting


