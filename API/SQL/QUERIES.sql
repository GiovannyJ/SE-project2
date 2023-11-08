--GET USER
SELECT * FROM ACCOUNT_INFO
--GET USER W PARAMS
SELECT * FROM ACCOUNT_INFO WHERE id = 1

--CREATE USER
INSERT INTO `ACCOUNT_INFO` (`fname`, `lname`, `fullname`, `email`, `pwd`, `pnum`, `age`, `username`)
VALUES ('Giovanny', 'Joseph', 'Giovanny Joseph', 'giovanny@example.com', 'password1', '1234567890', 25, 'giovannyj')

--UPDATE USER
UPDATE ACCOUNT_INFO 
SET 
id='4',
fname='madA',
lname='Jablonka',
fullname='Adam Jablonka',
email='ajjablonka@gmail.com',
pwd='password',
pnum='17187155801',
age='22',
username='WebsiteLover'
WHERE 
id='4' AND 
fname='Adam' AND
lname='Jablonka' AND 
fullname='Adam Jablonka' AND 
email='ajjablonka@gmail.com' AND 
pwd='password' AND 
pnum='17187155801' AND 
age='22' AND 
username='WebsiteLover'