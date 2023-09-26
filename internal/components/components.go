package components

type Component struct {
	Name  string
	Model string
	Event string
}

func GetAutoFixableComponents() []Component {
	return []Component{
		// src/app/component/form
		{
			Name:  "sw-url-field",
			Model: "value",
			Event: "input",
		},
		{
			Name:  "sw-textarea-field",
			Model: "value",
			Event: "input",
		},
		{
			Name:  "sw-text-field",
			Model: "value",
			Event: "input",
		},
		{
			Name:  "sw-text-editor",
			Model: "value",
			Event: "input",
		},
		{
			Name:  "sw-tagged-field",
			Model: "value",
			Event: "change",
		},
		{
			Name:  "sw-switch-field",
			Model: "value",
			Event: "change",
		},
		{
			Name:  "sw-select-number-field",
			Model: "value",
			Event: "change",
		},
		{
			Name:  "sw-select-field",
			Model: "value",
			Event: "change",
		},
		{
			Name:  "sw-radio-field",
			Model: "value",
			Event: "change",
		},
		{
			Name:  "sw-purchase-price-field",
			Model: "price",
			Event: "input",
		},
		{
			Name:  "sw-price-field",
			Model: "price",
			Event: "change",
		},
		{
			Name:  "sw-password-field",
			Model: "value",
			Event: "input",
		},
		{
			Name:  "sw-number-field",
			Model: "value",
			Event: "input",
		},
		{
			Name:  "sw-gtc-checkbox",
			Model: "value",
			Event: "change",
		},
		{
			Name:  "sw-form-field-renderer",
			Model: "value",
			Event: "input",
		},
		{
			Name:  "sw-file-input",
			Model: "value",
			Event: "change",
		},
		{
			Name:  "sw-email-field",
			Model: "value",
			Event: "input",
		},
		{
			Name:  "sw-dynamic-url-field",
			Model: "value",
			Event: "input",
		},
		{
			Name:  "sw-confirm-field",
			Model: "value",
			Event: "input",
		},
		{
			Name:  "sw-compact-colorpicker",
			Model: "value",
			Event: "input",
		},
		{
			Name:  "sw-colorpicker",
			Model: "value",
			Event: "input",
		},
		{
			Name:  "sw-checkbox-field",
			Model: "value",
			Event: "change",
		},
		{
			Name:  "sw-boolean-radio-group",
			Model: "value",
			Event: "change",
		},
		{
			Name:  "sw-single-select",
			Model: "value",
			Event: "change",
		},
		{
			Name:  "sw-multi-tag-select",
			Model: "value",
			Event: "change",
		},
		{
			Name:  "sw-multi-tag-ip-select",
			Model: "value",
			Event: "change",
		},
		{
			Name:  "sw-multi-select",
			Model: "value",
			Event: "change",
		},
		{
			Name:  "sw-grouped-single-select",
			Model: "value",
			Event: "change",
		},
		{
			Name:  "sw-entity-single-select",
			Model: "value",
			Event: "change",
		},
		{
			Name:  "sw-entity-multi-select",
			Model: "entityCollection",
			Event: "change",
		},
		{
			Name:  "sw-entity-tag-select",
			Model: "entityCollection",
			Event: "change",
		},
		{
			Name:  "sw-entity-multi-id-select",
			Model: "ids",
			Event: "change",
		},
		{
			Name:  "sw-entity-many-to-many-select",
			Model: "entityCollection",
			Event: "change",
		},
		// src/app/component/base
		{
			Name:  "sw-button-process",
			Model: "processSuccess", // There should be no v-model for this but the event needs to be changed
			Event: "process-finish",
		},
		{
			Name:  "sw-simple-search-field",
			Model: "value",
			Event: "input",
		},
		// src/app/component/utils
		{
			Name:  "sw-inherit-wrapper",
			Model: "value",
			Event: "input",
		},
	}
}
