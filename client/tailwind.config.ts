import type { Config } from "tailwindcss"

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: "#54b3e9",
        "primary-hover": "#3a9ad0",
        "color-error-primary": "rgba(208, 2, 27, 0.8)",
        "color-error-secondary": "rgba(208, 2, 27, 0.3)",
      },
    },
  },
  plugins: [],
}
export default config
