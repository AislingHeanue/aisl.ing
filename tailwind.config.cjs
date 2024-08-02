/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["index.html ","./src/**/*.{jsx,js}"],
  darkMode: "class",
  theme: {
    extend: {
      fontFamily: {
        inter: ["inter","serif"]
      }
    },
  },
  variants: {
    imageRendering: ['responsive'],
  },
  plugins: [
    require('tailwindcss-image-rendering')(),
  ],
}
