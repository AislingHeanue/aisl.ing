const portfolio = [
  {
    title: "TerraSchema: A Terraform to JSON Schema generator",
    imgURL: "/assets/TerraSchema.png",
    summary: "Open source CLI tool I created while working at HPE to help automate Terraform input validation.",
    stack: ["Go", "Terraform"],
    link: "https://github.com/HewlettPackard/terraschema/",
    internal: false
  },
  {
    title: "Rubik's Cube",
    imgURL: "/assets/Go Rubik's Cube.png",
    summary: "Written as a learning exercise for using shader graphics. Runs in the web browser.",
    stack: ["Go", "Wasm", "WebGL"],
    link: "/rubiks-cube/",
    internal: true
  },
  {
    title: "Particle Game (Android app)",
    imgURL: "/assets/Particle Game Screenshot.jpg",
    summary: "Toy app I made which uses the device's gyroscope to simulate particle movements and collisions. WIP.",
    stack: ["Flutter", "Flame"],
    link: "https://github.com/AislingHeanue/particle_game",
    internal: false
  },
  {
    title: "Mandelbrot Set Demo",
    imgURL: "/assets/Mandelbrot Screenshot.png",
    summary: "Browser demo for viewing and zooming into the Mandelbrot set.",
    stack: ["Go", "Wasm", "WebGL"],
    link: "/mandelbrot/",
    internal: true
  },
  {
    title: "Conway's Game of Life",
    imgURL: "/assets/Game of Life Screenshot.png",
    summary: "Browser demo which allows loading different Game of Life patterns. Interface WIP.",
    stack: ["Go", "Wasm", "WebGL"],
    link: "/game-of-life/",
    internal: true
  },
  {
    title: "Quantum Computing Circuits",
    imgURL: "/assets/Quantum Screenshot.png",
    summary: "Programming work done for my undergraduate thesis. Uses Qisket to implement Shor's algorithm.",
    stack: ["Python", "Qiskit"],
    link: "https://github.com/AislingHeanue/Quantum-Computing-Circuits",
    internal: false
  },
  {
    title: "Collatz Tree Generator",
    imgURL: "/assets/Collatz Screenshot.png",
    summary: "Interactively generates vector graphics based on a famous maths problem",
    stack: ["Typescript", "React"],
    link: "/collatz/",
    internal: true
  },
  {
    title: "Advent of Code Solutions",
    imgURL: "/assets/AoC Screenshot 2.png",
    summary: "Collection of my Advent of Code solutions. WIP: need to make a presentable README for that repo.",
    stack: ["Rust", "Go", "Python"],
    link: "https://github.com/AislingHeanue/Advent-Of-Code/blob/master/",
    internal: false
  }
]
export default portfolio
