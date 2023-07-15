/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
      './server/views/templates/*.html',
  ],
  theme: {
    extend: {},
  },
  plugins: [
      require('tailwindcss'),
      require('autoprefixer'),
  ],
}

