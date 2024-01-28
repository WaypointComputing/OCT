INSERT INTO user (name) VALUES ("Aiden");
INSERT INTO user (name) VALUES ("Linton");
INSERT INTO user (name) VALUES ("Kyle");
INSERT INTO user (name) VALUES ("Josh");
INSERT INTO user (name) VALUES ("Rhys");
INSERT INTO user (name) VALUES ("Caleb");
INSERT INTO user (name) VALUES ("Ben");

INSERT INTO quiz (id, name) VALUES (1, "Spiritual Gifts");
INSERT INTO quiz (id, name) VALUES (2, "Personality Test");

INSERT INTO trait (id, name, quiz_id) VALUES (1, "INFJ", 2);
INSERT INTO trait (id, name, quiz_id) VALUES (2, "INFP", 2);
INSERT INTO trait (id, name, quiz_id) VALUES (3, "INTJ", 2);
INSERT INTO trait (id, name, quiz_id) VALUES (4, "INTP", 2);
INSERT INTO trait (id, name, quiz_id) VALUES (5, "Giving", 1);
INSERT INTO trait (id, name, quiz_id) VALUES (6, "Serving", 1);
INSERT INTO trait (id, name, quiz_id) VALUES (7, "Sharing", 1);

INSERT INTO question (question, trait_id, quiz_id) VALUES ("How are you?", 1, 2);
INSERT INTO question (question, trait_id, quiz_id) VALUES ("Are you of death?", 2, 2);
INSERT INTO question (question, trait_id, quiz_id) VALUES ("Do you have of the stupid?", 3, 2);
INSERT INTO question (question, trait_id, quiz_id) VALUES ("Beans?", 4, 2);
INSERT INTO question (question, trait_id, quiz_id) VALUES ("Do you like giving gifts?", 5, 1);
INSERT INTO question (question, trait_id, quiz_id) VALUES ("Do you like serving?", 6, 1);
INSERT INTO question (question, trait_id, quiz_id) VALUES ("Do you like sharing?", 7, 1);
