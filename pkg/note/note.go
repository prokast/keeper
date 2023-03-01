// Configuration handler for active directory
package note

import (
	"github.com/infrastate/treasure/pkg/configReader"
)

// Active directory configuration structure
type config []struct {
	Alias       string
	ExpiredDate string
}

// Return activeDirectory.config structure with data
func Note() config {
	cfg := config(configReader.ReadConfig().Note)
	return cfg
}
