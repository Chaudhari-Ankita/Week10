# Week10
Creating portfolio using go and docker

# How to run the Go application:
1. git clone https://github.com/Chaudhari-Ankita/Week10.git
2. cd github.com/Chaudhari-Ankita/Week10
3. go run main.go
4. Add, commit and push the changes

# How to build the Cloud Engineer Portfolio Docker image
1. git pull
2. docker build -t myPortfolio:0.1 .

# How to run the Cloud Engineer Portfolio Docker image
1. docker run -p 80:80 myPortfolio:0.1

# how to push the image to the docker hub
1. docker login
2. docker tag myPortfolio:0.1 chaudhariankita/myPortfolio:0.1
3. docker push chaudhariankita/myPortfolio:0.1