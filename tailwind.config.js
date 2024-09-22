/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.{html,js,templ}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui"), require('@tailwindcss/forms')],
  daisyui: {
    themes: ["light", "dark", "cupcake", "dracula"],
  },
}
