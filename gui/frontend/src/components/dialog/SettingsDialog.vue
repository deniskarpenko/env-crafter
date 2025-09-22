<template>
  <dialog ref="dialogRef" class="settings-dialog" @close="handleClose">
    <form method="dialog" class="dialog-form">
      <header  class="dialog-header">
        <h2>{{ title }}</h2>
        <button  @click="handleClose" class="close-button">x</button>
      </header>
      <main class="dialog-content">
        <tabs :tabs="tabsConfig">
          <template #ports>
            <div><span>Ports</span></div>
            <input-plus></input-plus>
          </template>
          <template #volumes>???</template>
        </tabs>
      </main>
    </form>
  </dialog>
</template>
<script setup lang="ts">

import {nextTick, ref, watch} from "vue";
import Tabs from "../Tabs.vue";
import InputPlus from "../InputPlus.vue";

const props = withDefaults(defineProps<{
  title: string;
  modelValue: boolean;
}>(), {
  title: "",
  modelValue: false,
});
const emit = defineEmits<{
  (event: 'close' | 'handleClose'): void;
}>();

const tabsConfig = [
  { id: 'ports', title: 'Ports', slot: 'ports' },
  { id: "volumes", title: 'Volumes', slot: 'volumes' },
  { id: 'envs', title: 'Result', slot: 'result' },
]

const dialogRef = ref<null | HTMLDialogElement >(null);

const openDialog = async () => {
  await nextTick();
  if (dialogRef.value && !dialogRef.value.open) {
    dialogRef.value.showModal();
  }
}

const closeDialog = async () => {
  await nextTick();
  if (dialogRef.value && !dialogRef.value.open) {
    dialogRef.value.close();
  }
}

const handleClose = () => {
  emit("close");
}

watch(() => props.modelValue, (newValue: boolean) => {
  if (newValue) {
    openDialog();
    return;
  }

  closeDialog();
});
</script>
<style scoped>
.settings-dialog {
  border-radius: 12px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
  padding: 0;
  max-width: 480px;
  width: 90vw;
  max-height: 80vh;
  overflow: hidden;
  background: #bfbaba;
  border: #cccccc;
}

.settings-dialog::backdrop {
  background: #656262;;
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
  flex-shrink: 0;
}

.dialog-header h2 {
  margin: 0;
  color: #1f2937;
  font-size: 20px;
  font-weight: 600;
}

.dialog-content {
  flex: 1;
  padding: 24px;
  overflow-y: auto;
}


</style>
