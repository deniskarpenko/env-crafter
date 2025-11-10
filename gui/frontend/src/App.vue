<template>
  <Application @build="handleBuild"></Application>
</template>
<script lang="ts" setup>
import Application from "./layers/Application.vue";
import {Project} from "./types/Application";
import {Build} from "../wailsjs/go/main/App";

const handleBuild = (app: Project): void => {
  Build(app);
}
</script>

