attribute vec3 aVertexPosition;
attribute vec4 aVertexColour;
uniform mat4 modelView;
uniform mat4 perspectiveMatrix;
uniform bool useUniformColour;
uniform lowp vec4 uColour;

varying lowp vec4 vColour;

void main(void) {
	gl_Position = perspectiveMatrix * modelView * vec4(aVertexPosition,1.0);
	vColour = useUniformColour ? uColour : aVertexColour;
}
