###  Basics

---

We can create a different binary for a different OS and architecture using the GOOS and GOARCH environment variables:
```bash
GOOS=windows GOARCH=amd64 go build hello.
```

### Go worskspaces

When you install packages with `go install` they are installed in `$HOME/go` by default. You can change the `GOPATH` environment variable to isolate the libs you use for different projects
