CREATE TABLE users(
  id BIGINT NOT NULL AUTO_INCREMENT,
  serial VARCHAR(255) NOT NULL,
  username VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  CONSTRAINT UNIQUE (username),
  CONSTRAINT UNIQUE (serial),
  CONSTRAINT PRIMARY KEY (id)
);

CREATE TABLE attachments(
  id BIGINT NOT NULL AUTO_INCREMENT,
  serial VARCHAR(255) NOT NULL,
  path TEXT,
  name VARCHAR(255) NOT NULL,
  is_active TINYINT(1) NOT NULL DEFAULT 0,
  user_id BIGINT NOT NULL,
  CONSTRAINT FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT UNIQUE (serial),
  CONSTRAINT PRIMARY KEY (id)
);

DELIMITER ;;
CREATE TRIGGER `users_before_insert` 
BEFORE INSERT ON `users` FOR EACH ROW 
BEGIN
  IF NEW.serial IS NULL OR NEW.serial = '' THEN
    SET NEW.serial = SHA2(RAND(), 256);
    SET NEW.password = SHA2(NEW.password, 256);
  END IF;
END;;
DELIMITER ;

DELIMITER ;;
CREATE TRIGGER `users_before_update` 
BEFORE UPDATE ON `users` FOR EACH ROW
BEGIN
  IF (OLD.password <> NEW.password) THEN
    SET NEW.password = SHA2(NEW.password, 256);
  END IF;
END;;
DELIMITER ;

DELIMITER ;;
CREATE TRIGGER `attachments_before_insert` 
BEFORE INSERT ON `attachments` FOR EACH ROW 
BEGIN
  IF NEW.serial IS NULL OR NEW.serial = '' THEN
    SET NEW.serial = SHA2(RAND(), 256);
  END IF;
END;;
DELIMITER ;