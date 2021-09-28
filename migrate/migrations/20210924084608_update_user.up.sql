ALTER TABLE user ADD `age` int(2);
ALTER TABLE user ADD `sex` int(1);

START TRANSACTION;
insert into user (username, password) values ("1", "123456");
insert into user (username, password) values ("2", "123456");
insert into user (username, password) values ("3", "123456");
insert into user (username, password) values ("4", "123456");
insert into user (username, password) values ("5", "123456");
COMMIT;

DROP PROCEDURE IF EXISTS gorm;
DELIMITER
CREATE PROCEDURE gorm()
  BEGIN
    DECLARE t_error INTEGER;
    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION SET t_error = 1;
    START TRANSACTION;
            insert into user (username, password) values ("4", "123456");
            insert into user (username, password) values ("5", "123456");
         IF t_error = 1 THEN
             ROLLBACK;
         ELSE
             COMMIT;
         END IF;
END
CALL gorm();
