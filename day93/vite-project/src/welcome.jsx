import { useState } from "react"

const COLORS = ["blue", "green", "brown", "red", "indigo", "pink"]

export function App() {
    const [color, setColor] = useState(COLORS[0])

    const ChangeColor = (newColor) => () => {
        setColor(newColor);
    };

    return (
        <>
            <h1 style={{color}}>Hello, World</h1>
            <button type="button" style={{color}}>Click here</button>
            {COLORS.map((color) => (
                <button type="button"
                key={color}
                onClick={ChangeColor(color)}>
                    Click, Me
                </button>
            ))}
        </>
    )

}