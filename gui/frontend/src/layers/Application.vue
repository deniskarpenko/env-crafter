<template>
  <div class="wrapper">
    <Tabs :tabs="tabsConfig">
      <template #images>
        <images-row
            v-for="config in imageRowConfig"
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
import {reactive, ref, onMounted} from "vue";
import {GetAllImages} from "../../wailsjs/go/main/App";
import {Application, ContainerConfig, ImageWithTagConfig} from "../types/Application";
import {ImageOption} from "../types/ImageOption";
import Tabs from "../components/Tabs.vue";
import ImagesRow from "../components/ImagesRow.vue";
import {useImageManagers} from "../composables/ImageManager";
import SettingsDialog from "../components/dialog/SettingsDialog.vue";
import {ImageWithTag} from "../types/ImageWithTag";

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
const selectedRow = ref<string>('');

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

const showSettings = (title: string) => {
  selectedRow.value = title;
  isShowSettingsDialog.value = true;
};

const handleClose = (config: ContainerConfig) => {
  isShowSettingsDialog.value = false;
  console.log(selectedRow.value);
  console.log(config);
  console.log(appModel);
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