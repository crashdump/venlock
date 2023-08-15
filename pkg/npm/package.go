package npm

import "encoding/json"

// PackageJSON represents a package.json
type PackageJSON struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Description     string            `json:"description"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

func parse(payload []byte) (*PackageJSON, error) {
	var res *PackageJSON
	err := json.Unmarshal(payload, &res)
	return res, err
}
