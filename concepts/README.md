Install this module:

```sh
go install concepts
```

1. This command builds the concepts program, producing an executable binary. It then installs that binary as $HOME/go/bin/concepts

2. The install directory is controlled by the `GOPATH` and `GOBIN`
    - If GOBIN is set, binaries are installed to that directory.
    - If GOPATH is set, binaries are installed to the bin subdirectory of the first directory in the GOPATH list. 
    - Otherwise, binaries are installed to the bin subdirectory of the default GOPATH ($HOME/go or %USERPROFILE%\go)
