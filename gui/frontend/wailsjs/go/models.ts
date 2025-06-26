export namespace models {
	
	export class Image {
	    image_id: string;
	    image: string;
	    docker_image: string;
	    is_dockerfile: boolean;
	    image_type: string;
	
	    static createFrom(source: any = {}) {
	        return new Image(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.image_id = source["image_id"];
	        this.image = source["image"];
	        this.docker_image = source["docker_image"];
	        this.is_dockerfile = source["is_dockerfile"];
	        this.image_type = source["image_type"];
	    }
	}
	export class Tag {
	    tag_id: number;
	    image_id: number;
	    TagName: string;
	
	    static createFrom(source: any = {}) {
	        return new Tag(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.tag_id = source["tag_id"];
	        this.image_id = source["image_id"];
	        this.TagName = source["TagName"];
	    }
	}

}

