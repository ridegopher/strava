# Serverless Strava Webhooks

[![Go Report Card](https://goreportcard.com/badge/github.com/ridegopher/strava)](https://goreportcard.com/report/github.com/ridegopher/strava)  

This is an implementation of the Strava Webhook Events API  
See: https://developers.strava.com/docs/webhooks      

### Set up secrets in Parameter Store
```bash
aws ssm put-parameter --name STRAVA_SUBSCRIPTION_ID --type String --value=<your_subscription_id>
aws ssm put-parameter --name STRAVA_VERIFY_TOKEN --type String --value=<your_made_up_token>
```

### Install Serverless Framework
https://serverless.com/framework/docs/providers/aws/guide/installation/

### Go
Figure it out!

### Install and run serverless
```bash

npm install
sls deploy

```