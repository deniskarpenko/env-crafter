<template>
  <Application @build="handleBuild"></Application>
</template>
<script lang="ts" setup>
import Application from "./layers/Application.vue";
import {Build} from "../wailsjs/go/main/App";
import {main} from "../wailsjs/go/models";
import ApplicationConfig = main.ApplicationConfig;
import {Project} from "./types/Application";
import ProjectConfig = main.ProjectConfig;

const handleBuild = (project : Project): void => {

 const app = new ApplicationConfig();

 const projectConfig = new ProjectConfig({
   backend: {
     name: project.backend?.image?.image
   }
 });

 if (!app.projects) {
   app.projects = [];
 }

 app.projects.push(project);

  Build(app);
}
</script>

