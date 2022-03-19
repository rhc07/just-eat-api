
formSubmitEventHandler();

getRestaurants();

handleClearAll();


function formSubmitEventHandler() {
    let button = document.getElementById("submit");
    button.addEventListener('click', function(event) {
        event.preventDefault();
        getPostcode();
    })
};

function getPostcode(){
    let postcode = document.getElementById("postcode").value
    if (postcode !== ""){
        console.log(postcode);
        showRestaurants(postcode);
        return postcode;
    } else {
        console.log("Invalid input, please try again")
        return "Invalid input, please try again"
    }
}

function getRestaurants(){
fetch('http://localhost:8080/').then(function (response) {
    // The API call was successful!
    if (response.ok) {
        return response.json();
    } else {
        return Promise.reject(response);
    }
}).then(function (data) {
    // This is the JSON from our response
    console.log(JSON.stringify(data));
    showRestaurants(JSON.stringify(data));
}).catch(function (err) {
    // There was an error
    console.warn('Something went wrong.', err);
});

}

function showRestaurants(data){
    let restaurants = document.getElementById('restaurants')
    restaurants.innerHTML += data;
}


function handleClearAll() {
    let button = document.getElementById('clearAll');
    button.addEventListener('click', function(event) {
        event.preventDefault();
        document.getElementById('restaurants').innerHTML += '';
    })
}
