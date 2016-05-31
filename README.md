# wat-switch

A simple Go TCPIP server that receives a connection and start a mruby runtime for that socket.

## Installation

Follow the instructions to install the [go-mruby](github.com/mitchellh/go-mruby) library, copy
the libmruby.a to this directory and compile the server.go

## Usage

Just open a telnet connection to the port 31415 and send the name of a mruby script inside the executing
directory without `.rb`, a space, and the input buffer who will passed to the mruby script as a global
varible named `$INPUT`

## Warning

This is REALLY REALLY REALLY INSECURE!. Any atacker could input malicious mruby scripts to take over the
control of the server.go runtime. Use at your own risk. To avoid that you should control system calls
inside the mruby runtime through a pre processing of the script or sandboxed mruby runtime.
