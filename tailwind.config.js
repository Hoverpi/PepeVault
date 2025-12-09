import daisyui from "./assets/daisyui.mjs";

module.exports = {
  content: [ "./src/**/*.{rs, html}", "./index.html" ],

  plugins: [
    [require("@tailwindcss/forms"), require("@tailwindcss/typography"), require(daisyui)],
  ],
};
