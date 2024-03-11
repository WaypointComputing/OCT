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

INSERT INTO user VALUES (1, "Aiden",  "aiden@southern.edu",  "b36a2e39e6478858acb5bdb5ac238951e195e353f43b4ddc1e6e099e14345141cc799390af4a7ee24710deb6ee819c1994b47560414209b2f432e2b487f508f1", 3);
INSERT INTO user VALUES (2, "Linton", "Linton@southern.edu", "ee1067d2c54d8b095bb7b3937aa40968cc3475e4360433a8bf816217e823271fcc9e7222dd9e57afb5675d999b88f49574ed8e6a3833b1437910e9aba7b6575f", 3);
INSERT INTO user VALUES (3, "Kyle",   "Kyle@southern.edu",   "ee1067d2c54d8b095bb7b3937aa40968cc3475e4360433a8bf816217e823271fcc9e7222dd9e57afb5675d999b88f49574ed8e6a3833b1437910e9aba7b6575f", 3);
INSERT INTO user VALUES (4, "Rhys",   "Rhys@southern.edu",   "ee1067d2c54d8b095bb7b3937aa40968cc3475e4360433a8bf816217e823271fcc9e7222dd9e57afb5675d999b88f49574ed8e6a3833b1437910e9aba7b6575f", 2);
INSERT INTO user VALUES (5, "Caleb",  "Caleb@southern.edu",  "ee1067d2c54d8b095bb7b3937aa40968cc3475e4360433a8bf816217e823271fcc9e7222dd9e57afb5675d999b88f49574ed8e6a3833b1437910e9aba7b6575f", 2);
INSERT INTO user VALUES (6, "Josh",   "Josh@southern.edu",   "ee1067d2c54d8b095bb7b3937aa40968cc3475e4360433a8bf816217e823271fcc9e7222dd9e57afb5675d999b88f49574ed8e6a3833b1437910e9aba7b6575f", 1);
INSERT INTO user VALUES (7, "Ben",    "Ben@southern.edu",    "ee1067d2c54d8b095bb7b3937aa40968cc3475e4360433a8bf816217e823271fcc9e7222dd9e57afb5675d999b88f49574ed8e6a3833b1437910e9aba7b6575f", 1);

INSERT INTO blog VALUES (1, 4, "Python is the best language", "Learn why python is the greatest programming language of all time", "test.file");
INSERT INTO blog VALUES (2, 1, "Rust is the best language", "Learn why rust is the greatest programming language of all time", "test.file");
INSERT INTO blog VALUES (3, 6, "This is impossible", "Privilege lvl 1 users can't write blogs.", "test.file");
