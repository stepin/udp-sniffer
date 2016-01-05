# UDP COM (serial) port Sniffer/Proxy [![GitHub release](https://img.shields.io/github/release/stepin/udp-sniffer.svg)](https://github.com/stepin/udp-sniffer/releases) [![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/stepin/udp-sniffer/master/LICENSE) [![GoDoc](https://godoc.org/github.com/stepin/udp-sniffer?status.svg)](https://godoc.org/github.com/stepin/udp-sniffer) [![wercker status](https://app.wercker.com/status/52211c1af988eb5c9608e6c65d918642/s/master "wercker status")](https://app.wercker.com/project/bykey/52211c1af988eb5c9608e6c65d918642) [![Go Report Card](http://goreportcard.com/badge/stepin/udp-sniffer)](http://goreportcard.com/report/stepin/udp-sniffer)

UDP COM-порт сниффер/прокси для отладки взаимодействия приложения и устройства, общающихся по UDP COM-порту.

Принимает и отправляет трафик с распечаткой его в шестнадцатеричной форме в консоль.

## Использование

Для справки выполните `udp-sniffer -h`:

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

Пример работы:

    $ udp-sniffer
    2016/01/05 03:08:03 Started...
    2016/01/05 03:08:03 udpServerProxyConnection: >> listen on :11000 and send to unknown yet
    2016/01/05 03:08:03 udpClientProxyConnection: << listen on :11010 and send to 192.168.1.34:11000
    2016/01/05 03:08:05 last client address: 192.168.1.35:11000
    2016/01/05 03:08:05 >> 00 00 0E 12
    2016/01/05 03:08:05 << 00 00 11 00
    ^C2016/01/05 03:08:08 Stopped...

## Установка
Проще всего скачать [готовый дистрибьютив](https://github.com/stepin/udp-sniffer/releases) для своей операционной системы.

Для сборки из исходников нужно установить [Golang](https://golang.org) (разрабатывалось на версии 1.5.1) и выполнить `go build` в папке проекта. В текущей папке появится исполняемый файл upd-sniffer (на Windows udp-sniffer.exe).

## Статус
Проект закончен. Дальнейшей разработки не планируется.

## Поддержка
Утилита для меня разовая, поэтому активной поддержки не предполагается. Если будут ошибки/патчи, то при наличии времени постараюсь обновить.

## Еще программы для работы с COM-портами

- [RealTerm](http://realterm.sourceforge.net) -- отличная программа, но не работает именно с UDP COM-портами.
- [Null-modem emulator (com0com)](http://com0com.sourceforge.net) -- позволяет подключать к приложению программный эмулятор устройства.
- [Hercules](http://www.hw-group.com/products/hercules/index_en.html) -- поддерживает UDP COM соединения, но получаемые данные приходят в ASCII, а не HEX.
