# Exercise – Multiple HTTP downloads &amp; merging results

### The objective is to write a small software that:
- Takes a list of URLs, for example raw github readmes:
```
raw.githubusercontent.com/GoogleContainerTools/distroless/main/java/README.md
raw.githubusercontent.com/golang/go/master/README.md
```
(list can be hardcoded)
- Download each URL, checking for errors
- In case of any error with any URL, we can ignore other downloads and fail (exit by printing the error)
- If all downloads succeeded, the contents of all downloads are appended together and printed in *reverse order* from the list of URLs.

If working in Go, it’s a good idea to follow the Go Tour: A Tour of Go

Extra improvements:

- Make an HTTP server instead of printing to screen that when called (via GET) returns the downloaded texts as response.
- Download the responses in parallel, cancelling all downloads when any error is received.

The resulting code can be published on GitHub.

## Solution

I decided to put list of urls into a separate file named `urls.txt`. I also did the extra improvements. Port for HTTP server as well as name of file with urls can be configured in `.env` file.

You can run `main.go`. It creates `GET` endpoint `/files`, which would return the downloaded contents when called. Each time the endpoint is called the script reads the list of urls and fetches their content in parallel, after that it assigns each file unique uuid, for example, `c33e0d27-969f-4b26-a1bc-226cc197af6f`, and returns the files and their contents in reversed order in JSON format.

Finally, the script sets up an HTTP server.

The downloaded contents in reversed order by default are accessible at

```
GET: localhost:8080/files
```

If script encounters any errors, they are printed to the console using `panic` and execution is stopped.