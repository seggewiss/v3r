package processors

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/seggewiss/vue3-template-refactor/internal/components"
	"github.com/seggewiss/vue3-template-refactor/internal/options"
)

type ProcessorOptions struct {
	ComponentName string
	ModelName     string
	EventName     string
	Text          string
}

type Processor interface {
	Process(options *ProcessorOptions) string
}

func getAll() []Processor {
	return []Processor{
		NewModelProcessor(),
		NewEventProcessor(),
	}
}

func Run() {
	options := options.GetOptions()
	if options.ComponentName == "" && !options.Auto {
		fmt.Println("Please specify a component name or use the auto flag")
		return
	}

	var unprocessedFiles []string

	if options.Auto {
		fmt.Println("Running v3r in auto mode")
		components := components.GetAutoFixableComponents()
		for _, component := range components {
			fmt.Println("Processing component: " + component.Name)
			options.ComponentName = component.Name
			options.ModelName = component.Model
			options.EventName = component.Event

			processFiles(&options, unprocessedFiles)
		}
	} else {
		fmt.Println("Running v3r in manual mode")
		fmt.Println("Processing component: " + options.ComponentName)
		processFiles(&options, unprocessedFiles)
	}

	if len(unprocessedFiles) > 0 {
		fmt.Println("")
		fmt.Println("These files contained the component but where unable to convert:")
		for _, file := range unprocessedFiles {
			fmt.Println(filepath.Base(file))
		}
	}

	fmt.Println("")
	fmt.Println("All done!")
}

func processFiles(options *options.Options, unprocessedFiles []string) {
	filepath.WalkDir(options.WorkingDirectory, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) != options.FileExtension {
			return nil
		}

		b, err := os.ReadFile(s)
		if err != nil {
			panic(err)
		}

		fileContent := string(b)
		if !strings.Contains(fileContent, "<"+options.ComponentName) {
			return nil
		}

		var processorOptions = new(ProcessorOptions)
		processorOptions.ComponentName = options.ComponentName
		processorOptions.ModelName = options.ModelName
		processorOptions.EventName = options.EventName
		processorOptions.Text = fileContent

		processedText := ""
		processors := getAll()
		for _, processor := range processors {
			processedText = processor.Process(processorOptions)
			processorOptions.Text = processedText
		}

		if fileContent == processedText {
			unprocessedFiles = append(unprocessedFiles, s)
			return nil
		}

		err = os.WriteFile(s, []byte(processedText), 0644)
		if err != nil {
			panic(err)
		}

		return nil
	})
}
