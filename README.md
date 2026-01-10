# RatioArr

RatioArr is a personal project designed to help users monitor and manage their upload and download ratios across various torrent trackers. By periodically crawling these trackers, RatioArr fetches user statistics and provides insights into their sharing habits.

## Main Objectives

- [ ] **Tracker Crawling**: Automatically connect to multiple torrent trackers to retrieve user data.
- [ ] **Ratio Monitoring**: Calculate and display upload/download ratios to help users maintain good standing on
- [ ] **Arr integration**: Integrate with media management tools like Radarr and Sonarr to enable disable tracker based on ratio.
- [ ] **Prometheus Metrics**: Expose metrics for monitoring and alerting purposes.

Work in Progress

## Tool requires to dev

- air:  for live reloading during development
- sqlc: for generating type-safe database queries
- goose: for database migrations
- templ: for generating webpage files from templates

### todo list

- [ ] Add more trackers support
- [ ] Improve error handling and logging
- [X] Setup db migrations
- [ ] Rework context usage across the app
- [ ] Rework logs to use structured logging
- [ ] Write documentation to :
  - [ ] help users set up and use the tool
  - [ ] help developers understand the codebase and contribute to it

## Available Trackers

- [X] UNIT3D (unit3d is a popular open-source torrent tracker software used by various private torrent communities.)
- [ ] YggTorrent
- [ ] la-cale
