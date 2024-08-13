



function add(a, b) {
    return a + b 
}

function sub(a, b) {
    return a - b
}

function div(a, b) {
    return a / b
}

function mul(a, b) {
    return a * b
}


function calc(input) {

    let x = 0

    

    for(let i = 0; i < len(input); i++) {

        if (input[i] === "/") {
            x = input[i - 1] / input[i + 1]
        }

        if (input[i] === "*") {
            x = input[i - 1] * input[i + 1]
        }

        input.splice(i - 1) 

    }

    for(let i = 0; i < len(input); i++) {

        if (input[i] === "-") {
            x = input[i - 1] - input[i + 1]
        }

        if (input[i] === "+") {
            x = input[i - 1] + input[i + 1]
        }
    }

    console.log(x)

}

input = ["2", "+", "2", "*", "10"]


arr = [
    [2, 2, add],
    [4, 10, mul]
]

// input = "2 + 2 * 10"

calc(input)












