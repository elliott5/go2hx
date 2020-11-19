package go2hx.types;

typedef Func = {
	name:String,
	params:Array<String>,
	results:Array<String>,
	exported:Bool,
	body:Array<String>,
	doc:String,
	//recv:Array<String>,
}

typedef Var = {
	name:String,
	constant:Bool,
	exported:Bool,
	type:String,
	value:String
}

typedef Struct = {
	name:String,
	export:Bool,
	type:String,
	funcs:Array<Func>,
}

typedef Package = {
	packagepath:String,
	name:String,
	imports:Array<Array<String>>,
	funcs:Array<Func>,
	vars:Array<Var>,
	structs:Array<Struct>
}

typedef JsonData = {
	pkgs:Array<Package>,
}

typedef Decl = {
	name:String,
	args:Array<String>,
}
