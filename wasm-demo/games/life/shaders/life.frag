precision lowp float;
uniform sampler2D u_sampler;
uniform float u_decay;
uniform float u_initial_decay;
uniform vec3 u_new_dead_colour;
uniform vec2 u_size;
varying vec2 v_tex_coord;
uniform bool u_boundary_loop;
uniform bool u_paused;

float get_neighbour(vec2 offset) {

	// vec2 new_coord = v_tex_coord + offset/u_size;
	// return u_boundary_loop ? 
	// 	(texture2D(u_sampler,mod(new_coord, 1.0)).a == 1.0 ? 1.0 : 0.0) :
	// 	(
	// 		((new_coord.x > 0.0) && (new_coord.y > 0.0) && (new_coord.x <  1.0) && (new_coord.y < 1.0)) ?
	// 		(texture2D(u_sampler,new_coord).a == 1.0 ? 1.0 : 0.0) :
	// 		0.0
	// 	);


  vec2 new_coord = v_tex_coord + (offset/u_size);
  vec2 mod_new_coord = mod(new_coord, 1.0);
  if (texture2D(u_sampler, mod_new_coord).a == 1.0) {
    if (u_boundary_loop || new_coord == mod_new_coord) {
      return 1.0;
    }
  }

  return 0.0;
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
  vec3 new_colour;
  float new_alpha;
  if (current_alive) {
    // stay alive
    if ((neighbours == 2.0) || (neighbours == 3.0) || u_paused) {
      new_alpha = 1.0;
      new_colour = vec3(0.9,0.9,0.9);
      // die
    } else {
      // not using alpha in the traditional sense here, instead it's just a measure of how dark to make the cell.
      new_alpha = 1.0 - u_initial_decay;
      new_colour = u_new_dead_colour * new_alpha;
    }
  } else {
    // born
    if (neighbours == 3.0 && !u_paused) {
      new_alpha = 1.0;
      new_colour = vec3(0.9,0.9,0.9);
      // stay dead
    } else {
      new_alpha = current_cell.a - u_decay;
      // get unscaled rgb of the colour it had when it died.
      vec3 original_colour = current_cell.rgb / current_cell.a;
      new_colour = original_colour * new_alpha;
    }
  }

	gl_FragColor = vec4(new_colour,new_alpha);
}
