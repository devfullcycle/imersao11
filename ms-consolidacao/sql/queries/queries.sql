-- name: CreateAction :exec
INSERT INTO actions (id, match_id, team_id, player_id, action, score, minute) VALUES (?, ?, ?, ?, ?, ?,?);

-- name: CreatePlayer :exec
INSERT INTO players (id, name, price) VALUES (?, ?, ?);

-- name: UpdatePlayer :exec
UPDATE players SET name = ?, price = ? WHERE id = ?;

-- name: UpdateMyTeamsScore :exec
UPDATE my_team SET score = ? WHERE id IN (?);

-- name: AddScoreToTeam :exec
UPDATE my_team SET score = score + ? WHERE id = ?;

-- name: RemoveActionFromMatch :exec
DELETE FROM actions WHERE match_id = ?;

-- name: CreateMatch :exec
INSERT INTO matches (id, match_date, team_a_id, team_a_name, team_b_id, team_b_name, result) VALUES (?, ?, ?, ?, ?, ?, ?); 

-- name: GetMatchActionsForUpdate :many
SELECT * FROM actions WHERE match_id = ? FOR UPDATE;

-- name: GetMatchActions :many
SELECT * FROM actions WHERE match_id = ?;

-- name: FindMyTeamById :one
SELECT * FROM my_team WHERE id = ?;

-- name: FindMyTeamByIdForUpdate :one
SELECT * FROM my_team WHERE id = ? FOR UPDATE;

-- name: FindTeamById :one
SELECT * FROM teams WHERE id = ?;

-- name: FindPlayerById :one
SELECT * FROM players WHERE id = ?;

-- name: FindPlayerByIdForUpdate :one
SELECT * FROM players WHERE id = ? FOR UPDATE;

-- name: CreateMyTeam :exec
INSERT INTO my_team (id, name, score) VALUES (?, ?, ?);

-- name: FindAllPlayers :many
SELECT * FROM players;

-- name: GetPlayersByMyTeamID :many
SELECT p.* FROM players p INNER JOIN my_team_players mtp ON p.id = mtp.player_id WHERE mtp.my_team_id = ?;

-- name: FindAllPlayersByIDs :many
SELECT * FROM players WHERE id IN (?);

-- name: DeleteAllPlayersFromMyTeam :exec
DELETE FROM my_team_players WHERE my_team_id = ?;

-- name: AddPlayerToMyTeam :exec
INSERT INTO my_team_players (my_team_id, player_id) VALUES (?, ?);

-- name: FindMatchById :one
SELECT * FROM matches WHERE id = ?;

-- name: FindAllMatches :many
SELECT * FROM matches;

-- name: UpdateMatch :exec
UPDATE matches SET match_date = ?, team_a_id = ?, team_a_name = ?, team_b_id = ?, team_b_name = ?, result = ? WHERE id = ?;

-- name: GetMyTeamBalance :one
SELECT score as balance FROM my_team WHERE id = ?;

-- name: UpdateMyTeamScore :exec
UPDATE my_team SET score = ? WHERE id = ?;