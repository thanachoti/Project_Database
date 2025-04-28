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
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    subscription VARCHAR(50),
    registration TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    age INT,
    profile_pic VARCHAR(255)
);

CREATE TABLE CONTENT (
    content_id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    description TEXT,
    release_year INT,
    duration INT,
    content_type VARCHAR(10),
    total_seasons INT,
    thumbnail_url VARCHAR(255),
    video_url VARCHAR(255),
    rating FLOAT
);

CREATE TABLE CATEGORY (
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(255)
);

CREATE TABLE LANGUAGE (
    language_id SERIAL PRIMARY KEY,
    language_name VARCHAR(255)
);


-- Create Related Tables (Child After Parent)

CREATE TABLE WATCH_HISTORY (
    history_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES "user"(user_id) ON DELETE CASCADE,
    content_id INT REFERENCES CONTENT(content_id) ON DELETE CASCADE,
    progress TIMESTAMP,
    language_preference VARCHAR(255),
    cc_preference VARCHAR(255)
);


CREATE TABLE FAVORITE (
    favorite_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES "user"(user_id) ON DELETE CASCADE,
    content_id INT REFERENCES CONTENT(content_id) ON DELETE CASCADE
);

CREATE TABLE REVIEW (
    review_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES "user"(user_id) ON DELETE CASCADE,
    content_id INT REFERENCES CONTENT(content_id) ON DELETE CASCADE,
    rating INT,
    review_text TEXT,
    review_date TIMESTAMP
);

CREATE TABLE CONTENT_CATEGORY (
    content_category_id SERIAL PRIMARY KEY,
    content_id INT REFERENCES CONTENT(content_id) ON DELETE CASCADE,
    category_id INT REFERENCES CATEGORY(category_id) ON DELETE CASCADE
);

CREATE TABLE CONTENT_LANGUAGE (
    id SERIAL PRIMARY KEY,
    content_id INT REFERENCES CONTENT(content_id) ON DELETE CASCADE,
    language_id INT REFERENCES LANGUAGE(language_id) ON DELETE CASCADE
);

CREATE TABLE CONTENT_SUBTITLE (
    id SERIAL PRIMARY KEY,
    content_id INT REFERENCES CONTENT(content_id) ON DELETE CASCADE,
    language_id INT REFERENCES LANGUAGE(language_id) ON DELETE CASCADE
);


-- Insert Base Data

INSERT INTO CONTENT (title, description, release_year, duration, content_type, total_seasons, thumbnail_url, video_url, rating)
VALUES 
('Movie A', 'A great movie', 2021, 120, 'Movie', 0, 'movie_a.jpg', '', 4.5),
('Show B', 'An exciting show', 2022, 60, 'Show', 3, 'show_b.jpg', '', 4.2),
('Movie C', 'A comedy movie', 2023, 100, 'Movie', 0, 'movie_c.jpg', '', 4.8),
('Show D', 'A drama show', 2024, 45, 'Show', 2, 'show_d.jpg', '', 4.0),
('Movie E', 'A thriller movie', 2023, 110, 'Movie', 0, 'movie_e.jpg', '', 4.6);

INSERT INTO CATEGORY (category_id, category_name)
VALUES (1, 'Action'), (2, 'Drama'), (3, 'Comedy'), (4, 'Thriller');

INSERT INTO CONTENT_CATEGORY (content_id, category_id)
VALUES 
(1, 1),
(2, 2),
(3, 3),
(4, 2),
(5, 4),
(1, 1),
(3, 3);

INSERT INTO LANGUAGE (language_id, language_name)
VALUES (1, 'English'), (2, 'Spanish'), (3, 'French'), (4, 'German');

INSERT INTO CONTENT_LANGUAGE (content_id, language_id)
VALUES 
(1, 1), (2, 1), (3, 2), (4, 1), (5, 4),
(1, 2), (2, 3), (3, 1), (4, 3), (5, 1);

INSERT INTO CONTENT_SUBTITLE (content_id, language_id)
VALUES 
(1, 1), (2, 1), (3, 2), (4, 3), (5, 4),
(1, 2), (2, 3), (3, 1), (4, 1), (5, 2);
