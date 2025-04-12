package cv

import (
	"encoding/json"

	"github.com/wasmvision/wasmvision/config"

	"github.com/wasmvision/wasmvision/datastore"
	_ "github.com/wasmvision/wasmvision/datastore/boltdb"
	_ "github.com/wasmvision/wasmvision/datastore/memory"
)

// Context is the configuration for the cv package used when each call is made
// from the guest module.
type Context struct {
	ReturnDataPtr  uint32
	ModelsDir      string
	Config         *config.Store
	FrameStore     *datastore.Frames
	ProcessorStore *datastore.Processors
	EnableCUDA     bool
}

func NewContext(modelsDir string, conf *config.Store, enableCUDA bool) (*Context, error) {
	var datastoreBackend string
	var datastoreBackendConfigStr string
	var datastoreBackendConfig any

	datastoreBackend, beFound := conf.Get("datastore_backend")
	datastoreBackendConfigStr, beConfigFound := conf.Get("datastore_backend")

	if beFound && beConfigFound {
		cfgMap := make(map[string]any)
		err := json.Unmarshal([]byte(datastoreBackendConfigStr), &cfgMap)
		if err != nil {
			return nil, err
		}
		datastoreBackendConfig = cfgMap

	} else {
		datastoreBackend = "memory"
		datastoreBackendConfig = nil
	}

	frameStore, err := datastore.NewFrames(datastoreBackend, datastoreBackendConfig)
	if err != nil {
		return nil, err
	}
	processorStore, err := datastore.NewProcessors(datastoreBackend, datastoreBackendConfig)
	if err != nil {
		return nil, err
	}

	return &Context{
		ModelsDir:      modelsDir,
		Config:         conf,
		FrameStore:     frameStore,
		ProcessorStore: processorStore,
		EnableCUDA:     enableCUDA,
	}, nil
}
