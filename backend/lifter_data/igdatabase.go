package lifter_data

import (
	"embed"
)

//go:embed **/*.csv
var InstagramDatabase embed.FS
