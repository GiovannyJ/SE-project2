-- Disable foreign key checks (if necessary)
SET foreign_key_checks = 0;

-- Delete all rows from the tables
DELETE FROM ACCOUNTS;
DELETE FROM IMAGES;
DELETE FROM COMMENTS;
DELETE FROM POSTS;

-- Reset the seed for auto-incrementing columns
ALTER TABLE ACCOUNTS AUTO_INCREMENT = 1;
ALTER TABLE IMAGES AUTO_INCREMENT = 1;
ALTER TABLE COMMENTS AUTO_INCREMENT = 1;
ALTER TABLE POSTS AUTO_INCREMENT = 1;

-- Enable foreign key checks (if disabled)
SET foreign_key_checks = 1;

INSERT INTO ACCOUNTS (fname, lname, fullname, email, pwd, pnum, username, accesslvl)
VALUES(
'guest',
'guest',
'guest guest',
'guest@gmail.com',
'guest',
0000000000,
'guest',
)

INSERT INTO IMAGES(imgname, size, date)
VALUES(
'img1.png',
'0',
CURDATE()
)
