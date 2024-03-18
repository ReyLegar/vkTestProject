CREATE TYPE UserRole AS ENUM ('User', 'Admin');

CREATE TABLE Users (
    UserID SERIAL PRIMARY KEY,
    Username VARCHAR(50) UNIQUE NOT NULL,
    PasswordHash CHAR(256) NOT NULL,
    Role UserRole NOT NULL
);

CREATE TYPE Gender AS ENUM ('Male', 'Female');

CREATE TABLE Actors (
    ActorID SERIAL PRIMARY KEY,
    Name VARCHAR(100) NOT NULL,
    Gender Gender,
    BirthDate DATE
);

CREATE TABLE Movies (
    MovieID SERIAL PRIMARY KEY,
    Title VARCHAR(150) NOT NULL,
    Description VARCHAR(1000),
    ReleaseDate DATE,
    Rating DECIMAL(2,1) CHECK (Rating >= 0 AND Rating <= 10)
);

CREATE TABLE MovieActors (
    MovieID INT REFERENCES Movies(MovieID),
    ActorID INT REFERENCES Actors(ActorID),
    PRIMARY KEY(MovieID, ActorID)
);