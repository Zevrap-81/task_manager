SELECT * 
FROM users; 

SELECT * 
FROM task; 

INSERT INTO users (firebase_uid, username, email) VALUES ("dummyid", "zevrap", "parvez@gmail.com"); 
INSERT INTO task (user_id, task_name, task_detail) VALUES ( 1, "Frist task", "Dummy task that has to be done");

CALL create_task("dummyid", "Second task", "Dummy task that has to be done");
CALL get_all_tasks("dummyid");
CALL get_task("dummyid", 1);
CALL create_task("dummyid", "Third task", "Dummy task that has to be done");
CALL update_task("dummyid", 3, "Third task new", "Dummy task that has to be done");
CALL delete_task("dummyid", 3);
