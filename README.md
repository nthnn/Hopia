<p align="center">
    <img src="docs/src/assets/hopia.png" width="270px" />
    <h1 align="center">Hopia Wannabe-Hacker Tool</h1>
</p>

![Go Build](https://github.com/nthnn/Hopia/actions/workflows/go-build.yml/badge.svg)

Hopia is a powerful and versatile wannabe penetration testing tool designed for ethical and legitimate security testing purposes. It provides a range of wannabe security testing capabilities to help assess and enhance the security of systems and networks.

It is also important to note that Hopia is intended to be used responsibly, within the boundaries of the law, and with proper authorization from system owners or administrators. Any malicious or unauthorized use of this software is strictly prohibited and may result in legal consequences.

Please refer to [disclaimer](DISCLAIMER.md).

## Features

- **DDoS Testing**: Simulate Distributed Denial of Service (DDoS) attacks to assess the resilience of your systems against such threats. Only use this feature with proper authorization on systems you own or have explicit permission to test.

- **Hash Cracking**: Attempt to crack hash strings using popular algorithms like MD5, SHA-1, and SHA-256. This feature should only be used for legitimate purposes, such as password recovery on systems where you have explicit authorization.

- **Admin Page Finder**: Find and discover administrator login pages on websites using a file containing list of common URLs. Always use this feature responsibly and only on websites for which you have proper and legal authorization.

- **Port Scanning**: Scan for open ports from a given host URL in order to identify potential vulnerabilities. As always, ensure you have legal permission and authorization to scan the target system before using this feature.

## Usage

```bash
# Example command for DDoS testing
./hopia ddos http://localhost attack.properties

# Example command for hash cracking
./hopia crack 596a96cc7bf9108cd896f33c44aedc8a bruteforce.properties

# Example command for admin page discovery
./hopia af http://localhost admin-path.txt

# Example command for port scanning
./hopia ps http://localhost
```

## License

```
Copyright (c) 2023 Nathanne Isip

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```