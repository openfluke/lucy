# Publishing Lucy 0.8+

Channels share one version (`VERSION`). Publish **Go tag first**, then npm, then PyPI.

## Go module

```bash
git tag v0.8.0
git push origin v0.8.0
# hosts: go get github.com/openfluke/lucy@v0.8.0
```

Local Tide replace (dev):

```
replace github.com/openfluke/lucy => ../lucy/go
```

## npm `@openfluke/lucy`

```bash
./go/scripts/build-artifacts.sh
cd js && npm publish --access public
# dry-run: npm publish --dry-run
```

## PyPI `openfluke-lucy`

```bash
./go/scripts/build-artifacts.sh
cd python && python -m build && twine upload dist/*
# dry-run: twine upload --repository testpypi dist/*
```

Helper: `./scripts/publish-check.sh` (version alignment + dry-runs when tools exist).
