precision lowp float;
uniform sampler2D u_sampler;
uniform float u_decay;
uniform float u_initial_decay;
uniform vec3 u_new_dead_colour;
uniform vec2 u_size;
varying vec2 v_tex_coord;
uniform bool u_boundary_loop;


float get_neighbour(vec2 offset) {
	vec2 new_coord = v_tex_coord + offset/u_size;
	return u_boundary_loop ? 
		(texture2D(u_sampler,mod(new_coord, 1.0)).a == 1.0 ? 1.0 : 0.0) :
		(
			((new_coord.x > 0.0) && (new_coord.y > 0.0) && (new_coord.x <  1.0) && (new_coord.y < 1.0)) ?
			(texture2D(u_sampler,new_coord).a == 1.0 ? 1.0 : 0.0) :
			0.0
		);
}

void main() {
	float neighbours = 0.0;
	vec4 current_cell = texture2D(u_sampler,v_tex_coord);
	bool current_alive = (current_cell.a == 1.0);

	neighbours += get_neighbour(vec2( 1.0,  1.0));
	neighbours += get_neighbour(vec2( 1.0,  0.0));
	neighbours += get_neighbour(vec2( 1.0, -1.0));
	neighbours += get_neighbour(vec2( 0.0,  1.0));
	neighbours += get_neighbour(vec2( 0.0, -1.0));
	neighbours += get_neighbour(vec2(-1.0,  1.0));
	neighbours += get_neighbour(vec2(-1.0,  0.0));
	neighbours += get_neighbour(vec2(-1.0, -1.0));

	// alpha = 1: alive. 0 < alpha < 1: dead and decaying colour. alpha = 0: dead
	float new_alpha = (
		(current_alive) ?
		(
			((neighbours == 2.0) || (neighbours == 3.0)) ?
			1.0 :
			(1.0 - u_initial_decay)
		) : (
			(neighbours == 3.0) ?
			1.0 :
			(current_cell.a - u_decay)
		)
	);
	bool new_alive = (new_alpha == 1.0);

	vec3 new_colour = (
		(!new_alive && (current_cell.a == 1.0)) ?
		u_new_dead_colour*new_alpha :
		current_cell.rgb/current_cell.a*new_alpha
	);

	gl_FragColor = (
		new_alive ?
		vec4(0.9,0.9,0.9,1.0) :
		vec4(new_colour,new_alpha)
	);	
}
