package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"os"

	extism "github.com/extism/go-sdk"
	"github.com/tetratelabs/wazero"
)

func loadPluginFromFile(file string) *extism.Plugin {
	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: "../plugin/plugin.wasm",
			},
		},
	}

	ctx := context.Background()
	cfg := extism.PluginConfig{
		EnableWasi: true,
		ModuleConfig: wazero.NewModuleConfig().
			WithStdout(os.Stdout).
			WithStderr(os.Stderr),
	}
	plugin, err := extism.NewPlugin(ctx, manifest, cfg, []extism.HostFunction{})

	if err != nil {
		log.Fatalf("Failed to initialize plugin: %s\n", err)
	}
	return plugin
}

func main() {
	plugin := loadPluginFromFile("../plugin/plugin.wasm")

	nums := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}
	bytes := make([]byte, 8*len(nums))

	for id, n := range nums {
		bits := math.Float64bits(n)
		binary.LittleEndian.PutUint64(bytes[8*id:8*(id+1)], bits)
	}

	exit, outMean, err := plugin.Call("mean", bytes)
	if err != nil {
		log.Println(err)
		os.Exit(int(exit))
	}
	bitsMean := binary.LittleEndian.Uint64(outMean)
	mean := math.Float64frombits(bitsMean)
	fmt.Printf("Mean: %F\n", mean)

	exit, outStdDev, err := plugin.Call("stdDev", bytes)
	if err != nil {
		log.Println(err)
		os.Exit(int(exit))
	}
	bits := binary.LittleEndian.Uint64(outStdDev)
	stdDev := math.Float64frombits(bits)
	fmt.Printf("StdDev: %F\n", stdDev)
}
