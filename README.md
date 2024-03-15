## RedditDB
This system currently only stores the active and subscribe users of the top 3000 most subscribed subreddits. I plan on expanding the datas I store within the free reddit API restrictions. 
It's only charting the users for now, for the purpose of letting you know when to post in order to have the most feedback.
## Techs
- Go
- go-chi
- react.js (ts)
- chart.js
## Requirements
- Go 1.21
- npm
- node.js
## Running yourself
1. Clone the repo
2. Make your .env out of the `.env.example` (Need to setup reddit APIs, follow [this]([https://www.example.com](https://github.com/reddit-archive/reddit/wiki/OAuth2)https://github.com/reddit-archive/reddit/wiki/OAuth2))
3. Docker compose up!
## TODOs
- [ ] Unit tests
- [ ] Fix the query cronjob stopping when deploying
## Contribution
This is just a side project of mine. However, feel free to open an issue for any suggestions or bug fixes!
