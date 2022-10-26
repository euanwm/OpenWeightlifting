package database

import (
	"embed"
)

//go:embed **/*.csv
var Database embed.FS
