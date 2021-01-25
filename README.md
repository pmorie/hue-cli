# hue-cli

A hue CLI client.

## building

```
go build ./cmd/hue-cli
```

## setting up a bridge

Use the `bridge discover` command to find a bridge on your network, then `bridge setup` to set that bridge up with the CLI.

```
$ ./hue-cli bridge discover
IP             ID
192.168.50.100 001788fffeeeeeee

$ ./hue-cli bridge setup --bridgeIP 192.168.50.100 --user myuser --wait=true
# you will have to hit the button on the hue bridge after invoking this
```
