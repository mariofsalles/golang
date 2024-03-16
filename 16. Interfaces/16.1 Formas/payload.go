package domain

import "fmt"

type Payload interface {
	proxy() map[string]string
	sharedFlows() map[string]string
}

type Proxy struct {
	NAME            string   `json:"NAME"`
	ORGANIZATION    string   `json:"ORGANIZATION"`
	ENVIRONMENT     string   `json:"ENVIRONMENT"`
	ENVIRONMENTS    []string `json:"ENVIRONMENTS"`
	IMPORT_TYPE     string   `json:"IMPORT_TYPE,omitempty"`
	SERVICE_ACCOUNT string   `json:"SERVICE_ACCOUNT,omitempty"`
	VERSION         int      `json:"VERSION,omitempty"`
}

func (p *Proxy) proxy() map[string]string {
	return map[string]string{
		"NAME":            p.NAME,
		"ORGANIZATION":    p.ORGANIZATION,
		"ENVIRONMENT":     p.ENVIRONMENT,
		"ENVIRONMENTS":    fmt.Sprintf("%v", p.ENVIRONMENTS),
		"IMPORT_TYPE":     p.IMPORT_TYPE,
		"SERVICE_ACCOUNT": p.SERVICE_ACCOUNT,
		"VERSION":         fmt.Sprintf("%v", p.VERSION),
	}
}

type SharedFlows struct {
	NAME         string `json:"NAME,omitempty"`
	ORGANIZATION string `json:"ORGANIZATION"`
	ENVIRONMENT  string `json:"ENVIRONMENT"`
	VERSION      int    `json:"VERSION,omitempty"`
}

func (p *SharedFlows) sharedFlows() map[string]string {
	return map[string]string{
		"NAME":         p.NAME,
		"ORGANIZATION": p.ORGANIZATION,
		"ENVIRONMENT":  p.ENVIRONMENT,
		"VERSION":      fmt.Sprintf("%v", p.VERSION),
	}
}
