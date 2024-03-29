# Golang Package Version Dump

[![Tests](https://github.com/bobvawter/go-pvd/actions/workflows/tests.yaml/badge.svg)](https://github.com/bobvawter/go-pvd/actions/workflows/tests.yaml)

`go-pvd` crawls the transitive dependency graph of one or more Golang packages to report the module
versions used. This is helpful if you are trying to determine if a CVE in some package imported from
a module actually impacts your project.

Sample output for pvd itself:

```
% go-pvd .
Module                       Version  Imported package                               Via (at least...)
github.com/bobvawter/go-pvd           github.com/bobvawter/go-pvd                    Main
                                      github.com/bobvawter/go-pvd/pkg/tool           github.com/bobvawter/go-pvd
github.com/spf13/cobra       v1.7.0   github.com/spf13/cobra                         github.com/bobvawter/go-pvd/pkg/tool
github.com/spf13/pflag       v1.0.5   github.com/spf13/pflag                         github.com/spf13/cobra
golang.org/x/mod             v0.12.0  golang.org/x/mod/semver                        golang.org/x/tools/internal/gocommand
golang.org/x/sys             v0.10.0  golang.org/x/sys/execabs                       golang.org/x/tools/go/packages
golang.org/x/tools           v0.11.0  golang.org/x/tools/go/gcexportdata             golang.org/x/tools/go/packages
                                      golang.org/x/tools/go/internal/packagesdriver  golang.org/x/tools/go/packages
                                      golang.org/x/tools/go/packages                 github.com/bobvawter/go-pvd/pkg/tool
                                      golang.org/x/tools/internal/event              golang.org/x/tools/internal/gocommand
                                      golang.org/x/tools/internal/event/core         golang.org/x/tools/internal/event
                                      golang.org/x/tools/internal/event/keys         golang.org/x/tools/internal/gocommand
                                      golang.org/x/tools/internal/event/label        golang.org/x/tools/internal/event/keys
                                      golang.org/x/tools/internal/event/tag          golang.org/x/tools/internal/gocommand
                                      golang.org/x/tools/internal/gcimporter         golang.org/x/tools/go/gcexportdata
                                      golang.org/x/tools/internal/gocommand          golang.org/x/tools/internal/packagesinternal
                                      golang.org/x/tools/internal/packagesinternal   golang.org/x/tools/go/packages
                                      golang.org/x/tools/internal/pkgbits            golang.org/x/tools/internal/gcimporter
                                      golang.org/x/tools/internal/tokeninternal      golang.org/x/tools/internal/gcimporter
                                      golang.org/x/tools/internal/typeparams         golang.org/x/tools/go/packages
                                      golang.org/x/tools/internal/typesinternal      golang.org/x/tools/go/packages
```

# Installing

* You'll need to [download and install](https://golang.org/doc/install) go
* `go install github.com/bobvawter/go-pvd@latest`
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

The package pattern is anything accepted by the usual golang package parser, so `./pkg/foo` will
report on packages reachable exactly from foo (relative to the `--dir` flag), and `./pkg/foo/...`
would include `foo` and all sub-packages.
