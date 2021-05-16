import React, {Fragment, useRef} from 'react'
import {Menu, Transition} from '@headlessui/react'

const userNavigation = [
    {name: 'Your Profile', href: '#'},
    {name: 'Settings', href: '#'},
]

function classNames(...classes) {
    return classes.filter(Boolean).join(' ')
}

export type UserDropdownProps = {
    csrf: string
}

const UserDropdown: React.FC = (props) => {

    console.log(props)
    const formEl = useRef(null)
    const onLogoutClick = () => {
        // `current` points to the mounted text input element
        formEl.current.submit()
    };
    return (
        <>
            <Menu as="div" className="ml-4 relative flex-shrink-0">
                {({open}) => (
                    <>
                        <div>
                            <Menu.Button
                                className="bg-white rounded-full flex text-sm ring-2 ring-white ring-opacity-20 focus:outline-none focus:ring-opacity-100">
                                <span className="sr-only">Open user menu</span>
                                <img className="h-8 w-8 rounded-full" src="" alt=""/>
                            </Menu.Button>
                        </div>
                        <Transition
                            show={open}
                            as={Fragment}
                            leave="transition ease-in duration-75"
                            leaveFrom="transform opacity-100 scale-100"
                            leaveTo="transform opacity-0 scale-95"
                        >
                            <Menu.Items
                                static
                                className="origin-top-right z-40 absolute -right-2 mt-2 w-48 rounded-md shadow-lg py-1 bg-white ring-1 ring-black ring-opacity-5 focus:outline-none"
                            >
                                {userNavigation.map((item) => (
                                    <Menu.Item key={item.name}>
                                        {({active}) => (
                                            <a
                                                href={item.href}
                                                className={classNames(
                                                    active ? 'bg-gray-100' : '',
                                                    'block px-4 py-2 text-sm text-gray-700'
                                                )}
                                            >
                                                {item.name}
                                            </a>
                                        )}

                                    </Menu.Item>
                                ))}
                                <Menu.Item key="logout">
                                    <button type="button" className="block px-4 py-2 text-sm text-gray-700" onClick={onLogoutClick}>Logout</button>
                                </Menu.Item>
                            </Menu.Items>
                        </Transition>
                    </>
                )}
            </Menu>
            <form ref={formEl} method="post" action="/logout" id="logout-form">
                <input type="hidden" name="method" value="DELETE"/>
                <input type="hidden" name="_csrf" value={props.csrf}/>
            </form>
        </>
    )
}

export default UserDropdown