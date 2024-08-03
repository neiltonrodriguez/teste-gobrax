-- -----------------------------------------------------
-- Schema teste_gobrax
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `teste_gobrax` DEFAULT CHARACTER SET utf8 ;
USE `teste_gobrax` ;

-- -----------------------------------------------------
-- Table `teste_gobrax`.`driver`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `teste_gobrax`.`driver` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `drivers_license` VARCHAR(3) NOT NULL,
  `phone` VARCHAR(12) NULL,
  `age` INT NULL,
  `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `teste_gobrax`.`vehicle`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `teste_gobrax`.`vehicle` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `driver_id` INT NULL DEFAULT NULL,
  `plate` VARCHAR(7) NOT NULL,
  `brand` VARCHAR(45) NOT NULL,
  `model` VARCHAR(45) NOT NULL,
  `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
  UNIQUE INDEX `placa_UNIQUE` (`plate` ASC) VISIBLE,
  INDEX `fk_driver_id_idx` (`driver_id` ASC) VISIBLE,
  UNIQUE INDEX `driver_id_UNIQUE` (`driver_id` ASC) VISIBLE,
  CONSTRAINT `fk_driver_id`
    FOREIGN KEY (`driver_id`)
    REFERENCES `teste_gobrax`.`driver` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE)
ENGINE = InnoDB;
