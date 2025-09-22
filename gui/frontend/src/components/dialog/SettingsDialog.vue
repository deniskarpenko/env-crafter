<template>
  <dialog ref="dialogRef" class="settings-dialog" @close="handleClose">
    <form method="dialog" class="dialog-form">
      <header  class="dialog-header">
        <h2>{{ title }}</h2>
        <button  @click="handleClose" class="close-button">x</button>/
      </header>
    </form>
  </dialog>
</template>
<script setup lang="ts">

import {ref} from "vue";

const props = withDefaults(defineProps<{
  title: string;
}>(), {
  title: "",
});
const emit = defineEmits<{
  (event: 'close' | 'handleClose'): void;
}>();

const dialogRef = ref<null | HTMLElement>(null);

const handleClose = () => {
  emit("close");
}
</script>
<style scoped>
.settings-dialog {
  border: none;
  border-radius: 12px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
  padding: 0;
  max-width: 480px;
  width: 90vw;
  max-height: 80vh;
  overflow: hidden;
}

.settings-dialog::backdrop {
  background: rgba(0, 0, 0, 0.6);
  backdrop-filter: blur(4px);
}

.dialog-form {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.dialog-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 24px;
  border-bottom: 1px solid #e5e7eb;
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%);
  flex-shrink: 0;
}

.dialog-header h2 {
  margin: 0;
  color: #1f2937;
  font-size: 20px;
  font-weight: 600;
}


</style>
