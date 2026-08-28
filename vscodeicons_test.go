// Copyright (c) 2026 the go-widgets authors. All rights reserved.
// Use of this source code is governed by a BSD-3-Clause license that can be
// found in the LICENSE file at the root of this repository.

package vscodeicons

import (
	"strings"
	"testing"
)

func svgOK(s string) bool { return s != "" && strings.Contains(s, "<svg") }

// TestIconByExtension: a file's extension selects its icon.
func TestIconByExtension(t *testing.T) {
	for _, name := range []string{"paper.tex", "notes.MD", "main.go", "app.py", "lib.rs", "style.css", "logo.png", "report.pdf"} {
		if !svgOK(Icon(name)) {
			t.Errorf("Icon(%q) returned no SVG", name)
		}
	}
}

// TestTeXAndStyDiffer: a LaTeX document (.tex) and a LaTeX package/class
// (.sty/.cls) — read apart from a document — get distinct glyphs.
func TestTeXAndStyDiffer(t *testing.T) {
	tex := Icon("paper.tex")
	if !svgOK(tex) {
		t.Fatal("Icon(paper.tex) returned no SVG")
	}
	for _, name := range []string{"pkg.sty", "book.cls"} {
		if got := Icon(name); !svgOK(got) || got == tex {
			t.Errorf("Icon(%q) should be a distinct config glyph, not the .tex icon", name)
		}
	}
}

// TestIconByName: a full base name (no useful extension) selects its icon.
func TestIconByName(t *testing.T) {
	for _, name := range []string{"LICENSE", "go.mod", ".gitignore", "path/to/go.sum", "package.json", "Dockerfile", "Makefile"} {
		if !svgOK(Icon(name)) {
			t.Errorf("Icon(%q) returned no SVG", name)
		}
	}
}

// TestIconDefaultFallback: an unknown extension falls back to the default glyph.
func TestIconDefaultFallback(t *testing.T) {
	if !svgOK(Icon("mystery.zzz")) {
		t.Error("unknown extension should fall back to the default icon")
	}
	if !svgOK(Icon("noext")) {
		t.Error("an extensionless unknown name should fall back to the default icon")
	}
}

// TestFolder: the folder glyph is present.
func TestFolder(t *testing.T) {
	if !svgOK(Folder()) {
		t.Error("Folder() returned no SVG")
	}
}

// TestName: the pack exposes a human label.
func TestName(t *testing.T) {
	if Name == "" {
		t.Error("Name should not be empty")
	}
}

// TestEveryMappedIconEmbedded guards the maps against a typo: every icon name
// they reference (plus the defaults) must resolve to an embedded SVG.
func TestEveryMappedIconEmbedded(t *testing.T) {
	seen := map[string]bool{"default_file": true, "default_folder": true}
	for _, m := range []map[string]string{byName, byExt} {
		for _, ic := range m {
			seen[ic] = true
		}
	}
	for ic := range seen {
		if read(ic) == "" {
			t.Errorf("mapped icon %q is not embedded", ic)
		}
	}
}

// TestReadMissing covers read's not-found branch directly.
func TestReadMissing(t *testing.T) {
	if read("definitely-not-an-icon") != "" {
		t.Error("read of a missing icon should return an empty string")
	}
}
