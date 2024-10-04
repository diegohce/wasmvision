package guest

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	//"github.com/hashicorp/go-getter"
	getter "github.com/diegohce/simple-getter"
)

type ProcessorFile struct {
	Alias    string
	Filename string
	URL      string
}

var KnownProcessors = map[string]ProcessorFile{
	"asciify": {
		Alias:    "asciify",
		Filename: "asciify.wasm",
		URL:      "https://github.com/wasmvision/wasmvision/raw/refs/heads/main/processors/asciify.wasm",
	},
	"blur": {
		Alias:    "blur",
		Filename: "blur.wasm",
		URL:      "https://github.com/wasmvision/wasmvision/raw/refs/heads/main/processors/blur.wasm",
	},
	"blurrs": {
		Alias:    "blurrs",
		Filename: "blurrs.wasm",
		URL:      "https://github.com/wasmvision/wasmvision/raw/refs/heads/main/processors/blurrs.wasm",
	},
	"candy": {
		Alias:    "candy",
		Filename: "candy.wasm",
		URL:      "https://github.com/wasmvision/wasmvision/raw/refs/heads/main/processors/candy.wasm",
	},
	"gaussianblur": {
		Alias:    "gaussianblur",
		Filename: "gaussianblur.wasm",
		URL:      "https://github.com/wasmvision/wasmvision/raw/refs/heads/main/processors/gaussianblur.wasm",
	},
	"hello": {
		Alias:    "hello",
		Filename: "hello.wasm",
		URL:      "https://github.com/wasmvision/wasmvision/raw/refs/heads/main/processors/hello.wasm",
	},
	"mosaic": {
		Alias:    "mosaic.wasm",
		Filename: "mosaic.wasm",
		URL:      "https://github.com/wasmvision/wasmvision/raw/refs/heads/main/processors/mosaic.wasm",
	},
	"pointilism": {
		Alias:    "pointilism",
		Filename: "pointilism.wasm",
		URL:      "https://github.com/wasmvision/wasmvision/raw/refs/heads/main/processors/pointilism.wasm",
	},
	"rainprincess": {
		Alias:    "rainprincess",
		Filename: "rainprincess.wasm",
		URL:      "https://github.com/wasmvision/wasmvision/raw/refs/heads/main/processors/rainprincess.wasm",
	},
	"udnie": {
		Alias:    "udnie",
		Filename: "udnie.wasm",
		URL:      "https://github.com/wasmvision/wasmvision/raw/refs/heads/main/processors/udnie.wasm",
	},
}

func DownloadProcessor(processor ProcessorFile, targetDir string) error {
	opts := []getter.ClientOption{}
	client := &getter.Client{
		Ctx:     context.Background(),
		Src:     processor.URL,
		Dst:     filepath.Join(targetDir, filepath.Base(processor.Filename)),
		Mode:    getter.ClientModeFile,
		Options: opts,
	}

	if err := client.Get(); err != nil {
		return err
	}

	return nil
}

func ProcessorExists(processor string) bool {
	if _, err := os.Stat(processor); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func ProcessorWellKnown(processor string) bool {
	if _, ok := KnownProcessors[processor]; ok {
		return true
	}

	return false
}
