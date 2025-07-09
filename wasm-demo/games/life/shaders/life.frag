precision lowp float;
uniform sampler2D uSampler;
uniform float uDecay;
uniform float uInitialDecay;
uniform vec3 uNewDeadColour;
uniform vec2 uSize;
varying vec2 vTexCoord;
uniform bool uBoundaryLoop;
uniform bool uPaused;

float getNeighbour(vec2 offset) {
  vec2 newCoord = vTexCoord + (offset/uSize);
  vec2 modNewCoord = mod(newCoord, 1.0);
  if (texture2D(uSampler, modNewCoord).a == 1.0) {
    if (uBoundaryLoop || newCoord == modNewCoord) {
      return 1.0;
    }
  }

  return 0.0;
}



void main() {
  float neighbours = 0.0;
  vec4 currentCell = texture2D(uSampler,vTexCoord);
  bool currentAlive = (currentCell.a == 1.0);
  neighbours += getNeighbour(vec2( 1.0,  1.0));
  neighbours += getNeighbour(vec2( 1.0,  0.0));
  neighbours += getNeighbour(vec2( 1.0, -1.0));
  neighbours += getNeighbour(vec2( 0.0,  1.0));
  neighbours += getNeighbour(vec2( 0.0, -1.0));
  neighbours += getNeighbour(vec2(-1.0,  1.0));
  neighbours += getNeighbour(vec2(-1.0,  0.0));
  neighbours += getNeighbour(vec2(-1.0, -1.0));

  // alpha = 1: alive. 0 < alpha < 1: dead and decaying colour. alpha = 0: dead
  vec3 newColour;
  float newAlpha;
  if (currentAlive) {
    // stay alive
    if ((neighbours == 2.0) || (neighbours == 3.0) || uPaused) {
      newAlpha = 1.0;
      newColour = vec3(0.9,0.9,0.9);
      // die
    } else {
      // not using alpha in the traditional sense here, instead it's just a measure of how dark to make the cell.
      newAlpha = 1.0 - uInitialDecay;
      newColour = uNewDeadColour * newAlpha;
    }
  } else {
    // born
    if (neighbours == 3.0 && !uPaused) {
      newAlpha = 1.0;
      newColour = vec3(0.9,0.9,0.9);
      // stay dead
    } else {
      newAlpha = currentCell.a - uDecay;
      // get unscaled rgb of the colour it had when it died.
      vec3 originalColour = currentCell.rgb / currentCell.a;
      newColour = originalColour * newAlpha;
    }
  }
  gl_FragColor = vec4(newColour,newAlpha);
}
