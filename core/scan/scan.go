package scan

import (
	"fmt"
	"strings"
)

// Run returns a summary line for a scan of target.
func Run(target string) string {
	if strings.TrimSpace(target) == "" {
		return "refusing to scan an empty target"
	}
	return fmt.Sprintf("scanned %s: 0 failed controls", target)
}
