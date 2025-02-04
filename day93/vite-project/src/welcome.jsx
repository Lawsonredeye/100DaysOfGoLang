import './App.css'

export function Welcome(props) {
    if (!props.name) {
        return <h1>Hello, Stranger!</h1>
    }

    return <h1>Welcome, {props.name}!</h1>
}

export function WelcomeMsg() {
    return (
        <>
            <nav>
                <ul>
                    <li>head</li>
                    <li>contact</li>
                    <li>menu</li>
                </ul>
            </nav>
        </>
    )
}
