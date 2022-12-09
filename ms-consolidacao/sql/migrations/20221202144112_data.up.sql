-- INSERT INTO players (id, name, price) VALUES ('4876d14f-d998-4abf-96ef-89fd53185464', 'Cristiano Ronaldo', 10.00);
-- INSERT INTO players (id, name, price) VALUES ('0b8f08d8-d871-4a42-b395-17d698f477db', 'De Bruyne', 10.00);
-- INSERT INTO players (id, name, price) VALUES ('0c9ba4fb-4609-464d-9845-421ca1e1e3bd', 'Harry Kane', 10.00);
-- INSERT INTO players (id, name, price) VALUES ('67fbf409-d94f-4858-8423-8043576cda05', 'Lewandowski', 10.00);
-- INSERT INTO players (id, name, price) VALUES ('c7830b65-cf79-49b7-a878-82250fec1d94', 'Maguirre', 10.00);
-- INSERT INTO players (id, name, price) VALUES ('64fb9c2f-a45b-4f96-9d8b-b127878ca6f3', 'Messi', 10.00);
-- INSERT INTO players (id, name, price) VALUES ('0f463bea-1dbd-4765-b080-9f5f170b6ded', 'Neymar', 10.00);
-- INSERT INTO players (id, name, price) VALUES ('5ce233a8-5cd8-4a85-8156-9ac255cf909e', 'Richarlison', 10.00);
-- INSERT INTO players (id, name, price) VALUES ('c707bfa9-074e-4636-8772-633e4b56248d', 'Vinicius Junior', 10.00);

INSERT INTO teams (id, name) VALUES ('76ef612d-acdf-48cd-af1c-9313d16ab642', 'Argentina');
INSERT INTO teams (id, name) VALUES ('b30cf8e9-d02e-408e-9494-c4ae95b85b49', 'Alemanha');
INSERT INTO teams (id, name) VALUES ('4da19d91-c700-4865-8333-80efc789a70f', 'Brasil');
INSERT INTO teams (id, name) VALUES ('52f02f8d-bf04-4bd2-8b3d-28ec97e562a8', 'Bélgica');
INSERT INTO teams (id, name) VALUES ('b503dc80-1f8c-4c8e-8218-a37049676815', 'Portugal');
INSERT INTO teams (id, name) VALUES ('322dc42e-33a0-42cb-8c50-5d1e5fc0a308', 'Polônia');
INSERT INTO teams (id, name) VALUES ('ace3abc1-003d-4baf-8d93-51135c437c7a', 'Inglaterra');

-- /** fazer implementação com verificação da relação de teams com players **/
-- INSERT INTO team_players (team_id, player_id) VALUES ('b503dc80-1f8c-4c8e-8218-a37049676815','4876d14f-d998-4abf-96ef-89fd53185464');
-- INSERT INTO team_players (team_id, player_id) VALUES ('52f02f8d-bf04-4bd2-8b3d-28ec97e562a8','0b8f08d8-d871-4a42-b395-17d698f477db');
-- INSERT INTO team_players (team_id, player_id) VALUES ('ace3abc1-003d-4baf-8d93-51135c437c7a','0c9ba4fb-4609-464d-9845-421ca1e1e3bd');
-- INSERT INTO team_players (team_id, player_id) VALUES ('322dc42e-33a0-42cb-8c50-5d1e5fc0a308','67fbf409-d94f-4858-8423-8043576cda05');
-- INSERT INTO team_players (team_id, player_id) VALUES ('ace3abc1-003d-4baf-8d93-51135c437c7a','c7830b65-cf79-49b7-a878-82250fec1d94');
-- INSERT INTO team_players (team_id, player_id) VALUES ('76ef612d-acdf-48cd-af1c-9313d16ab642','64fb9c2f-a45b-4f96-9d8b-b127878ca6f3');
-- INSERT INTO team_players (team_id, player_id) VALUES ('4da19d91-c700-4865-8333-80efc789a70f','0f463bea-1dbd-4765-b080-9f5f170b6ded');
-- INSERT INTO team_players (team_id, player_id) VALUES ('4da19d91-c700-4865-8333-80efc789a70f','5ce233a8-5cd8-4a85-8156-9ac255cf909e');
-- INSERT INTO team_players (team_id, player_id) VALUES ('4da19d91-c700-4865-8333-80efc789a70f','c707bfa9-074e-4636-8772-633e4b56248d');

