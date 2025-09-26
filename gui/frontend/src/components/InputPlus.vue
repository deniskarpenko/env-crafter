<template>
<div v-for="(input, index) in inputs" :key="input.id" class="inputs-wrapper">
  <input
      v-if="type !== 'file'"
    v-model="input.value"
    :type="type"
    :class="['input-field', { 'input-error': !input.isValid && input.value !== '' }]"
    @blur="updateInputValue(index)"
/>
  <input
      v-else
      v-model="input.value"
      :type="type"
      :class="['input-field', { 'input-error': !input.isValid && input.value !== '' }]"
      @blur="updateInputValue(index)"
      accept=".env"
      multiple
  />
  <button v-if="index > 0" type="button" @click="removeInput(index)">x</button>
  <div v-if="!input.isValid && input.value !== ''" class="error-message">
    {{ input.errorMessage }}
  </div>
</div>
  <button type="button" @click="addInput" >+</button>
</template>
<script setup lang="ts">
import {ref} from "vue";
import {InputItem, InputType} from "../types/InputItem";

const props = defineProps<{
  type: InputType;
  errorMessage?: string;
  validateInput?: (value: string) => boolean;
}>();

const emit = defineEmits<{
  (event: 'updateInputs', values: string[]): void;
}>();

const inputs = ref<InputItem[]>([
  {id: 0, value:'', isValid: true, errorMessage: ''}
]);

const nextId = ref(1);

const addInput = () => {
  inputs.value.push({id: nextId.value, value:'', isValid: true, errorMessage: ''});
  nextId.value++;
};

const removeInput = (index: number): void => {
  inputs.value.splice(index, 1);
}

const updateInputValue = (index: number): void => {
  const input = inputs.value[index];

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

  emit('updateInputs', inputs.value
      .filter(input => input.value.trim() !== '')
      .map((input: InputItem) => input.value)
  );
}

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