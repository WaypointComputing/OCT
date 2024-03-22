DROP TABLE IF EXISTS response;
DROP TABLE IF EXISTS session;
DROP TABLE IF EXISTS question;
DROP TABLE IF EXISTS trait;
DROP TABLE IF EXISTS quiz;
DROP TABLE IF EXISTS blog;
DROP TABLE IF EXISTS user;

CREATE TABLE IF NOT EXISTS quiz (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(32) NOT NULL,
    desc VARCHAR(128)
);

CREATE TABLE IF NOT EXISTS trait (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(32) NOT NULL,
    desc VARCHAR(128),
    quiz_id INTEGER DEFAULT 5,
    FOREIGN KEY(quiz_id) REFERENCES quiz(id)
);

CREATE TABLE IF NOT EXISTS question (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    question VARCHAR(128) NOT NULL,
    trait_id INTEGER NOT NULL,
    quiz_id INTEGER NOT NULL,
    FOREIGN KEY(trait_id) REFERENCES trait(id)
    FOREIGN KEY(quiz_id) REFERENCES quiz(id)
);

CREATE TABLE IF NOT EXISTS session (
    quiz_id INTEGER,
    user_id INTEGER,
    quiz_completed SMALLINT DEFAULT 0 NOT NULL,
    CONSTRAINT CHK_quiz_completed CHECK (quiz_completed == 0 OR quiz_completed == 1),
    PRIMARY KEY(quiz_id, user_id)
);

CREATE TABLE IF NOT EXISTS response (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    question_id INTEGER NOT NULL,
    session_id INTEGER NOT NULL,
    weight INTEGER NOT NULL,
    CONSTRAINT CHK_Weight CHECK (weight > 0 AND weight <= 5),
    FOREIGN KEY(question_id) REFERENCES question(id)
    FOREIGN KEY(session_id) REFERENCES session(id)
);

CREATE TABLE IF NOT EXISTS user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(32) NOT NULL,
    email VARCHAR(32) NOT NULL UNIQUE,
    pwd_hash VARCHAR(128) NOT NULL,
    privileges INTEGER DEFAULT 1 NOT NULL,
    CONSTRAINT CHK_Privileges CHECK (privileges > 0 AND privileges <= 3)
);

CREATE TABLE IF NOT EXISTS blog (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    author_id INTEGER NOT NULL,
    title VARCHAR(32) NOT NULL,
    tagline VARCHAR(128) NULL DEFAULT NULL,
    blog_filepath VARCHAR(32) NOT NULL,
    FOREIGN KEY(author_id) REFERENCES user(id)
);
