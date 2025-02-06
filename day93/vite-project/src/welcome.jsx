import { useState } from "react"

const COLORS = ["blue", "green", "brown", "red", "indigo", "pink"]

export function App() {
    const [color, setColor] = useState(COLORS[0])

    const ChangeColor = (newColor) => () => {
        setColor(newColor);
    };

    return (
        <div style={{"background-color": color }}>
            <h1 style={{color}}>Hello, World</h1>
            <button type="button" style={{color}}>Click here</button>
            {COLORS.map((color) => (
                <button type="button"
                key={color}
                onClick={ChangeColor(color)}>
                    Click, Me
                </button>
            ))}
        </div>
    )

}

// When working with States like array and objects its good to copy the object into a new object while updating the property
export function Person() {
    const [person, setPerson] = useState({ name: "Lawson", age: 100 });
    
    const handleIncreaseAge = () => {
        const newPerson = { ...person, age: person.age + 1 };
        setPerson(newPerson);
    }

    const handleDecreaseAge = () => {
        const newPerson = { ...person, age: person.age - 1 };
        setPerson(newPerson);
    };

    return (
        <>
            <h1>{person.name}</h1>
            <h2>{person.age}</h2>
            <button onClick={handleIncreaseAge}>Increase age</button>
            <br />
            <button onClick={handleDecreaseAge}>Decrease age</button>
        </>
    );
}