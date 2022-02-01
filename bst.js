

var bst = { firstDepth: 0 };

// e.g. root:{ key:1, left:0, right:2 }
bst.merge = function( inputs ) {
	this.root = { key:inputs[0] };
	let deepest = { keys:[inputs[0]], depth:this.firstDepth }
	for( let i = 1; i < inputs.length; i += 1 ) {
		let newKey = inputs[i];
		let depth = this.insert( this.firstDepth, this.root, newKey );
		if( depth > deepest.depth ) {
			deepest.keys = [newKey];
			deepest.depth = depth;
		}
		else if( depth == deepest.depth ) deepest.keys.push(newKey); 
	}
	console.log("Deepest: "+deepest.keys+"; Depth: ",deepest.depth);
}

bst.insert = function( depth, node, num ) {
	if( num == node.key ) return depth;
	let dir = ( num < node.key ) ? 'left' : 'right';

	if( node[dir] == undefined ) {
		node[dir] = { key:num };
		return depth+1;
	}
	else {
		return this.insert( depth+1, node[dir], num );
	}
}

bst.merge([12,11,90,82,7,9]);
bst.merge([26, 82, 16, 92, 33]);
//console.log(JSON.stringify(bst.root,null,'  '));

