package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

func convertRoot(path string, srcRoot string, dstRoot string) string {
	relativePath, err := filepath.Rel(srcRoot, path)
	if err != nil {
		return path
	}

	return filepath.Join(dstRoot, relativePath)
}

func convertJSONToYAML(jsonBytes []byte) ([]byte, error) {
	var jsonObj interface{}
	err := json.Unmarshal(jsonBytes, &jsonObj)
	if err != nil {
		return nil, err
	}

	yamlBytes, err := yaml.Marshal(jsonObj)
	if err != nil {
		return nil, err
	}

	return yamlBytes, nil
}


func genConvertJSONFileToYamlWorkDirFunc(convertPath func(string) string) fs.WalkDirFunc {
	return func(path string, d os.DirEntry, err error) error {
		dstPath := convertPath(path)

		wrapErr := func(err error) error {
			if err == nil {
				return nil
			} else {
				return fmt.Errorf("error processing %s: %w", path, err)
			}
		}

		if err != nil {
			return wrapErr(err)
		}

		if d.IsDir() {
			err := os.MkdirAll(dstPath, 0755)
			return wrapErr(err)
		}

		fileBytes, err := os.ReadFile(path)
		if err != nil {
			return wrapErr(err)
		}

		if strings.HasSuffix(d.Name(), ".json") {
			dstYamlPath := strings.TrimSuffix(dstPath, ".json") + ".yaml"

			yamlBytes, err := convertJSONToYAML(fileBytes)
			if err != nil {
				return wrapErr(err)
			}

			err = os.WriteFile(dstYamlPath, yamlBytes, 0644)
			return wrapErr(err)
		} else {
			err := os.WriteFile(dstPath, fileBytes, 0644)
			return wrapErr(err)
		}
	}
}


func convertAllJsonFilesToYaml(inputRoot string, outputRoot string) error {
	convertPath := func(path string) string {
		return convertRoot(path, inputRoot, outputRoot)
	}

	err := os.RemoveAll(outputRoot)
	if err != nil {
		return fmt.Errorf("error cleaning up output directory: %w", err)
	}

	fileSystem := os.DirFS(".")

	err = fs.WalkDir(fileSystem, inputRoot, genConvertJSONFileToYamlWorkDirFunc(convertPath))
  
	return err

}
