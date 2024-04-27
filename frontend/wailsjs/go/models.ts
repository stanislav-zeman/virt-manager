export namespace domain {
	
	export class Domain {
	    name: string;
	    status: string;
	
	    static createFrom(source: any = {}) {
	        return new Domain(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.status = source["status"];
	    }
	}

}

