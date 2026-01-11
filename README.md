# RatioArr

RatioArr is a project designed to help users monitor and manage their upload and download ratios across various torrent trackers.
By periodically crawling these trackers, RatioArr fetches user statistics and provides insights into their sharing habits to help maintain a healthy ratio.

## Main Objectives

- [X] **Tracker Crawling**: Automatically connect to multiple torrent trackers to retrieve user data.
- [X] **Ratio Monitoring**: Calculate and display upload/download ratios to help users maintain good standing on
- [ ] **Arr integration**: Integrate with media management tools like Radarr and Sonarr to enable disable tracker based on ratio.
- [ ] **Prometheus Metrics**: Expose metrics for monitoring and alerting purposes.
- [ ] **Rules**: Implement customizable rules to automate actions based on ratio thresholds.

Work in Progress

## Tech Stack

- Go
- Air (live reloading)
- SQLC (type-safe queries)
- Goose (migrations)
- Templ (templates)
- Tailwind CSS
- DaisyUI (UI components)

### todo list

- [ ] Improve error handling and logging
- [ ] Rework context usage across the app
- [ ] Write documentation to :
  - [ ] help users set up and use the tool
  - [ ] help developers understand the codebase and contribute to it

## Available Trackers / To Implement

- [X] UNIT3D (unit3d is a popular open-source torrent tracker software used by various private torrent communities.)
- [X] YggTorrent
- [ ] la-cale (cant implement it with api token need to login get cookies and use them)
