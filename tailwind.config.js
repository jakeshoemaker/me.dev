/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
      './server/views/templates/*.html',
  ],
  darkMode: 'class',
  theme: {
    fontFamily: {
        'sans': ['Open Sans'],
    },
    extend: {},
  },
  plugins: [
      require('tailwindcss'),
      require('autoprefixer'),
  ],
}

