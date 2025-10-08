<template>
  <Application @build="handleBuild"></Application>
</template>
<script lang="ts" setup>
import Application from "./layers/Application.vue";
import {Project} from "./types/Application";

const handleBuild = (app: Project): void => {
  console.log(app);
}
</script>

