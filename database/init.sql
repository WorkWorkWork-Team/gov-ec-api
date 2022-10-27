CREATE DATABASE devDB;
USE devDB

CREATE TABLE Province
(
    ProvinceID INTEGER,
    Name VARCHAR(255),
    PRIMARY KEY (ProvinceID)
);

CREATE TABLE District
(
    DistrictID INTEGER,
    Name VARCHAR(255),
    ProvinceID INTEGER,
    CONSTRAINT fk_province FOREIGN KEY (ProvinceID) REFERENCES Province(ProvinceID),
    PRIMARY KEY (DistrictID)
);

CREATE TABLE Population
(
    CitizenID BIGINT,
    LazerID VARCHAR(255),
    Name VARCHAR(255),
    Lastname VARCHAR(255),
    Birthday DATE,
    Nationality VARCHAR(255),
    DistrictID INTEGER,
    CONSTRAINT fk_district FOREIGN KEY (DistrictID) REFERENCES District(DistrictID),
    PRIMARY KEY (CitizenID)
);

CREATE TABLE Candidate
(
    ID INTEGER,
    CitizenID BIGINT,
    CONSTRAINT fk_citizen_candidate FOREIGN KEY (CitizenID) REFERENCES Population(CitizenID),
    PRIMARY KEY (ID)
);

CREATE TABLE Mp
(
    ID INTEGER,
    CitizenID BIGINT,
    CONSTRAINT fk_citizen_mp FOREIGN KEY (CitizenID) REFERENCES Population(CitizenID),
    PRIMARY KEY (ID)
);