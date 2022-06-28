# channelsTest

Simple project to handle processing multiple lines from a log file using go routines and channels. Each entry from the log file represents an HTTP request. Following is an example of different entries:
```
213.30.63.110 - - [14/Jan/2020:14:17:07 +0000] "POST /wp-admin/admin-ajax.php HTTP/1.1" 200 116 "https://meujornal.pt/wp-admin/post.php?post=3635570&action=edit" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36"
95.92.237.182 - - [14/Jan/2020:14:17:07 +0000] "GET /wp-json/obs_api/v4/news/widget HTTP/1.1" 301 178 "-" "Today/201908022200 CFNetwork/978.0.7 Darwin/18.7.0"
```

## Run

```
$ go build -o channels
$ ./channels  log/file/path
```


## Makefile

A simple makefile can also be used instead of go commands. If you use only make or make all, passing the log file path in the log variable is required. 

```
$ make LOG=files/golang.log

```


## Files

The log files used during tests can be found in files/golang.log 

## Important

This project is inspired by a practical test for a senior job position. The full details can be accessed here: https://www.youtube.com/watch?v=u3GAYK8hvTI&ab_channel=Filhodanuvem (Language: Portuguese)




