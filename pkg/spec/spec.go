package spec

import "sync"

const (
	// DefaultVersion is the default version of the specification.
	DefaultVersion = 1
)

// Spec is the specification for the scaffolding tool.
type Spec struct {
	// Version is the version of the specification.
	Version int `validate:"required" yaml:"version"`
	// Name is the name of the project.
	Name string `validate:"required" yaml:"name"`
	// Templates is the list of templates to use.
	Templates []Template `validate:"required" yaml:"templates"`

	sync.Mutex `yaml:"-"`
}

// Template is a template to use.
type Template struct {
	// Source is the source of the template.
	Source string `validate:"required" yaml:"source"`
	// Destination is the destination of the template.
	Destination string `validate:"required" yaml:"destination"`
}

// Default returns the default specification.
func Default() *Spec {
	return &Spec{
		Version: DefaultVersion,
	}
}
