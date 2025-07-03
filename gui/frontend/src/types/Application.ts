import {ImageWithTag} from "./ImageWithTag";

export interface Application {
    name: string;
    backend: ImageWithTag | null;
    sql: ImageWithTag | null;
    nosql: ImageWithTag | null;
    web: ImageWithTag | null;

}
