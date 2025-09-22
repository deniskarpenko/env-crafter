<template>
<div v-for="(input, index) in inputs" :key="input.id" class="inputs-wrapper">
<input
    v-model="input.value"
    type="text"
    class="input-field"
/>
  <button v-if="index > 0" type="button">x</button>
</div>
  <button @click="addInput" type="button">+</button>
</template>
<script setup lang="ts">

import {ref} from "vue";

interface InputItem {
  id: number;
  value: string;
  isValid: boolean;
  errorMessage: string;
}

const inputs = ref<InputItem[]>([
  {id: 0, value:'', isValid: true, errorMessage: ''}
]);

const nextId = ref(1);

const addInput = () => {
  console.log("addInput");
  inputs.value.push({id: nextId.value, value:'', isValid: true, errorMessage: ''});
  nextId.value++;
};

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
</style>