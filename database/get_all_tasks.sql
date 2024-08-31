USE `task_manager`;
DROP procedure IF EXISTS `get_all_tasks`;

DELIMITER $$
USE `task_manager`$$
CREATE PROCEDURE `get_all_tasks`(IN firebase_uid VARCHAR(255))
BEGIN
SELECT task.task_id, task.user_id, task.task_name, task.task_detail, task.created_at
FROM task
LEFT JOIN users
	ON task.user_id = users.id
WHERE users.firebase_uid = firebase_uid;
END$$

DELIMITER ;