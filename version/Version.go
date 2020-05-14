package version

import "fmt"

type Version struct {
	Number     float32
	PatchLevel int
}

func (v Version) ToString() string {
	return version(v.Number, v.PatchLevel)
}
func version(version float32, patchLevel int) string {
	return "GDSP version: " + fmt.Sprintf("%.1f.%d", version, patchLevel)
}
