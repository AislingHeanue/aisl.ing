attribute vec2 aPosition;
attribute vec2 aTexCoord;

varying highp vec2 vTexCoord;

void main() {
	gl_Position = vec4(aPosition, 0.0, 1.0);
	vTexCoord = aTexCoord;
}
