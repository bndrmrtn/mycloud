/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./components/**/*.{js,vue,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./app.vue",
    "./error.vue",
  ],
  theme: {
    extend: {
      colors: {
         main: {
          from: "#353a65",
          to: "#333862",
        },
        input: "#282c4f",
        widget: "#3c416e",
      },
    },
  },
  plugins: [],
}

