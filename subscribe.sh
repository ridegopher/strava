#/bin/bash

echo curl -X POST https://api.strava.com/api/v3/push_subscriptions \
      -F client_id=${STRAVA_CLIENT_ID} \
      -F client_secret=${STRAVA_CLIENT_SECRET} \
      -F 'callback_url=https://api.ridegopher.com/strava/event' \
      -F verify_token=${STRAVA_VERIFY_TOKEN}

