<template>
  <div class="tabs">
    <div class="tab-headers">
      <button
          v-for="(tab, index) in tabs"
          :key="tab.id"
          :class="{ active: activeTab === index }"
          @click="setActiveTab(index)"
          :disabled="tab.disabled"
      >
        {{ tab.title }}
      </button>
    </div>
    <div class="tab-content">
      <slot :name="tabs[activeTab].slot" />
    </div>
  </div>
</template>
<script setup lang="ts">
import {Tab} from "../types/Tabs/Tab";
import {computed, ref} from "vue";

interface Props {
  tabs: Tab[],
  defaultTab?: number
}

interface Emits {
  tabChanged: [index: number, tab: Tab]
}

const props = withDefaults(defineProps<Props>(), {
  defaultTab: 0
});

const emit = defineEmits<Emits>();

const activeTab = ref(props.defaultTab);

const setActiveTab = (index: number) => {
  if (props.tabs[index]?.disabled) return
  activeTab.value = index
  emit('tabChanged', index, props.tabs[index])
}

const currentTab = computed(() => props.tabs[activeTab.value])
</script>
<style></style>