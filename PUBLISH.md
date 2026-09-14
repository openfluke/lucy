# Publishing Lucy 1.0+

Channels share one version (`VERSION`). **v1.0.0** is the stable portable cut.

Publish **Go tag first**, then npm, then PyPI.

## Go module

```bash
git tag v1.0.0
git push origin v1.0.0
go get github.com/openfluke/lucy@v1.0.0
```

Local monorepo replace (Tide):

```
require github.com/openfluke/lucy v1.0.0
replace github.com/openfluke/lucy => ../lucy/go
```

## npm `@openfluke/lucy`

```bash
./go/scripts/build-artifacts.sh
cd js && npm publish --access public
```

## PyPI `openfluke-lucy`

```bash
./go/scripts/build-artifacts.sh
cd python && python -m build && twine upload dist/*
```

Check alignment: `./scripts/publish-check.sh`  
Portable claim: [`PORTABLE.md`](PORTABLE.md)

## GitHub release (tag + notes + binaries)

```bash
./scripts/release.sh           # dry-run: checks, builds, example smoke, prints notes
./scripts/release.sh --push    # git tag vX.Y.Z + gh release + upload artifacts
```

Release body includes **what Lucy is/isn’t**, quickstart/example pointers, and
attaches linux/darwin/windows binaries + `lucy.wasm`.

