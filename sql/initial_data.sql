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
    profile_pic BYTEA
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
VALUES
    ('Inception', 'Description of Inception, an exciting and engaging storyline.', 2002, 46, 'TV Show', 5, 'https://example.com/thumbnails/inception.jpg', 'https://https://youtu.be/gQ6cdfiIoiQ?si=SXT3KKa9Y1LM7FXK', 8.5),
    ('The Matrix', 'A mind-bending sci-fi adventure exploring simulated reality.', 1999, 136, 'Movie', 0, 'https://example.com/thumbnails/matrix.jpg', 'https://youtu.be/Qobz4DZ_ofs?si=a3MHz8KU_E1o0nWp', 8.7),
    ('Stranger Things', 'A supernatural mystery involving a group of kids in the 80s.', 2016, 50, 'TV Show', 4, 'https://example.com/thumbnails/stranger_things.jpg', 'https://youtu.be/KM2XiKcYJ10?si=GT5fC9M5nj3zuUjF', 8.8),
    ('Interstellar', 'A team of explorers travel through a wormhole in space.', 2014, 169, 'Movie', 0, 'https://example.com/thumbnails/interstellar.jpg', 'https://youtu.be/kfQW2orL0hg?si=KTT61bnmKw3753yp', 8.6),
    ('Breaking Bad', 'A chemistry teacher turns to making meth to secure his family''s future.', 2008, 47, 'TV Show', 5, 'https://example.com/thumbnails/breaking_bad.jpg', 'https://youtu.be/DYYDxLIUASE?si=ezqzJFfltQcG4XGG', 9.5),
    ('The Office', 'A hilarious mockumentary about office workers at Dunder Mifflin.', 2005, 22, 'TV Show', 9, 'https://example.com/thumbnails/the_office.jpg', 'https://youtu.be/Y205hADGEyY?si=_mX_CgrXYWo_sJUM', 8.9),
    ('The Godfather', 'An epic tale of a mafia family and their rise to power.', 1972, 175, 'Movie', 0, 'https://example.com/thumbnails/godfather.jpg', 'https://youtu.be/TxanlC8ZnTY?si=bRNXVPa9aKecybH_', 9.2),
    ('The Crown', 'A historical drama about the reign of Queen Elizabeth II.', 2016, 58, 'TV Show', 6, 'https://example.com/thumbnails/the_crown.jpg', 'https://youtu.be/kOaJmFaNt6o?si=zBaMogGkCr3e1gf7', 8.6),
    ('Avengers: Endgame', 'The Avengers assemble for a final battle to save the universe.', 2019, 181, 'Movie', 0, 'https://example.com/thumbnails/endgame.jpg', 'https://youtu.be/6MHS_BmxZMw?si=Bbxa2S-ue-Vo2S4E', 8.4),
    ('Sherlock', 'Modern-day version of the classic detective stories.', 2010, 90, 'TV Show', 4, 'https://example.com/thumbnails/sherlock.jpg', 'https://youtu.be/2hggbGAwQ2I?si=5x_vvj79BR-fuq1b', 9.1),
    ('Dark', 'A mysterious tale of time travel and a small German town.', 2017, 60, 'TV Show', 3, 'https://example.com/thumbnails/dark.jpg', 'https://youtu.be/UAv6vo8kLqs?si=CBDIpDBqlRJmwvxQ', 8.8),
    ('Pulp Fiction', 'A crime drama with interwoven stories and iconic characters.', 1994, 154, 'Movie', 0, 'https://example.com/thumbnails/pulp_fiction.jpg', 'https://youtu.be/R0dyyTFHXcA?si=tPe5bnEsL8p0O0Kd', 8.9),
    ('Friends', 'Six friends navigate life and love in New York City.', 1994, 22, 'TV Show', 10, 'https://example.com/thumbnails/friends.jpg', 'https://youtu.be/PnWtLyCMCMU?si=pqPuXVINAIPUu00o', 8.9),
    ('Parasite', 'A dark comedy thriller about class and deception.', 2019, 132, 'Movie', 0, 'https://example.com/thumbnails/parasite.jpg', 'https://youtu.be/B_O7dEt08zA?si=49xL0RVQ1DPD3yCv', 8.6),
    ('Black Mirror', 'Anthology series exploring the dark side of technology.', 2011, 60, 'TV Show', 6, 'https://example.com/thumbnails/black_mirror.jpg', 'https://youtu.be/fxXT1UZ8t1A?si=rmQAAw3jZwH2edNg', 8.8),
    ('The Witcher', 'A monster hunter navigates a world of magic and politics.', 2019, 60, 'TV Show', 3, 'https://example.com/thumbnails/witcher.jpg', 'https://youtu.be/4bxb7gdMdLw?si=66PIeu4rBQKSqJo5', 8.2),
    ('The Mandalorian', 'A lone bounty hunter in the outer reaches of the galaxy.', 2019, 40, 'TV Show', 3, 'https://example.com/thumbnails/mandalorian.jpg', 'https://youtu.be/8C1OOyaxthI?si=E7QgFKU6kAdBRbje', 8.7),
    ('Titanic', 'A love story aboard the ill-fated RMS Titanic.', 1997, 195, 'Movie', 0, 'https://example.com/thumbnails/titanic.jpg', 'https://youtu.be/fuSRjyR_ZJU?si=F27eLudq-I4ao-6O', 7.9),
    ('Money Heist', 'Criminals plan the perfect heist on the Spanish Royal Mint.', 2017, 45, 'TV Show', 5, 'https://example.com/thumbnails/money_heist.jpg', 'https://youtu.be/futT5F-qx_A?si=EMbQ_jopcuWm6tDs', 8.3),
    ('The Social Network', 'The rise of Facebook and the legal battles that followed.', 2010, 120, 'Movie', 0, 'https://example.com/thumbnails/social_network.jpg', 'https://youtu.be/tSfoYYdTOvE?si=C28rGNShGclursq5', 7.7),
    ('Game of Thrones', 'Noble families vie for control of the Iron Throne.', 2011, 55, 'TV Show', 8, 'https://example.com/thumbnails/got.jpg', 'https://youtu.be/eDOPuOsnyUg?si=fvS7Mm39GpYVPszz', 9.3),
    ('Loki', 'The God of Mischief explores alternate timelines.', 2021, 50, 'TV Show', 2, 'https://example.com/thumbnails/loki.jpg', 'https://youtu.be/jEWaLRi_G60?si=aqI_ENKTwBpKldsf', 8.4),
    ('Joker', 'An origin story of the iconic villain in a gritty setting.', 2019, 122, 'Movie', 0, 'https://example.com/thumbnails/joker.jpg', 'https://www.youtube.com/live/0KDTneGVIhY?si=475jllWQ-zElS6ZS', 8.5),
    ('Narcos', 'The rise and fall of drug kingpin Pablo Escobar.', 2015, 50, 'TV Show', 3, 'https://example.com/thumbnails/narcos.jpg', 'https://youtu.be/T1Etiu31aK8?si=x7qHLYpJbZ6Tl_gl', 8.8),
    ('The Queen''s Gambit', 'A young chess prodigy battles addiction and rivals.', 2020, 55, 'TV Show', 1, 'https://example.com/thumbnails/queens_gambit.jpg', 'https://youtu.be/b_LJol1RF1I?si=tjQahlJKvhmKQKJa', 8.6),
    ('Dune', 'A noble family''s son becomes a messiah on a desert planet.', 2021, 155, 'Movie', 0, 'https://example.com/thumbnails/dune.jpg', 'https://youtu.be/Q56PMJbCFXQ?si=ioejSfdizZczgWSG', 8.1),
    ('House of the Dragon', 'Prequel to Game of Thrones about House Targaryen.', 2022, 60, 'TV Show', 1, 'https://example.com/thumbnails/house_dragon.jpg', 'https://youtu.be/Q0zMgGZCkpk?si=waxBOj2enlgGmB0R', 8.6),
    ('Ford v Ferrari', 'Rivalry between Ford and Ferrari at Le Mans 1966.', 2019, 152, 'Movie', 0, 'https://example.com/thumbnails/ford_ferrari.jpg', 'https://youtu.be/VUmmhslxg9s?si=tYVQg2B1ffE8GCwA', 8.1),
    ('Chernobyl', 'The true story of the 1986 nuclear disaster.', 2019, 65, 'TV Show', 1, 'https://example.com/thumbnails/chernobyl.jpg', 'https://youtu.be/Iaj_T7k8FfE?si=iBeJIVKf9zz5vt9g', 9.4),
    ('Bohemian Rhapsody', 'A biographical film about Queen’s lead singer Freddie Mercury.', 2018, 134, 'Movie', 0, 'https://example.com/thumbnails/bohemian.jpg', 'https://youtu.be/WdyFrGMfEkk?si=tvbvDWc06yjJVY2b', 8.0);

INSERT INTO CATEGORY (category_id, category_name)
    VALUES 
    (1, 'Action'),
    (2, 'Adventure'),
    (3, 'Comedy'),
    (4, 'Crime'),
    (5, 'Drama'),
    (6, 'Fantasy'),
    (7, 'Horror'),
    (8, 'Mystery'),
    (9, 'Sci-Fi'),
    (10, 'Thriller');

INSERT INTO CONTENT_CATEGORY (content_id, category_id)
    VALUES (1, 1),
    (2, 2),
    (3, 3),
    (4, 2),
    (5, 4),
    (1, 1),
    (3, 3);

INSERT INTO LANGUAGE (language_id, language_name)
    VALUES
(1, 'English'),
(2, 'French'),
(3, 'Japanese'),
(4, 'Spanish');


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
