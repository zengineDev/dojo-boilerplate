module.exports = {
    purge: [
        './asset/src/**/*.tsx',
        './views/**/*.gohtml',
    ],
    darkMode: false, // or 'media' or 'class'
    theme: {
        extend: {},
    },
    variants: {
        extend: {},
    },
    plugins: [
        require('@tailwindcss/forms')
    ],
}