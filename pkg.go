package depth

import (
	"go/build"
)

// Pkg represents a Go source package, and its dependencies.
type Pkg struct {
	Name   string `json:"name"`
	SrcDir string `json:"-"`

	Internal bool `json:"internal"`
	Resolved bool `json:"resolved"`
	Test     bool `json:"-"`

	Tree   *Tree `json:"-"`
	Parent *Pkg  `json:"-"`
	Deps   []Pkg `json:"deps"`

	Raw *build.Package `json:"-"`
}

// Resolve recursively finds all dependencies for the Pkg and the packages it depends on.
func (p *Pkg) Resolve(i Importer) {
	_ = "STUB: not implemented"
	// Resolved is always true, regardless of if we skip the import,
	// it is only false if there is an error while importing.
	return
}

// Stop resolving imports if we've reached max depth or found a duplicate.

// TODO: Check the error type?

// Update the name with the fully qualified import path.

// If this is an internal dependency, we may need to skip it.

//first we set the regular dependencies, then we add the test dependencies
//sharing the same set. This allows us to mark all test-only deps linearly

// setDeps takes a slice of import paths and the source directory they are relative to,
// and creates the Deps of the Pkg. Each dependency is also further resolved prior to being added
// to the Pkg.
func (p *Pkg) setDeps(i Importer, imports []string, srcDir string, unique map[string]struct{}, isTest bool) {
	_ = "STUB: not implemented"
	return
}

// Mostly for testing files where cyclic imports are allowed.

// Skip duplicates.

// addDep creates a Pkg and it's dependencies from an imported package name.
func (p *Pkg) addDep(i Importer, name string, srcDir string, isTest bool) {
	_ = "STUB: not implemented"
	return
}

// isParent goes recursively up the chain of Pkgs to determine if the name provided is ever a
// parent of the current Pkg.
func (p *Pkg) isParent(name string) bool { _ = "STUB: not implemented"; return false }

// depth returns the depth of the Pkg within the Tree.
func (p *Pkg) depth() int { _ = "STUB: not implemented"; return 0 }

// cleanName returns a cleaned version of the Pkg name used for resolving dependencies.
//
// If an empty string is returned, dependencies should not be resolved.
func (p *Pkg) cleanName() string {
	_ = "STUB: not implemented"

	// C 'package' cannot be resolved.
	return ""
}

// Internal golang_org/* packages must be prefixed with vendor/
//
// Thanks to @davecheney for this:
// https://github.com/davecheney/graphpkg/blob/master/main.go#L46

// String returns a string representation of the Pkg containing the Pkg name and status.
func (p *Pkg) String() string { _ = "STUB: not implemented"; return "" }

// byInternalAndName ensures a slice of Pkgs are sorted such that the internal stdlib
// packages are always above external packages (ie. github.com/whatever).
type byInternalAndName []Pkg

func (b byInternalAndName) Len() int { _ = "STUB: not implemented"; return 0 }

func (b byInternalAndName) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (b byInternalAndName) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
