# Challenge Summary

Marvel Entertainment exposes a REST api which can be used to retrieve all kinds of data about their comic universe. We will write a simple client to fetch data from this service. To complete the challenge, you will need to write code to make basic HTTP GET requests to the proper URLs, deserialize JSON responses into appropriate data structures, and return requested elements from the responses.

There are four exercises to complete. Each has a corresponding function in the provided code for you to fill out. Automated tests are provided to test your implementation. Passing the tests is important, but code will be reviewed for correctness and style, so returning the right answer without making the correct API call does not count 😃

## Assumptions

1. Accessing the API requires an API key which you can get by signing up for a free developer account. The testcode includes functionality to automatically sign requests (see provided skeleton code for an example), but it can be helpfulto sign up for your own free account. This will let you write and test code locally, and also lets you use their interactive API docs to build and send requests from the browser.
2. To get an API key, go to the API docs page and click "Get a Key Now". This will take you to a signin page, where you can create an account, agree to their TOS, and get an API key (and/or use the interactive docs to make requests). If you have signed up but still get an "missing API key" error in the interactive docs, click on the Sign In link at the top right (again).
3. Except where noted, for this exercise, we are not concerned about error handling, timeouts, etc., although those would be important in real life.

## Tasks

1. Create a function that retrieves the first 25 characters (ordered by name), and return their names as a slice of strings, and return their names as a slice of strings like []string{"Captain America", "Hulk", "Thor"}.
2. Create a function that accepts a character ID and returns its name. It should return the provided ErrNotFound if the character does not exist.
3. Create a function that takes a character's name and returns its id, description, thumbnail URL, and number of comics it appears in in the provided characterDetail struct. It should return ErrNotFound if no character is found.
4. Get a list of all the characters in the comic with id 41112 and return their names as a []string.
