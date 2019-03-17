# UDP COM (serial) port Sniffer/Proxy
[![GitHub release](https://img.shields.io/github/release/stepin/udp-sniffer.svg)](https://github.com/stepin/udp-sniffer/releases) [![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/stepin/udp-sniffer/master/LICENSE) [![GoDoc](https://godoc.org/github.com/stepin/udp-sniffer?status.svg)](https://godoc.org/github.com/stepin/udp-sniffer) [![Go Report Card](https://goreportcard.com/badge/github.com/stepin/udp-sniffer)](https://goreportcard.com/report/stepin/udp-sniffer)

UDP COM port Sniffer/Proxy is used for debugging communication between apps and devices.

Tool receives and sends UDP traffic logging it to the console in the HEX format.


    .-----.                        .-------.               .--------.                      .---------.         .--------.
    | App |------UDP---->[proxyApp]|  App  |--+-channel--->| Device |------UDP---->[device]| com2tcp |---COM-->| Device |
    |     |[app]<----UDP-----------| proxy |<-+-channel-+--| proxy  |[proxyDevice]<--UDP---|         |<--COM---|        |
    '-----'                        '-------'  |         |  '--------'                      '---------'         '--------'
                                              v         v
                                             .-----------.
                                             | Standard  |
                                             | text      |
                                             | console   |
                                             '-----------'

App and Device proxies are in this tool.

App, proxyApp, device, and proxyDevice are UDP endpoints (IP + port).
App endpoint is autodetected from the first packet as in most cases it's dynamic.
Other endpoints are input for the tool.

Com2tcp is used to add UPD interface to COM device. Manual: http://com0com.sourceforge.net/doc/UsingCom0com.pdf

Communication between App proxy and Device proxy is done using Go channels (toDevice, fromDevice in the code).


## Usage
For usage information execute `udp-sniffer -h`:

    $ udp-sniffer -h
    Usage of udp-sniffer:
	  -deviceIP string
			IP of device (default "127.0.0.1")
	  -devicePort int
			UPD port of device (default 11000)
	  -proxyAppIP string
			IP for this tool to communicate with device app
	  -proxyAppPort int
			UPD port for this tool to communicate with app (default 11000)
	  -proxyDeviceIP string
			IP for this tool to communicate with device
	  -proxyDevicePort int
			UPD port for this tool to communicate with device (default 11010)
	  -v	Show app version

Here is an example of session:

    $ udp-sniffer
    2016/01/05 03:08:03 Started...
    2016/01/05 03:08:03 appProxy: >> listen on :11000 and send to unknown yet
    2016/01/05 03:08:03 appProxy: app's IP and port will be autodetected from the first data packet
    2016/01/05 03:08:03 deviceProxy: << listen on :11010 and send to 192.168.1.34:11000
    2016/01/05 03:08:05 Last app address: 192.168.1.35:11000
    2016/01/05 03:08:05 >> 00 00 0E 12
    2016/01/05 03:08:05 << 00 00 11 00
    ^C2016/01/05 03:08:08 Stopped...


## Installation
You can download [binaries and packages](https://github.com/stepin/udp-sniffer/releases) compatible with your operation system.

To compile you need the [Golang v1.12](https://golang.org). Run `make` (or just `go build`) command in the project folder.


## Status
This project is feature complete. No future development is expected.


## Support
For me it's one time tool. So, no active support or development is planned. I will try to apply PRs if they will be created.


## Similar apps for COM ports
- [Null-modem emulator (com0com)](http://com0com.sourceforge.net) - it converts COM port to UDP port.
- [RealTerm](http://realterm.sourceforge.net) - a good app though it does not support UDP.
- [Hercules](http://www.hw-group.com/products/hercules/index_en.html) - it has a limited UDP support: received data is in the ASCII format only.
