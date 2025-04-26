DROP TABLE IF EXISTS CONTENT_SUBTITLE;

DROP TABLE IF EXISTS CONTENT_LANGUAGE;

DROP TABLE IF EXISTS CONTENT_CATEGORY;

DROP TABLE IF EXISTS CONTENT;

DROP TABLE IF EXISTS LANGUAGE;

DROP TABLE IF EXISTS CATEGORY;

DROP TABLE IF EXISTS "user";

DROP TABLE IF EXISTS REVIEW;

DROP TABLE IF EXISTS FAVORITE;

CREATE TABLE FAVORITE (
    favorite_id serial PRIMARY KEY,
    user_id int REFERENCES "user"(user_id) ON DELETE CASCADE,
    content_id int REFERENCES CONTENT(content_id) ON DELETE CASCADE
);

CREATE TABLE REVIEW (
    review_id serial PRIMARY KEY,
    user_id int REFERENCES user (user_id) ON DELETE CASCADE,
    content_id int REFERENCES content (content_id) ON DELETE CASCADE
    rating int,
    review_text text,
    review_date timestamp
);

CREATE TABLE "user" (
    user_id serial PRIMARY KEY,
    username varchar(255),
    email varchar(255),
    password VARCHAR(255)
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

INSERT INTO CONTENT (title, description, release_year, duration, content_type, total_seasons, thumbnail_url, video_url, rating)
    VALUES ('Movie A', 'A great movie', 2021, 120, 'Movie', 0, 'movie_a.jpg', '', 4.5),
    ('Show B', 'An exciting show', 2022, 60, 'Show', 3, 'show_b.jpg', '', 4.2),
    ('Movie C', 'A comedy movie', 2023, 100, 'Movie', 0, 'movie_c.jpg', '', 4.8),
    ('Show D', 'A drama show', 2024, 45, 'Show', 2, 'show_d.jpg', '', 4.0),
    ('Movie E', 'A thriller movie', 2023, 110, 'Movie', 0, 'movie_e.jpg', '', 4.6);

CREATE TABLE CATEGORY (
    category_id serial PRIMARY KEY,
    category_name varchar(255)
);

INSERT INTO CATEGORY (category_id, category_name)
    VALUES (1, 'Action'),
    (2, 'Drama'),
    (3, 'Comedy'),
    (4, 'Thriller');

CREATE TABLE CONTENT_CATEGORY (
    content_category_id serial PRIMARY KEY,
    content_id int,
    category_id int,
    FOREIGN KEY (content_id) REFERENCES CONTENT (content_id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES CATEGORY (category_id) ON DELETE CASCADE
);

INSERT INTO CONTENT_CATEGORY (content_id, category_id)
    VALUES (1, 1), -- Movie A - Action
    (2, 2), -- Show B - Drama
    (3, 3), -- Movie C - Comedy
    (4, 2), -- Show D - Drama
    (5, 4), -- Movie E - Thriller
    (1, 1), -- Movie A - Action.  Fix:  Using 1 again.
    (3, 3);

-- Movie C - Comedy.  Fix: Using 3 again.
CREATE TABLE
LANGUAGE (
    language_id serial PRIMARY KEY,
    language_name varchar(255)
);

INSERT INTO
LANGUAGE (language_id, language_name)
    VALUES (1, 'English'),
    (2, 'Spanish'),
    (3, 'French'),
    (4, 'German');

CREATE TABLE CONTENT_LANGUAGE (
    id serial PRIMARY KEY,
    content_id int,
    language_id int,
    FOREIGN KEY (content_id) REFERENCES CONTENT (content_id) ON DELETE CASCADE,
    FOREIGN KEY (language_id) REFERENCES
    LANGUAGE (language_id) ON DELETE CASCADE
);

INSERT INTO CONTENT_LANGUAGE (content_id, language_id)
    VALUES (1, 1), -- Movie A - English
    (2, 1), -- Show B - English
    (3, 2), -- Movie C - Spanish
    (4, 1), -- Show D - English
    (5, 4), -- Movie E - German
    (1, 2), -- Movie A - Spanish. Fix: Changed from 6 to 1
    (2, 3), -- Show B - French.  Fix: Changed from 7 to 2
    (3, 1), -- Movie C - English. Fix: Changed from 8 to 3
    (4, 3), -- Show D - French.  Fix: Changed from 9 to 4
    (5, 1);

-- Movie E - English. Fix: Changed from 10 to 5
CREATE TABLE CONTENT_SUBTITLE (
    id serial PRIMARY KEY,
    content_id int,
    language_id int,
    FOREIGN KEY (content_id) REFERENCES CONTENT (content_id) ON DELETE CASCADE,
    FOREIGN KEY (language_id) REFERENCES
    LANGUAGE (language_id) ON DELETE CASCADE
);

INSERT INTO CONTENT_SUBTITLE (content_id, language_id)
    VALUES (1, 1), -- Movie A - English
    (2, 1), -- Show B - English
    (3, 2), -- Movie C - Spanish
    (4, 3), -- Show D - French
    (5, 4), -- Movie E - German
    (1, 2), -- Movie A - Spanish. Fix: Changed from 6 to 1
    (2, 3), -- Show B - French. Fix: Changed from 7 to 2
    (3, 1), -- Movie C - English. Fix: Changed from 8 to 3
    (4, 1), -- Show D - English. Fix: Changed from 9 to 4
    (5, 2);

-- Movie E - Spanish. Fix: Changed from 10 to 5
