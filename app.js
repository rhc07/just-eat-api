
function formSubmitEventHandler() {
    let button = document.getElementById("submit");
    button.addEventListener('click', function(event) {
        event.preventDefault();
        getPostcode();
    })
}

function getPostcode(){
    let postcode = document.getElementById("postcode").value
    if (postcode !== ""){
        console.log(postcode);
        getRestaurants();
        return postcode;
    } else {
        console.log("Invalid input, please try again")
        return "Invalid input, please try again"
    }
}

async function getRestaurants() {
    try {
        const response = await fetch('http://localhost:8080/')

        if (response.ok) {
            const data = await response.json();
            console.log(data);
            display(data)
            return data;
        }
    } catch (err) {
        let restaurantHeader = document.getElementById('restaurants')
        restaurantHeader.innerHTML = `<h4>⚠️Server failure: Failed to find restaurants⚠️</h4>`
        console.log(err);

    }
}

function display(data) {
    let items = data
    console.log(items)
    for (let key in items) {
        items[key].forEach( function(item) {
            console.log(item)
            for (const [key, value] of Object.entries(item)) {
                let newData = (`<br>${key}: ${value}<br>`)
                showRestaurants(newData)
            }
        })
    }
}

function showRestaurants(data){
    let restaurants = document.getElementById('restaurants')
    let restaurantHeader = document.getElementById('restaurant-header')
    restaurantHeader.innerHTML = `<h4>Restaurants</h4>`
    restaurants.innerHTML += data;
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
