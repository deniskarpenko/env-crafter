import {ImageWithTag} from "./ImageWithTag";


export interface ContainerConfig {
    ports: string[];
    volumes: string[];
    envFiles: string[];
    envs: string[];

}

export interface ImageWithTagConfig {
    image: ImageWithTag | null;
    config: ContainerConfig;

}

export interface Application {
    backend: ImageWithTagConfig | null;
    sql: ImageWithTagConfig | null;
    nosql: ImageWithTagConfig | null;
    web: ImageWithTagConfig | null;

}
