# Google Wifi Dump

googlewifidump will download, parse, and print a JSON version of the Google Wifi
diagnostic report. It's a command from the
[googlewifi](https://github.com/bengalaviz/googlewifi) Go package.

This repo has been updated to read the new protobuf output from the Google Wifi pucks. Not
sure if it works for the OnHub version. 

## Installation

```sh
go get github.com/bengalaviz/googlewifi/cmd/googlewifidump
```

## Running

```sh
googlewifidump
```

This returns a JSON dump of the data from `http://192.168.86.1/api/v1/diagnostic-report`.

If you use a different subnet for your router:

```sh
googlewifidump http://192.168.85.1/api/v1/diagnostic-report
```

If you want to run against a local, already downloaded dump:

```sh
googlewifidump ./path/to/diagnostic-report
```

You will likely want to either save a copy of the diagnostic report or save a
copy of the parsed output from googlewifidump rather than re-running the command as
generating the report takes a few seconds each time.

```sh
curl -o diagnostic-report http://192.168.86.1/api/v1/diagnostic-report
googlewifidump ./diagnostic-report
```

Alternatively:

```sh
googlewifidump > diagnostic-report.json
```

Use the [jq](https://stedolan.github.io/jq/) tool to format and manipulate
output on the command line.

Unbuffered:

```sh
googlewifidump | jq
```

Buffered report:

```sh
curl -o diagnostic-report http://192.168.86.1/api/v1/diagnostic-report
googlewifidump ./diagnostic-report | jq
```

Buffered output:

```sh
onhubdump > diagnostic-report.json
jq < diagnostic-report.json
```

I have included a script (getInfo.sh) that will try and fix the infoJSON output and make it valid JSON.
Make sure the output binary is in your path before running.

Then use [hjson](https://hjson.org) tool to correct the JSON so jq can read it.

```sh
hjson -j final.json | jq '.'

hjson -j final.json | jq '.network_service_state.station_state_updates | .[] | .station_state_update'
```
