### How long did you spend on the technical test?
I spent my Friday evening and some of my Saturday (10:30am - 2pm) on this project.

### What would you add to your solution if you had more time?
**Testing:** I would like to add more tests throughout the stack. I would like to add more tests to the `Validate()` function to check more edge cases around the postcode. The tests are currently failing even though the expected output is correct. I'd perhaps look to use the `assert` package instead of using Golangs table unit tests template.

I'd also like to test the `ApiHandler()` function. I did try to create some tests around this function, but I found myself blocked while creating mocks for the api calls (looked into the `go mock` package). To counter this, I was manually testing as I went along.

I'd add tests to the front-end, making sure that each function is outputting the correct data and manipulating the parts of the DOM on the UI.

**Error-Handling:** There are a couple of error messages in the `APIHandler()` function which need to be handled. By adding these in, it would make it easier to de-bug in the future.

### What do you think you could research or learn more about to help you build a better solution, given more time?
**Testing:** As mentioned previously, I'd look into using the go `mock` package to generate the `APIHandler()` mocks. Currently, there are no unit or e2e tests that test the connectivity between the API requests. I want to make sure that each request is producing the expected response.

In order to successfully test the Frontend, I'd look into the JS Testing Framework and apply it towards the current code base. It's been a while since I've used this framework, so I'd probably need a bit more time to re-fresh myself.

**API:** Even though I used the `Mux` package to hit the API endpoint, I would have liked to compare its capabilities to the `Echo` framework.

### How would you go about researching or learning the above?

My way of learning about new stacks, technologies, frameworks and libraries is usually through alot of **doing** rather than reading. I'd read around the topic, watch a few videos, read the technology documentation (if there is any) and then set up a mini project that will help me understand how it works.

Alternatively, I'd find a blog or guide that implements the technology I'm learning and follow it. I'd set up my own repo, type out each line, and make sure the code is working as it should be. Following that, I'd break and fix each line, to give me a good understanding of how the code works with each other. I find this less time-costuming than above, and helps with tight deadlines.