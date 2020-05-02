# go-httpEmail-server

Sends email from server endpoint /send-email via go routine (No delay on send)

## Instructions

```bash
go build
```

Note your smtp configurations are defined in config.json so set them accordingly.
Run:

```bash
./sendEmail
listening on : 9090
```

Now visit localhost:9090 and fill out contact from. You will be redirected immediately and the sending process will run as go routing on server's side (localhost).

## Works on top of gmail (smtp.gmail port 587). In order to change to other smtp services change smtp.SendMail.addr

# Tests

First let's run the non-concurrent test with sends 5 emails one by one and then compare to our concurrent solution:
(Slow non-concurrent solution)

```bash
$ go test -run TestNonConcurrentSendEmail -v
=== RUN   TestNonConcurrentSendEmail
2018/10/02 20:21:20 Sent email...
2018/10/02 20:21:22 Sent email...
2018/10/02 20:21:24 Sent email...
2018/10/02 20:21:25 Sent email...
2018/10/02 20:21:27 Sent email...
--- PASS: TestNonConcurrentSendEmail (10.17s)
    send_email_test.go:20: Go routine num. 0
    send_email_test.go:20: Go routine num. 1
    send_email_test.go:20: Go routine num. 2
    send_email_test.go:20: Go routine num. 3
    send_email_test.go:20: Go routine num. 4
PASS
ok      github.com/go-httpEmail-server  10.184s
```

Which takes 10 seconds.
Now let's run our concurrent solution:

```bash
$ go test -run TestConcurrentSendEmail -v
=== RUN   TestConcurrentSendEmail
2018/10/02 20:26:21 Sent email...
2018/10/02 20:26:21 Sent email...
2018/10/02 20:26:21 Sent email...
2018/10/02 20:26:22 Sent email...
2018/10/02 20:26:22 Sent email...
--- PASS: TestConcurrentSendEmail (3.65s)
    send_email_test.go:12: Go routine num. 2
    send_email_test.go:12: Go routine num. 3
    send_email_test.go:12: Go routine num. 4
    send_email_test.go:12: Go routine num. 1
    send_email_test.go:12: Go routine num. 0
PASS
ok      github.com/go-httpEmail-server  3.708s
```

Almost 3 times faster!!
We can also notice the emails we receive come in different order than sent (subject 4 before subject 1) which assures our endpoint runs concurrently.
