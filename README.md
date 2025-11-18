# oivan_url_shortener

## Author: Tran Ha Nam

Since this can't be paginated so you can copy the content's name to go to it's section

## Table of contents

### 1. Overview

### 2. Problem

### 3. Solutions

### 4. Final thoughts

##

## 1. Overview

Hi I'm Nam, I've work as a Software Engineer for like more than 3 years. I've been mainly working with Golang, and for ReactJS I just know how to put stuff in the UI rather than optimizing the FE codes, like tree-shaking, building webpack <br>
This document is intended for you to understand my aspect/view when approaching your assignment
<br>
I was really happy for getting a chance to be a part of Oivan by doing assignment
<br>
Beware that this document is not written or generated in AI haha so I just want you to know that maybe I wasn't that bright but I am serious when doing this assignment.<br>


## 2. Problem

So when approaching and reading the exercise, I know I have to built some kind of a service. <br>
Where the server would have to receive an input as an url, validate it, and then use some kind of algorithms or stuff doing whatever it takes to shorten that url, create a mapper between the original url and the shortened url and response it to the client.
<br>
The main problem is I've never worked with this issue before so I was kinda inexperience not knowing how should I write my code so that it can actually encode without being collision tho. <br>
Also I don't usually work with EC2, my main working experience is mainly with Lambda and stuffs related to it such as Cloudwatch, S3,... so there is stuff to learn when doing this assignment haha

## 3. Solution

For codebase, I've decided to built a simple yet effective clean architect. <br>
The codebase will mainly consists of: <br>
- main.go -> which is where the requests is received, there we can loading configs, declares routes, middlewares,... <br>
- tools -> tools for the sourcecode, like configs, dbutil, utils
- tests -> intergration tests. Noted that the unit tests will sit beside to it's module parent directory
- docs -> swagger API documentation
- .github -> for github action only
- internal -> consists of what will be the internal logics, business layer of the code
- interal/api -> responsible for handling declaring routes which receive from parent group declared in the main.go
- internal/model -> responsible for declaring model that interact directly with the database
- internal/dto > responsible for declaring structs, mapping between the http layer into the input of the service's layer and the response from the service into the http layer
- internal/repository -> responsible for connecting to the database, insert, delete, etc manipulate with the database, the repository will not handle business logic
- internal/service -> responsible for handling the business logic talks with the repository, and receive the input that was mapped in http layer and response to it <br>

So to solve this exercise we must build 3 apis: 1 to encode, 1 to decode and the last one is for redirection <br>
There is some approaches which comes in my mind and when I googling stuff:<br>

1. ### Worker pool <br>
You can build some kind of a worker pool, where you actually already encode says like 10~50 encoded urls. Then if the requests hit the server, it just auto assigned to an encoded url in the pool. <br>
With this approach there will be an issue of what if the requests per seconds is too high, will the worker pool able to generate fastly enough to be ready to be assigned back for the request? <br>
Imo it will hard to ensure the speed of the encoded url generation racing with the speed of the request per second made by the users tho... and the worker have to run 24/7 <br> 
But the tradeoff is that you don't have to wait to encode the url, which the speed of the encode api will be extremely fast

2. ### Auto-increment id and base62 encoding <br>
Really straight forward solution where you will insert a record into the database gets the id and then use that id to encode with base62 -> returns the shorten url
With this approach there will be some issues <br>
Firstly is the db auto increment will be bottle-neck at scaling depending on the numbers of requests per section = query per second. and then also we have to calculate based on the CPU provided for the database based on the below formular<br>
```
Queries per Second (QPS) ≈ (1 / Average Query Runtime in Seconds) × Number of Cores
```
So this means that if the query per second is too high + the number of CPU cores didn't meet the demand traffic that the server is receiving, the api encode will be slow to response to the user since it is bottle-necking at the creating of the record <br>
Will be hard for you to scale this solution with a "horizontal" way means that you just simply add more db instaces, but the thing you would need a global ID generator otherwise there will be collision with the same id between instances of db. <br>
And lastly if the db dead -> that api is dead too. <br>
But the trade off is that the api is very simple, clean, and most important it won't have collision. With the postgres, a sql db, it have ACID in it, which means that even with millions of concurrent insert, the database will always guarrantee a unique database. <br>

3. ### SnowflakeID generator <br>
I know for a fact that this approach is used by Twitter, Tiktok,... themself for shorten the url. Mainly for sharing stuff, you can try this by going Tiktok, click on sharing a video and then choose messenger, or zalo stuff to share. There you can see the shorten url when you send to your friend. <br>
Their algorithm to generate is some kind like this <br>
![alt text](https://images.viblo.asia/ecb36cf4-3bec-43a7-bfed-304f7cdede2b.png) <br>
So it will be combination like this <br>
`timestamp | datacenter | worker | sequence` <br>
and then it will convert that 64 bit ID to Base62 as in shorten url <br>
With this approach, this is I think this is the best to built a shorten url service since it will always generate a globally unique id between nodes, it is perfect for scaling, sharding,... anything you can think of, no need database locking, bottlenecking,...  <br>
The downside is that to implement this shorten algorithm, it is really complex since it will need <br>
Firstly is the clock drift, clock drift means that the host pc, server can run faster or slower compared to the actual time due to hardware's issue or OS, host issues, time adjustment... etc <br>
For example with even the server's clock drift backward by 1 millisecond the timestamp part becomes smaller than the previous IDs which can cause collision, breaking the uniqueness <br>
Same with the multi-regional, their clock can have a slightly difference <br>
The snowflake generates 4096 IDs per millisecond (says for the 12-bit sequence), then what if there is a super high traffic at that node? Says 50K-70K at 1ms (I know it is really hard to happen due to load balancer but just want to say) then there will be no new next ids until the very next millisecond which can cause panic to those requests.

### Conclusion

So by overviewing these solutions, I choose to implement the auto-increment id and base62 encoding since it will be really simple and straight-forward. Would be easy to read too, and of course "demo state" <br>
And won't be bottle-neck issue at the db <br>

## 4. Final thoughts

If there is more time, I would build like to build migrations, currently it's just a simple auto-migrate to migrate to models into database <br>
I would usually use gormigrate for this, or just write by myself mi-micking that library <br>
And also adding metrics, more meaningful loggings <br>
This assignment is a fun trip for me to experience building shortener service <br>
If you truly read all of these, thank you for spending your time reading the document that I wrote <br>