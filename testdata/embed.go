package testdata

import _ "embed"

// PluginStdin is binary data from os.Stdin for protoc plugin.
// stdin.bin.pb is generated by `make testdata` command.
//
//go:embed stdin.bin.pb
var PluginStdin []byte

//go:embed stdout.bin.pb
var ExpectedStdout []byte
