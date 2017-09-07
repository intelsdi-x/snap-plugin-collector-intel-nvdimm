# snap collector plugin - Intel NVDIMM

This plugin is able to gather metrics from Intel NVDIMMs related infrastructure. It is collecting the data from Intel management software: https://github.com/01org/ixpdimm_sw and Non-Volatile Memory Libraries http://pmem.io/

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license-and-authors)
6. [Acknowledgements](#acknowledgements)

## Getting Started


### System Requirements
* [golang 1.6+](https://golang.org/dl/)  (needed only for building)
* [ixpdimm_sw](https://github.com/01org/ixpdimm_sw)
* [glide]()

### Operating systems
All OSs currently supported by plugin:
* Fedora (Server) 25 or higher

### Configuration and Usage
* Set up the [SNAP framework](https://github.com/intelsdi-x/snap#getting-started)
* Load the plugin: ` $ snaptel plugin load snap-plugin-collector-intel-nvdimm `

### Installation
Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:
```
$ git clone https://github.com/intelsdi-x/snap-plugin-collector-intel-nvdimm.git
```
And run `$ make`

## Documentation
To learn more about ixpdimm_sw visit [ixpdimm_sw](https://github.com/01org/ixpdimm_sw).

### Collected Metrics

List of collected metrics is described in [METRICS.md](METRICS.md).

## Community Support
This repository is one of **many** plugins in **Snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap
To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support) or visit [Slack](http://slack.snap-telemetry.io).

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[Snap](http://github.com:intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
Authors: 
* [Maciej Maciejewski](https://github.com/Maciuch)
* [Krzysztof Filipek](https://github.com/KFilipek)
* [Adrian Nidzgorski](https://github.com/anidzgor)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.
