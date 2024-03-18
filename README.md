# walkdir-example

walkdir を使ってみるサンプル。
このプログラムでは、 srcDir 以下を dstDir 以下にコピーし、その際にコピー元で json ファイルだったものはコピー先で yaml ファイルになっているような関数 `convertAllJsonFilesToYaml`が、walkdir-example.go に定義されている。

# 使い方

`go run .` をすると sampleInputRoot 以下のファイルが処理されて、generatedOutputRoot ディレクトリ以下にファイルが生成される。
`go test` でテストがされる。
