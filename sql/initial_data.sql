DROP TABLE IF EXISTS CONTENT_SUBTITLE;

DROP TABLE IF EXISTS CONTENT_LANGUAGE;

DROP TABLE IF EXISTS CONTENT_CATEGORY;

DROP TABLE IF EXISTS CONTENT;

DROP TABLE IF EXISTS LANGUAGE;

DROP TABLE IF EXISTS CATEGORY;

DROP TABLE IF EXISTS "user";

DROP TABLE IF EXISTS REVIEW;

DROP TABLE IF EXISTS FAVORITE;

DROP TABLE IF EXISTS WATCH_HISTORY;

-- Create Base Tables (Parent First)
CREATE TABLE "user" (
    user_id serial PRIMARY KEY,
    username varchar(255) NOT NULL,
    email varchar(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    subscription varchar(50),
    registration timestamp DEFAULT CURRENT_TIMESTAMP,
    age int,
    profile_pic varchar(255)
);

CREATE TABLE CONTENT (
    content_id serial PRIMARY KEY,
    title varchar(255),
    description text,
    release_year int,
    duration int,
    content_type varchar(10),
    total_seasons int,
    thumbnail_url varchar(255),
    video_url varchar(255),
    rating float
);

CREATE TABLE CATEGORY (
    category_id serial PRIMARY KEY,
    category_name varchar(255)
);

CREATE TABLE
LANGUAGE (
    language_id serial PRIMARY KEY,
    language_name varchar(255)
);

-- Create Related Tables (Child After Parent)
CREATE TABLE WATCH_HISTORY (
    history_id serial PRIMARY KEY,
    user_id int REFERENCES "user" (user_id) ON DELETE CASCADE,
    content_id int REFERENCES CONTENT (content_id) ON DELETE CASCADE,
    progress timestamp,
    language_preference varchar(255),
    cc_preference varchar(255)
);

CREATE TABLE FAVORITE (
    favorite_id serial PRIMARY KEY,
    user_id int REFERENCES "user" (user_id) ON DELETE CASCADE,
    content_id int REFERENCES CONTENT (content_id) ON DELETE CASCADE
);

CREATE TABLE REVIEW (
    review_id serial PRIMARY KEY,
    user_id int REFERENCES "user" (user_id) ON DELETE CASCADE,
    content_id int REFERENCES CONTENT (content_id) ON DELETE CASCADE,
    rating int,
    review_text text,
    review_date timestamp
);

CREATE TABLE CONTENT_CATEGORY (
    content_category_id serial PRIMARY KEY,
    content_id int REFERENCES CONTENT (content_id) ON DELETE CASCADE,
    category_id int REFERENCES CATEGORY (category_id) ON DELETE CASCADE
);

CREATE TABLE CONTENT_LANGUAGE (
    id serial PRIMARY KEY,
    content_id int REFERENCES CONTENT (content_id) ON DELETE CASCADE,
    language_id int REFERENCES
    LANGUAGE (language_id) ON DELETE CASCADE
);

CREATE TABLE CONTENT_SUBTITLE (
    id serial PRIMARY KEY,
    content_id int REFERENCES CONTENT (content_id) ON DELETE CASCADE,
    language_id int REFERENCES
    LANGUAGE (language_id) ON DELETE CASCADE
);

-- Insert Base Data
INSERT INTO CONTENT (title, description, release_year, duration, content_type, total_seasons, thumbnail_url, video_url, rating)
    VALUES ('Movie A', 'A great movie', 2021, 120, 'Movie', 0, 'movie_a.jpg', '', 4.5),
    ('Show B', 'An exciting show', 2022, 60, 'Show', 3, 'show_b.jpg', '', 4.2),
    ('Movie C', 'A comedy movie', 2023, 100, 'Movie', 0, 'movie_c.jpg', '', 4.8),
    ('Show D', 'A drama show', 2024, 45, 'Show', 2, 'show_d.jpg', '', 4.0),
    ('Movie E', 'A thriller movie', 2023, 110, 'Movie', 0, 'movie_e.jpg', '', 4.6);

INSERT INTO CATEGORY (category_id, category_name)
    VALUES (1, 'Action'),
    (2, 'Drama'),
    (3, 'Comedy'),
    (4, 'Thriller');

INSERT INTO CONTENT_CATEGORY (content_id, category_id)
    VALUES (1, 1),
    (2, 2),
    (3, 3),
    (4, 2),
    (5, 4),
    (1, 1),
    (3, 3);

INSERT INTO
LANGUAGE (language_id, language_name)
    VALUES (1, 'English'),
    (2, 'Spanish'),
    (3, 'French'),
    (4, 'German');

INSERT INTO CONTENT_LANGUAGE (content_id, language_id)
    VALUES (1, 1),
    (2, 1),
    (3, 2),
    (4, 1),
    (5, 4),
    (1, 2),
    (2, 3),
    (3, 1),
    (4, 3),
    (5, 1);

INSERT INTO CONTENT_SUBTITLE (content_id, language_id)
    VALUES (1, 1),
    (2, 1),
    (3, 2),
    (4, 3),
    (5, 4),
    (1, 2),
    (2, 3),
    (3, 1),
    (4, 1),
    (5, 2);
