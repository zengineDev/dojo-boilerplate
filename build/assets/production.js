process.env.NODE_ENV = "production"
const postcssPlugin = require('esbuild-plugin-postcss2').default;

const esbuild = require('esbuild')
const autoprefixer = require('autoprefixer')
const tailwind = require('tailwindcss')
const pluginClear = require('./plugin_clear')

esbuild.build({
    platform: 'browser',
    entryPoints: ['./assets/src/app.tsx','./assets/src/main.css'],
    entryNames: '[name]-[hash]',
    assetNames: 'assets/[name]-[hash]',
    bundle: true,
    plugins: [
        pluginClear,
        postcssPlugin({
            plugins: [autoprefixer, tailwind]
        }),
    ],
    define: {
        "process.env.NODE_ENV": '\"production\"'
    },
    outdir: './assets/dist',
}).catch(() => process.exit(1))