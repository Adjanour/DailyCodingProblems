

const callFunction = (func,time) => {
    setTimeout(func,time)
}

function getTime() {
    return console.log(new Date().toLocaleTimeString());
}

// Example usage:
callFunction(getTime, 1000);