/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
      './server/components/templates/*.html',
  ],
  daisyui: {
    themes: [
      "light", 
      "dark", 
      "night", 
      "dracula", 
      "acid",
      {
        "soft-lite" : {
          "primary": "#baffaa",
          "secondary": "#2b57c6",
          "accent": "#075a8e",
          "neutral": "#2c323a",
          "base-100": "#dee3ed",
          "info": "#84c2e6",
          "success": "#5beccf",
          "warning": "#f3b539",
          "error": "#f46775",
        },
      }
    ],
  },
  theme: {
    fontFamily: {
        body: ["JetBrains Mono"]
    },
    extend: {
    },
  },
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
}