INSERT INTO my_team (id, name, score) VALUES ('22087246-01bc-46ad-a9d9-a99a6d734167', 'Meu Time FC', 100);

-- INSERT INTO my_team_players (my_team_id, player_id) VALUES ('22087246-01bc-46ad-a9d9-a99a6d734167', '4876d14f-d998-4abf-96ef-89fd53185464');
-- INSERT INTO my_team_players (my_team_id, player_id) VALUES ('22087246-01bc-46ad-a9d9-a99a6d734167', '0b8f08d8-d871-4a42-b395-17d698f477db');
-- INSERT INTO my_team_players (my_team_id, player_id) VALUES ('22087246-01bc-46ad-a9d9-a99a6d734167', '0c9ba4fb-4609-464d-9845-421ca1e1e3bd');
-- INSERT INTO my_team_players (my_team_id, player_id) VALUES ('22087246-01bc-46ad-a9d9-a99a6d734167', '67fbf409-d94f-4858-8423-8043576cda05');

-- INSERT INTO matches (id, match_date, team_a_id, team_a_name, team_b_id, team_b_name, result) VALUES ('97ed7b97-ea2f-4bab-80e6-e14800e4bfcc', '2021-12-02 00:00:00', '76ef612d-acdf-48cd-af1c-9313d16ab642', 'Argentina', 'b30cf8e9-d02e-408e-9494-c4ae95b85b49', 'Alemanha', '1-0');
-- INSERT INTO matches (id, match_date, team_a_id, team_a_name, team_b_id, team_b_name, result) VALUES ('a9d93ef1-7f1a-489f-89a3-d0321735c326', '2021-12-02 00:00:00', '4da19d91-c700-4865-8333-80efc789a70f', 'Brasil', '52f02f8d-bf04-4bd2-8b3d-28ec97e562a8', 'Bélgica', '1-0');

-- -- insert match actions
-- INSERT INTO actions (id, match_id, team_id, player_id, action, score, minute) VALUES ('ea32249c-4474-4e56-b6b6-747573c43552', '97ed7b97-ea2f-4bab-80e6-e14800e4bfcc', '64fb9c2f-a45b-4f96-9d8b-b127878ca6f3', '4da19d91-c700-4865-8333-80efc789a70f', 'goal', 1, 20);
-- INSERT INTO actions (id, match_id, team_id, player_id, action, score, minute) VALUES ('24420fc7-9e3f-4b9a-af50-025f4b857c9e', 'a9d93ef1-7f1a-489f-89a3-d0321735c326', '4da19d91-c700-4865-8333-80efc789a70f', '5ce233a8-5cd8-4a85-8156-9ac255cf909e', 'yellow card', 1, 5);
-- INSERT INTO actions (id, match_id, team_id, player_id, action, score, minute) VALUES ('b2c3e094-0c73-4a18-a1aa-b91b42124627', 'a9d93ef1-7f1a-489f-89a3-d0321735c326', '4da19d91-c700-4865-8333-80efc789a70f', 'c707bfa9-074e-4636-8772-633e4b56248d', 'assist', 1, 10);
-- INSERT INTO actions (id, match_id, team_id, player_id, action, score, minute) VALUES ('a315fee9-55de-4b5b-aed7-9ff7dee268dc', 'a9d93ef1-7f1a-489f-89a3-d0321735c326', '4da19d91-c700-4865-8333-80efc789a70f', '0f463bea-1dbd-4765-b080-9f5f170b6ded', 'goal', 1, 10);