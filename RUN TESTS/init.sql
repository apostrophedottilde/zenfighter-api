CREATE TABLE IF NOT EXISTS `zenfightertest`.`fighters` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(16) NOT NULL,
  `strength` INT(11) NOT NULL,
  `weaponpower` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `name_UNIQUE` (`name` ASC),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC))
ENGINE = InnoDB
AUTO_INCREMENT = 159
DEFAULT CHARACTER SET = latin1;