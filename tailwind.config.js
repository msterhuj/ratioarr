module.exports = {
    content: [
        "./internal/views/**/*.templ",
        "./internal/views/**/*_templ.go",
        "./test.html"
    ],
    theme: {
        extend: {},
    },
    plugins: [require("@tailwindcss/typography"), require("daisyui")],
    daisyui: {
        themes: ["light"],
    },
};
