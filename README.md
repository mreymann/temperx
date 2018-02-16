# temperx :bar_chart:

Munin plugin that monitors temperature and humidity as measured by the TEMPerHUM/TEMPerX USB dongle (413d:2107)

## Info ##

This will only work for devices with ID 413d:2107. Mine identifies itself as "TEMPerX_V3.1".

## Prerequesites ##

This plugin relies entirely on the hid-query binary provided by https://github.com/edorfaus/TEMPered .
Make sure the binary uses libusb not hidraw. On Ubuntu 16.04 I did the following:

* apt install libhidapi-dev
* edit CMakeLists.txt:
```
-       find_library(HIDAPI_LIB NAMES hidapi-hidraw hidapi-libusb
+       find_library(HIDAPI_LIB NAMES hidapi-libusb
```
* continue with compilation

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

![Munin Example](https://github.com/mreymann/temperx/blob/master/example.png)

## Troubleshooting

My dongle reports two USB paths:
```
$ hid-query -e
0002:0002:00 : 413d:2107 interface 0 : (null) (null)
0002:0002:01 : 413d:2107 interface 1 : (null) (null)
```
I had to use the path ending with "01". To try the "00" path, change the regex in temperx like this:
```
-       preg_match_all( '|(.*?01) : 413d:2107.*|', $raw, $matches );
+       preg_match_all( '|(.*?00) : 413d:2107.*|', $raw, $matches );
```
