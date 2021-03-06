// +build darwin

package procfs

import (
	"encoding/json"
	"errors"
)

// NewProcLoadAvg is ProcLoadAvg constructor.
func NewProcLoadAvg() *ProcLoadAvg {
	p := &ProcLoadAvg{}
	p.Data = make(map[string][]interface{})
	return p
}

// ProcLoadAvg is a reader that scrapes /proc/loadavg data.
type ProcLoadAvg struct {
	Data map[string][]interface{}
}

func (p *ProcLoadAvg) Run() error {
	return errors.New("/proc/loadavg is only available on Linux.")
}

// ToJson serialize Data field to JSON.
func (p *ProcLoadAvg) ToJson() ([]byte, error) {
	return json.Marshal(p.Data)
}
