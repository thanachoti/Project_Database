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

CREATE TABLE "user" (
    user_id serial PRIMARY KEY,
    username varchar(255) NOT NULL UNIQUE,
    email varchar(255) UNIQUE NOT NULL,
    password varchar(255) NOT NULL,
    subscription varchar(50) DEFAULT 'Free' CHECK (subscription IN ('Free', 'Basic', 'Premium')),
    registration timestamp DEFAULT CURRENT_TIMESTAMP,
    age int CHECK (age IS NULL OR age > 0),
    "role" varchar(10) DEFAULT 'User' CHECK ("role" IN ('User', 'Admin')),
    profile_pic bytea
);

CREATE TABLE CONTENT (
    content_id serial PRIMARY KEY,
    title varchar(255),
    description text,
    release_year int CHECK (release_year <= EXTRACT(YEAR FROM CURRENT_DATE)),
    duration int CHECK (duration > 0),
    content_type varchar(10) CHECK (content_type IN ('Movie', 'TV Show')),
    total_seasons int,
    thumbnail_url varchar(255),
    video_url varchar(255),
    rating float CHECK (rating >= 0 AND rating <= 10),
    director varchar(255),
    CHECK ((content_type = 'TV Show' AND total_seasons IS NOT NULL AND total_seasons > 0) OR (content_type = 'Movie' AND total_seasons = 0)),
    UNIQUE (title, release_year, content_type, director)
);

CREATE TABLE CATEGORY (
    category_id serial PRIMARY KEY,
    category_name varchar(255) UNIQUE
);

CREATE TABLE
LANGUAGE (
    language_id serial PRIMARY KEY,
    language_name varchar(255) UNIQUE
);

CREATE TABLE WATCH_HISTORY (
    history_id serial PRIMARY KEY,
    user_id int REFERENCES "user" (user_id) ON DELETE CASCADE,
    content_id int REFERENCES CONTENT (content_id) ON DELETE CASCADE,
    watched_timestamp timestamp DEFAULT CURRENT_TIMESTAMP,
    progress timestamp CHECK (progress IS NULL OR progress <= CURRENT_TIMESTAMP),
    language_preference varchar(255),
    cc_preference varchar(255),
    UNIQUE (user_id, content_id)
);

CREATE TABLE FAVORITE (
    favorite_id serial PRIMARY KEY,
    user_id int REFERENCES "user" (user_id) ON DELETE CASCADE,
    content_id int REFERENCES CONTENT (content_id) ON DELETE CASCADE,
    UNIQUE (user_id, content_id)
);

CREATE TABLE REVIEW (
    review_id serial PRIMARY KEY,
    user_id int REFERENCES "user" (user_id) ON DELETE CASCADE,
    content_id int REFERENCES CONTENT (content_id) ON DELETE CASCADE,
    rating int CHECK (rating >= 1 AND rating <= 5),
    review_text text,
    review_date timestamp,
    UNIQUE (user_id, content_id)
);

CREATE TABLE CONTENT_CATEGORY (
    content_category_id serial PRIMARY KEY,
    content_id int REFERENCES CONTENT (content_id) ON DELETE CASCADE,
    category_id int REFERENCES CATEGORY (category_id) ON DELETE CASCADE,
    UNIQUE (content_id, category_id)
);

CREATE TABLE CONTENT_LANGUAGE (
    id serial PRIMARY KEY,
    content_id int REFERENCES CONTENT (content_id) ON DELETE CASCADE,
    language_id int REFERENCES
    LANGUAGE (language_id) ON DELETE CASCADE,
    UNIQUE (content_id, language_id)
);

CREATE TABLE CONTENT_SUBTITLE (
    id serial PRIMARY KEY,
    content_id int REFERENCES CONTENT (content_id) ON DELETE CASCADE,
    language_id int REFERENCES
    LANGUAGE (language_id) ON DELETE CASCADE,
    UNIQUE (content_id, language_id)
);

INSERT INTO "user" (username, email, PASSWORD, subscription, age, ROLE)
    VALUES ('admin', 'admin@example.com', '$2a$10$V2z/xMaErmhg0eYMhDedruY32RYyT9FjUKsV.MCCn1ZLfpmHPBZJ6', 'Premium', 30, 'Admin');

