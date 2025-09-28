<template>
  <div class="wrapper">
    <Tabs :tabs="tabsConfig">
      <template #images>
        <div  v-for="config in imageRowConfig">
          <images-row
              :key="config.title"
              :title="config.title"
              :images="config.images"
              :tags="config.tags"
              :type="config.type"
              @update:model-value="handleImageRowUpdate(config, $event)"
              @click:settings="showSettings"
          ></images-row>
          <settings-dialog
              title="Settings"
              :model-value="isShowSettingsDialog"
              @close="handleClose"
          ></settings-dialog>
        </div>
      </template>
      <template #result>
        <div>!!!!</div>
      </template>
    </Tabs>
    <div class="buttons-containers">
      <button>Build</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import {reactive, ref, onMounted, nextTick} from "vue";
import {GetAllImages} from "../../wailsjs/go/main/App";
import {Application, ContainerConfig, ImageWithTagConfig} from "../types/Application";
import {ImageOption} from "../types/ImageOption";
import Tabs from "../components/Tabs.vue";
import ImagesRow from "../components/ImagesRow.vue";
import {useImageManagers} from "../composables/ImageManager";
import SettingsDialog from "../components/dialog/SettingsDialog.vue";
import {ImageWithTag} from "../types/ImageWithTag";
import {ImageTypes} from "../types/ImageTypes";

const tabsConfig = [
  { id: 'images', title: 'Images', slot: 'images' },
  { id: 'result', title: 'Result', slot: 'result' },
  { id: 'settings', title: 'Settings', slot: 'settings' },
];

const appModel = reactive<Application>({
  backend: null,
  sql: null,
  nosql: null,
  web: null,
});

const imagesOptions = ref<ImageOption[]>([]);
const isShowSettingsDialog = ref(false);
const selectedRow = ref<ImageTypes | null>(null);

const { configs: imageRowConfig } = useImageManagers(imagesOptions);

const handleImageRowUpdate = (config: any, value: ImageWithTag) => {
  if (config.updateHandler) {
    config.updateHandler(value);
  }

  if (!config.type || !(config.type in appModel)) {
    return;
  }

  const propertyName = config.type as keyof Application;

  if (appModel[propertyName] === null) {
    appModel[propertyName] = {
      image: value,
      config: {
        ports: [],
        volumes: [],
        envFiles: [],
        envs: []
      }
    };
  } else {
    appModel[propertyName]!.image = value;
  }
};

const showSettings = (type: ImageTypes) => {
  selectedRow.value = type;
  isShowSettingsDialog.value = true;
};

const handleClose = async (config: ContainerConfig) => {
  await nextTick();

  console.log("APP CLOSE");

  if (selectedRow.value === null || !(selectedRow.value in appModel)) {
    isShowSettingsDialog.value = false;
    return;
  }

  const propertyName = selectedRow.value as keyof Application;

  if (appModel[propertyName] === null) {
    isShowSettingsDialog.value = false;
    return;
  }

  appModel[propertyName]!.config = config;
  isShowSettingsDialog.value = false;
};

onMounted(async () => {
  imagesOptions.value = await GetAllImages();
});
</script>

<style lang="css" scoped>
.wrapper {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 20px;
  gap: 20px;
}

.buttons-containers {
  display: flex;
  justify-content: center;
}
</style>