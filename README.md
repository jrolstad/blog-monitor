
# Blog Monitor
Checking for new posts on blogging sites can be tedious and when found many are full of ads them make them difficult to read.  This project regularly polls specific sites and notifies one or more recipients with the full blog post contents so they can be read in a timely manner and without ads.

# Supported Platforms
The following platforms are supported
|Name|Status|
|---|---|
|[Blogger](https://www.blogger.com/)|Available|

## Primary Use Case
Inspired by https://cliffmass.blogspot.com/ , I want to be actively notified of new content and read without any advertisements

# How it Works
For every subscribed blog, posts are queried on a regular basis and recipients notified via email with the full blog post contents in the message.  Each post is only notified once - when published.

This is implemented using an AWS Lambda function that is invoked on a regular basis who then calls the Google Blogger APIs and queries for blog / post data.  Notifications are sent using the Amazon Simple Email Service with subscriptions and notification histories held in Amazon DynamoDb.

# How to Use
## Components
 ![Component Diagram](/docs/components.png)
 
## Instructions
To install, configure, and subscribe to 1..n blogs follow information in the [src_ReadMe](./src_README.md)

# License
This projects is made available under the [MIT License](LICENSE).
