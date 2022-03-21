
function formSubmitEventHandler() {
    let button = document.getElementById("submit");
    button.addEventListener('click', function(event) {
        event.preventDefault();
        document.getElementById("restaurants").innerHTML = "";
        validatePostcode();
    })
}

function validatePostcode(){
    let input = document.getElementById("postcode").value
    let postcode = input.replace(/\s+/g, '')
    if (postcode === ""){
        alert("Invalid input, please try again")
        console.log("Invalid input, please try again")
    } else if (postcode.length > 7) {
        alert("invalid postcode: too many characters")
        console.log("Invalid postcode: too many characters")
    } else if (postcode.length < 3) {
        alert("invalid postcode: not enough characters")
        console.log("Invalid postcode: not enough characters")
    } else {
        console.log(postcode);
        apiRequest(postcode)
        return postcode;
    }
}

async function apiRequest(postcode) {
    try {
        const response = await fetch(`http://localhost:8080/postcode/${postcode}`)
        if (response.ok) {
            const data = await response.json();
            console.log(data);
            validate(data)
        }
    } catch (err) {
        let restaurantHeader = document.getElementById('restaurants')
        restaurantHeader.innerHTML = `<h4>⚠️Server failure: Failed to find restaurants⚠️</h4>`
        console.log(err);
    }
}

function validate(data){
    if (data["restaurants"].length === 0){
        alert("invalid postcode. please try again")
        console.log("invalid postcode. please try again")
    } else {
        iterate(data)
    }
}

function iterate(data) {
    for (let key in data) {
        data[key].forEach( function(restaurant) {
            let restaurants = (`<br>${restaurant["name"]}: ${restaurant["ratingAverage"]}<br>`)
            display(restaurants)
        })
    }
}

function display(restaurants){
    let list = document.getElementById('restaurants')
    let restaurantHeader = document.getElementById('restaurant-header')
    restaurantHeader.innerHTML = `<h3>Restaurants 🍔: Rating ⭐️</h3>`
    list.innerHTML += restaurants;
}

function handleClearAll() {
    let button = document.getElementById('clear');
    button.addEventListener('click', function(event) {
        event.preventDefault()
        document.getElementById("restaurants").innerHTML = "";
    })
}

handleClearAll();

formSubmitEventHandler();
