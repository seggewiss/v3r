# Warning

This program will change all files with the specified extension. This software is best used with a GIT repository.
Make sure you don't have any uncomitted and pushed changes. Then and only then run the software.

# About this project
This tool will refactor component usages in templates according to Vue 3.
If your component uses `model` to define different props as the default model and change events, this is no longer supported with Vue 3.

Take the `sw-entity-single-select` for example it defines `value` as the prop and `change` as the event. Default would be `value` and `input`. Behind the `VUE3` feature flag the new event is `update:value` this goes with the Vue 3 standard.

So this programm will change all component usages in templates as follows.

## Example

From this Vue 2 template:
```twig
<sw-entity-single-select
    ref="selectedField"
    v-model="selectedField"
    entity="custom_field"
    :criteria="customFieldCriteria"
    :placeholder="$tc('...')"
    :disabled="disabled"
    size="medium"
    show-clearable-button
    @change="onFieldChange"
>
```

To this bridge template:

```twig
<sw-entity-single-select
    ref="selectedField"
    {% if VUE3 %}
    v-model:value="selectedField"
    {% else %}
    v-model="selectedField"
    {% endif %}
    entity="custom_field"
    :criteria="customFieldCriteria"
    :placeholder="$tc('...')"
    :disabled="disabled"
    size="medium"
    show-clearable-button
    {% if VUE3 %}
    @update:value="onFieldChange"
    {% else %}
    @change="onFieldChange"
    {% endif %}
>
```

# Why are there multiple if statements

We need to keep in mind that this template gets reduced to a Vue 2 template again before linting by our `Vue twig plugin`.
If we combine the event listeners with v-models this would trigger some eslint rules about the attribute occurence not being in order.

# Usage
It's recommended that you add the released binary for your os to your `$PATH`.

## Flags

| Flag       | Type    | Default value       | Description                                                                           |
| ---------- | ------- | ------------------- | ------------------------------------------------------------------------------------- |
| -auto      | boolean | false               | Automatically fix all base components. Overrides `-component`, `-model` and `-event`! |
| -component | string  | ""                  | The component you want to fix up cross file                                           |
| -model     | string  | "value"             | The model used by the component                                                       |
| -event     | string  | "input"             | The event name used to update the model                                               |
| -directory | string  | Current working dir | The directory being searched                                                          |
| -extension | string  | ".twig"             | The extension of files being searched for components                                  |
| -help      | boolean | false               | Displays the above list excluding -help                                               |

### Examples

```shell
$ cd <shopwareRoot>/src/Administration/Resources/app/administration
$ v3r -component=sw-component-name-without -model=modelName -event=eventName
$ v3r -auto
```

## Check the results

This program is rather simple and relies on regular expressions internally. At least do a visual check of the result.