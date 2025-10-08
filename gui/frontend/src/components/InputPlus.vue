<template>
<div v-for="(input, index) in localInputs" :key="input.id" class="inputs-wrapper">
  <input
      v-if="type !== 'file'"
    v-model="input.value"
    :type="type"
    :class="['input-field', { 'input-error': !input.isValid && input.value !== '' }]"
    @input="updateInputValue(index)"
/>
  <input
      v-else
      :type="type"
      class="input-field"
      @input="updateInputValue(index)"
      accept=".env"
      multiple
      :ref="(el) => fileInputRefs[index] = el as HTMLInputElement | null"
  />
  <button v-if="index > 0" type="button" @click="removeInput(index)">x</button>
  <div v-if="!input.isValid && input.value !== ''" class="error-message">
    {{ input.errorMessage }}
  </div>
</div>
  <button type="button" @click="addInput" >+</button>
</template>
<script setup lang="ts">
import {onMounted, ref} from "vue";
import {InputItem, InputType} from "../types/InputItem";

interface Props {
  type: InputType;
  errorMessage?: string;
  validateInput?: (value: string) => boolean;
  inputs?: string[];
}

const props = withDefaults(defineProps<Props>(),{
  inputs: () => []
});

const emit = defineEmits<{
  (event: 'updateInputs', values: string[]): void;
}>();

const localInputs = ref<InputItem[]>([]);
const fileInputRefs = ref<(HTMLInputElement | null)[]>([]);

const nextId = ref(
    Math.max(...localInputs.value.map(item => item.id), 0) + 1
);

const addInput = () => {
  localInputs.value.push({id: nextId.value, value:'', isValid: true, errorMessage: ''});
  nextId.value++;
};

const removeInput = (index: number): void => {
  localInputs.value.splice(index, 1);
}

const updateInputValue = (index: number): void => {
  const input = localInputs.value[index];

  if (
      input === undefined ||
      props.validateInput === undefined ||
      props.errorMessage === undefined
  ) {
    return;
  }

  input.isValid = props.validateInput(input.value);
  input.errorMessage = props.errorMessage;

  if (!input.isValid) {
    return;
  }

  emit('updateInputs', localInputs.value
      .filter(input => input.value.trim() !== '')
      .map((input: InputItem) => input.value)
  );
}

const handleFileChanged = (files: File[]) => {}

onMounted(() => {
  if (props.inputs.length === 0) {
    localInputs.value.push({id: nextId.value, value: '', isValid: false, errorMessage: props.errorMessage});
    return;
  }

  props.inputs.forEach(input => {
    const isValid = props.validateInput !== undefined ? props.validateInput(input) : true;
    localInputs.value.push({id: nextId.value, value: input, isValid: isValid, errorMessage: props.errorMessage});
  });
})

</script>
<style scoped>
.inputs-wrapper {
  display: flex;
  padding: 10px;
}
.input-field {
  flex: 1;
  padding: 10px;
  border: 2px solid #e1e5e9;
  border-radius: 6px;
  font-size: 14px;
  transition: border-color 0.2s ease;
  background: transparent;
}

.input-field:focus {
  outline: none;
  border-color: #007bff;
}

.input-field.input-error {
  border-color: #dc3545;
}

.error-message {
  color: #dc3545;
  font-size: 12px;
  margin-top: 4px;
  margin-bottom: 8px;
}
</style>