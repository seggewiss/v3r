package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Find all html.twig files using the component
func findTemplatesUsingComponent(root, ext, componentName string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) != ext {
			return nil
		}

		b, err := os.ReadFile(s)
		if err != nil {
			panic(err)
		}

		fileContent := string(b)
		if !strings.Contains(fileContent, "<"+componentName) {
			return nil
		}

		a = append(a, s)
		return nil
	})
	return a
}

func main() {
	if len(os.Args) < 3 {
		panic("Please provide a component and model name!")
	}

	var path string
	if len(os.Args) == 5 {
		path = os.Args[4]
	} else {
		path, _ = os.Getwd()
	}
	fmt.Println(path)

	componentName := os.Args[1]
	modelName := os.Args[2]
	eventName := os.Args[3]
	files := findTemplatesUsingComponent(path, ".twig", componentName)

	fmt.Println("Vue3 template refactoring \"" + componentName + "\" with model \"" + modelName + "\" and event \"" + eventName + "\"")
	fmt.Println("")

	for _, file := range files {
		fmt.Printf("Processing %s ...\n", filepath.Base(file))
		b, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}

		fileContent := string(b)
		// https://regex101.com/r/Q3VTcx/1
		modelRegex := regexp.MustCompile(`(<` + componentName + `\n( *)[a-zA-Z0-9.\-:_$(')@!=>"{%}\[\]? \n]*)v-model="(.*)"`)
		modelFixed := modelRegex.ReplaceAllString(fileContent, "${1}{% if VUE3 %}\n${2}v-model:"+modelName+"=\"$3\"\n$2{% else %}\n${2}v-model=\"${3}\"\n${2}{% endif %}")

		// https://regex101.com/r/VkXBIA/1
		eventRegex := regexp.MustCompile(`(<` + componentName + `\n( *)[a-zA-Z0-9.\-:_$(')@!=>"{%}\[\]? \n]*)@` + eventName + `="(.*)"`)
		eventFixed := eventRegex.ReplaceAllString(modelFixed, "${1}{% if VUE3 %}\n${2}@update:"+modelName+"=\"${3}\"\n${2}{% else %}\n${2}@"+eventName+"=\"${3}\"\n${2}{% endif %}")

		err = os.WriteFile(file, []byte(eventFixed), 0644)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("")
	fmt.Println("All done!")
}
