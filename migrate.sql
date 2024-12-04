INSERT INTO Division (Division) 
VALUES ('Kahima HMTC'), 
       ('Kabem FTEIC'), 
       ('Presbem ITS');

INSERT INTO Candidate (Name, Division_ID) 
VALUES ('John Doe', 1), -- John Doe untuk divisi Finance
       ('Jane Smith', 2), -- Jane Smith untuk divisi Human Resources
       ('Mike Johnson', 3), -- Mike Johnson untuk divisi Marketing
       ('Emily Davis', 3); -- Emily Davis untuk divisi IT

INSERT INTO Users (Name, IdentityNumber, Password, RoleID) 
VALUES ('admin', 'admin', 'password', 1), 
       ('Alice Brown', '1234567890', 'password1', 2), 
       ('Bob White', '0987654321', 'password2', 2);
