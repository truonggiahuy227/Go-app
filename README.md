# Prepare for running this project
Install Golang
Follow this [guide](https://go.dev/doc/install)
# Running project
Run command bellow to build:
<pre><code>$ go build main.go
</code></pre>
It will create binary file inside project folder with name: main
Run command bellow for running:
<pre><code>$ ./main.go
</code></pre>

Output:

<pre><code>
➜  arrow_service_akw git:(master) ✗ go run main.go  
Service RUN on DEBUG mode
INFO[0000] [BACKEND][arrow-service][Main_Info]Cache initializing.... 
INFO[0000] [BACKEND][arrow-service][Main_Info]Cache finished... 

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v3.3.10-dev
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:9090
</code></pre>

Go to Browser with [this link](http://localhost:9090/arrow/api/v1.0/greeting) for check web working:

Output:
<pre><code>{"Message":"Success","Data":{"Color":"Hello member of training Docker, K8s organize by xPlat Team!!!!!!"}}
</code></pre>

Require for tranning Docker
- 1. Need to size of image docker less than: 30 MB
- 2. Need use non-root account
  + Check by call to [API](http://localhost:9090/arrow/api/v1.0/user)

  Bad Answer: Good luck to you next time!!!!!!
  
  Good Answer: Congratulation!!!!!!

## Optional 
- 3. Need to connect with another container Redis Server (API: is working)
  + First run Redis container with command:
  + Next, Add Redis URL to REDIS_URL env in Docker image env
  + Check by call to [API](http://localhost:9090/arrow/api/v1.0/communicate)
    + Good Result: {"Message":"Success","Data":{"Flag":"2024"}}
    + Bad Result: {"Message":"Success","Data":{"Flag":""}}
