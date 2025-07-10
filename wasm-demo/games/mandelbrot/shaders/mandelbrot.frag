precision highp float;
uniform float uZoom;
uniform vec2 uCentre;
varying vec2 vTexCoord;

vec3 colour(float number) {
    const float PI = 3.14159265359;
    return vec3(
        0.5 + 0.5 * sin(number / 50. + 5. * PI / 3.),
        0.5 + 0.5 * sin(number / 50.),
        0.5 + 0.5 * sin(number / 50. + 3. * PI / 3.));
}

void main() {
    float x0 = (vTexCoord.x * 2.5 / uZoom) - (1.25 / uZoom) + uCentre.x;
    float y0 = (vTexCoord.y * 2.5 / uZoom) - (1.25 / uZoom) + uCentre.y;
    float x = 0.;
    float y = 0.;
    float temp = 0.;
    int iteration = 0;
    const int maxIterations = 3000;

    for (int i = 0; i <= maxIterations; i++) {
        temp = x * x - y * y + x0;
        y = 2. * x * y + y0;
        x = temp;

        iteration = i;
        if (x * x + y * y > pow(2., 8.)) {
            break;
        }
    }

    float iterationFloat = float(iteration);
    if (iteration < maxIterations) {
        float logZn = log(x * x + y * y) / 2.;
        float nu = log(logZn / log(2.)) / log(2.); // nu is between 0 and 1
        iterationFloat = float(iteration) + 1. - nu;
        gl_FragColor = vec4(colour(iterationFloat), 1.);
    } else {
        gl_FragColor = vec4(0., 0., 0., 1.);
    }

    // vec3 colour1 = colour(floor(iterationFloat));

    // vec3 colour2 = colour(floor(iterationFloat) + 1.);

    // gl_FragColor = vec4(colour1 + (colour2 - colour1) * mod(iterationFloat, 1.), 1.0);
}
