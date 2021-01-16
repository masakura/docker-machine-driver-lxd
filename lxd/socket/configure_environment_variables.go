package socket

import "os"

func ConfigureEnvironmentVariables() error {
	path := DefaultUnixSocketResolver().Resolve()

	if path == "" {
		return nil
	}

	return os.Setenv("LXD_SOCKET", path)
}
