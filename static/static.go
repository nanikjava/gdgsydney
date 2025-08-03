package static

import "embed"

//go:embed *.html
var StaticFS embed.FS
