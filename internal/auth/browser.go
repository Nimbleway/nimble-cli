package auth

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func OpenBrowser(url string) error {
	if cmd := os.Getenv("NIMBLE_BROWSER"); cmd != "" {
		parts := strings.Fields(cmd)
		parts = append(parts, url)
		return exec.Command(parts[0], parts[1:]...).Start()
	}
	switch runtime.GOOS {
	case "darwin":
		return exec.Command("open", url).Start()
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}
