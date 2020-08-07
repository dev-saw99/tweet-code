# Tweet-Code

Tweet-code is bot which uses twitter API to share progress of your *#100daysleetcode* challenge.  It shares your progress on twitter as a tweet and motivates you to keep doing the goodwork. 

## Requirements

* Golang
* git
* Heroku CLI
* docker

## How to integrate this bot to your twitter account

To integrate this bot to your twitter you need to follow these four simple steps:

* Make a developer account on twitter
* Clone this project
* Configure the BOT 
* Host it in cloud

### Step 1 : Developer account

* Go to [this](https://developer.twitter.com/) link and login with twitter account.
* Create a app from Projects and Apps section
* Generate following keys:
    * Consumer Key
    * Consumer Secret
    * Access Token
    * Access Secret 

**Note** Please make sure that you dont share / commit these key in any public reporitory.

### Step 2 : Clone this project

You can directly download this project from [here](https://github.com/dev-saw99/tweet-code/archive/master.zip)

or

You can use following commands to from your command line:

(Make sure you have git installed in your 
system)

```
git init

git clone "https://github.com/dev-saw99/tweet-code"

```
### Step 3 : Configure the BOT

For this step you need a account at [leetcode](https://leetcode.com)

* Create your account at leetcode
* Go to you profile and copy the URL of your profile page, It should look like :
``` https://leetcode.com/your_user_name```
* Open project you just downloaded in Step 2 and perform following operations
    * Go to `main.go`
    * Search for main function (should be at line 166)
    * You will see `url` variable in main function at line `170`
    * Replace the content of `url` variable with the url to you profile. Make sure they are in inverted commas.
    * Open `credentials.json` and replace all the keys with the keys you genereated in Step 1

**Note** If you aren't new a user please go to your profile check the numberofquestion you have solved. Go to `data.json` and replace `0` with the correct detail



#### Now you are good to go, you can run your bot using `go run main.go`

### Step 4 : Host in cloud

We will use [heroku](https://www.heroku.com/) to host this project, but you can opt for any cloud service  provider.

* Create a account in `heroku`
* Open terminal and run following command

```
heroku container:login

heroku create

heroku container:push web

heroku container:release web

```






