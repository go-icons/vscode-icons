// Copyright (c) 2026 the go-widgets authors. All rights reserved.
// Use of this source code is governed by a BSD-3-Clause license that can be
// found in the LICENSE file at the root of this repository.

// Package vscodeicons serves file-type icons from the vscode-icons icon set (by
// Roberto Huertas and contributors, MIT — see VSCODE-ICONS-LICENSE), as SVG
// documents keyed by file name.
//
// It is a data package: a curated subset of the icons is embedded, and [Icon]
// maps a file name to the best-matching SVG (falling back to a generic
// document), [Folder] returns the folder glyph. A renderer such as
// go-widgets/toolkit's SVGIcon turns the returned SVG into a drawn glyph; this
// package draws nothing itself.
package vscodeicons

import (
	"embed"
	"path"
	"strings"
)

// Name is the human label a picker shows for this pack.
const Name = "VSCode Icons"

//go:embed svg/*.svg
var files embed.FS

// byName maps a lower-cased base file name to an icon, for names that carry more
// meaning than their extension (or have none).
var byName = map[string]string{
	"license":        "file_type_license",
	"license.md":     "file_type_license",
	"license.txt":    "file_type_license",
	"copying":        "file_type_license",
	".gitignore":     "file_type_git",
	".gitattributes": "file_type_git",
	".gitmodules":    "file_type_git",
	"go.mod":         "file_type_go",
	"go.sum":         "file_type_go",
	"package.json":   "file_type_npm",
	"dockerfile":     "file_type_docker",
	"makefile":       "file_type_config",
}

// byExt maps a lower-cased extension (with the dot) to an icon.
var byExt = map[string]string{
	".tex": "file_type_tex", ".sty": "file_type_tex", ".cls": "file_type_tex",
	".bib": "file_type_tex", ".dtx": "file_type_tex", ".ins": "file_type_tex",
	".md": "file_type_markdown", ".markdown": "file_type_markdown",
	".json": "file_type_json",
	".yml":  "file_type_yaml", ".yaml": "file_type_yaml",
	".xml": "file_type_xml",
	".svg": "file_type_svg",
	".css": "file_type_css",
	".js":  "file_type_js", ".mjs": "file_type_js", ".cjs": "file_type_js", ".jsx": "file_type_js",
	".ts": "file_type_typescript", ".tsx": "file_type_typescript",
	".py": "file_type_python",
	".go": "file_type_go",
	".rs": "file_type_rust",
	".c":  "file_type_c", ".h": "file_type_c",
	".cpp": "file_type_cpp", ".cc": "file_type_cpp", ".cxx": "file_type_cpp",
	".hpp": "file_type_cpp", ".hh": "file_type_cpp",
	".cs": "file_type_csharp",
	".sh": "file_type_shell", ".bash": "file_type_shell", ".zsh": "file_type_shell",
	".pdf": "file_type_pdf",
	".png": "file_type_image", ".jpg": "file_type_image", ".jpeg": "file_type_image",
	".gif": "file_type_image", ".webp": "file_type_image", ".bmp": "file_type_image", ".eps": "file_type_image",
	".lua":  "file_type_lua",
	".rb":   "file_type_ruby",
	".java": "file_type_java",
	".php":  "file_type_php",
	".toml": "file_type_toml",
	".ini":  "file_type_ini", ".cfg": "file_type_ini", ".conf": "file_type_ini",
	".lock": "file_type_config",
}

// Icon returns the SVG document for the file named filename (any path — only the
// base name matters), matching by exact name first, then by extension, then a
// generic document. It never returns "" for a normal name: the default icon is
// embedded.
func Icon(filename string) string {
	base := strings.ToLower(path.Base(filename))
	if ic, ok := byName[base]; ok {
		return read(ic)
	}
	if ic, ok := byExt[strings.ToLower(path.Ext(base))]; ok {
		return read(ic)
	}
	return read("default_file")
}

// Folder returns the SVG document for a directory.
func Folder() string { return read("default_folder") }

// read returns the embedded SVG for an icon name, or "" when absent.
func read(name string) string {
	b, err := files.ReadFile("svg/" + name + ".svg")
	if err != nil {
		return ""
	}
	return string(b)
}
