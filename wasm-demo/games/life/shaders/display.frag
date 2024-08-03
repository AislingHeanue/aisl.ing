precision lowp float;
uniform sampler2D u_sampler;
varying highp vec2 v_tex_coord;
void main() {
    gl_FragColor = texture2D(u_sampler, v_tex_coord);
    // gl_FragColor = vec4(float(floor(gl_FragCoord.x) == 1.0),0.0,0.0,1.0);
}