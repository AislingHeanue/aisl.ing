precision highp float;
uniform float uZoom;
uniform int uMaxIterations;
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
    float x2 = 0.;
    float y2 = 0.;
    float w = 0.;
    int iteration = 0;

    for (int i = 0; i <= 3000; i++) {
        y = 2. * x * y + y0;
        x = x2 - y2 + x0;
        x2 = x * x;
        y2 = y * y;

        iteration = i;
        if (x2 + y2 > pow(2., 8.) || iteration >= uMaxIterations) {
            break;
        }
    }

    float iterationFloat = float(iteration);
    if (iteration < uMaxIterations) {
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

