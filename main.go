package main

func main() {
    inputRoot := "sampleInputRoot"
    outputRoot := "generatedOutputRoot" 

	err := convertAllJsonFilesToYaml(inputRoot, outputRoot)
	if err != nil {
		panic(err)
	}
}
