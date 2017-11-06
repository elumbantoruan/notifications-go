# notifications-go

A set of framework to send SMS message via AWS SNS.  The component 
includes AWS Lambda (go), AWS SNS, and MongoDB for storage

### Notes:
1. handler.go contains Handler function which is the entry point 
for AWS Lambda
2. The lambda will be invoked by CloudWatch scheduler
http://docs.aws.amazon.com/AmazonCloudWatch/latest/events/RunLambdaSchedule.html


### Setup:

1. AWS Lambda go shim
   https://github.com/eawsy/aws-lambda-go-shim
2. AWS SDK Go 
   https://docs.aws.amazon.com/sdk-for-go/api/
3. MongoDB driver
   http://gopkg.in/mgo.v2
4. MongoDB Atlas
   1. https://www.mongodb.com/cloud/atlas
   2. Get connection string from Atlas cluster connection 
5. UUID
   https://github.com/google/uuid
   
   
### References:

1. MongoDB examples:
    1. https://gist.github.com/border/3489566
    2. http://labix.org/mgo
    