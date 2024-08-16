precision lowp float;
uniform sampler2D u_sampler;
uniform float u_decay;
varying highp vec2 v_tex_coord;
void main() {
	float neighbours = 0.0;
	vec4 current_cell = texture2D(u_sampler,v_tex_coord);
	bool current_alive = (current_cell.a == 1.0);

	// alpha = 1: alive. 0 < alpha < 1: dead and decaying colour. alpha = 0: dead
	float new_alpha = (
		(current_alive) ?
		(
			1.0
		) : (
			(current_cell.a - u_decay)
		)
	);
	bool new_alive = current_alive;

	vec3 new_colour = (
		current_cell.rgb/current_cell.a*new_alpha
	);

	gl_FragColor = (
		new_alive ?
		vec4(1.0,1.0,1.0,1.0) :
		vec4(new_colour,new_alpha)
	);	
}