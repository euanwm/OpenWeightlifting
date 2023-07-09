package lifterdata

import (
	"embed"
)

//go:embed *.csv
var InstagramDatabase embed.FS
