<template>
  <Application @build="handleBuild"></Application>
</template>
<script lang="ts" setup>
import Application from "./layers/Application.vue";
import {Build} from "../wailsjs/go/main/App";
import {main} from "../wailsjs/go/models";
import ApplicationConfig = main.ApplicationConfig;
import Project = main.ProjectConfig;

const handleBuild = (project : Project): void => {

 const app = new ApplicationConfig();

 if (!app.projects) {
   app.projects = [];
 }

 app.projects.push(project);

  Build(app);
}
</script>

