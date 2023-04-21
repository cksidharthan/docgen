# Docgen

NOTE: This is a fork of [docgen](https://github.com/thedevsaddam/docgen) with some minor changes.

Transform your postman collection to HTML/Markdown documentation

![Task screenshot](screenshot.png)

#### Installation/Update on Mac/Linux

```bash
curl https://raw.githubusercontent.com/cksidharthan/docgen/v3/install.sh -o install.sh \
&& sudo chmod +x install.sh \
&& sudo ./install.sh \
&& rm install.sh
```

#### Uninstallation

```bash
curl https://raw.githubusercontent.com/cksidharthan/docgen/v3/uninstall.sh -o uninstall.sh \
&& sudo chmod +x uninstall.sh \
&& sudo ./uninstall.sh \
&& rm uninstall.sh
```

#### Windows

**For windows download the binary and set environment variable so that you can access the binary from terminal**

#### Binary link

[Download binary](https://github.com/cksidharthan/docgen/releases)

### Available features

- Live postman collection to documentation
- Build postman collection to html/markdown documentation
- Supports multi-level collection build

### Usage

- To view live HTML documentation from postman collection use `docgen server -f input-postman-collection.json -p 8000` This will open the html version of postman collection to the defined port
- To view live Markown documentation from postman collection use `docgen server -f input-postman-collection.json -p 8000 -m` This will open the markdown version of postman collection to the defined port
- To make HTML documentation use `docgen build -i input-postman-collection.json -o ~/Downloads/index.html`
- To make Markdown documentation use `docgen build -i input-postman-collection.json -o ~/Downloads/index.md -m`

**_[Demo markdown API documentation](_examples/example-doc.md)_**

### Contributor

1. [Anondo](https://github.com/Anondo)

### Contribution

Your suggestions will be more than appreciated.
[Read the contribution guide here](CONTRIBUTING.md)

### See all [contributors](https://github.com/cksidharthan/docgen/graphs/contributors)

### **License**

The **docgen** is an open-source software licensed under the [MIT License](LICENSE.md).
