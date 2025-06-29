<template>
  <div class="wrapper">
  <Tabs :tabs="tabsConfig">
    <template #images>
      <images-row
          title="Backend"
          :images="backendOptions"
          :tags="backendTags"
          @update:model-value="updateBackendModelHandler"
      ></images-row>
      <images-row
          title="SQL"
          :images="sqlOptions"
          :tags="sqlTags"
          @update:model-value="updateSQLModelHandler"
      ></images-row>
      <images-row
          title="NoSQL"
          :images="noSqlOptions"
          :tags="noSqlTags"
          @update:model-value="updateNoSQLModelHandler"
      ></images-row>
      <images-row
          title="Web"
          :images="webOptions"
          :tags="webTags"
          @update:model-value="updateWebModelHandler"
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
import {ImageTypes} from "../types/ImageTypes";
import {ImageOption} from "../types/ImageOption";
import {TagOption} from "../types/TagOption";
import {ImageWithTag} from "../types/ImageWithTag";
import Tabs from "../components/Tabs.vue";
import ImagesRow from "../components/ImagesRow.vue";

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

const backendModel = ref<ImageWithTag>({
  image: null,
  tag: null
});

const sqlModel = ref<ImageWithTag>({
  image: null,
  tag: null
});

const noSQLModel = ref<ImageWithTag>({
  image: null,
  tag: null
});

const webModel = ref<ImageWithTag>({
  image: null,
  tag: null
});

const backendOptions =computed(() => {
  return  imagesOptions.value.filter((image) =>  image.image_type === ImageTypes.Backend);
});

const sqlOptions = computed(() => {
  return  imagesOptions.value.filter((image) =>  image.image_type === ImageTypes.SQL);
});

const noSqlOptions = computed(() => {
  return  imagesOptions.value.filter((image) =>  image.image_type === ImageTypes.NOSQL);
});

const webOptions = computed(() => {
  return  imagesOptions.value.filter((image) =>  image.image_type === ImageTypes.WEB);
});

const backendTags = ref<TagOption[]>([]);

const sqlTags = ref<TagOption[]>([]);

const noSqlTags = ref<TagOption[]>([]);

const webTags = ref<TagOption[]>([]);

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

watch(() => noSQLModel.value?.image, async  (newImage) => {
  if (!newImage) {
    return;
  }

  try {
    noSqlTags.value = await GetTagByImageId(parseInt(newImage.image_id));
  } catch (error) {
    console.error('Error fetching nosql tags:', error);
    noSqlTags.value = [];
  }
});

watch(() => webModel.value?.image, async  (newImage) => {
  if (!newImage) {
    return;
  }

  try {
    webTags.value = await GetTagByImageId(parseInt(newImage.image_id));
  } catch (error) {
    console.error('Error fetching nosql tags:', error);
    webTags.value = [];
  }
});

const updateBackendModelHandler = (value: ImageWithTag) => {
  backendModel.value = value;
};

const updateSQLModelHandler = (value: ImageWithTag) => {
  sqlModel.value = value;
};

const updateNoSQLModelHandler = (value: ImageWithTag) => {
  noSQLModel.value = value;
};

const updateWebModelHandler = (value: ImageWithTag) => {
  webModel.value = value;
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