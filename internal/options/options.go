package options

import (
	"flag"
	"os"
)

type Options struct {
	Auto             bool
	WorkingDirectory string
	ComponentName    string
	ModelName        string
	EventName        string
	FileExtension    string
}

func GetOptions() Options {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	autoPtr := flag.Bool("auto", false, "Automatically fix all base components. Overrides -component, -model and -event!")
	componentNamePtr := flag.String("component", "", "The name of the component to refactor")
	modelNamePtr := flag.String("model", "value", "The name of the model to use")
	eventNamePtr := flag.String("event", "input", "The name of the event to use")
	direcotryPtr := flag.String("directory", cwd, "The directory to search for components")
	extensionPtr := flag.String("extension", ".twig", "The file extension to search for")

	flag.Parse()

	return Options{
		Auto:             *autoPtr,
		WorkingDirectory: *direcotryPtr,
		ComponentName:    *componentNamePtr,
		ModelName:        *modelNamePtr,
		EventName:        *eventNamePtr,
		FileExtension:    *extensionPtr,
	}
}
