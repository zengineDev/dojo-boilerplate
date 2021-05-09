const fs = require('fs')

export default {
    name: 'clear',
    setup(build) {
        build.onStart(() => {
            fs.rm('./assets/dist', { recursive: true }, () =>{});
        })
    },
}