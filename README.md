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

## Arguments

The binary needs at least 3 arguments.

1. The component name without the leading carret e.g. `sw-entity-single-select`
2. The model prop name e.g. `value`
3. The previously emitted even name e.g. `change`
4. Optional: Specify the working directory, where the binary searches for `.twig` files

If the fourth parameter is left out the current working dir is taken instead.

```shell
$ cd <shopwareRoot>/src/Administration/Resources/app/administration
$ v3r sw-component-name-without modelName eventName
```

## Check the results

This program is rather simple and relies on regular expressions internally. At least do a visual check of the result.