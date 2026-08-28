# vscode-icons

[![ci](https://github.com/go-icons/vscode-icons/actions/workflows/ci.yml/badge.svg)](https://github.com/go-icons/vscode-icons/actions/workflows/ci.yml)
![coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-icons/vscode-icons.svg)](https://pkg.go.dev/github.com/go-icons/vscode-icons)
[![License](https://img.shields.io/badge/license-BSD--3--Clause-blue.svg)](LICENSE)

File-type icons from the [vscode-icons](https://github.com/vscode-icons/vscode-icons)
icon set (by Roberto Huertas and contributors, MIT), as embedded SVG documents
keyed by file name — for pure-Go UIs that render their own icons.

```go
import vscodeicons "github.com/go-icons/vscode-icons"

svg := vscodeicons.Icon("paper.tex") // the vscode-icons .tex glyph, as an SVG string
dir := vscodeicons.Folder()          // the folder glyph
```

`Icon(filename)` matches by exact base name first (`go.mod`, `LICENSE`,
`.gitignore`, `package.json`, `Dockerfile`), then by extension, then falls back
to a generic document (`default_file`). It is a **data package**: it returns SVG
strings and draws nothing. A renderer such as
[go-widgets/toolkit](https://github.com/go-widgets/toolkit)'s `SVGIcon` turns
the SVG into a drawn glyph.

A curated subset of the vscode-icons artwork is embedded (common source, markup,
data and image types). Every embedded icon is verified to rasterise under the
[go-gfx/gfx](https://github.com/go-gfx/gfx) SVG renderer; icons that rely on SVG
features the renderer does not support (e.g. `<polygon>`, masks, filters,
gradients) are omitted. Contributions adding more are welcome.

## Licence

The Go code is BSD-3-Clause (`LICENSE`). The embedded vscode-icons artwork is
MIT, © 2016 Roberto Huertas (`VSCODE-ICONS-LICENSE`) — redistributed unmodified.
