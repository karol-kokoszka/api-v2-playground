package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	var (
		flags      flag.FlagSet
		importPath = flags.String("import_path", "github.com/scylladb/scylla-cloud", "")
		source     = flags.String("source", "gen", "")
		buildTag   = flags.String("build_tag", "mock", "")
	)
	opts := protogen.Options{
		ParamFunc: flags.Set,
	}

	fileImportPath := func(f *protogen.File) string {
		d := path.Dir(f.GeneratedFilenamePrefix)
		d = strings.TrimPrefix(d, "./")
		return path.Join(*importPath, *source, d)
	}

	opts.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			out := fmt.Sprintf("%s_mock.go", f.GeneratedFilenamePrefix)
			g := gen.NewGeneratedFile(out, f.GoImportPath)
			fmt.Fprintf(g, "//go:build %s\n\n", *buildTag)

			cmd := exec.Command("mockgen",
				"-package", string(f.GoPackageName),
				"-self_package", fileImportPath(f),
				"-source", fmt.Sprintf("./%s/%s.twirp.go", *source, f.GeneratedFilenamePrefix),
			)
			cmd.Stdout = g
			cmd.Stderr = os.Stderr
			return cmd.Run()
		}

		return nil
	})
}
