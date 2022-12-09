# Oakley 192/256-bit
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/oakley/blob/master/LICENSE.md) 
[![GoDoc](https://godoc.org/github.com/pedroalbanese/oakley?status.png)](http://godoc.org/github.com/pedroalbanese/oakley)
[![GitHub downloads](https://img.shields.io/github/downloads/pedroalbanese/oakley/total.svg?logo=github&logoColor=white)](https://github.com/pedroalbanese/oakley/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/oakley)](https://goreportcard.com/report/github.com/pedroalbanese/oakley)
[![Lines of code](https://img.shields.io/tokei/lines/github/pedroalbanese/oakley)](https://github.com/pedroalbanese/oakley)

### Parameters for the [Oakley](https://www.rfc-editor.org/rfc/rfc5114#section-2.6) Random ECP Group Elliptic curves IETF in RFC5114
### Command-line Oakley Diffie-Hellman Tool:
<pre>Usage of oakley:
  -key string
        Private key.
  -keygen
        Generate asymmetric keypair.
  -pub string
        Remote's side Public key.</pre>

### Examples:
```sh
./oakley -keygen // 2x
./oakley -key $2ndprivatekey -pub $1stpublickey
./oakley -key $1stprivatekey -pub $2ndpublickey
```
## License

This project is licensed under the ISC License.

##### Industrial-Grade Reliability. Copyright (c) 2020-2023 ALBANESE Research Lab.
