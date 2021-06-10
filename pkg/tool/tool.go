// Copyright 2021 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tool

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"golang.org/x/tools/go/packages"
)

func Cmd() *cobra.Command {
	cfg := &packages.Config{
		Mode: packages.NeedModule | packages.NeedImports,
	}
	ret := &cobra.Command{
		Use:  filepath.Base(os.Args[0]) + " <go package pattern> ...",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return exec(cmd.Context(), cfg, args)
		},
	}
	ret.Flags().StringVarP(&cfg.Dir, "dir", "d", ".", "the source directory")
	ret.Flags().StringArrayVarP(&cfg.BuildFlags, "build", "b", []string{"-mod=mod"},
		"arguments to pass to the golang build tool")

	return ret
}

func exec(ctx context.Context, cfg *packages.Config, pattern []string) error {
	cfg.Context = ctx
	pkgs, err := packages.Load(cfg, pattern...)
	if err != nil {
		return err
	}

	seen := make(map[string]bool)
	var out [][]string
	for _, pkg := range pkgs {
		if len(pkg.Errors) > 0 {
			return pkg.Errors[0]
		}
		out = crawl(pkg, pkg, seen, out)
	}

	sort.Slice(out, func(i, j int) bool {
		return strings.Compare(out[i][0], out[j][0]) < 0
	})

	tw := tabwriter.NewWriter(os.Stdout, 2, 8, 2, ' ', 0)
	fmt.Fprintln(tw, "Package\tModule\tVersion\tVia (at least...)")
	for _, data := range out {
		for idx := range data {
			if idx > 0 {
				fmt.Fprint(tw, "\t")
			}
			fmt.Fprint(tw, data[idx])
		}
		fmt.Fprintln(tw)
	}
	tw.Flush()

	return nil
}

func crawl(pkg, via *packages.Package, seen map[string]bool, out [][]string) [][]string {
	if seen[pkg.ID] {
		return out
	}
	seen[pkg.ID] = true

	if mod := pkg.Module; mod != nil {
		line := []string{pkg.ID, pkg.Module.Path, pkg.Module.Version}
		if pkg == via {
			line = append(line, "Main")
		} else {
			line = append(line, via.ID)
		}
		out = append(out, line)
	}
	for _, i := range pkg.Imports {
		out = crawl(i, pkg, seen, out)
	}
	return out
}
