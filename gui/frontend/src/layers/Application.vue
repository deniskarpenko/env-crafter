<template>
  <div class="wrapper">
  <Tabs :tabs="tabsConfig">
    <template #images>
      <images-row
          v-for="config in imageRowConfig"
          :title="config.title"
          :images="config.images"
          :tags="config.tags"
          @update:model-value="config.updateHandler"
      ></images-row>
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
  { id: 'settings', title: 'Settings', slot: 'settings' },
]

const appModel = reactive<Application>({
  name: "",
  sql: "",
  sqlTag: "",
  nosql: "",
  nosqlTag: "",
  frontend: "",
  frontendTag: "",
  webserver: "",
  webserverTag: "",
  utility: "",
  utilityTag: "",
});

const imagesOptions = ref<ImageOption[]>([]);

const { configs: imageRowConfig } = useImageManagers(imagesOptions);

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