INSERT INTO CONTENT (title, description, release_year, duration, content_type, total_seasons, thumbnail_url, video_url, rating, director)
    VALUES ('Inception', 'Description of Inception, an exciting and engaging storyline.', 2002, 46, 'TV Show', 5, 'https://encrypted-tbn3.gstatic.com/images?q=tbn:ANd9GcQovCe0H45fWwAtV31ajOdXRPTxSsMQgPIQ3lcZX_mAW0jXV3kH', 'https://youtu.be/gQ6cdfiIoiQ?si=SXT3KKa9Y1LM7FXK', 8.5, 'Christopher Nolan'),
    ('The Matrix', 'A mind-bending sci-fi adventure exploring simulated reality.', 1999, 136, 'Movie', 0, 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQCWXVvfvZR3oe7PCMM0exwV0dObOTKvLfSM-bjvKpQ1VegKXuCtq6aBrxqbIgUNxMbfavy', 'https://youtu.be/Qobz4DZ_ofs?si=a3MHz8KU_E1o0nWp', 8.7, 'Lana Wachowski, Lilly Wachowski'),
    ('Stranger Things', 'A supernatural mystery involving a group of kids in the 80s.', 2016, 50, 'TV Show', 4, 'https://m.media-amazon.com/images/M/MV5BMjg2NmM0MTEtYWY2Yy00NmFlLTllNTMtMjVkZjEwMGVlNzdjXkEyXkFqcGc@._V1_.jpg', 'https://youtu.be/KM2XiKcYJ10?si=GT5fC9M5nj3zuUjF', 8.8, 'Matt Duffer, Ross Duffer'),
    ('Interstellar', 'A team of explorers travel through a wormhole in space.', 2014, 169, 'Movie', 0, 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSngBJ0B7UDrLUkDlp6DCQLsEYuWR-DiHwbnxFFCniB3HiP3f3NZmR1-lKSC34ge6YXu4LX', 'https://youtu.be/kfQW2orL0hg?si=KTT61bnmKw3753yp', 8.6, 'Christopher Nolan'),
    ('Breaking Bad', 'A chemistry teacher turns to making meth to secure his family''s future.', 2008, 47, 'TV Show', 5, 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTOcWkpWG_NRrU2M8-WB8EbEcJk7smhdrY1eO0ttKXm0bo2ooOEWxk3zBSbsFrSgSJh2OEKOQ', 'https://youtu.be/DYYDxLIUASE?si=ezqzJFfltQcG4XGG', 9.5, 'Vince Gilligan'),
    ('The Office', 'A hilarious mockumentary about office workers at Dunder Mifflin.', 2005, 22, 'TV Show', 9, 'https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcQ00mxCZs-9YF6bPtPBe-dlgo6cnxI6klO2hHMtiQJpC_vHKQuNQKEB627kLUXoHXhFnwdBVw', 'https://youtu.be/Y205hADGEyY?si=_mX_CgrXYWo_sJUM', 8.9, 'Greg Daniels'),
    ('The Godfather', 'An epic tale of a mafia family and their rise to power.', 1972, 175, 'Movie', 0, 'https://encrypted-tbn2.gstatic.com/images?q=tbn:ANd9GcQAY2xsJVIZxm3K0gNtOMr9CSCvLdr5kdo3V3pv2HMuUkTBhFzRe5-b8NDRmO1mt5S5Xp_YyQ', 'https://youtu.be/TxanlC8ZnTY?si=bRNXVPa9aKecybH_', 9.2, 'Francis Ford Coppola'),
    ('The Crown', 'A historical drama about the reign of Queen Elizabeth II.', 2016, 58, 'TV Show', 6, 'https://stanforddaily.com/wp-content/uploads/2021/01/the-crown.png', 'https://youtu.be/kOaJmFaNt6o?si=zBaMogGkCr3e1gf7', 8.6, 'Peter Morgan'),
    ('Avengers: Endgame', 'The Avengers assemble for a final battle to save the universe.', 2019, 181, 'Movie', 0, 'https://upload.wikimedia.org/wikipedia/th/0/0d/Avengers_Endgame_poster.jpg', 'https://youtu.be/6MHS_BmxZMw?si=Bbxa2S-ue-Vo2S4E', 8.4, 'Anthony Russo, Joe Russo'),
    ('Sherlock', 'Modern-day version of the classic detective stories.', 2010, 90, 'TV Show', 4, 'https://encrypted-tbn2.gstatic.com/images?q=tbn:ANd9GcSKAffHEy-QCEkW_rqDKlHTcHELw7CBuxM3hi-vE7LFJM7yFZ4Msgeg75Kh1988BBF4Lrf21A', 'https://youtu.be/2hggbGAwQ2I?si=5x_vvj79BR-fuq1b', 9.1, 'Steven Moffat, Mark Gatiss'),
    ('Dark', 'A mysterious tale of time travel and a small German town.', 2017, 60, 'TV Show', 3, 'https://encrypted-tbn3.gstatic.com/images?q=tbn:ANd9GcS6ApenbR-AZkjJySI-VBzjMJYPWFoqUgCxfkGyvXpCru89imX7jzAdmaDSgEVOY4MIDnR_', 'https://youtu.be/UAv6vo8kLqs?si=CBDIpDBqlRJmwvxQ', 8.8, 'Baran bo Odar'),
    ('Pulp Fiction', 'A crime drama with interwoven stories and iconic characters.', 1994, 154, 'Movie', 0, 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSCqzGSUVlP74iyuFujryxWBYV6yqGirkn7BFHIJXEMzS4gNI-Z2wEKZsW6dLYA9B77BgyPdg', 'https://youtu.be/R0dyyTFHXcA?si=tPe5bnEsL8p0O0Kd', 8.9, 'Quentin Tarantino'),
    ('Friends', 'Six friends navigate life and love in New York City.', 1994, 22, 'TV Show', 10, 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRcG-LN-6lRuoMhV8Fvh9AXAm63iwuEdQPasRnB-0Bgjyt2RL5DA3M2jcOZ0x45pRxGb4YeSg', 'https://youtu.be/PnWtLyCMCMU?si=pqPuXVINAIPUu00o', 8.9, 'David Crane, Marta Kauffman'),
    ('Parasite', 'A dark comedy thriller about class and deception.', 2019, 132, 'Movie', 0, 'https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcTTRGm5Vxt-AKoe72ASaC0F1w58TkuIQTuYrjrzhHkAcZYXXUS9WQdAuaikkuRMX50MWN01lw', 'https://youtu.be/B_O7dEt08zA?si=49xL0RVQ1DPD3yCv', 8.6, 'Bong Joon Ho'),
    ('Black Mirror', 'Anthology series exploring the dark side of technology.', 2011, 60, 'TV Show', 6, 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSuehP2oimC52zxXcw2S2P1DDJ2TJcBhcqgjYd0NBz54rEAY8KC3krqy8R_HbRSi0bmZsxQ', 'https://youtu.be/fxXT1UZ8t1A?si=rmQAAw3jZwH2edNg', 8.8, 'Charlie Brooker'),
    ('The Witcher', 'A monster hunter navigates a world of magic and politics.', 2019, 60, 'TV Show', 3, 'https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcQrpQa4kPsu6vyXAGfOOHSuscIGwdQuxyOs0Lp-EndlqkhFdcxNPg0kt-lws9e49GAqx2wxRg', 'https://youtu.be/4bxb7gdMdLw?si=66PIeu4rBQKSqJo5', 8.2, 'Lauren Schmidt Hissrich'),
    ('The Mandalorian', 'A lone bounty hunter in the outer reaches of the galaxy.', 2019, 40, 'TV Show', 3, 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTPz4zyHFlnCZr0RuXDKfJOhPB83w0jh_RJ1utNuTGMd1-apSkHVSGEsfZ17_D31rtkhQpZ', 'https://youtu.be/8C1OOyaxthI?si=E7QgFKU6kAdBRbje', 8.7, 'Jon Favreau'),
    ('Titanic', 'A love story aboard the ill-fated RMS Titanic.', 1997, 195, 'Movie', 0, 'https://m.media-amazon.com/images/I/41MVlLiRqWL._AC_UF1000,1000_QL80_.jpg', 'https://youtu.be/fuSRjyR_ZJU?si=F27eLudq-I4ao-6O', 7.9, 'James Cameron'),
    ('Money Heist', 'Criminals plan the perfect heist on the Spanish Royal Mint.', 2017, 45, 'TV Show', 5, 'https://m.media-amazon.com/images/I/81WSF164GOL._AC_UF1000,1000_QL80_.jpg', 'https://youtu.be/futT5F-qx_A?si=EMbQ_jopcuWm6tDs', 8.3, '╬ô├▓┬╝Γö£Γöñ╬ô├╢┬ú╬ô├▓├│╬ô├╢┬╝Γö£ΓòæΓò¼├┤Γö£ΓòóΓö¼├║Γò¼├┤Γö£ΓûôΓö¼├ælex Pina'),
    ('The Social Network', 'The rise of Facebook and the legal battles that followed.', 2010, 120, 'Movie', 0, 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSEawbHeCDXCGG-gxJQFrsX3FfZbyEkCpidng&s', 'https://youtu.be/tSfoYYdTOvE?si=C28rGNShGclursq5', 7.7, 'David Fincher'),
    ('Game of Thrones', 'Noble families vie for control of the Iron Throne.', 2011, 55, 'TV Show', 8, 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSTDou8fFUXJ86ACUbf7--1Ek1dLEtpN7WQwA&s', 'https://youtu.be/eDOPuOsnyUg?si=fvS7Mm39GpYVPszz', 9.3, 'David Benioff, D. B. Weiss'),
    ('Loki', 'The God of Mischief explores alternate timelines.', 2021, 50, 'TV Show', 2, 'https://m.media-amazon.com/images/M/MV5BYzA2YjM2ZWQtYTZhMS00OTI3LTlhYzQtZjBiZWZkMDdlNjA5XkEyXkFqcGc@._V1_FMjpg_UX1000_.jpg', 'https://youtu.be/jEWaLRi_G60?si=aqI_ENKTwBpKldsf', 8.4, 'Kate Herron'),
    ('Joker', 'An origin story of the iconic villain in a gritty setting.', 2019, 122, 'Movie', 0, 'https://encrypted-tbn2.gstatic.com/images?q=tbn:ANd9GcRkNeYGwWeQEwOoPhxW93QIeNUWnLmEvMPwTw9AlDBGN4uXjIAcOEwz2z2yZL8BpXHp3ZYyjQ', 'https://www.youtube.com/live/0KDTneGVIhY?si=475jllWQ-zElS6ZS', 8.5, 'Todd Phillips'),
    ('Narcos', 'The rise and fall of drug kingpin Pablo Escobar.', 2015, 50, 'TV Show', 3, 'https://m.media-amazon.com/images/M/MV5BNzQwOTcwMzIwN15BMl5BanBnXkFtZTgwMjYxMTA0NjE@._V1_FMjpg_UX1000_.jpg', 'https://youtu.be/T1Etiu31aK8?si=x7qHLYpJbZ6Tl_gl', 8.8, 'Various Directors'),
    ('The Queen''s Gambit', 'A young chess prodigy battles addiction and rivals.', 2020, 55, 'TV Show', 1, 'https://m.media-amazon.com/images/M/MV5BMmRlNjQxNWQtMjk1OS00N2QxLTk0YWQtMzRhYjY5YTFhNjMxXkEyXkFqcGc@._V1_FMjpg_UX1000_.jpg', 'https://youtu.be/b_LJol1RF1I?si=tjQahlJKvhmKQKJa', 8.6, 'Scott Frank'),
    ('Dune', 'A noble family''s son becomes a messiah on a desert planet.', 2021, 155, 'Movie', 0, 'https://encrypted-tbn2.gstatic.com/images?q=tbn:ANd9GcSt7xlJEzb-xopDqcQ6iw9SbY8PAlJN8H7DYUzTqmZkwLT9o8JXv6YWvDMGRKwkyRnf6RHzGg', 'https://youtu.be/Q56PMJbCFXQ?si=ioejSfdizZczgWSG', 8.1, 'Denis Villeneuve'),
    ('House of the Dragon', 'Prequel to Game of Thrones about House Targaryen.', 2022, 60, 'TV Show', 1, 'https://m.media-amazon.com/images/M/MV5BM2QzMGVkNjUtN2Y4Yi00ODMwLTg3YzktYzUxYjJlNjFjNDY1XkEyXkFqcGc@._V1_.jpg', 'https://youtu.be/Q0zMgGZCkpk?si=waxBOj2enlgGmB0R', 8.6, 'Various Directors'),
    ('Ford v Ferrari', 'Rivalry between Ford and Ferrari at Le Mans 1966.', 2019, 152, 'Movie', 0, 'https://m.media-amazon.com/images/M/MV5BOTBjNTEyNjYtYjdkNi00YzE5LTljYzUtZjVlYmYwZmJmZWYxXkEyXkFqcGc@._V1_.jpg', 'https://youtu.be/VUmmhslxg9s?si=tYVQg2B1ffE8GCwA', 8.1, 'James Mangold'),
    ('Chernobyl', 'The true story of the 1986 nuclear disaster.', 2019, 65, 'TV Show', 1, 'https://upload.wikimedia.org/wikipedia/en/a/a7/Chernobyl_2019_Miniseries.jpg', 'https://youtu.be/Iaj_T7k8FfE?si=iBeJIVKf9zz5vt9g', 9.4, 'Johan Renck'),
    ('Bohemian Rhapsody', 'A biographical film about Queen lead singer Freddie Mercury.', 2018, 134, 'Movie', 0, 'https://m.media-amazon.com/images/M/MV5BMTA2NDc3Njg5NDVeQTJeQWpwZ15BbWU4MDc1NDcxNTUz._V1_.jpg', 'https://youtu.be/WdyFrGMfEkk?si=tvbvDWc06yjJVY2b', 8.0, 'Bryan Singer');

-- Insert Base Data
INSERT INTO CATEGORY (category_name)
    VALUES ('Action'),
    ('Adventure'),
    ('Comedy'),
    ('Crime'),
    ('Drama'),
    ('Fantasy'),
    ('Horror'),
    ('Mystery'),
    ('Sci-Fi'),
    ('Thriller');

INSERT INTO CONTENT_CATEGORY (content_id, category_id)
    VALUES (1, 1),
    (2, 2),
    (3, 3),
    (4, 2),
    (5, 4),
    (1, 2),
    (3, 4),
    (6, 10),
    (7, 9),
    (8, 8),
    (9, 7),
    (10, 6),
    (11, 5),
    (12, 4),
    (13, 3),
    (14, 2),
    (15, 1),
    (16, 10),
    (17, 9),
    (18, 8),
    (19, 7),
    (20, 6),
    (21, 5),
    (22, 4),
    (23, 3),
    (24, 2),
    (25, 1),
    (26, 10),
    (27, 9),
    (28, 8),
    (29, 7),
    (30, 6);

INSERT INTO
LANGUAGE (language_name)
    VALUES ('English'),
    ('French'),
    ('Japanese'),
    ('Spanish');

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
    (5, 1),
    (6, 1),
    (7, 2),
    (8, 3),
    (9, 4),
    (10, 1),
    (11, 2),
    (12, 3),
    (13, 4),
    (14, 1),
    (15, 2),
    (16, 3),
    (17, 4),
    (18, 1),
    (19, 2),
    (20, 3),
    (21, 4),
    (22, 1),
    (23, 2),
    (24, 3),
    (25, 4),
    (26, 1),
    (27, 2),
    (28, 3),
    (29, 4),
    (30, 1);

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
    (5, 2),
    (6, 4),
    (7, 3),
    (8, 2),
    (9, 1),
    (10, 4),
    (11, 3),
    (12, 2),
    (13, 1),
    (14, 4),
    (15, 3),
    (16, 2),
    (17, 1),
    (18, 4),
    (19, 3),
    (20, 2),
    (21, 1),
    (22, 4),
    (23, 3),
    (24, 2),
    (25, 1),
    (26, 4),
    (27, 3),
    (28, 2),
    (29, 1),
    (30, 4);
