USE `task_manager`;
DROP procedure IF EXISTS `update_task`;

USE `task_manager`;
DROP procedure IF EXISTS `task_manager`.`update_task`;
;

DELIMITER $$
USE `task_manager`$$
CREATE DEFINER=`root`@`localhost` PROCEDURE `update_task`(IN firebase_uid VARCHAR(255), IN task_id INT, IN task_name VARCHAR(255), IN task_detail TEXT)
BEGIN
UPDATE task
SET task_name = task_name, task_detail = task_detail, created_at = CURRENT_TIMESTAMP
WHERE task.task_id = task_id AND task.user_id = (
									SELECT users.id
                                    FROM users
                                    WHERE users.firebase_uid = firebase_uid)
;
END$$

DELIMITER ;
;

