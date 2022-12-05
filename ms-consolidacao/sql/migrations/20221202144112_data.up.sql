INSERT INTO players (id, name, price) VALUES ('1', 'Neymar 1', 10.00);
INSERT INTO players (id, name, price) VALUES ('2', 'Pelé', 10.00);
INSERT INTO players (id, name, price) VALUES ('3', 'Ronaldo', 10.00);
INSERT INTO players (id, name, price) VALUES ('4', 'Messi', 10.00);
INSERT INTO players (id, name, price) VALUES ('5', 'Cristiano Ronaldo', 10.00);
INSERT INTO players (id, name, price) VALUES ('6', 'Neymar 2', 10.00);
INSERT INTO players (id, name, price) VALUES ('7', 'Zico', 10.00);
INSERT INTO players (id, name, price) VALUES ('8', 'Romário', 10.00);
INSERT INTO players (id, name, price) VALUES ('9', 'Ronaldinho', 10.00);
INSERT INTO players (id, name, price) VALUES ('10', 'Rivaldo', 10.00);

INSERT INTO teams (id, name) VALUES ('1', 'Flamengo');
INSERT INTO teams (id, name) VALUES ('2', 'Vasco');
INSERT INTO teams (id, name) VALUES ('3', 'Botafogo');
INSERT INTO teams (id, name) VALUES ('4', 'Fluminense');
INSERT INTO teams (id, name) VALUES ('5', 'Corinthians');
INSERT INTO teams (id, name) VALUES ('6', 'São Paulo');
INSERT INTO teams (id, name) VALUES ('7', 'Palmeiras');
INSERT INTO teams (id, name) VALUES ('8', 'Santos');
INSERT INTO teams (id, name) VALUES ('9', 'Grêmio');

INSERT INTO team_players (team_id, player_id) VALUES ('1', '1');
INSERT INTO team_players (team_id, player_id) VALUES ('1', '2');
INSERT INTO team_players (team_id, player_id) VALUES ('1', '3');
INSERT INTO team_players (team_id, player_id) VALUES ('1', '4');
INSERT INTO team_players (team_id, player_id) VALUES ('2', '5');
INSERT INTO team_players (team_id, player_id) VALUES ('2', '6');
INSERT INTO team_players (team_id, player_id) VALUES ('2', '7');

INSERT INTO my_team (id, name, score) VALUES ('1', 'Time 1', 100);

INSERT INTO my_team_players (my_team_id, player_id) VALUES ('1', '1');
INSERT INTO my_team_players (my_team_id, player_id) VALUES ('1', '2');
INSERT INTO my_team_players (my_team_id, player_id) VALUES ('1', '3');
INSERT INTO my_team_players (my_team_id, player_id) VALUES ('1', '4');

INSERT INTO matches (id, match_date, team_a_id, team_a_name, team_b_id, team_b_name, result) VALUES ('1', '2021-12-02 00:00:00', '1', 'Flamengo', '2', 'Vasco', '1-0');
INSERT INTO matches (id, match_date, team_a_id, team_a_name, team_b_id, team_b_name, result) VALUES ('2', '2021-12-02 00:00:00', '3', 'Botafogo', '4', 'Fluminense', '1-0');

-- insert match actions
INSERT INTO actions (id, match_id, team_id, player_id, action, score, minute) VALUES ('1', '1', '1', '1', 'goal', 1, 10);
INSERT INTO actions (id, match_id, team_id, player_id, action, score, minute) VALUES ('2', '1', '1', '2', 'goal', 1, 20);
INSERT INTO actions (id, match_id, team_id, player_id, action, score, minute) VALUES ('3', '1', '1', '3', 'goal', 1, 30);
INSERT INTO actions (id, match_id, team_id, player_id, action, score, minute) VALUES ('4', '1', '1', '4', 'goal', 1, 40);