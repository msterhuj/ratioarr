-- name: GetAllTrackerStats :many
SELECT * FROM tracker_stats;

-- name: InsertTrackerStat :exec
INSERT INTO tracker_stats (name, type, uploaded, downloaded, ratio)
VALUES (?, ?, ?, ?, ?);
