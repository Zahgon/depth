package main

import (
	"io"
	"os"

	"github.com/KyleBanks/depth"
)

const (
	outputClosedPadding = "  "
	outputOpenPadding   = "│ "
	outputPrefix        = "├ "
	outputPrefixLast    = "└ "
)

var outputJSON bool
var explainPkg string

type summary struct {
	numInternal int
	numExternal int
	numTesting  int
}

func main() {
	t, pkgs := parse(os.Args[1:])
	if err := handlePkgs(t, pkgs, outputJSON, explainPkg); err != nil {
		os.Exit(1)
	}
}

// parse constructs a depth.Tree from command-line arguments, and returns the
// remaining user-supplied package names
func parse(args []string) (*depth.Tree, []string) { _ = "STUB: not implemented"; return nil, nil }

// handlePkgs takes a slice of package names, resolves a Tree on them,
// and outputs each Tree to Stdout.
func handlePkgs(t *depth.Tree, pkgs []string, outputJSON bool, explainPkg string) error {
	_ = "STUB: not implemented"
	return nil
}

// writePkgSummary writes a summary of all packages in a tree
func writePkgSummary(w io.Writer, pkg depth.Pkg) { _ = "STUB: not implemented"; return }

func collectSummary(sum *summary, pkg depth.Pkg, nameSet map[string]struct{}) {
	_ = "STUB: not implemented"
	return
}

// writePkgJSON writes the full Pkg as JSON to the provided Writer.
func writePkgJSON(w io.Writer, p depth.Pkg) { _ = "STUB: not implemented"; return }

func writePkg(w io.Writer, p depth.Pkg) { _ = "STUB: not implemented"; return }

// writePkg recursively prints a Pkg and its dependencies to the Writer provided.
func writePkgRec(w io.Writer, p depth.Pkg, closed []bool, isLast bool) {
	_ = "STUB: not implemented"
	return
}

// writeExplain shows possible paths for a given package.
func writeExplain(w io.Writer, pkg depth.Pkg, stack []string, explain string) {
	_ = "STUB: not implemented"
	return
}
