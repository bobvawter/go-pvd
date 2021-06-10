# Golang Package Version Dump

This utility will crawl the transitive dependency graph of one or more
Golang packages to report the module versions used.

Sample output for pvd itself:

```
% go-pvd .
Package                                        Module                       Version                             Via (at least...)
github.com/bobvawter/go-pvd                    github.com/bobvawter/go-pvd                                      Main
github.com/bobvawter/go-pvd/pkg/tool           github.com/bobvawter/go-pvd                                      github.com/bobvawter/go-pvd
github.com/spf13/cobra                         github.com/spf13/cobra       v1.1.3                              github.com/bobvawter/go-pvd/pkg/tool
github.com/spf13/pflag                         github.com/spf13/pflag       v1.0.5                              github.com/spf13/cobra
golang.org/x/mod/semver                        golang.org/x/mod             v0.4.2                              golang.org/x/tools/internal/gocommand
golang.org/x/sys/execabs                       golang.org/x/sys             v0.0.0-20210510120138-977fb7262007  golang.org/x/tools/internal/gocommand
golang.org/x/tools/go/gcexportdata             golang.org/x/tools           v0.1.2                              golang.org/x/tools/go/packages
golang.org/x/tools/go/internal/gcimporter      golang.org/x/tools           v0.1.2                              golang.org/x/tools/go/gcexportdata
golang.org/x/tools/go/internal/packagesdriver  golang.org/x/tools           v0.1.2                              golang.org/x/tools/go/packages
golang.org/x/tools/go/packages                 golang.org/x/tools           v0.1.2                              github.com/bobvawter/go-pvd/pkg/tool
golang.org/x/tools/internal/event              golang.org/x/tools           v0.1.2                              golang.org/x/tools/internal/gocommand
golang.org/x/tools/internal/event/core         golang.org/x/tools           v0.1.2                              golang.org/x/tools/internal/event
golang.org/x/tools/internal/event/keys         golang.org/x/tools           v0.1.2                              golang.org/x/tools/internal/event/core
golang.org/x/tools/internal/event/label        golang.org/x/tools           v0.1.2                              golang.org/x/tools/internal/event
golang.org/x/tools/internal/gocommand          golang.org/x/tools           v0.1.2                              golang.org/x/tools/internal/packagesinternal
golang.org/x/tools/internal/packagesinternal   golang.org/x/tools           v0.1.2                              golang.org/x/tools/go/packages
golang.org/x/tools/internal/typesinternal      golang.org/x/tools           v0.1.2                              golang.org/x/tools/go/packages
golang.org/x/xerrors                           golang.org/x/xerrors         v0.0.0-20200804184101-5ec99f83aff1  golang.org/x/tools/go/packages
golang.org/x/xerrors/internal                  golang.org/x/xerrors         v0.0.0-20200804184101-5ec99f83aff1  golang.org/x/xerrors
```

# Installing

* You'll need to [download and install](https://golang.org/doc/install) go
* `go install github.com/bobvawter/go-pvd`
* Ensure that `$HOME/go/bin` is in your `$PATH`

# Usage

```
Usage:
  go-pvd <go package pattern> ... [flags]

Flags:
  -b, --build stringArray   arguments to pass to the golang build tool (default [-mod=mod])
  -d, --dir string          the source directory (default ".")
  -h, --help                help for go-pvd
  -t, --tests               include test code
```

The package pattern is anything accepted by the usual golang package parser, so `./pkg/foo` will report on packages reachable exactly from foo (relative to the `--dir` flag), and `./pkg/foo/...` would include `foo` and all sub-packages.
