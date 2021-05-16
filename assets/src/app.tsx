import React from 'react'
import ReactDOM from 'react-dom'

import UserDropdown from "./components/UserDropdown";

const components = {
    "user-dropdown": <UserDropdown/>
}


for (let c in components) {
    const elements: HTMLCollection = document.getElementsByTagName("user-dropdown")
    for (let element in elements ){
        const attr = elements[element].attributes
        ReactDOM.render(<UserDropdown csrf={attr.getNamedItem("csrf").value}/>, elements[element])
    }
}