import {Project, ImageWithTagConfig} from "../types/Application";
import {main} from "../../wailsjs/go/models";
import ProjectConfig = main.ProjectConfig;
import GoImageWithTagConfig = main.ImageWithTagConfig;
import GoImageWithTag= main.ImageWithTag;
import {ImageWithTag} from "../types/ImageWithTag";

export const toProjectConfig = (config: Project): ProjectConfig => {
    const projectConfig = new ProjectConfig();


    projectConfig.backend = toGoImageWithTagConfig(config.backend);

    return projectConfig;
}

export const toGoImageWithTagConfig = (config: ImageWithTagConfig): GoImageWithTagConfig => {
    const imageWithTagConfig = new GoImageWithTagConfig();

    return imageWithTagConfig;
}

export const toGoImageWithTag = (imageWithTag: ImageWithTag): GoImageWithTag => {
    const imageWithTagForGo = new GoImageWithTag(imageWithTag);

    return imageWithTagForGo;
}

export const to

