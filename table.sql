CREATE TABLE Division (
    ID int NOT NULL AUTO_INCREMENT,
    Division varchar(50) NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE Candidates (
    ID int NOT NULL AUTO_INCREMENT,
    Name varchar(50) NOT NULL,
    Division_ID int NOT NULL,
    PRIMARY KEY (ID),
    FOREIGN KEY (Division_ID) REFERENCES Division(ID)
);

CREATE TABLE Users (
    ID int NOT NULL AUTO_INCREMENT,
    Name varchar(50) NOT NULL,
    IdentityNumber varchar(30) NOT NULL,
    Password varchar(255) NOT NULL,  -- Hashed password
    RoleID int NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE User_Candidate (
    User_ID int NOT NULL,
    Candidate_ID int NOT NULL,
    Image BLOB,
    PRIMARY KEY (User_ID, Candidate_ID),
    FOREIGN KEY (User_ID) REFERENCES User(ID),
    FOREIGN KEY (Candidate_ID) REFERENCES Candidate(ID)
);
