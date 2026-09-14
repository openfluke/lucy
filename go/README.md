# go/ — Lucy measuring core

**Version:** `0.5.0`

```go
pack := lucy.BuildBoardCharts(board, 8)
_ = pack.ConsciousnessJPG
_ = lucy.WriteReportDir("out", board, 8)
http.ListenAndServe(":7474", lucy.Handler())
```

CLI: `build-lpd` · `chart-*[-png|-jpg]` · `chart-pack` · `report` · `serve`
