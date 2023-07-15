/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
      './server/views/templates/*.html',
  ],
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

