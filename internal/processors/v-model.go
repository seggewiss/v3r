package processors

import "regexp"

type vModelProcessor struct{}

func NewModelProcessor() Processor {
	return &vModelProcessor{}
}

func (p *vModelProcessor) Process(options *ProcessorOptions) string {
	// https://regex101.com/r/Q3VTcx/1
	modelRegex := regexp.MustCompile(`(<` + options.ComponentName + `\n( *)[a-zA-Z0-9.\-+*#:,_$(')@!=>"{%}\[\]? \n|/` + "`" + `&]*)v-model="(.*)"`)
	modelFixed := modelRegex.ReplaceAllString(options.Text, "${1}{% if VUE3 %}\n${2}v-model:"+options.ModelName+"=\"$3\"\n$2{% else %}\n${2}v-model=\"${3}\"\n${2}{% endif %}")

	return modelFixed
}
