# dw
A super simple CLI that - given a database url - will wait for that database to become available.

### Installation

1. Download the latest version of dw for your OS from the [Releases](https://github.com/codingconcepts/dw/releases) page.

2. Extract the binary:

``` sh
tar -xvf dw_YOUR_OS.tar.gz
```

3. Move the binary to your PATH.

### Usage

``` sh
dw "DATABASE_CONNECTION_STRING"
```

dw will continually attempt to connect to the database until it achieves a connection, at which point it will terminate with a successfully exit code.