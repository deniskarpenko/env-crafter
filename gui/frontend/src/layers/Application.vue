<template>
  <div class="wrapper">
  <Tabs :tabs="tabsConfig">
    <template #images>
      <images-row
          v-for="config in imageRowConfig"
          :title="config.title"
          :images="config.images"
          :tags="config.tags"
          @update:model-value="handleImageRowUpdate(config, $event)"
      ></images-row>
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

import {reactive, ref, onMounted, computed, watch} from "vue";
import {GetAllImages, GetTagByImageId} from "../../wailsjs/go/main/App";
import {Application} from "../types/Application";
import {ImageOption} from "../types/ImageOption";
import Tabs from "../components/Tabs.vue";
import ImagesRow from "../components/ImagesRow.vue";
import {useImageManagers} from "../composables/ImageManager";

const tabsConfig = [
  { id: 'images', title: 'Images', slot: 'images' },
  { id: 'result', title: 'Result', slot: 'result' },
  { id: 'settings', title: 'Settings', slot: 'settings' },
]

const appModel = reactive<Application>({
  name: "",
  backend: null,
  sql: null,
  nosql: null,
  web: null,
});

const imagesOptions = ref<ImageOption[]>([]);

const { configs: imageRowConfig } = useImageManagers(imagesOptions);

const handleImageRowUpdate = (config: any, value: any) => {
  console.log(value);

  if (config.updateHandler) {
    config.updateHandler(value);
  }
};


onMounted(async () => {
  imagesOptions.value = await GetAllImages();

})
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