function Button({ text = "Click Me!", color = "blue", fontSize = 12, handleClick}) {
    const buttonStyle = {
        color: color,
        fontSize: fontSize + 'px'
    };

    return (
        <button  onClick={handleClick} style={buttonStyle}>{text}</button>
    );
}

export default function App() {
    const handleButtonClick = () => {
        window.location.href = "https://www.google.com"
    };

    return (
        <div>
            {/* <Button />
            <Button text="Don't Click Me!" color="red" />
            <Button fontSize={12}/> */}

            <div>
                <Button handleClick={handleButtonClick} />
            </div>
        </div>
    );
}