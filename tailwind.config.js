/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
      './server/views/templates/*{.html,js}',
  ],
  darkMode: 'class',
  theme: {
    fontFamily: {
        body: ["JetBrains Mono"]
    },
    extend: {
    },
  },
  plugins: [require("daisyui")],
}

