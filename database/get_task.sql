USE `task_manager`;
DROP procedure IF EXISTS `get_task`;

DELIMITER $$
USE `task_manager`$$
CREATE PROCEDURE `get_task`(IN firebase_uid VARCHAR(255), IN task_id INT)
BEGIN
WITH user_tasks (task_id, user_id, task_name, task_detail, created_at) AS
(
SELECT task.task_id, task.user_id, task.task_name, task.task_detail, task.created_at
FROM task
LEFT JOIN users
	ON task.user_id = users.id
WHERE users.firebase_uid = firebase_uid) 

SELECT *
FROM user_tasks
WHERE user_tasks.task_id = task_id
;
END$$

DELIMITER ;