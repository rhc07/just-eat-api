### How long did you spend on the technical test?
I spent my Friday evening and some of my Saturday (10:30am - 2pm) on this project.

### What would you add to your solution if you had more time?
**Testing:** I would like to add more tests throughout the stack. The backend requires more coverage around the `GetPostcode()` and `GetInput()` functions - I would add more tests cases to check whether the functions throw errors correctly.

I'd also like to test the `ApiHandler()` function. I did try to create some tests around this function, but I found myself blocked while creating mocks for the api calls (looked into the `go mock` package). To counter this, I was manually testing as I went along.

I'd add tests to the front-end, making sure that each function is outputting the correct data and manipulating the parts of the DOM on the UI.

**Frontend:** I used Vanilla JS on the Front-end because this application didn't require anything too complex. I would like to make some tweaks to the `display(data)` function. The `for` loop is not assigning the key/value pair for the `name` or `ratingAverage` quite how I'd like it. If I had much more time I would have sent the values into an empty array so the data could be presented much cleaner on the UI once the user clicks submit.

**API:** I wanted to create a fullstack application that sent several API requests between the JS Front-end, Golang Back-end and Just Eat API. I wanted to add a `POST` request from the Front-end to the Back-end which would contain postcode data. However, I was having issues with the Cross-Origin Data Transfer using Vanilla JS's `fetch` function, so I was unable to implement it.

**Error-Handling:** There are a couple of calls in the server which need more error handling. By adding these in, it would make it easier to de-bug in the future.

### What do you think you could research or learn more about to help you build a better solution, given more time?
**Testing:** As mentioned previously, I'd look into using the go `mock` package to generate the `APIHandler()` mocks. Currently, there are no unit tests or e2e tests that tests the connectivity between these API requests. I want to make sure that each request is producing the expected response.

In order to successfully test the Frontend, I'd look into the JS Testing Framework and apply it towards the current code base.

**API:** I would like to understand the `Mux` package a little more. I've used it as a `GET` Request in order to hit the `Just Eat API` endpoint. I wanted to use a `POST` request to fetch the users' postcode input from the UI. I looked at either sending it through as a query string on the URL or as a JSON response, but I feel like I needed more time to fully understand and successfully implement this feature. Failing that, I know Golang offers other frameworks/libraries that carry out similar functions, such as the `Echo` framework.

I was also getting blocked on the Front-end because of Cross-Origin Data Transfer issues. After doing some research, this problem could be resolved by deploying the application onto Heroku. Ideally, I didn't want to do this as I wanted to keep this application private, so I would like to find an alternative solution to this issue.

### How would you go about researching or learning the above?

My way of learning about new stacks, technologies, frameworks and libraries is usually through alot of **doing** rather than reading. I'd read around the topic, watch a few videos, read the technology documentation (if there is any) and then set up a mini project that will help me understand how it works.

Alternatively, I'd find a blog or guide that implements the technology I'm learning and follow it. I'd set up my own repo, type out each line and make sure the code is working as it should be. I'd then break and fix each line, to give me a good understanding of how the code works with each other. I find this less time-costuming than above, and helps with tight deadlines.