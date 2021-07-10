# xendit

Software Engineer - Technical Assessment

## Running program

1. clone this repo
2. make run
3. use postman collection included to test it

## Caching all characters strategy

1. Init (hit marvel api to get all characters) on binary start (async)
2. Cron hourly to refresh the data cache (incase of new characters)
3. If endpoint get all character is hit before cache is complete, it will return error
