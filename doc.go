// Copyright 2016 Igor Stepin. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
UDP COM port Sniffer/Proxy is used for debugging communication between apps and devices.

Tool receives and sends UDP traffic logging it to the console in the HEX format.


Usage

For usage information execute `udp-sniffer -h`:
    $ udp-sniffer -h
    Usage of udp-sniffer:
      -localIP string
        	Local IP port to send data from
      -localPort int
        	Local UPD port to send data from (default 11010)
      -receiveIP string
        	Local IP port to listen connection
      -receivePort int
        	Local UPD port to listen connection (default 11000)
      -remoteIP string
        	Remote IP port to send data (default "127.0.0.1")
      -remotePort int
        	Remote UPD port to send data (default 11000)

Here is an example of session:
    $ udp-sniffer
    2016/01/05 03:08:03 Started...
    2016/01/05 03:08:03 udpServerProxyConnection: >> listen on :11000 and send to unknown yet
    2016/01/05 03:08:03 udpClientProxyConnection: << listen on :11010 and send to 192.168.1.34:11000
    2016/01/05 03:08:05 last client address: 192.168.1.35:11000
    2016/01/05 03:08:05 >> 00 00 0E 12
    2016/01/05 03:08:05 << 00 00 11 00
    ^C2016/01/05 03:08:08 Stopped...


Installation

You can download binaries (https://github.com/stepin/udp-sniffer/releases)
compatible with your operation system.

To compile you need the Golang(https://golang.org) (v1.5.1). Run `go build`
command in the project folder.


Status

This project is feature complete. No future development is expected.


Similar apps for COM ports

- RealTerm (http://realterm.sourceforge.net) -- a good app though it does
not support UDP.
- Null-modem emulator (com0com) (http://com0com.sourceforge.net) -- it allows
debugging apps with device emulators.
- Hercules (http://www.hw-group.com/products/hercules/index_en.html) -- it has
a limited UDP support: received data is in the ASCII format only.
*/
package main