USE `task_manager`;
DROP procedure IF EXISTS `create_task`;

DELIMITER $$
USE `task_manager`$$
CREATE PROCEDURE `create_task`(IN 	firebase_uid VARCHAR(255), IN task_name VARCHAR(255), IN task_detail TEXT)
BEGIN
INSERT INTO task (user_id, task_name, task_detail) 
SELECT id, task_name, task_detail
FROM users
WHERE users.firebase_uid = firebase_uid;
END$$

DELIMITER ;