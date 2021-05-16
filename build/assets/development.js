process.env.NODE_ENV = "dev"
const postcssPlugin = require('esbuild-plugin-postcss2').default;

const esbuild = require('esbuild')
const autoprefixer = require('autoprefixer')
const tailwind = require('tailwindcss')
const fs = require('fs')

esbuild.build({
    platform: 'browser',
    entryPoints: ['./assets/src/app.tsx','./assets/src/main.css'],
    entryNames: '[name]-[hash]',
    assetNames: 'assets/[name]-[hash]',
    bundle: true,
    plugins: [
        {
            name: 'clear',
            setup(build) {
                build.onStart(() => {
                    fs.rm('./assets/dist', { recursive: true }, () =>{});
                })
            },
        },
        postcssPlugin({
            plugins: [autoprefixer, tailwind]
        }),
    ],
    define: {
        "process.env.NODE_ENV": '\"development\"'
    },
    outdir: './assets/dist',
}).catch(() => process.exit(1))