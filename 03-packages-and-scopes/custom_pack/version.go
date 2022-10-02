package custom_pack

import "runtime"

func Version() string {
	return runtime.Version()
}
