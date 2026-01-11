-- name: GetAllTrackerStats :many
SELECT * FROM tracker_stats;

-- name: GetLatestTrackerStats :many
SELECT * FROM tracker_stats 
WHERE (name, timestamp) IN (
    SELECT name, MAX(timestamp) 
    FROM tracker_stats 
    GROUP BY name
);

-- name: InsertTrackerStat :exec
INSERT INTO tracker_stats (name, type, uploaded, downloaded, ratio)
VALUES (?, ?, ?, ?, ?);

