package processors

import "regexp"

type eventProcessor struct{}

func (p *eventProcessor) Process(options *ProcessorOptions) string {
	// https://regex101.com/r/VkXBIA/1
	eventRegex := regexp.MustCompile(`(<` + options.ComponentName + `\n( *)[a-zA-Z0-9.\-+*#:_$(')@!=>"{%}\[\]? \n|/` + "`" + `&]*)@` + options.EventName + `="(.*)"`)
	eventFixed := eventRegex.ReplaceAllString(options.Text, "${1}{% if VUE3 %}\n${2}@update:"+options.ModelName+"=\"${3}\"\n${2}{% else %}\n${2}@"+options.EventName+"=\"${3}\"\n${2}{% endif %}")

	return eventFixed
}

func NewEventProcessor() Processor {
	return &eventProcessor{}
}
