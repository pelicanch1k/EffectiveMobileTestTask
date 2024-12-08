CREATE TABLE songs (
    id SERIAL PRIMARY KEY,

    genre VARCHAR(30),
    song VARCHAR(30),

    releaseDate VARCHAR(30),
    text TEXT,
    link VARCHAR(255)
);