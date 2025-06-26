<template>
  <div class="wrapper">
  <Tabs :tabs="tabsConfig">
    <template #images>
      <images-row
          title="Backend"
          :images="backendOptions"
          :tags="backendTags"
          @image-selected="imageSelectedHandler"
      ></images-row>
    </template>
  </Tabs>
  </div>
</template>
<script setup lang="ts">

import {reactive, ref, onMounted, computed, watch} from "vue";
import {GetAllImages, GetTagByImageId} from "../../wailsjs/go/main/App";
import {Application} from "../types/Application";
import {ImageTypes} from "../types/ImageTypes";
import {ImageOption} from "../types/ImageOption";
import {TagOption} from "../types/TagOption";
import {ImageWithTag} from "../types/ImageWithTag";
import Tabs from "../components/Tabs.vue";
import ImagesRow from "../components/ImagesRow.vue";

const tabsConfig = [
  { id: 'images', title: 'Images', slot: 'images' },
  { id: 'settings', title: 'Настройки', slot: 'settings' },
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

const backendModel = ref<ImageWithTag>({
  image: null,
  tag: null
});

const sqlModel = ref<ImageWithTag>({
  image: null,
  tag: null
});

const backendOptions =computed(() => {
  return  imagesOptions.value.filter((image) =>  image.image_type === ImageTypes.Backend);
});

const sqlOptions = computed(() => {
  return  imagesOptions.value.filter((image) =>  image.image_type === ImageTypes.SQL);
})

const backendTags = ref<TagOption[]>([]);

const sqlTags = ref<TagOption[]>([]);

const imagesOptions = ref<ImageOption[]>([]);


watch(() => backendModel.value?.image, async (newImage) => {
  if (!newImage) {
    return;
  }

  try {
    backendTags.value = await GetTagByImageId(parseInt(newImage.image_id));
  } catch (error) {
    console.error('Error fetching backend tags:', error);
    backendTags.value = [];
  }
});

watch(() => sqlModel.value?.image, async  (newImage) => {
  if (!newImage) {
    return;
  }

  try {
    sqlTags.value = await GetTagByImageId(parseInt(newImage.image_id));
  } catch (error) {
    console.error('Error fetching sql tags:', error);
    sqlTags.value = [];
  }
});

const imageSelectedHandler = (value: ImageOption) => {
  backendModel.value.image = value;
}

onMounted(async () => {
  imagesOptions.value = await GetAllImages();

})
</script>
<style lang="css" scoped>
.wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 20px;
}
</style>