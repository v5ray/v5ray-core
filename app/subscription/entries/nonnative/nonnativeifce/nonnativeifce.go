package nonnativeifce

import (
	"io/fs"

	"github.com/v4fly/v4ray-core/v0/app/subscription/entries"
)

type NonNativeConverterConstructorT func(fs fs.FS) (entries.Converter, error)

var NewNonNativeConverterConstructor NonNativeConverterConstructorT
