/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
      './server/views/templates/*.html',
  ],
  darkMode: 'class',
  theme: {
    fontFamily: {
        body: ["JetBrains Mono"]
    },
    extend: {},
  },
  plugins: [
      require('tailwindcss'),
      require('autoprefixer'),
  ],
}

