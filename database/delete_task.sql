USE `task_manager`;
DROP procedure IF EXISTS `task_manager`.`delete_task`;
;

DELIMITER $$
USE `task_manager`$$
CREATE PROCEDURE `delete_task`(IN firebase_uid VARCHAR(255), IN task_id INT)
BEGIN
DELETE FROM task
WHERE task.task_id = task_id AND task.user_id = (
												SELECT id
                                                FROM users
                                                WHERE users.firebase_uid = firebase_uid);
END$$

DELIMITER ;
;
